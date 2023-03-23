package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSesV2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"github.com/BlackHole1/sesmate/pkg/template"
)

func UpdateTemplate(body *template.SchemaBody) error {
	_, err := client.UpdateEmailTemplate(context.Background(), &awsSesV2.UpdateEmailTemplateInput{
		TemplateName: aws.String(body.TemplateName),
		TemplateContent: &types.EmailTemplateContent{
			Html:    body.HtmlPart,
			Text:    body.TextPart,
			Subject: body.SubjectPart,
		},
	})

	return err
}
