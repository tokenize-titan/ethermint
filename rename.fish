#!/usr/local/bin/fish

# ensure script receives one arguments
if test (count $argv) -ne 1
  echo "Usage: rename.fish <new module name>"
  exit 1
end

# get current module name
set -l old (go list -m)

echo old module name : $old
echo new module name : $argv[1]

# replace module name in go.mod
go mod edit -module $argv[1]

# replace module name in all go files
find . -name '*.go' -type f -exec sed -i '' -e 's|'$old'|'$argv[1]'|g' {} \;

# replace module name in all proto files
find . -name '*.proto' -type f -exec sed -i '' -e 's|'$old'|'$argv[1]'|g' {} \;