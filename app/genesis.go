package app

import (
	"encoding/json"
	"testing"

	"cosmossdk.io/log"

	dbm "github.com/cosmos/cosmos-db"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// GenesisState of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
// Deprecated: use wasmApp.DefaultGenesis() instead
func NewDefaultGenesisState(t *testing.T) GenesisState {
	// we "pre"-instantiate the application for getting the injected/configured encoding configuration
	// note, this is not necessary when using app wiring, as depinject can be directly used (see root_v2.go)
	tempApp := NewWasmApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, wasmtypes.EnableAllProposals, simtestutil.NewAppOptionsWithFlagHome(t.TempDir()), []wasm.Option{})
	return tempApp.DefaultGenesis()
}
