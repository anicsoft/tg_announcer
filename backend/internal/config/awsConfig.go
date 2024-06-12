package config

import (
	"os"
)

const (
	awsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	awsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	awsRegion          = "AWS_REGION"
)

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
	BaseURL         string
}

func NewAwsConfig() *AWSConfig {
	accessKeyID := os.Getenv(awsAccessKeyId)
	accessKeySecret := os.Getenv(awsSecretAccessKey)
	region := os.Getenv(awsRegion)
	uploadTimeout := 10
	return &AWSConfig{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		Region:          region,
		UploadTimeout:   uploadTimeout,
	}
}

/*func CreateSession(awsConfig AWSConfig) *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(awsConfig.Region),
			Credentials: credentials.NewStaticCredentials(
				awsConfig.AccessKeyID,
				awsConfig.AccessKeySecret,
				"",
			),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}*/
