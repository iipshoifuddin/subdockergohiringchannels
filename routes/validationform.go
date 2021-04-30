package routes

import (
	"regexp"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var bookableGeneral validator.Func = func(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 0

}

func validateEmail(fl validator.FieldLevel) bool {
	nationalId := fl.Field().String()

	regex, _ := regexp.Compile("^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$")
	result := regex.MatchString(nationalId)
	// fmt.Println("from email Validatiin :", nationalId)
	return result
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func bookablePassword(fl validator.FieldLevel) bool {
	pass := fl.Field().String()
	letters := 0
	number := false
	upper := false
	lower := false
	special := false
	valBack := true
	for _, c := range pass {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsLower(c):
			lower = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}

	if !number || !upper || !lower || !special || letters <= 7 {
		valBack = false
	}
	return valBack
}

func bookableUsername(fl validator.FieldLevel) bool {
	pass := fl.Field().String()
	letters := 0
	upper := false
	lower := false
	notSpace := true
	valBack := true
	for _, c := range pass {
		switch {
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsLower(c):
			lower = true
			letters++
		case unicode.IsLetter(c):
			letters++
		case c == ' ':
			notSpace = false
		default:
			//return false, false, false, false
		}
	}
	// fmt.Println("from username :", notSpace)
	if !upper || !lower || !notSpace || letters <= 7 {
		valBack = false
	}
	return valBack
}
