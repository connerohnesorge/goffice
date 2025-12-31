#!/usr/bin/env bash
set -euo pipefail

# Build the C# DocxBridge project

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$SCRIPT_DIR/DocxBridge"

echo "Building DocxBridge..."
cd "$PROJECT_DIR"

dotnet restore
dotnet build --configuration Release --no-restore

echo "Build complete. Binary at: bin/Release/net9.0/DocxBridge"
