package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awsSesV2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
)

var client *awsSesV2.Client

func Setup(ak, sk, endpoint, region string) error {
	r := config.WithRegion(region)

	var customResolver aws.EndpointResolverWithOptions
	var customCredentials aws.CredentialsProvider

	if endpoint != "" {
		customResolver = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: endpoint,
			}, nil
		})
	}

	if ak != "" && sk != "" {
		customCredentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     ak,
				SecretAccessKey: sk,
			}, nil
		})
	}

	c, err := config.LoadDefaultConfig(context.Background(), r, config.WithEndpointResolverWithOptions(customResolver), config.WithCredentialsProvider(customCredentials))
	if err != nil {
		return err
	}

	client = awsSesV2.NewFromConfig(c)

	return nil
}

func Client() *awsSesV2.Client {
	return client
}
