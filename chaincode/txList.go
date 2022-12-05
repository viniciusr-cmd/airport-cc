package main

import (
	txdefs "github.com/goledgerdev/airport-cc/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.CreateNewAirport,
	txdefs.GetNumberOfFlightsFromAirport,
	txdefs.UpdateFlightAirline,
	txdefs.GetFlightsByAirline,
}
