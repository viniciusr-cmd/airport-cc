package datatypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

// CustomDataTypes contain the user-defined primary data types
var CustomDataTypes = map[string]assets.DataType{
	"cnpj":         cnpj,
	"flightcode":   flightcode,
	"icao":         icao,
	"registration": registration,
}
