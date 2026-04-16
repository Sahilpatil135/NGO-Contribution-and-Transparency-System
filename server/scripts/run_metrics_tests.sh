#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
REPORT_DIR="$ROOT_DIR/reports/perf"

mkdir -p "$REPORT_DIR"

echo "Running API benchmark tests..."
(
  cd "$ROOT_DIR"
  go test ./internal/server -run '^$' -bench 'BenchmarkHelloWorldHandler' -benchmem -count=3
) | tee "$REPORT_DIR/api_benchmark.txt"

echo "Running smart contract gas tests..."
(
  cd "$ROOT_DIR/contracts"
  REPORT_GAS=true npx hardhat test
) | tee "$REPORT_DIR/contracts_gas_report.txt"

echo "Summarizing basic quantitative metrics..."
(
  cd "$ROOT_DIR"
  python3 ./scripts/summarize_metrics.py
)

echo "Metrics reports generated in $REPORT_DIR"
