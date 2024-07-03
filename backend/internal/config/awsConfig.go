package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	awsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	awsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	awsRegion          = "AWS_REGION"
	awsBucket          = "S3_BUCKET_NAME"
)

type AWSConfig interface {
	CreateSession() *session.Session
	CreateS3Session(sess *session.Session) *s3.S3
	Bucket() string
}

type awsConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
}

func NewAwsConfig() AWSConfig {
	accessKeyID := os.Getenv(awsAccessKeyId)
	accessKeySecret := os.Getenv(awsSecretAccessKey)
	region := os.Getenv(awsRegion)
	bucket := os.Getenv(awsBucket)
	uploadTimeout := 10

	return &awsConfig{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		Region:          region,
		BucketName:      bucket,
		UploadTimeout:   uploadTimeout,
	}
}

func (c *awsConfig) CreateSession() *session.Session {
	return session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(c.Region),
			Credentials: credentials.NewStaticCredentials(
				c.AccessKeyID,
				c.AccessKeySecret,
				"",
			),
		},
	))
}

func (c *awsConfig) CreateS3Session(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func (c *awsConfig) Bucket() string {
	return c.BucketName
}
