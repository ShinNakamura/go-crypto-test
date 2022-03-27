#! /bin/bash
unalias -a

## go build -o ./bin/go-crypto-test

# mydir -p myconf
# chmod 600 myconf
# chmod 400 myconf/passwd

# 暗号化
k="$(cat myconf/passwd)"
t="Hello, Cipher."
echo "$t を myconf/passwd 内の文字列をパスワードとして使い暗号化"
./bin/go-crypto-test enc "$k" "$t" | base64

echo ""

# 復号化
echo "myconf/passwd 内の文字列をパスワードとして使い暗号化した $t を復号化"
./bin/go-crypto-test dec "$k" "$(./bin/go-crypto-test enc "$k" "$t")"
# Hello, Cipher.

echo ""
