package helpers

import "github.com/asaskevich/govalidator"

func ValidateStruct(payload interface{}) error {
	_, err := govalidator.ValidateStruct(payload)

	return err
}
