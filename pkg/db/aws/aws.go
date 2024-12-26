package aws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewAwsClient(endpoint, region, accesskey, secretkey string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
		config.WithBaseEndpoint(endpoint), // can be used with any s3 compatible storage providers
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accesskey,
				secretkey,
				""),
		),
	)
	if err != nil {
		return nil, errors.New("failed to load configuration, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
