package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var cnpj = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		cnpj, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		cnpj = strings.ReplaceAll(cnpj, ".", "")
		cnpj = strings.ReplaceAll(cnpj, "-", "")
		cnpj = strings.ReplaceAll(cnpj, "/", "")
		cnpj = strings.ReplaceAll(cnpj, " ", "")

		if len(cnpj) != 14 {
			return "", nil, errors.NewCCError("CNPJ must have 14 digits", 400)
		}

		// var vd0 int
		// for i, d := range cnpj {
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
		// if int(cnpj[9])-'0' != vd0 {
		// 	return "", nil, errors.NewCCError("Invalid CNPJ", 400)
		// }

		// var vd1 int
		// for i, d := range cnpj {
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
		// if int(cnpj[10])-'0' != vd1 {
		// 	return "", nil, errors.NewCCError("Invalid CNPJ", 400)
		// }

		return cnpj, cnpj, nil
	},
}
