package config_extractor

import (
	"configs-saver/files"
	"configs-saver/types"
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	Dirs  []string
	Files []string
}

func LoadConfigs(filePath string) (configVars types.ConfigVars, err error) {
	// Load and parse the INI file
	cfg, err := ini.Load(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Config struct to store the parsed values
	config := Config{}

	// Parse the 'dirs' section
	if pathsSection, err := cfg.GetSection("paths"); err == nil {
		dirsValue := pathsSection.Key("dirs").String()
		config.Dirs = parsePaths(dirsValue)
		filesValue := pathsSection.Key("files").String()
		config.Files = parsePaths(filesValue)
	}

	configVars = types.ConfigVars{
		Dirs: config.Dirs,
		Files: config.Files,
		OutputFilename: cfg.Section("configs").Key("output_filename").String(),
		AwsCredentials: types.AwsCredentials{
			AwsAccessKeyId: cfg.Section("credentials").Key("aws_access_key_id").String(),
			AwsSecretAccessKey: cfg.Section("credentials").Key("aws_secret_access_key").String(),
			AwsDefaultRegion: cfg.Section("credentials").Key("aws_default_region").String(),
			AwsEndpointUrl: cfg.Section("credentials").Key("aws_endpoint_url").String(),
			AwsBucketName: cfg.Section("credentials").Key("aws_bucket_name").String(),
			AwsBucketFolderPrefix: cfg.Section("credentials").Key("aws_bucket_folder_prefix").String(),
		},
	}

	return configVars, nil
}

func parsePaths(s string) []string {
	values := strings.Split(s, ",")
	for i := range values {
		values[i] = files.ParsePath(strings.TrimSpace(values[i]))
	}

	return values
}

