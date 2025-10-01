package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// ValidatorRule type untuk custom rules
type ValidatorRule struct {
	Tag     string
	Pattern string
	Message string
}

func InitValidator() {
	Validate = validator.New()

	// List custom rules → tinggal tambah di sini
	customRules := []ValidatorRule{
		{"alphaonly", "^[a-zA-Z]+$", "must contain only letters"},
		{"alphanumonly", "^[a-zA-Z0-9]+$", "must contain only letters and numbers"},
		{"numericonly", "^[0-9]+$", "must contain only numbers"},
		{"email", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, "must be a valid email address"},
	}

	// Register semua custom rules
	for _, rule := range customRules {
		re := regexp.MustCompile(rule.Pattern)

		Validate.RegisterValidation(rule.Tag, func(fl validator.FieldLevel) bool {
			return re.MatchString(fl.Field().String())
		})
	}
}
