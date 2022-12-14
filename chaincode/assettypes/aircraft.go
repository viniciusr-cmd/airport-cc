package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a aircraft
var Aircraft = assets.AssetType{
	Tag:         "aircraft",
	Label:       "aircraft",
	Description: "aircraft",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "aircraft registration",
			Label:    "Aircraft registration",
			DataType: "registration",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			// Composite Key
			Required: true,
			Tag:      "model",
			Label:    "Aircraft model",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			Required: true,
			Tag:      "airline",
			Label:    "Aircraft Airline",
			DataType: "->airline",
		},
	},
}
