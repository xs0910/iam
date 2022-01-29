package validation

import (
	"fmt"
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	_ "github.com/go-playground/validator/v10/translations/en"
	"github.com/xs0910/iam/pkg/component-base/validation/field"
	"os"
	"reflect"
)

const (
	maxDescriptionLength = 255
)

type Validator struct {
	val   *validator.Validate
	data  interface{}
	trans ut.Translator
}

// NewValidator create a new validator
func NewValidator(data interface{}) *Validator {
	validate := validator.New()

	// independent validators
	validate.RegisterValidation("dir", validateDir)
	validate.RegisterValidation("file", validateFile)
	validate.RegisterValidation("description", validateDescription)
	validate.RegisterValidation("name", validateName)

	// default translations
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	err := en.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	// additional translations
	translations := []struct {
		tag         string
		translation string
	}{
		{tag: "dir", translation: "{0} must point to an existing directory, but found '{1}'"},
		{tag: "file", translation: "{0} must point to an existing file, but found '{1}'"},
		{tag: "description", translation: fmt.Sprintf("must be less than %d", maxDescriptionLength)},
		{tag: "name", translation: "is not a invalid name"},
	}

	for _, t := range translations {
		err = validate.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation), translateFunc)
		if err != nil {
			panic(err)
		}
	}

	return &Validator{
		val:   validate,
		data:  data,
		trans: trans,
	}
}

// Validate validates config for errors and returns an error (it can be cast to ValidationErrors, containing a list of errors inside).
// When error is printed as string, it will automatically contain the full list of validation errors.
func (v *Validator) Validate() field.ErrorList {
	// validate policy
	err := v.val.Struct(v.data)
	if err == nil {
		return nil
	}

	// this check is only needed when your code could produce an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return field.ErrorList{field.Invalid(field.NewPath(""), err.Error(), "")}
	}

	allErrs := field.ErrorList{}

	// collect human-readable errors
	vErrors, _ := err.(validator.ValidationErrors)
	for _, vErr := range vErrors {
		allErrs = append(allErrs, field.Invalid(field.NewPath(vErr.Namespace()), vErr.Translate(v.trans), ""))
	}
	return allErrs
}

func registrationFunc(tag string, translation string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, true); err != nil {
			return
		}
		return
	}
}
func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field(), reflect.ValueOf(fe.Value()).String())
	if err != nil {
		return fe.(error).Error()
	}
	return t
}

func validateDir(fl validator.FieldLevel) bool {
	path := fl.Field().String()
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
func validateFile(fl validator.FieldLevel) bool {
	path := fl.Field().String()
	if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
		return true
	}
	return false
}

// validateDescription checks if a given description is illegal.
func validateDescription(fl validator.FieldLevel) bool {
	description := fl.Field().String()

	return len(description) <= maxDescriptionLength
}

// validateName checks if a given name is illegal.
func validateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	if errs := IsQualifiedName(name); len(errs) > 0 {
		return false
	}

	return true
}
