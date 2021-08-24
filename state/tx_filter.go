package state

import (
	mempl "github.com/tendermint/tendermint/internal/mempool"
	"github.com/tendermint/tendermint/pkg/block"
)

// TxPreCheck returns a function to filter transactions before processing.
// The function limits the size of a transaction to the block's maximum data size.
func TxPreCheck(state State) mempl.PreCheckFunc {
	maxDataBytes := block.MaxDataBytesNoEvidence(
		state.ConsensusParams.Block.MaxBytes,
		state.Validators.Size(),
	)
	return mempl.PreCheckMaxBytes(maxDataBytes)
}

// TxPostCheck returns a function to filter transactions after processing.
// The function limits the gas wanted by a transaction to the block's maximum total gas.
func TxPostCheck(state State) mempl.PostCheckFunc {
	return mempl.PostCheckMaxGas(state.ConsensusParams.Block.MaxGas)
}
