package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const kRawBody = "sesmate-raw-body"

func RawBody() func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set(kRawBody, string(data))
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		c.Next()
	}
}

func MustRawBody(c *gin.Context) string {
	return c.MustGet(kRawBody).(string)
}
