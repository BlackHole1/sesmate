package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BlackHole1/sesmate/internal/server/svc"
)

func APIClearEmail(c *gin.Context) {
	err := svc.Email.DeleteAllRecord()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, "ok")
}
