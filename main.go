package main

import (
	"bytes"
	"encoding/hex"
	"github.com/edgelesssys/ego/enclave"
	"log"
)

func uniqkey() {
	for i := 0; i < 10; i++ {
		k, kinfo, err := enclave.GetUniqueSealKey()
		if err != nil {
			log.Println("get uniq key failed", err)
			return
		}
		log.Printf("got key(%s), kinfo(%s)\n", hex.EncodeToString(k1), hex.EncodeToString(kinfo))

		k2, err := enclave.GetSealKey(kinfo)
		if err != nil {
			log.Println("reget seal key failed", err)
			return
		}
		if bytes.Compare(k, k2) == 0 {
			log.Println("uniqkey and reget is equal")
		}
	}
}

func productKey() {

	k1, kinfo, err := enclave.GetProductSealKey()
	if err != nil {
		log.Println("get product key failed", err)
		return
	}
	log.Printf("got key(%s), kinfo(%s)\n", hex.EncodeToString(k1), hex.EncodeToString(kinfo))

	k2, err := enclave.GetSealKey(kinfo)
	if err != nil {
		log.Println("reget seal key failed", err)
		return
	}
	log.Printf("get key(%s) with kinfo\n", hex.EncodeToString(k2))
}

func main() {
	uniqkey()
	productKey()
}
