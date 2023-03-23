package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSesV2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func RemoveTemplate(name string) error {
	_, err := client.DeleteEmailTemplate(context.Background(), &awsSesV2.DeleteEmailTemplateInput{
		TemplateName: aws.String(name),
	})

	return err
}
