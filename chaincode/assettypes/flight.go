package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a flight
var Flight = assets.AssetType{
	Tag:         "flight",
	Label:       "Flight",
	Description: "Flight",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "code",
			Label:    "Flight code",
			DataType: "flightcode",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			// Composite Key
			Required: true,
			Tag:      "airline",
			Label:    "Flight Airline",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			Required: true,
			Tag:      "origin",
			Label:    "Flight Origin",
			DataType: "->airport",
		},
		{
			Required: true,
			Tag:      "destination",
			Label:    "Flight Destination",
			DataType: "->airport",
		},
		{
			// Date property
			Required: true,
			Tag:      "Estimated flight departure",
			Label:    "Estimated flight Departure",
			DataType: "datetime",
		},
		{
			// Status property
			Required:     true,
			Tag:          "Flight Status",
			Label:        "Confirmed?",
			DefaultValue: true,
			DataType:     "boolean",
		},
	},
}
