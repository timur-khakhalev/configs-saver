package main

import (
	"configs-saver/archivator"
	configVarsExtractor "configs-saver/config-extractor"
	uploader "configs-saver/s3-uploader"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	configSaverIniPath := flag.String("c", "", "Path to config *.ini")
	flag.Parse()

	if *configSaverIniPath == "" {
		log.Fatalln("Config file not specified")
	}

	configVars, err := configVarsExtractor.LoadConfigs(*configSaverIniPath)
	if err != nil {
		log.Fatal(err)
	}

	archivePath := archivator.ArchiveFiles(append(configVars.Dirs, configVars.Files...), configVars.OutputFilename)

	objectKey, err := uploader.UploadFile(archivePath, configVars.AwsCredentials)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(archivePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Archive successfully uploaded to s3. Object key:", objectKey)
}
