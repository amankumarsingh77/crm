package repository

import (
	"context"
	"fmt"
	"github.com/amankumarsingh77/cmr/internal/auth"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"regexp"
)

type awsRepository struct {
	client *s3.Client
}

func NewAwsRepository(awsClient *s3.Client) auth.AWSRepository {
	return &awsRepository{
		client: awsClient,
	}
}

func (a *awsRepository) PutObject(ctx context.Context, input models.UploadInput) (*s3.PutObjectOutput, error) {
	pattern := `^.+\.(doc|docx|pdf|txt|xlsx|xls|ppt|pptx|odt|ods|odp)$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(input.Name) {
		return nil, fmt.Errorf("invalid file format: %s", input.Name)
	}
	res, err := a.client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:        &input.BucketName,
			Key:           &input.Name,
			ContentType:   &input.ContentType,
			ContentLength: &input.Size,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file : %w", err)
	}
	return res, nil
}

func (a *awsRepository) GetObject(ctx context.Context, bucket, filename string) (*s3.GetObjectOutput, error) {
	res, err := a.client.GetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &filename,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to download file : %w", err)
	}
	return res, nil
}

func (a *awsRepository) RemoveObject(ctx context.Context, bucket, filename string) error {
	_, err := a.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &filename,
	})
	if err != nil {
		return fmt.Errorf("failed to remove file : %w", err)
	}
	return nil
}
