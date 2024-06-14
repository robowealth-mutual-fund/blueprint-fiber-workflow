package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	errs "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils/fiber/errors"
)

// CustomValidator ...
type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

// Configure ...
func (cv *CustomValidator) Configure() {
	// Setup validator and initial language
	v := validator.New()

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, trans)

	// Assign the Validate and Trans
	cv.Validator = v
	cv.Trans = trans

	cv.RegisterRegexp()
	// Start register a custom rules here ...
	//cv.RegisterRule(validatorRule.NewMobileNumberRule())
}

// RegisterRule ...
// FIXME Please Enable when you want to register custom rules
//func (cv *CustomValidator) RegisterRule(rule validatorRule.Rule) {
//	validatorRule.RegisterValidationRule(rule, cv.Validator, cv.Trans)
//}

func (cv *CustomValidator) RegisterRegexp() {
	err := cv.Validator.RegisterValidation("regexp", Regexp)
	if err != nil {
		return
	}
}

func Regexp(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(fl.Param())
	return re.MatchString(fl.Field().String())
}

// Validate ...
func (cv *CustomValidator) Validate(structRule interface{}) *errs.GoError {
	if err := cv.Validator.Struct(structRule); err != nil {
		formError := errs.New("Wrong Input")
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)

		for _, e := range validationErrors {
			jsonFieldName := getJSONFieldName(structRule, e.Field())

			param := e.Param()
			if param != "" {
				param = "=" + param
			}

			errorMsg := fmt.Sprintf("ERROR_%s%s: %s", strings.ToUpper(e.Tag()), param, e.Translate(cv.Trans))

			formError.AddErrorField(jsonFieldName, errorMsg, e.Tag())
		}

		return formError
	}

	return nil
}

// getJSONFieldName extracts the JSON field name from the struct tag
func getJSONFieldName(structRule interface{}, fieldName string) string {
	rv := reflect.ValueOf(&structRule).Elem()
	if field, ok := rv.Elem().Type().FieldByName(fieldName); ok {
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			return strings.Split(jsonTag, ",")[0]
		}
	}
	return fieldName
}

// NewCustomValidator ...
func NewCustomValidator() *CustomValidator {
	var cv = &CustomValidator{}
	cv.Configure()
	return cv
}
