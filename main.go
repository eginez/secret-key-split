package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"os"

	"fmt"
	"math/big"
)

const primeBitSize = 1024

//P Large prime number, safe to share
var P *big.Int

func init() {
	P, _ = rand.Prime(rand.Reader, primeBitSize)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Split2(secret []byte) (parts [][]byte, err error) {
	r, _ := rand.Int(rand.Reader, P)
	doubler := big.NewInt(0).Mul(big.NewInt(2), r)
	secInt := big.NewInt(0).SetBytes(secret)

	if P.Cmp(secInt) < 1 {
		panic(fmt.Errorf("P is not large enough"))
	}

	x1 := big.NewInt(0).Mod(big.NewInt(0).Add(secInt, r), P)
	x2 := big.NewInt(0).Mod(big.NewInt(0).Add(secInt, doubler), P)

	res := make([][]byte, 2)
	res[0] = x1.Bytes()
	res[1] = x2.Bytes()
	return res, nil
}

func Combine(parts [][]byte) (secret []byte) {
	x1 := big.NewInt(0).SetBytes(parts[0])
	x2 := big.NewInt(0).SetBytes(parts[1])

	// S = (2P1 - P2) MOD P
	diff := big.NewInt(0).Sub(big.NewInt(0).Mul(big.NewInt(2), x1), x2)
	sec := big.NewInt(0).Mod(diff, P)
	return sec.Bytes()
}

func main() {

	input := os.Args[1]
	hash := sha1.New()
	hash.Write([]byte(input))
	secret := hash.Sum(nil)

	parts, err := Split2(secret[:])
	panicIfErr(err)

	for _, part := range parts {
		fmt.Println(hex.EncodeToString(part))
	}
}
