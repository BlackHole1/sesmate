package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/BlackHole1/sesmate/internal/server/middleware"
	"github.com/BlackHole1/sesmate/internal/server/model"
	"github.com/BlackHole1/sesmate/internal/server/schema"
	"github.com/BlackHole1/sesmate/internal/server/svc"
)

type awsV2SendEmailRequest struct {
	ConfigurationSetName                      string                           `binding:"omitempty"`
	Content                                   *schema.AWSEmailContent          `binding:"required"`
	Destination                               *schema.AWSDestination           `binding:"omitempty"`
	AWSMessageTag                             []*schema.AWSMessageTag          `binding:"omitempty"`
	FeedbackForwardingEmailAddress            string                           `binding:"omitempty"`
	FeedbackForwardingEmailAddressIdentityArn string                           `binding:"omitempty"`
	FromEmailAddress                          string                           `binding:"omitempty"`
	FromEmailAddressIdentityArn               string                           `binding:"omitempty"`
	ListManagementOptions                     *schema.AWSListManagementOptions `binding:"omitempty"`
	ReplyToAddresses                          []string                         `binding:"omitempty"`
}

func AWSV2SendEmail(c *gin.Context) {
	var req awsV2SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Header("X-Amzn-Errortype", "BadRequestException")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	templateName := req.Content.Template.TemplateName
	if templateName != "" {
		if !middleware.HasTemplate(c, templateName) {
			c.Header("X-Amzn-Errortype", "NotFoundException")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Template " + templateName + " does not exist",
			})
			return
		}
	}

	messageId := uuid.NewString()

	data, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err = svc.Email.Record(&model.EmailRecord{
		Data:      string(data),
		RawData:   middleware.MustRawBody(c),
		MessageId: messageId,
		RequestId: middleware.MustAWSRequestId(c),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"MessageId": messageId,
	})
}
