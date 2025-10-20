package validation

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var IDTrans ut.Translator
var ENTrans ut.Translator
var Validator *validator.Validate

func init() {
	Setup()
}

func Setup() {
	Validator = validator.New()

	uni := ut.New(id.New(), en.New())
	IDTrans, _ = uni.GetTranslator("id")
	ENTrans, _ = uni.GetTranslator("en")

	id_translations.RegisterDefaultTranslations(Validator, IDTrans)
	en_translations.RegisterDefaultTranslations(Validator, ENTrans)

	Validator.RegisterValidation("enum", Enum)
	AddTranslationID("enum", ErrEnumStrID, Validator)
	AddTranslationEN("enum", ErrEnumStrEN, Validator)

	Validator.RegisterValidation("date", Date)
	AddTranslationID("date", ErrDateStrID, Validator)
	AddTranslationEN("date", ErrDateStrEN, Validator)
}

func AddTranslationID(tag string, errMessage string, validate *validator.Validate) {
	registerID := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	trans := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = validate.RegisterTranslation(tag, IDTrans, registerID, trans)
}

func AddTranslationEN(tag string, errMessage string, validate *validator.Validate) {
	registerEN := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	trans := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = validate.RegisterTranslation(tag, ENTrans, registerEN, trans)
}
