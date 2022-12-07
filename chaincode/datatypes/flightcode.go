package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var flightcode = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		flightcode, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		flightcode = strings.ReplaceAll(flightcode, ".", "")
		flightcode = strings.ReplaceAll(flightcode, "-", "")
		flightcode = strings.ReplaceAll(flightcode, "/", "")
		flightcode = strings.ReplaceAll(flightcode, " ", "")

		if len(flightcode) != 6 {
			return "", nil, errors.NewCCError("Flight Code must have 6 digits", 400)
		}

		return flightcode, flightcode, nil
	},
}
