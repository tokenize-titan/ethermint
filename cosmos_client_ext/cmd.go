package cosmos_client_ext

import (
	"crypto/tls"
	"errors"

	"github.com/cometbft/cometbft/libs/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

const ClientContextKey = sdk.ContextKey("cosmos_client_ext.context")

// SetCmdClientContextHandler is to be used in a command pre-hook execution to
// read flags that populate a Context and sets that to the command's Context.
func SetCmdClientContextHandler(clientCtx Context, cmd *cobra.Command) (err error) {
	clientCtx, err = ReadPersistentCommandFlags(clientCtx, cmd.Flags())
	if err != nil {
		return err
	}

	return SetCmdClientContext(cmd, clientCtx)
}

// GetClientQueryContext returns a Context from a command with fields set based on flags
// defined in AddQueryFlagsToCmd. An error is returned if any flag query fails.
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func GetClientQueryContext(cmd *cobra.Command) (Context, error) {
	ctx := GetClientContextFromCmd(cmd)
	return readQueryCommandFlags(ctx, cmd.Flags())
}

// GetClientContextFromCmd returns a Context from a command or an empty Context
// if it has not been set.
func GetClientContextFromCmd(cmd *cobra.Command) Context {
	if v := cmd.Context().Value(ClientContextKey); v != nil {
		clientCtxPtr := v.(*Context)
		return *clientCtxPtr
	}

	return Context{}
}

// readQueryCommandFlags returns an updated Context with fields set based on flags
// defined in AddQueryFlagsToCmd. An error is returned if any flag query fails.
//
// Note, the provided clientCtx may have field pre-populated. The following order
// of precedence occurs:
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func readQueryCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	if clientCtx.Height == 0 || flagSet.Changed(flags.FlagHeight) {
		height, _ := flagSet.GetInt64(flags.FlagHeight)
		clientCtx = clientCtx.WithHeight(height)
	}

	if !clientCtx.UseLedger || flagSet.Changed(flags.FlagUseLedger) {
		useLedger, _ := flagSet.GetBool(flags.FlagUseLedger)
		clientCtx = clientCtx.WithUseLedger(useLedger)
	}

	return ReadPersistentCommandFlags(clientCtx, flagSet)
}

// ReadPersistentCommandFlags returns a Context with fields set for "persistent"
// or common flags that do not necessarily change with context.
//
// Note, the provided clientCtx may have field pre-populated. The following order
// of precedence occurs:
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func ReadPersistentCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	if clientCtx.OutputFormat == "" || flagSet.Changed(cli.OutputFlag) {
		output, _ := flagSet.GetString(cli.OutputFlag)
		clientCtx = clientCtx.WithOutputFormat(output)
	}

	if clientCtx.HomeDir == "" || flagSet.Changed(flags.FlagHome) {
		homeDir, _ := flagSet.GetString(flags.FlagHome)
		clientCtx = clientCtx.WithHomeDir(homeDir)
	}

	if !clientCtx.Simulate || flagSet.Changed(flags.FlagDryRun) {
		dryRun, _ := flagSet.GetBool(flags.FlagDryRun)
		clientCtx = clientCtx.WithSimulation(dryRun)
	}

	if clientCtx.KeyringDir == "" || flagSet.Changed(flags.FlagKeyringDir) {
		keyringDir, _ := flagSet.GetString(flags.FlagKeyringDir)

		// The keyring directory is optional and falls back to the home directory
		// if omitted.
		if keyringDir == "" {
			keyringDir = clientCtx.HomeDir
		}

		clientCtx = clientCtx.WithKeyringDir(keyringDir)
	}

	if clientCtx.ChainID == "" || flagSet.Changed(flags.FlagChainID) {
		chainID, _ := flagSet.GetString(flags.FlagChainID)
		clientCtx = clientCtx.WithChainID(chainID)
	}

	if clientCtx.Keyring == nil || flagSet.Changed(flags.FlagKeyringBackend) {
		keyringBackend, _ := flagSet.GetString(flags.FlagKeyringBackend)

		if keyringBackend != "" {
			kr, err := cosmosclient.NewKeyringFromBackend(clientCtx.Context, keyringBackend)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithKeyring(kr)
		}
	}

	if clientCtx.Client == nil || flagSet.Changed(flags.FlagNode) {
		rpcURI, _ := flagSet.GetString(flags.FlagNode)
		if rpcURI != "" {
			clientCtx = clientCtx.WithNodeURI(rpcURI)

			client, err := cosmosclient.NewClientFromNode(rpcURI)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithClient(client)
			clientCtx = clientCtx.WithRPCClient(client)
		}
	}

	if clientCtx.GRPCClient == nil || flagSet.Changed(flags.FlagGRPC) {
		grpcURI, _ := flagSet.GetString(flags.FlagGRPC)
		if grpcURI != "" {
			var dialOpts []grpc.DialOption

			useInsecure, _ := flagSet.GetBool(flags.FlagGRPCInsecure)
			if useInsecure {
				dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
			} else {
				dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
					MinVersion: tls.VersionTLS12,
				})))
			}

			grpcClient, err := grpc.Dial(grpcURI, dialOpts...)
			if err != nil {
				return Context{}, err
			}
			clientCtx = clientCtx.WithGRPCClient(grpcClient)
		}
	}

	return clientCtx, nil
}

// SetCmdClientContext sets a command's Context value to the provided argument.
func SetCmdClientContext(cmd *cobra.Command, clientCtx Context) error {
	v := cmd.Context().Value(ClientContextKey)
	if v == nil {
		return errors.New("client context not set")
	}

	clientCtxPtr := v.(*Context)
	*clientCtxPtr = clientCtx

	v2 := cmd.Context().Value(cosmosclient.ClientContextKey)
	if v == nil {
		return errors.New("client context not set")
	}

	clientCtxPtr2 := v2.(*cosmosclient.Context)
	*clientCtxPtr2 = clientCtx.Context

	return nil
}
