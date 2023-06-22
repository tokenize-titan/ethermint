package cosmos_client_ext

import (
	"bufio"
	"io"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Context struct {
	cosmosclient.Context

	RPCClient rpcclient.Client
}

// WithKeyring returns a copy of the context with an updated keyring.
func (ctx Context) WithKeyring(k keyring.Keyring) Context {
	ctx.Keyring = k
	return ctx
}

// WithKeyringOptions returns a copy of the context with an updated keyring.
func (ctx Context) WithKeyringOptions(opts ...keyring.Option) Context {
	ctx.KeyringOptions = opts
	return ctx
}

// WithInput returns a copy of the context with an updated input.
func (ctx Context) WithInput(r io.Reader) Context {
	// convert to a bufio.Reader to have a shared buffer between the keyring and the
	// the Commands, ensuring a read from one advance the read pointer for the other.
	// see https://github.com/cosmos/cosmos-sdk/issues/9566.
	ctx.Input = bufio.NewReader(r)
	return ctx
}

// WithCodec returns a copy of the Context with an updated Codec.
func (ctx Context) WithCodec(m codec.Codec) Context {
	ctx.Codec = m
	return ctx
}

// WithLegacyAmino returns a copy of the context with an updated LegacyAmino codec.
// TODO: Deprecated (remove).
func (ctx Context) WithLegacyAmino(cdc *codec.LegacyAmino) Context {
	ctx.LegacyAmino = cdc
	return ctx
}

// WithOutput returns a copy of the context with an updated output writer (e.g. stdout).
func (ctx Context) WithOutput(w io.Writer) Context {
	ctx.Output = w
	return ctx
}

// WithFrom returns a copy of the context with an updated from address or name.
func (ctx Context) WithFrom(from string) Context {
	ctx.From = from
	return ctx
}

// WithOutputFormat returns a copy of the context with an updated OutputFormat field.
func (ctx Context) WithOutputFormat(format string) Context {
	ctx.OutputFormat = format
	return ctx
}

// WithNodeURI returns a copy of the context with an updated node URI.
func (ctx Context) WithNodeURI(nodeURI string) Context {
	ctx.NodeURI = nodeURI
	return ctx
}

// WithHeight returns a copy of the context with an updated height.
func (ctx Context) WithHeight(height int64) Context {
	ctx.Height = height
	return ctx
}

// WithClient returns a copy of the context with an updated RPC client
// instance.
func (ctx Context) WithClient(client cosmosclient.TendermintRPC) Context {
	ctx.Client = client
	return ctx
}

// WithClient returns a copy of the context with an updated RPC client
// instance.
func (ctx Context) WithRPCClient(client rpcclient.Client) Context {
	ctx.RPCClient = client
	ctx.Client = client
	return ctx
}

// WithGRPCClient returns a copy of the context with an updated GRPC client
// instance.
func (ctx Context) WithGRPCClient(grpcClient *grpc.ClientConn) Context {
	ctx.GRPCClient = grpcClient
	return ctx
}

// WithUseLedger returns a copy of the context with an updated UseLedger flag.
func (ctx Context) WithUseLedger(useLedger bool) Context {
	ctx.UseLedger = useLedger
	return ctx
}

// WithChainID returns a copy of the context with an updated chain ID.
func (ctx Context) WithChainID(chainID string) Context {
	ctx.ChainID = chainID
	return ctx
}

// WithHomeDir returns a copy of the Context with HomeDir set.
func (ctx Context) WithHomeDir(dir string) Context {
	if dir != "" {
		ctx.HomeDir = dir
	}
	return ctx
}

// WithKeyringDir returns a copy of the Context with KeyringDir set.
func (ctx Context) WithKeyringDir(dir string) Context {
	ctx.KeyringDir = dir
	return ctx
}

// WithGenerateOnly returns a copy of the context with updated GenerateOnly value
func (ctx Context) WithGenerateOnly(generateOnly bool) Context {
	ctx.GenerateOnly = generateOnly
	return ctx
}

// WithSimulation returns a copy of the context with updated Simulate value
func (ctx Context) WithSimulation(simulate bool) Context {
	ctx.Simulate = simulate
	return ctx
}

// WithOffline returns a copy of the context with updated Offline value.
func (ctx Context) WithOffline(offline bool) Context {
	ctx.Offline = offline
	return ctx
}

// WithFromName returns a copy of the context with an updated from account name.
func (ctx Context) WithFromName(name string) Context {
	ctx.FromName = name
	return ctx
}

// WithFromAddress returns a copy of the context with an updated from account
// address.
func (ctx Context) WithFromAddress(addr sdk.AccAddress) Context {
	ctx.FromAddress = addr
	return ctx
}

// WithFeePayerAddress returns a copy of the context with an updated fee payer account
// address.
func (ctx Context) WithFeePayerAddress(addr sdk.AccAddress) Context {
	ctx.FeePayer = addr
	return ctx
}

// WithFeeGranterAddress returns a copy of the context with an updated fee granter account
// address.
func (ctx Context) WithFeeGranterAddress(addr sdk.AccAddress) Context {
	ctx.FeeGranter = addr
	return ctx
}

// WithBroadcastMode returns a copy of the context with an updated broadcast
// mode.
func (ctx Context) WithBroadcastMode(mode string) Context {
	ctx.BroadcastMode = mode
	return ctx
}

// WithSignModeStr returns a copy of the context with an updated SignMode
// value.
func (ctx Context) WithSignModeStr(signModeStr string) Context {
	ctx.SignModeStr = signModeStr
	return ctx
}

// WithSkipConfirmation returns a copy of the context with an updated SkipConfirm
// value.
func (ctx Context) WithSkipConfirmation(skip bool) Context {
	ctx.SkipConfirm = skip
	return ctx
}

// WithTxConfig returns the context with an updated TxConfig
func (ctx Context) WithTxConfig(generator cosmosclient.TxConfig) Context {
	ctx.TxConfig = generator
	return ctx
}

// WithAccountRetriever returns the context with an updated AccountRetriever
func (ctx Context) WithAccountRetriever(retriever cosmosclient.AccountRetriever) Context {
	ctx.AccountRetriever = retriever
	return ctx
}

// WithInterfaceRegistry returns the context with an updated InterfaceRegistry
func (ctx Context) WithInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) Context {
	ctx.InterfaceRegistry = interfaceRegistry
	return ctx
}

