package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return the number of flights of a airport
// GET method
var GetNumberOfFlightsFromAirport = tx.Transaction{
	Tag:         "getNumberOfFlightsFromAirport",
	Label:       "Get Number Of Flights From Airport",
	Description: "Return the number of flights of a airport",
	Method:      "GET",
	Callers:     []string{"$org2MSP", "$orgMSP"}, // Only org2 can call this transactions

	Args: []tx.Argument{
		{
			Tag:         "airport",
			Label:       "Airport",
			Description: "Airport",
			DataType:    "->airport",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		airportKey, _ := req["airport"].(assets.Key)

		// Returns Airport from channel
		airportMap, err := airportKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		numberOfFlights := 0
		flights, ok := airportMap["flights"].([]interface{})
		if ok {
			numberOfFlights = len(flights)
		}

		returnMap := make(map[string]interface{})
		returnMap["numberOfFlights"] = numberOfFlights

		// Marshal asset back to JSON format
		returnJSON, nerr := json.Marshal(returnMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return returnJSON, nil
	},
}
