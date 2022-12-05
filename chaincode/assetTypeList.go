package main

import (
	"github.com/goledgerdev/airport-cc/chaincode/assettypes"
	"github.com/goledgerdev/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Airline,
	assettypes.Flight,
	assettypes.Airport,
	assettypes.Secret,
}
