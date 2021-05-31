// https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/09.6.html
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"os"

	"fmt"
	"log"
)

// command line
// arg1: enc|dec
// arg2: 鍵文字列。16、24または32桁。
// arg3: 暗号化したい文字列(encの場合) または 復号化したい文字列(decの場合)
func main() {
	// コマンドライン引数の処理 ここから
	if len(os.Args) < 4 {
		log.Fatal("len(os.Args) < 4")
	}

	subcmd := os.Args[1]
	keyText := []byte(os.Args[2]) // 16、24または32桁の[]byte
	if kl := len(keyText); !(kl == 16 || kl == 24 || kl == 32) {
		log.Fatalf("keyText length must be 16 || 24 || 32 bytes. got=%d", kl)
	}
	text := []byte(os.Args[3])
	// コマンドライン引数の処理 ここまで

	// 暗号化アルゴリズムaesを作成
	c, err := aes.NewCipher(keyText)
	if err != nil {
		log.Fatal(err)
	}

	var iv = []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	switch subcmd {
	case "enc":
		enc := encoder(c, iv, text)
		fmt.Printf("%s", enc)
	case "dec":
		dec := decoder(c, iv, text)
		fmt.Printf("%s", string(dec))
	default:
		log.Fatalf("unsupported sub command %q", subcmd)
	}
}

func encoder(c cipher.Block, iv []byte, plainText []byte) []byte {
	cfb := cipher.NewCFBEncrypter(c, iv)
	ciph := make([]byte, len(plainText))
	cfb.XORKeyStream(ciph, []byte(plainText))
	return ciph
}

func decoder(c cipher.Block, iv []byte, cipheredText []byte) []byte {
	cfb := cipher.NewCFBDecrypter(c, iv)
	dec := make([]byte, len(cipheredText))
	cfb.XORKeyStream(dec, []byte(cipheredText))
	return dec
}
