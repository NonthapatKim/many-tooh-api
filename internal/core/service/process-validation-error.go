package service

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

func ProcessValidationError(err error) domain.ValidationErrorResponse {
	var validationErrors domain.ValidationErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Field() {
		case "Email":
			if err.Tag() == "customEmail" {
				message := "รูปแบบอีเมลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
				validationErrors.Email = &message
			} else {
				message := "กรุณากรอกอีเมลของคุณ"
				validationErrors.Email = &message
			}
		case "Password":
			if err.Tag() == "customPassword" {
				message := "รหัสผ่านต้องมีอย่างน้อย 8 ตัว และมีตัวพิมพ์เล็ก, ตัวพิมพ์ใหญ่, ตัวเลข, และอักขระพิเศษ"
				validationErrors.Password = &message
			} else {
				message := "กรุณากรอกรหัสผ่านของคุณ"
				validationErrors.Password = &message
			}
		}
	}
	return validationErrors
}
