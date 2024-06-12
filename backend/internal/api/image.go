package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi/v5"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func (a *BaseApi) UploadImage(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*	file, header, err := r.FormFile("image")
			if err != nil {
				a.Error(w, http.StatusBadRequest, err)
				return
			}
			defer file.Close()

			keys := r.MultipartForm.Value["path"][0]
			if len(keys) == 0 {
				a.Error(w, http.StatusBadRequest, fmt.Errorf("path variables are not provided"))
				return
			}

			contentType := header.Header.Get("Content-Type")
			if contentType != "image/jpeg" && contentType != "image/png" {
				a.Error(w, http.StatusBadRequest, fmt.Errorf("only JPG or PNG files are allowed"))
				return
			}

			fileBytes, err := io.ReadAll(file)
			if err != nil {
				a.Error(w, http.StatusBadRequest, err)
				return
			}

			uploader := s3manager.NewUploader(sess)
			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(awsConf.BucketName),
				Key:    aws.String("awsConf."),
				Body:   file,
			})

			if err != nil {
				a.Error(w, http.StatusBadRequest, err)
				return
			}

			a.Respond(w, http.StatusOK, nil)*/
		announcementId := chi.URLParam(r, "parentId")
		file, header, err := r.FormFile("image")
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}
		defer file.Close()

		s3URL, err := uploadToS3(file, header)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}
		idInt, _ := strconv.Atoi(announcementId)
		err = a.imageService.Upload(ctx, idInt, []string{s3URL})
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, Response{Data: s3URL})
	}
}

func (a *BaseApi) FetchImage(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parentId := chi.URLParam(r, "parentId")
		if parentId == "" {
			a.Error(w, http.StatusBadRequest, fmt.Errorf("empty parentId"))
			return
		}
		intId, _ := strconv.Atoi(parentId)

		key, err := a.imageService.Get(ctx, intId)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}
		log.Println("KEy", key)
		image, err := fetchFromS3(key[0])
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(image)
	}

}

func uploadToS3(file multipart.File, header *multipart.FileHeader) (string, error) {
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

func fetchFromS3(key string) ([]byte, error) {
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
		Key:    aws.String("uploads/8_1sasa11.jpg"),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	// Read the content of the image
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)

	return buf.Bytes(), nil
}
