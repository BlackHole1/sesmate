package validator

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	allowedCharsRegex = "^[a-zA-Z0-9_]+$"
)

var allowedCharsCompile = regexp.MustCompile(allowedCharsRegex)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("allowedChars", allowedChars); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("validator is not *validator.Validate")
	}
}

func allowedChars(fl validator.FieldLevel) bool {
	return allowedCharsCompile.MatchString(fl.Field().String())
}
