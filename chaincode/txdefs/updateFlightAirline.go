package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Updates the airline of a Flight
// POST Method
var UpdateFlightAirline = tx.Transaction{
	Tag:         "updateFlightAirline",
	Label:       "Update Flight Airline",
	Description: "Change the airline of a flight",
	Method:      "PUT",
	Callers:     []string{`$org\dMSP`, "orgMSP"}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "flight",
			Label:       "Flight",
			Description: "Flight",
			DataType:    "->flight",
			Required:    true,
		},
		{
			Tag:         "airline",
			Label:       "airline",
			Description: "New airline of the flight",
			DataType:    "->airline",
		},
		{
			Tag:         "flightcode",
			Label:       "flightcode",
			Description: "New flightcode",
			DataType:    "flightcode",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		flightKey, ok := req["flight"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter flight must be an asset")
		}
		airlineKey, ok := req["airline"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter airline must be an asset")
		}
		flightcodeKey, ok := req["flightcode"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter flightcode must be a string")
		}
		
		// Returns Flight from channel
		flightAsset, err := flightKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		flightMap := (map[string]interface{})(*flightAsset)

		// Returns airline from channel
		airlineAsset, err := airlineKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		airlineMap := (map[string]interface{})(*airlineAsset)

		// Returns flightcode from channel
		flightcodeAsset, err := flightcodeKey.Get(stub)

		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		flightcodeMap := (map[string]interface{})(*flightcodeAsset)

		updatedAirlineKey := make(map[string]interface{})
		updatedAirlineKey["@assetType"] = "airline"
		updatedAirlineKey["@key"] = airlineMap["@key"]

		updatedflightcodeKey := make(map[string]interface{})
		updatedflightcodeKey["@assetType"] = "flightcode"
		updatedflightcodeKey["@key"] = flightcodeMap["@key"]

		// Update data
		flightMap["airline"] = updatedAirlineKey
		flightMap["flightcode"] = updatedflightcodeKey

		flightMap, err = flightAsset.Update(stub, flightMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		// Marshal asset back to JSON format
		flightJSON, nerr := json.Marshal(flightMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return flightJSON, nil
	},
}
