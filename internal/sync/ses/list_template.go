package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSesV2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func ListTemplateName() ([]string, error) {
	var result []string

	var nextToken *string = nil
	for {
		templates, err := client.ListEmailTemplates(context.Background(), &awsSesV2.ListEmailTemplatesInput{
			NextToken: nextToken,
			PageSize:  aws.Int32(50),
		})
		if err != nil {
			return result, err
		}

		for _, template := range templates.TemplatesMetadata {
			result = append(result, *template.TemplateName)
		}
		if templates.NextToken == nil {
			break
		}
		nextToken = templates.NextToken
	}

	return result, nil
}
