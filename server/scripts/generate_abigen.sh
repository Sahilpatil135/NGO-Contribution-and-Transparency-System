#!/bin/env sh

shopt -s extglob

mkdir -p internal/blockchain/{abi,contracts}

for compiledJson in contracts/artifacts/contracts/*.sol/!(*.dbg).json; do
    echo "ABIGEN from $compiledJson"

    filename=$(basename $compiledJson .json)

    ABI=internal/blockchain/abi/$filename.abi
    OUT=internal/blockchain/contracts/$filename.go

    jq .abi $compiledJson > $ABI

    abigen \
        --abi $ABI \
        --out $OUT \
        --type $filename \
        --pkg contracts 
done
