package state

import (
	"github.com/tendermint/tendermint/pkg/abci"
	"github.com/tendermint/tendermint/pkg/consensus"
	"github.com/tendermint/tendermint/pkg/metadata"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"
)

//
// TODO: Remove dependence on all entities exported from this file.
//
// Every entity exported here is dependent on a private entity from the `state`
// package. Currently, these functions are only made available to tests in the
// `state_test` package, but we should not be relying on them for our testing.
// Instead, we should be exclusively relying on exported entities for our
// testing, and should be refactoring exported entities to make them more
// easily testable from outside of the package.
//

// UpdateState is an alias for updateState exported from execution.go,
// exclusively and explicitly for testing.
func UpdateState(
	state State,
	blockID metadata.BlockID,
	header *metadata.Header,
	abciResponses *tmstate.ABCIResponses,
	validatorUpdates []*consensus.Validator,
) (State, error) {
	return updateState(state, blockID, header, abciResponses, validatorUpdates)
}

// ValidateValidatorUpdates is an alias for validateValidatorUpdates exported
// from execution.go, exclusively and explicitly for testing.
func ValidateValidatorUpdates(abciUpdates []abci.ValidatorUpdate, params consensus.ValidatorParams) error {
	return validateValidatorUpdates(abciUpdates, params)
}
