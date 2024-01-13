package types

type AwsCredentials struct {
	AwsAccessKeyId string
	AwsSecretAccessKey string
	AwsDefaultRegion string
	AwsEndpointUrl string
	AwsBucketName  string
	AwsBucketFolderPrefix string
}

type ConfigVars struct {
	Dirs []string
	Files []string
	OutputFilename string
	AwsCredentials AwsCredentials
}
