#!/usr/bin/env bash
# Run go test in every workspace module (repo root is not a Go module).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

modules=(
  gateway
  shared
  services/admin
  services/merchant
  services/customer
  services/ledger
  services/report
)

for m in "${modules[@]}"; do
  echo "==> go test ./... ($m)"
  (cd "$m" && go test ./...)
done

echo "All modules OK"
