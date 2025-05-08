package uploadprovider

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"g05-food-delivery/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
)

type S3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secretKey  string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName, region, apiKey, domain, secretKey string) *S3Provider {
	// Validate required parameters
	if bucketName == "" {
		log.Println("Warning: S3 bucket name is empty")
	}

	if region == "" {
		log.Println("Warning: S3 region is empty")
	}

	provider := &S3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secretKey:  secretKey,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey,
			provider.secretKey,
			""),
	})

	if err != nil {
		log.Fatal(err)
	}

	provider.session = s3Session

	return provider
}

func (provider *S3Provider) SaveFileUploaded(ctx context.Context, dataBytes []byte, dst string) (*common.Image, error) {
	// Debug logs to help diagnose issues
	log.Printf("S3 Upload - Bucket: '%s', Region: '%s', Path: '%s'",
		provider.bucketName, provider.region, dst)

	// Validate bucket name before attempting upload
	if provider.bucketName == "" {
		return nil, errors.New("S3 bucket name cannot be empty")
	}

	fileBytes := bytes.NewReader(dataBytes)
	fileType := http.DetectContentType(dataBytes)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		Body:        fileBytes,
		ContentType: aws.String(fileType),
	})

	if err != nil {
		log.Printf("S3 upload error: %v", err)
		return nil, err
	}

	//req, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
	//	Bucket: aws.String(provider.bucketName),
	//	Key:    aws.String(dst),
	//	ACL:    aws.String("private"),
	//})
	//
	//req.Presign(15 * time.Minute) ->

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
