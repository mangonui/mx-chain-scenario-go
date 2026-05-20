package scenexec

import (
	"testing"

	scenmodel "github.com/multiversx/mx-chain-scenario-go/scenario/model"
	worldmock "github.com/multiversx/mx-chain-scenario-go/worldmock"
	"github.com/stretchr/testify/require"
)

func TestExecuteSetStateStepAddsAuthorizedDRWASyncCallers(t *testing.T) {
	t.Parallel()

	caller := []byte("authorized-drwa-caller")
	executor := &ScenarioExecutor{World: worldmock.NewMockWorld()}

	err := executor.ExecuteSetStateStep(&scenmodel.SetStateStep{
		AuthorizedDRWASyncCallers: []scenmodel.JSONBytesFromString{
			scenmodel.NewJSONBytesFromString(caller, "''authorized-drwa-caller"),
		},
	})

	require.NoError(t, err)
	require.True(t, executor.World.IsAuthorizedDRWASyncCaller(caller))
	require.False(t, executor.World.IsAuthorizedDRWASyncCaller([]byte("other-caller")))
}

func TestExecuteSetStateStepInitializesAuthorizedDRWASyncCallersMap(t *testing.T) {
	t.Parallel()

	caller := []byte("authorized-drwa-caller")
	executor := &ScenarioExecutor{World: &worldmock.MockWorld{}}

	err := executor.ExecuteSetStateStep(&scenmodel.SetStateStep{
		AuthorizedDRWASyncCallers: []scenmodel.JSONBytesFromString{
			scenmodel.NewJSONBytesFromString(caller, "''authorized-drwa-caller"),
		},
	})

	require.NoError(t, err)
	require.True(t, executor.World.IsAuthorizedDRWASyncCaller(caller))
}
