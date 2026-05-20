package scenjsonparse

import (
	"testing"

	scenmodel "github.com/multiversx/mx-chain-scenario-go/scenario/model"
	"github.com/stretchr/testify/require"
)

func TestParseSetStateStepAuthorizedDRWASyncCallers(t *testing.T) {
	t.Parallel()

	snippet := `
	{
		"step": "setState",
		"authorizedDRWASyncCallers": [
			"''caller-one",
			"0x63616c6c65722d74776f"
		]
	}`

	p := Parser{}
	step, err := p.ParseScenarioStep(snippet)

	require.NoError(t, err)
	setStateStep, ok := step.(*scenmodel.SetStateStep)
	require.True(t, ok)
	require.Len(t, setStateStep.AuthorizedDRWASyncCallers, 2)
	require.Equal(t, []byte("caller-one"), setStateStep.AuthorizedDRWASyncCallers[0].Value)
	require.Equal(t, []byte("caller-two"), setStateStep.AuthorizedDRWASyncCallers[1].Value)
}
