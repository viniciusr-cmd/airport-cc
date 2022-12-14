package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var status = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		status, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		status = strings.ReplaceAll(status, " ", "")
		status = strings.ToUpper(status)

		// status must not be empty
		if status == "" {
			return "", nil, errors.NewCCError("status Airline Code must not be empty", 400)
		}

		// status must have 3 digits
		if len(status) != 3 {
			return "", nil, errors.NewCCError("status Airline Code must have 3 digits", 400)
		}

		// status must not have any numbers
		if strings.ContainsAny(status, "0123456789") {
			return "", nil, errors.NewCCError("status Airline Code must not have any numbers", 400)
		}

		return status, status, nil
	},
}
