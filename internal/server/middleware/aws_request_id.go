package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const kAWSRequestId = "sesmate-aws-request-id"

func AWSRequestId() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set(kAWSRequestId, uuid.NewString())
		c.Header("X-Amzn-Requestid", MustAWSRequestId(c))
		c.Next()
	}
}

func MustAWSRequestId(c *gin.Context) string {
	return c.MustGet(kAWSRequestId).(string)
}
