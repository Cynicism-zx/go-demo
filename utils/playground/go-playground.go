package main

import (
	"fmt"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
测试 validator库对结构体和map等的校验功能
https://www.liwenzhou.com/posts/Go/validator_usages/
*/

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var (
	validate *validator.Validate
)

func main() {
	validate = validator.New()
	validateStruct()
	validateVariable()
}

func validateStruct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}
	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}
	z := zh.New()
	uni := ut.New(z, z)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, ok := uni.GetTranslator("zh")
	if !ok {
		fmt.Printf("uni.GetTranslator(%s) failed \n", "zh")
		return
	}
	//验证器注册翻译器
	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		fmt.Printf("注册中文翻译失败 err:%s \n", err.Error())
		return
	}
	err := validate.Struct(user)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		for _, ee := range err.(validator.ValidationErrors) {
			fmt.Println("-----", ee.Translate(trans))
			fmt.Println(ee.Namespace())
			fmt.Println(ee.Field())
			fmt.Println(ee.StructNamespace())
			fmt.Println(ee.StructField())
			fmt.Println(ee.Tag())
			fmt.Println(ee.ActualTag())
			fmt.Println(ee.Kind())
			fmt.Println(ee.Type())
			fmt.Println(ee.Value())
			fmt.Println(ee.Param())
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
