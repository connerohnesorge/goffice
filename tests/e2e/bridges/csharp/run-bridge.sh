#!/usr/bin/env bash
# Wrapper script to run the C# bridge with dotnet
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec dotnet "${SCRIPT_DIR}/DocxBridge/bin/Release/net9.0/DocxBridge.dll" "$@"
