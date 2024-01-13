package archivator

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func ArchiveFiles(filesPaths []string, outputFilename string) string {
	// Files which to include in the tar.gz archive

	currentTime := time.Now()
	timestamp := fmt.Sprintf("%v-%v-%v_%v-%v-%v",
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Day(),
		int(currentTime.Month()),
		currentTime.Year())

	outputArchivePath := fmt.Sprintf("%v_%v.tar.gz", outputFilename, timestamp)
	// Create output file
	out, err := os.Create(outputArchivePath)
	if err != nil {
		log.Fatalln("Error writing archive:", err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatalln("Error closing archive:", err)
		}
	}(out)

	// Create the archive and write the output to the "out" Writer
	err = createArchive(filesPaths, out)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting working directory:", err)
	}

	return fmt.Sprintf("%s/%s", wd ,outputArchivePath)
}

func createArchive(filesPaths []string, buf io.Writer) error {
	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	gw := gzip.NewWriter(buf)
	defer func(gw *gzip.Writer) {
		err := gw.Close()
		if err != nil {
			log.Fatalln("Error closing gzip writer:", err)
		}
	}(gw)
	tw := tar.NewWriter(gw)
	defer func(tw *tar.Writer) {
		err := tw.Close()
		if err != nil {
			log.Fatalln("Error closing tar writer:", err)
		}
	}(tw)

	// Iterate over files and add them to the tar archive

	err := handleFileOrFolder(filesPaths, tw)
	if err != nil {
		return err
	}
	return nil
}

func handleFileOrFolder(filesPaths []string, tw *tar.Writer) error {
	for _, file := range filesPaths {
		fileInfo, err := os.Stat(file)
		if err != nil {
			return fmt.Errorf("error getting file info for %s: %w", file, err)
		}
		if fileInfo.IsDir() {
			filesInDir, err := os.ReadDir(file)
			if err != nil {
				return fmt.Errorf("error reading directory %s: %w", file, err)
			}

			var arrayOfFilePaths []string
			for _, fileName := range filesInDir {
				fullFilePath := fmt.Sprintf("%s/%s", file, fileName.Name())
				arrayOfFilePaths = append(arrayOfFilePaths, fullFilePath)
			}

			if err := handleFileOrFolder(arrayOfFilePaths, tw); err != nil {
				return err
			}
		} else {
			if err := addToArchive(tw, file); err != nil {
				return fmt.Errorf("error adding %s to archive: %w", file, err)
			}
		}
	}

	return nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	// Open the file which will be written into the archive
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("Error closing file:", err)
		}
	}(file)

	// Get FileInfo about our file providing file size, mode, etc.
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create a tar Header from the FileInfo data
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	// Use full path as name (FileInfoHeader only takes the basename)
	// If we don't do this the directory structure would
	// not be preserved
	// https://golang.org/src/archive/tar/common.go?#L626
	header.Name = filename

	// Write file header to the tar archive
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// Copy file content to tar archive
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
