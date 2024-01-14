package main

import (
	"configs-saver/archivator"
	"configs-saver/config-extractor"
	"configs-saver/files"
	uploader "configs-saver/s3-uploader"
	"fmt"
	"log"
)

func main() {
	configSaverIniPath := "~/.configs-saver.ini"

	exists := files.EnsurePathExists(configSaverIniPath, false)
	if !exists {
		log.Fatalln("Config file configs-saver.ini not found:", configSaverIniPath)
	}

	configVars, err := config_extractor.LoadConfigs(configSaverIniPath)
	if err != nil {
		log.Fatal(err)
	}

	archivePath := archivator.ArchiveFiles(append(configVars.Dirs, configVars.Files...), configVars.OutputFilename)

	objectKey, err := uploader.UploadFile(archivePath, configVars.AwsCredentials)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Object key:", objectKey)
}
