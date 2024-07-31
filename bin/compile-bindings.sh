#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

# Constants
readonly BINDING_DIR="./pkg/bindings"
readonly JSON_DIR="./out"

die() {
    echo "Error: $*" >&2
    exit 1
}

assert_command() {
    command -v "$1" >/dev/null 2>&1 || die "$1 is required but not installed."
}

assert_dir() {
    [[ -d "$1" ]] || die "Directory $1 does not exist."
}

create_binding() {
    local contract_name="$1"
    local contract_json_path="${JSON_DIR}/${contract_name}.sol/${contract_name}.json"
    local binding_out_dir="${BINDING_DIR}/${contract_name}"

    [[ -f "$contract_json_path" ]] || die "Contract JSON file not found: $contract_json_path"

    mkdir -p "$binding_out_dir" || die "Failed to create directory: $binding_out_dir"

    jq -er '.abi' "$contract_json_path" > "$binding_out_dir/tmp.abi" || die "Failed to extract ABI from $contract_json_path"
    jq -er '.bytecode.object' "$contract_json_path" > "$binding_out_dir/tmp.bin" || die "Failed to extract bytecode from $contract_json_path"

    if ! abigen --bin="$binding_out_dir/tmp.bin" \
                --abi="$binding_out_dir/tmp.abi" \
                --pkg="$contract_name" \
                --out="$BINDING_DIR/$contract_name/binding.go" \
                > /dev/null 2>&1; then
        die "Failed to generate binding for $contract_json_path"
    fi

    rm "$binding_out_dir/tmp.abi" "$binding_out_dir/tmp.bin"
}

main() {
    assert_command jq
    assert_command abigen
    assert_dir "$JSON_DIR"

    while IFS= read -r -d '' contract_file; do
        contract_name=$(basename "$contract_file" .sol)
        create_binding "$contract_name"
    done < <(find src/contracts -type f -name "*.sol" -print0)
}

main "$@"
