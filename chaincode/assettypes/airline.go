package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Airline = assets.AssetType{
	Tag:         "airline",
	Label:       "Airline",
	Description: "Airline data",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "CNPJ (Brazil National Registry of Legal Entities ID)",
			DataType: "cnpj",                        // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the airline",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "ICAO",
			Label:    "ICAO Code (International Civil Aviation Organization)",
			DataType: "string",
			// Validate funcion
			Validate: func(icao interface{}) error {
				icaoStr := icao.(string)
				if icaoStr == "" {
					return fmt.Errorf("ICAO must be non-empty")
				}
				if len(icaoStr) != 3 {
					return fmt.Errorf("ICAO must be 3 characters long")
				}
				return nil
			},
		},
	},
}
