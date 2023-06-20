package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	// "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client"
	cosmosclientconfig "github.com/cosmos/cosmos-sdk/client/config"
	cosmosclientext "github.com/evmos/ethermint/cosmos_client_ext"
	"github.com/spf13/viper"
)

// Default constants
const (
	chainID        = ""
	keyringBackend = "os"
	output         = "text"
	node           = "tcp://localhost:26657"
	broadcastMode  = "sync"
)

const defaultConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

###############################################################################
###                           Client Configuration                            ###
###############################################################################

# The network chain ID
chain-id = "{{ .ChainID }}"
# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)
keyring-backend = "{{ .KeyringBackend }}"
# CLI output format (text|json)
output = "{{ .Output }}"
# <host>:<port> to Tendermint RPC interface for this chain
node = "{{ .Node }}"
# Transaction broadcasting mode (sync|async)
broadcast-mode = "{{ .BroadcastMode }}"
`

// defaultClientConfig returns the reference to ClientConfig with default values.
func defaultClientConfig() *cosmosclientconfig.ClientConfig {
	return &cosmosclientconfig.ClientConfig{ChainID: chainID, KeyringBackend: keyringBackend, Output: output, Node: node, BroadcastMode: broadcastMode}
}

// ensureConfigPath creates a directory configPath if it does not exist
func ensureConfigPath(configPath string) error {
	return os.MkdirAll(configPath, os.ModePerm)
}

// writeConfigToFile parses defaultConfigTemplate, renders config using the template and writes it to
// configFilePath.
func writeConfigToFile(configFilePath string, config *cosmosclientconfig.ClientConfig) error {
	var buffer bytes.Buffer

	tmpl := template.New("clientConfigFileTemplate")
	configTemplate, err := tmpl.Parse(defaultConfigTemplate)
	if err != nil {
		return err
	}

	if err := configTemplate.Execute(&buffer, config); err != nil {
		return err
	}

	return os.WriteFile(configFilePath, buffer.Bytes(), 0o600)
}

// getClientConfig reads values from client.toml file and unmarshalls them into ClientConfig
func getClientConfig(configPath string, v *viper.Viper) (*cosmosclientconfig.ClientConfig, error) {
	v.AddConfigPath(configPath)
	v.SetConfigName("client")
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(cosmosclientconfig.ClientConfig)
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// ReadFromClientConfig reads values from client.toml file and updates them in client Context
func ReadFromClientConfig(ctx cosmosclientext.Context) (cosmosclientext.Context, error) {
	configPath := filepath.Join(ctx.HomeDir, "config")
	configFilePath := filepath.Join(configPath, "client.toml")
	conf := defaultClientConfig()

	// if config.toml file does not exist we create it and write default ClientConfig values into it.
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		if err := ensureConfigPath(configPath); err != nil {
			return ctx, fmt.Errorf("couldn't make client config: %v", err)
		}

		if ctx.ChainID != "" {
			conf.ChainID = ctx.ChainID // chain-id will be written to the client.toml while initiating the chain.
		}

		if err := writeConfigToFile(configFilePath, conf); err != nil {
			return ctx, fmt.Errorf("could not write client config to the file: %v", err)
		}
	}

	conf, err := getClientConfig(configPath, ctx.Viper)
	if err != nil {
		return ctx, fmt.Errorf("couldn't get client config: %v", err)
	}
	// we need to update KeyringDir field on Client Context first cause it is used in NewKeyringFromBackend
	ctx = ctx.WithOutputFormat(conf.Output).
		WithChainID(conf.ChainID).
		WithKeyringDir(ctx.HomeDir)

	keyring, err := cosmosclientext.NewKeyringFromBackend(ctx, conf.KeyringBackend)
	if err != nil {
		return ctx, fmt.Errorf("couldn't get key ring: %v", err)
	}

	ctx = ctx.WithKeyring(keyring)

	// https://github.com/cosmos/cosmos-sdk/issues/8986
	client, err := client.NewClientFromNode(conf.Node)
	if err != nil {
		return ctx, fmt.Errorf("couldn't get client from nodeURI: %v", err)
	}

	ctx = ctx.WithNodeURI(conf.Node).
		WithClient(client).
		WithBroadcastMode(conf.BroadcastMode)

	return ctx, nil
}
