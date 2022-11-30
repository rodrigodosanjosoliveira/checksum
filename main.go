package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const sha256HashFunction string = "sha256"

func main() {
	var hashFunction string
	var fileName string
	var providedHash string
	var result bool
	var err error

	fmt.Print("Enter the hash function: ")
	fmt.Scanf("%s", &hashFunction)

	fmt.Print("Enter the provided hash: ")
	fmt.Scanf("%s", &providedHash)

	fmt.Print("Enter the file name to check: ")
	fmt.Scanf("%s", &fileName)

	switch strings.ToLower(hashFunction) {
	case sha256HashFunction:
		result, err = checkSha256(fileName, providedHash)
	default:
		log.Fatal(fmt.Errorf("hash function invalid: %s", hashFunction))
	}

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if result {
		fmt.Println("File is valid!")
	}
}

func checkSha256(filename string, providedHash string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}

	result := hex.EncodeToString(hasher.Sum(nil))

	return result == providedHash, nil
}
