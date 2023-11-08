let
  pkgs = import ../../../nix { };
  fetchEthermint = rev: builtins.fetchTarball "https://github.com/tokenize-titan/ethermint/archive/${rev}.tar.gz";
  released = pkgs.buildGoApplication rec {
    pname = "ethermintd";
    version = "v0.23.0-t.1";
    # src = fetchEthermint "d29cdad6e667f6089dfecbedd36bb8d3a2a7d025";
    src = pkgs.lib.sourceByRegex ../../../. [
      "^(x|app|cmd|client|server|crypto|rpc|types|encoding|ethereum|indexer|testutil|version|go.mod|go.sum|gomod2nix.toml)($|/.*)"
      "^tests(/.*[.]go)?$"
    ];
    pwd = src;
    # modules = gomod2nix.toml;
    subPackages = [ "cmd/ethermintd" ];
    # vendorSha256 = "sha256-cQAol54b6hNzsA4Q3MP9mTqFWM1MvR5uMPrYpaoj3SY=";
    doCheck = false;
    CGO_ENABLED = "1";
  };
  current = pkgs.callPackage ../../../. { };
in
pkgs.linkFarm "upgrade-test-package" [
  { name = "genesis"; path = released; }
  { name = "integration-test-upgrade"; path = current; }
]
