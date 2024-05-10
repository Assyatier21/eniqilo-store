package pkg

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func SetupValidator() *validator.Validate {
	v := validator.New()

	v.RegisterValidation("validateCategory", customProductCategoryEnum)
	v.RegisterValidation("validatePhoneNumber", validatePhoneNumber)
	v.RegisterValidation("validateImageURL", validateImageURL)

	return v
}

func BindValidate(c echo.Context, req interface{}) (err error) {
	if err = c.Bind(req); err != nil {
		err = fmt.Errorf("[Utils][Pkg][Validator] failed to bind request, err: %s", err.Error())
		return
	}

	if err = c.Validate(req); err != nil {
		err = fmt.Errorf("[Utils][Pkg][Validator] failed to validate request, err: %s", err.Error())
		return
	}

	return
}

func customProductCategoryEnum(fl validator.FieldLevel) bool {
	allowedValues := []string{
		"Clothing",
		"Accessories",
		"Footwear",
		"Beverages",
	}

	value := fl.Field().String()
	for _, v := range allowedValues {
		if strings.EqualFold(value, v) {
			return true
		}
	}
	return false
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	// Regular expression to match international calling codes
	regex := `^\+\d{1,3}[-\d]*$`
	match, _ := regexp.MatchString(regex, value)
	return match
}

func validateImageURL(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^https?://(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(?:/[^/?#]+)+\.(?:jpg|jpeg|png|gif|bmp)(?:\?[^\s]*)?$`

	re := regexp.MustCompile(pattern)

	return re.MatchString(value)
}
