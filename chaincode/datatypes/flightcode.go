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

		// var vd0 int
		// for i, d := range flightcode {
		// 	if i >= 9 {
		// 		break
		// 	}
		// 	dnum := int(d) - '0'
		// 	vd0 += (10 - i) * dnum
		// }
		// vd0 = 11 - vd0%11
		// if vd0 > 9 {
		// 	vd0 = 0
		// }
		// if int(flightcode[9])-'0' != vd0 {
		// 	return "", nil, errors.NewCCError("Invalid Flight Code", 400)
		// }

		// var vd1 int
		// for i, d := range flightcode {
		// 	if i >= 10 {
		// 		break
		// 	}
		// 	dnum := int(d) - '0'
		// 	vd1 += (11 - i) * dnum
		// }
		// vd1 = 11 - vd1%11
		// if vd1 > 9 {
		// 	vd1 = 0
		// }
		// if int(flightcode[10])-'0' != vd1 {
		// 	return "", nil, errors.NewCCError("Invalid Flight Code", 400)
		// }

		return flightcode, flightcode, nil
	},
}
