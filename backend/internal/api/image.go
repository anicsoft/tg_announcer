package api

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// UploadImage godoc
//
//	@Summary		Upload an image
//	@Description	Uploads an image for an announcement to S3 and updates the entity's record with the S3 URL.
//	@Tags			announcements
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id				path		string				true	"announcements ID"
//	@Param			Authorization	header		string				true	"Authorization token"
//	@Param			image			formData	file				true	"Logo image file"
//	@Success		200				{object}	model.S3Response	"Successfully uploaded"
//	@Failure		400				{object}	HttpError			"Bad request"
//	@Failure		500				{object}	HttpError			"Internal server error"
//	@Router			/announcements/{id}/image [post]
func (a *BaseApi) UploadImage(ctx *gin.Context) {
	id := ctx.Param("id")
	header, err := ctx.FormFile("image")
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	s3URL, err := uploadToS3(header)
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	err = a.imageService.UploadLogo(ctx, id, s3URL)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, s3URL)
}

// UploadLogo godoc
//
//	@Summary		Upload a logo image
//	@Description	Uploads a logo image for a company to S3 and updates the entity's record with the S3 URL.
//	@Tags			companies
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id				path		string				true	"company ID"
//	@Param			Authorization	header		string				true	"Authorization token"
//	@Param			image			formData	file				true	"Logo image file"
//	@Success		200				{object}	model.S3Response	"Successfully uploaded"
//	@Failure		400				{object}	HttpError			"Bad request"
//	@Failure		401				{object}	HttpError			"Unauthorized"
//	@Failure		403				{object}	HttpError			"Forbidden"
//	@Failure		404				{object}	HttpError			"Entity not found"
//	@Failure		500				{object}	HttpError			"Internal server error"
//	@Router			/companies/{id}/logo [post]
func (a *BaseApi) UploadLogo(ctx *gin.Context) {
	id := ctx.Param("id")
	header, err := ctx.FormFile("image")
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	s3URL, err := uploadToS3(header)
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	err = a.imageService.UploadLogo(ctx, id, s3URL)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, s3URL)
}

func uploadToS3(header *multipart.FileHeader) (string, error) {
	file, err := header.Open()
	if err != nil {
		return "", err
	}

	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")), // change to your region
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return "", err
	}

	s3Client := s3.New(awsSession)

	fileKey := "uploads/" + header.Filename
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(os.Getenv("S3_BUCKET_NAME")), // change to your bucket name
		Key:                aws.String(fileKey),
		Body:               file,
		ContentType:        aws.String(header.Header.Get("Content-Type")),
		ContentDisposition: aws.String(fmt.Sprintf("%s; %s", "inline", "filename="+header.Filename)),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("S3_BUCKET_NAME"), fileKey), nil
}

/*func fetchFromS3(key string) ([]byte, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")), // change to your region
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(awsSession)

	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")), // change to your bucket name
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	// Read the content of the image
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)

	return buf.Bytes(), nil
}*/
