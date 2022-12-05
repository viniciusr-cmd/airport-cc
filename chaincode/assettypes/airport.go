package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a Library as a collection of books
var Airport = assets.AssetType{
	Tag:         "airport",
	Label:       "Airport",
	Description: "Airport as a collection of Flights",

	Props: []assets.AssetProp{
		{
			// Primary Key
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Airport Name",
			DataType: "string",
			Writers:  []string{`org3MSP`, "orgMSP"}, // This means only org3 can create the asset (others can edit)
		},
		{
			// Asset reference list
			Tag:      "flights",
			Label:    "Flight Collection",
			DataType: "[]->flight",
		},
	},
}
