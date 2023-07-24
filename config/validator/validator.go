package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

var (
	validate *validator.Validate
	lock     = &sync.Mutex{}
)

func GetValidator() *validator.Validate {
	lock.Lock()
	defer lock.Unlock()

	if validate != nil {
		return validate
	}

	validate = validator.New()
	return validate
}

func ValidateStruct(s interface{}) (errs map[string]string) {
	err := GetValidator().Struct(s)
	message := make(map[string]string)
	if err != nil {
		castedObject := err.(validator.ValidationErrors)
		for _, v := range castedObject {
			valueType := fmt.Sprintf("%v", reflect.TypeOf(v.Value()))
			switch v.Tag() {
			case "required":
				message[v.Field()] = fmt.Sprintf("%s is required", v.Field())
			case "email":
				message[v.Field()] = fmt.Sprintf("%s is not valid", v.Field())
			case "gte":
				message[v.Field()] = fmt.Sprintf("%s value must be greater than or equal to %s", v.Field(), v.Param())
			case "lte":
				message[v.Field()] = fmt.Sprintf("%s value must be lower than or equal to %s", v.Field(), v.Param())
			case "gt":
				message[v.Field()] = fmt.Sprintf("%s value must be greater than %s", v.Field(), v.Param())
			case "lt":
				if valueType == "string" {
					message[v.Field()] = fmt.Sprintf("%s length must be lower than %s", v.Field(), v.Param())
				} else {
					message[v.Field()] = fmt.Sprintf("%s value must be lower than %s", v.Field(), v.Param())
				}
			case "max":
				message[v.Field()] = fmt.Sprintf("%s value max %s", v.Field(), v.Param())
			case "min":
				message[v.Field()] = fmt.Sprintf("%s value min %s", v.Field(), v.Param())
			case "oneof":
				message[v.Field()] = fmt.Sprintf("%s value must be one of %s", v.Field(), v.Param())
			}
		}
		return message
	}
	return nil
}
