package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var icao = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		icao, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		icao = strings.ReplaceAll(icao, " ", "")
		icao = strings.ToUpper(icao)

		// ICAO must not be empty
		if icao == "" {
			return "", nil, errors.NewCCError("ICAO Airline Code must not be empty", 400)
		}

		// ICAO must have 3 digits
		if len(icao) != 3 {
			return "", nil, errors.NewCCError("ICAO Airline Code must have 3 digits", 400)
		}

		// ICAO must not have any numbers
		if strings.ContainsAny(icao, "0123456789") {
			return "", nil, errors.NewCCError("ICAO Airline Code must not have any numbers", 400)
		}

		return icao, icao, nil
	},
}
