package worldmock

import (
	"testing"

	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/stretchr/testify/require"
)

type drwaAuthHookStub struct {
	vmcommon.BlockchainHook
	authorized bool
}

func (stub *drwaAuthHookStub) IsAuthorizedDRWASyncCaller(_ []byte) bool {
	return stub.authorized
}

func TestAccountValidateReportsCodeWithoutSmartContractAddress(t *testing.T) {
	address := make([]byte, 32)
	address[0] = 1
	account := &Account{
		Address: address,
		Code:    []byte("contract"),
	}

	err := account.Validate()
	require.EqualError(t, err, "account has code but not a smart contract address: 0x0100000000000000000000000000000000000000000000000000000000000000")
}

func TestAccountValidateReportsSmartContractAddressWithoutCode(t *testing.T) {
	address := make([]byte, 32)
	address[8] = 1
	account := &Account{
		Address: address,
	}

	err := account.Validate()
	require.EqualError(t, err, "account has a smart contract address, but has no code: 0x0000000000000000010000000000000000000000000000000000000000000000")
}

func TestGetBuiltinFunctionNamesReturnsNilWhenWrapperNotInitialized(t *testing.T) {
	world := NewMockWorld()

	require.Nil(t, world.GetBuiltinFunctionNames())
}

func TestApplyDRWASyncEnvelopeBytesRequiresProvidedHook(t *testing.T) {
	world := NewMockWorld()

	err := world.ApplyDRWASyncEnvelopeBytes([]byte("payload"), []byte("caller"))
	require.ErrorIs(t, err, ErrProvidedBlockchainHookNotInitialized)
}

func TestIsAuthorizedDRWASyncCallerUsesSetStateWhitelist(t *testing.T) {
	world := NewMockWorld()
	caller := []byte("authorized-drwa-caller")
	world.AuthorizedDRWASyncCallers[string(caller)] = struct{}{}

	require.True(t, world.IsAuthorizedDRWASyncCaller(caller))
	require.False(t, world.IsAuthorizedDRWASyncCaller([]byte("other-caller")))
}

func TestIsAuthorizedDRWASyncCallerKeepsWhitelistWhenHookIsProvided(t *testing.T) {
	world := NewMockWorld()
	caller := []byte("authorized-drwa-caller")
	world.AuthorizedDRWASyncCallers[string(caller)] = struct{}{}
	world.ProvidedBlockchainHook = &drwaAuthHookStub{authorized: false}

	require.True(t, world.IsAuthorizedDRWASyncCaller(caller))
	require.False(t, world.IsAuthorizedDRWASyncCaller([]byte("other-caller")))
}

func TestIsAuthorizedDRWASyncCallerFallsBackToProvidedHook(t *testing.T) {
	world := NewMockWorld()
	world.ProvidedBlockchainHook = &drwaAuthHookStub{authorized: true}

	require.True(t, world.IsAuthorizedDRWASyncCaller([]byte("caller-from-hook")))
}

func TestClearResetsAuthorizedDRWASyncCallers(t *testing.T) {
	world := NewMockWorld()
	caller := []byte("authorized-drwa-caller")
	world.AuthorizedDRWASyncCallers[string(caller)] = struct{}{}

	world.Clear()

	require.False(t, world.IsAuthorizedDRWASyncCaller(caller))
	require.NotNil(t, world.AuthorizedDRWASyncCallers)
}
