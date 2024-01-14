package s3_uploader

import (
	"configs-saver/types"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"os"
	filepathTool "path/filepath"

	s3Config "github.com/aws/aws-sdk-go-v2/config"
	s3Credentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFile(filepath string, credentials types.AwsCredentials) (string, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: credentials.AwsEndpointUrl,
		}, nil
	})

	cfg, err := s3Config.LoadDefaultConfig(context.TODO(),
		s3Config.WithEndpointResolverWithOptions(customResolver),
		s3Config.WithCredentialsProvider(s3Credentials.NewStaticCredentialsProvider(
			credentials.AwsAccessKeyId,
			credentials.AwsSecretAccessKey, "")),
			s3Config.WithRegion("auto"),
	)

	if err != nil {
		log.Fatal(err)
	}

	if credentials.AwsBucketFolderPrefix == ""{
		log.Fatal("AwsBucketFolderPrefix is not valid")
	}

	objectKey := fmt.Sprintf("%v/%v", credentials.AwsBucketFolderPrefix, filepathTool.Base(filepath))

	s3Client := s3.NewFromConfig(cfg)
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", filepath, err)
	} else {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Printf("Couldn't close file %v. Here's why: %v\n", filepath, err)
			}
		}(file)

		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(credentials.AwsBucketName),
			Key: aws.String(objectKey),
			Body: file,
		})

		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n", filepath, credentials.AwsBucketName, objectKey, err)
		}
	}

	return objectKey, nil
}
