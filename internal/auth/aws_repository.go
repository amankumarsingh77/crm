package auth

import (
	"context"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSRepository interface {
	PutObject(ctx context.Context, input models.UploadInput) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, bucket, filename string) (*s3.GetObjectOutput, error)
	RemoveObject(ctx context.Context, bucket, filename string) error
}
