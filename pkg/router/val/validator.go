package val

import (
	"time"

	"assemble/pkg/datajson"

	"github.com/go-playground/validator/v10"
)

func Date(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}
	_, err := time.Parse(datajson.FormatDate, s)
	return err == nil
}
