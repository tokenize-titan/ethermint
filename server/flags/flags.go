// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/tokenize-titan/ethermint/blob/main/LICENSE
package flags

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
)

// Tendermint/cosmos-sdk full-node start flags
const (
	WithTendermint = "with-tendermint"
	Address        = "address"
	Transport      = "transport"
	TraceStore     = "trace-store"
	CPUProfile     = "cpu-profile"
	// The type of database for application and snapshots databases
	AppDBBackend = "app-db-backend"
)

// GRPC-related flags.
const (
	GRPCOnly       = "grpc-only"
	GRPCEnable     = "grpc.enable"
	GRPCAddress    = "grpc.address"
	GRPCWebEnable  = "grpc-web.enable"
	GRPCWebAddress = "grpc-web.address"
)

// Cosmos API flags
const (
	RPCEnable         = "api.enable"
	EnabledUnsafeCors = "api.enabled-unsafe-cors"
)

// JSON-RPC flags
const (
	JSONRPCEnable              = "json-rpc.enable"
	JSONRPCAPI                 = "json-rpc.api"
	JSONRPCAddress             = "json-rpc.address"
	JSONWsAddress              = "json-rpc.ws-address"
	JSONRPCGasCap              = "json-rpc.gas-cap"
	JSONRPCEVMTimeout          = "json-rpc.evm-timeout"
	JSONRPCTxFeeCap            = "json-rpc.txfee-cap"
	JSONRPCFilterCap           = "json-rpc.filter-cap"
	JSONRPCLogsCap             = "json-rpc.logs-cap"
	JSONRPCBlockRangeCap       = "json-rpc.block-range-cap"
	JSONRPCHTTPTimeout         = "json-rpc.http-timeout"
	JSONRPCHTTPIdleTimeout     = "json-rpc.http-idle-timeout"
	JSONRPCAllowUnprotectedTxs = "json-rpc.allow-unprotected-txs"
	JSONRPCMaxOpenConnections  = "json-rpc.max-open-connections"
	JSONRPCEnableIndexer       = "json-rpc.enable-indexer"
	// JSONRPCEnableMetrics enables EVM RPC metrics server.
	// Set to `metrics` which is hardcoded flag from go-ethereum.
	// https://github.com/ethereum/go-ethereum/blob/master/metrics/metrics.go#L35-L55
	JSONRPCEnableMetrics            = "metrics"
	JSONRPCFixRevertGasRefundHeight = "json-rpc.fix-revert-gas-refund-height"
)

// EVM flags
const (
	EVMTracer         = "evm.tracer"
	EVMMaxTxGasWanted = "evm.max-tx-gas-wanted"
)

// TLS flags
const (
	TLSCertPath = "tls.certificate-path"
	TLSKeyPath  = "tls.key-path"
)

type FlagOverride struct {
	Value interface{}
	Usage string
}

// AddTxFlags adds common flags for commands to post tx
func UpdateFlags(cmd *cobra.Command) (*cobra.Command, error) {
	overrides := map[string]FlagOverride{
		flags.FlagFees: {
			Value: "",
			Usage: "Fees to pay along with transaction; eg: 10aphoton",
		},
		flags.FlagGasPrices: {
			Value: "",
			Usage: "Gas prices to determine the transaction fee (e.g. 10aphoton)",
		},
		flags.FlagKeyringBackend: {
			Usage: "Select keyring's backend (os|file|kwallet|pass|test|memory) (default in client.toml)",
		},
		genutilcli.FlagDefaultBondDenom: {
			Value: "aphoton",
			Usage: "genesis file default denomination, if left blank default value is 'aphoton'",
		},
	}

	return OverrideFlags(cmd, overrides)
}

func OverrideFlags(cmd *cobra.Command, overrides map[string]FlagOverride) (*cobra.Command, error) {
	for name, override := range overrides {
		persistentflag := cmd.PersistentFlags().Lookup(name)
		flag := cmd.Flags().Lookup(name)

		if persistentflag != nil {
			if override.Value != nil {
				err := persistentflag.Value.Set(fmt.Sprintf("%v", override.Value))
				if err != nil {
					return nil, err
				}
			}
			if override.Usage != "" {
				persistentflag.Usage = override.Usage
			}
		}
		if flag != nil {
			if override.Value != nil {
				err := flag.Value.Set(fmt.Sprintf("%v", override.Value))
				if err != nil {
					return nil, err
				}
			}
			if override.Usage != "" {
				flag.Usage = override.Usage
			}
		}
	}

	for _, childCmd := range cmd.Commands() {
		_, err := OverrideFlags(childCmd, overrides)
		if err != nil {
			return nil, err
		}
	}

	return cmd, nil
}
