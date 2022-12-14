package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var registration = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		registration, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}
		//registration must be uppercase
		registration = strings.ToUpper(registration)

		//registration must contains 3 to 6 characters
		if len(registration) < 3 || len(registration) > 6 {
			return "", nil, errors.NewCCError("aircraft registration code must have 3 to 6 characters", 400)
		}

		// registration must not be empty
		if registration == "" {
			return "", nil, errors.NewCCError("aircraft registration code must not be empty", 400)
		}

		return registration, registration, nil
	},
}
