package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

const (
	RSAPublicKey = "-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCKFctVrhfF3m2Kes0FBL/JFeO\n" +
		"cmNg9eJz8k/hQy1kadD+XFUpluRqa//Uxp2s9W2qE0EoUCu59ugcf/p7lGuL99Uo\n" +
		"SGmQEynkBvZct+/M40L0E0rZ4BVgzLOJmIbXMp0J4PnPcb6VLZvxazGcmSfjauC7\n" +
		"F3yWYqUbZd/HCBtawwIDAQAB\n" +
		"-----END PUBLIC KEY-----"

	RSAPublicKey2 = "h0HtbcA_ud27f5vc4U_9OsB2fn3Ar5QD6bpuHB1VGTXDB_zIko2ENmtHQmJAZJEEGJxA5v1fzs7v3Yk6WRY7XbJFYvKWr8A7_txUgwPCFaR0eH1HpiCbldw4X6Y690O75ksoSepbyYwmdi5u2JqX1lz3a2O5taYdBYC0pO6gaNfgT-lYSf4Ws5CAZND3qhMLD8Cnby4n0Hxj6xnpr8ODAnVNbWQ0JECthfjolCI026t87kC7S5hHSnd2DFvM4arHE7TRj__3SrBKzcJZxM70ApNkAwytOUgLbHKmL9x2IXW5x650mqloaR0ZHiizD9vjvzFm42D9OqDYcAaywsZotQ"
)

func RSAEncryptString(s string) string {
	byteString := []byte(s)
	block, _ := pem.Decode([]byte(RSAPublicKey))
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)
	encBytes, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, byteString)
	return string(encBytes)
}

func parseKey(s string) *rsa.PublicKey {
	p, _ := pem.Decode([]byte(s))
	k, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		panic(err)
	}
	return k.(*rsa.PublicKey)
}

// func parseKey2(s string) *rsa.PublicKey {
// 	p, _ := pem.Decode(base64urluint.Decode(s))
// 	k, err := x509.ParsePKCS1PublicKey(p.Bytes)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return k
// }

func main() {
	// stringToEncrypt := "this_is_my_string"
	// encryptedString := RSAEncryptString(stringToEncrypt)
	// doWork(encryptedString)
	k := parseKey(RSAPublicKey)
	// k := parseKey2(RSAPublicKey2)
	fmt.Println(k)
}

func doWork(s string) {
	fmt.Println(s)
}
