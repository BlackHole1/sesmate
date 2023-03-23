package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSesV2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"github.com/BlackHole1/sesmate/internal/sync/template"
)

func CreateTemplate(body *template.SchemaBody) error {
	_, err := client.CreateEmailTemplate(context.Background(), &awsSesV2.CreateEmailTemplateInput{
		TemplateName: aws.String(body.TemplateName),
		TemplateContent: &types.EmailTemplateContent{
			Html:    body.HtmlPart,
			Text:    body.TextPart,
			Subject: body.SubjectPart,
		},
	})

	return err
}
