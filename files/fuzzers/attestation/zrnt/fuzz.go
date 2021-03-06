package fuzz

import (
	"github.com/protolambda/zrnt/eth2/phase0"
	"helper"
)

func init() {
	helper.SetInputType(helper.INPUT_TYPE_ATTESTATION)
}

func Fuzz(data []byte) []byte {
	input, err := helper.DecodeAttestation(data, false)
	if err != nil {
		// Assumes preprocessing ensures data is decodable
		panic("Decoding failed - bug in preprocessing.")
	}

	ffstate := phase0.NewFullFeaturedState(&input.Pre)
	ffstate.LoadPrecomputedData()

    // TODO(gnattishness) disable validation and sig verification (once supported)
	if err := ffstate.ProcessAttestation(&input.Attestation); err != nil {
		return []byte{}
	}

	return helper.EncodePoststate(&input.Pre)
}
