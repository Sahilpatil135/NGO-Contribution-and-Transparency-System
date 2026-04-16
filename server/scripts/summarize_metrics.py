#!/usr/bin/env python3
import csv
import json
import re
from pathlib import Path

ANSI_ESCAPE = re.compile(r"\x1B\[[0-?]*[ -/]*[@-~]")


def parse_api_benchmarks(path: Path):
    pattern = re.compile(
        r"^(BenchmarkHelloWorldHandler(?:Parallel)?)-\d+\s+\d+\s+(\d+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op$"
    )
    grouped = {}
    for line in path.read_text().splitlines():
        m = pattern.match(line.strip())
        if not m:
            continue
        name, ns, bytes_op, allocs = m.group(1), int(m.group(2)), int(m.group(3)), int(m.group(4))
        grouped.setdefault(name, []).append((ns, bytes_op, allocs))

    output = {}
    for name, runs in grouped.items():
        count = len(runs)
        avg_ns = sum(r[0] for r in runs) / count
        avg_bytes = sum(r[1] for r in runs) / count
        avg_allocs = sum(r[2] for r in runs) / count
        output[name] = {
            "samples": count,
            "avg_latency_ns": round(avg_ns, 2),
            "avg_latency_ms": round(avg_ns / 1_000_000, 6),
            "avg_throughput_rps": round(1_000_000_000 / avg_ns, 2),
            "avg_bytes_per_op": round(avg_bytes, 2),
            "avg_allocs_per_op": round(avg_allocs, 2),
        }
    return output


def parse_contract_gas(path: Path):
    contract_header_pattern = re.compile(r"^\|\s*([A-Za-z][A-Za-z0-9]*)\s*·")
    method_row_pattern = re.compile(r"^\|\s*([A-Za-z][A-Za-z0-9_]*)\s*·\s*([0-9,\-]+)\s*·\s*([0-9,\-]+)\s*·\s*([0-9,\-]+)\s*·")
    methods = {}
    current_contract = None
    for raw_line in path.read_text().splitlines():
        line = ANSI_ESCAPE.sub("", raw_line).strip()
        if not line.startswith("|"):
            continue
        if "Contracts / Methods" in line or "Solidity and Network Configuration" in line:
            continue

        header_match = contract_header_pattern.match(line)
        if header_match and "recordDonation" not in line and "registerOrUpdateCause" not in line:
            value = header_match.group(1)
            if value not in ("Methods", "Deployments", "Key", "Solidity"):
                current_contract = value
            continue

        method_match = method_row_pattern.match(line)
        if not method_match or current_contract is None:
            continue

        method, min_gas, max_gas, avg_gas = method_match.groups()
        avg_clean = avg_gas.replace(",", "").strip()
        if not avg_clean.isdigit():
            continue

        key = f"{current_contract}.{method}"
        methods[key] = {
            "avg_gas": int(avg_clean),
            "min_gas": None if min_gas.strip() == "-" else int(min_gas.replace(",", "").strip()),
            "max_gas": None if max_gas.strip() == "-" else int(max_gas.replace(",", "").strip()),
        }
    return methods


def main():
    root = Path(__file__).resolve().parent.parent
    report_dir = root / "reports" / "perf"
    api_file = report_dir / "api_benchmark.txt"
    gas_file = report_dir / "contracts_gas_report.txt"
    json_out = report_dir / "basic_metrics.json"
    csv_out = report_dir / "basic_metrics.csv"

    api_metrics = parse_api_benchmarks(api_file)
    gas_metrics = parse_contract_gas(gas_file)

    data = {"api": api_metrics, "blockchain": gas_metrics}
    json_out.write_text(json.dumps(data, indent=2))

    with csv_out.open("w", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(["category", "metric", "value", "unit"])
        for bench_name, metrics in api_metrics.items():
            writer.writerow(["api", f"{bench_name}.avg_latency_ms", metrics["avg_latency_ms"], "ms"])
            writer.writerow(["api", f"{bench_name}.avg_throughput_rps", metrics["avg_throughput_rps"], "rps"])
            writer.writerow(["api", f"{bench_name}.avg_bytes_per_op", metrics["avg_bytes_per_op"], "bytes/op"])
            writer.writerow(["api", f"{bench_name}.avg_allocs_per_op", metrics["avg_allocs_per_op"], "allocs/op"])
        for method, metrics in gas_metrics.items():
            writer.writerow(["blockchain", f"{method}.avg_gas", metrics["avg_gas"], "gas"])
            if metrics["min_gas"] is not None:
                writer.writerow(["blockchain", f"{method}.min_gas", metrics["min_gas"], "gas"])
            if metrics["max_gas"] is not None:
                writer.writerow(["blockchain", f"{method}.max_gas", metrics["max_gas"], "gas"])

    print(f"Wrote {json_out}")
    print(f"Wrote {csv_out}")


if __name__ == "__main__":
    main()
