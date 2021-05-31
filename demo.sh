#! /bin/bash
unalias -a

## go build -o ./bin/go-crypto-test

# 暗号化
k="0123456789abcdef;[]+-|()"
t="Hello, Cipher."
./bin/go-crypto-test enc "$k" "$t" | base64
# rrDuwXGmMB+rr8mqgU0=

# 復号化
./bin/go-crypto-test dec "$k" "$(./bin/go-crypto-test enc "$k" "$t")"
# Hello, Cipher.

echo ""
