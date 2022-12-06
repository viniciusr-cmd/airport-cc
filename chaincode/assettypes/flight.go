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
			IsKey:    true,
			Tag:      "airline",
			Label:    "Flight Airline",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			/// Reference to another asset
			Tag:      "currentAirline",
			Label:    "Current Airline Flight",
			DataType: "->airline",
		},
		{
			// Date property
			Tag:      "flight date",
			Label:    "Flight Date",
			DataType: "datetime",
		},
		// {
		// 	// Custom data type
		// 	Tag:      "flightType",
		// 	Label:    "Flight Type",
		// 	DataType: "flightType",
		// },
	},
}