// WithViper returns the context with Viper field. This Viper instance is used to read
// client-side config from the config file.
func (ctx Context) WithViper(prefix string) Context {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	ctx.Viper = v
	return ctx
}

// WithAux returns a copy of the context with an updated IsAux value.
func (ctx Context) WithAux(isAux bool) Context {
	ctx.IsAux = isAux
	return ctx
}

// WithLedgerHasProto returns the context with the provided boolean value, indicating
// whether the target Ledger application can support Protobuf payloads.
func (ctx Context) WithLedgerHasProtobuf(val bool) Context {
	ctx.LedgerHasProtobuf = val
	return ctx
}

// WithPreprocessTxHook returns the context with the provided preprocessing hook, which
// enables chains to preprocess the transaction using the builder.
func (ctx Context) WithPreprocessTxHook(preprocessFn cosmosclient.PreprocessTxFn) Context {
	ctx.PreprocessTxHook = preprocessFn
	return ctx
}

// NewKeyringFromBackend gets a Keyring object from a backend
func NewKeyringFromBackend(ctx Context, backend string) (keyring.Keyring, error) {
	if ctx.Simulate {
		backend = keyring.BackendMemory
	}

	return keyring.New(sdk.KeyringServiceName(), backend, ctx.KeyringDir, ctx.Input, ctx.Codec, ctx.KeyringOptions...)
}
