package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"
)

func oneTimePassword(key []byte, interval int64) (uint32, int64) {
	epochSeconds := time.Now().Unix()
	secondsRemaining := interval - (epochSeconds % interval)

	value := make([]byte, 8)
	binary.PutVarint(value, epochSeconds/interval)

	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	bytes := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	bytes[0] = bytes[0] & 0x7F

	number := binary.BigEndian.Uint32(bytes)
	return number % 1000000, secondsRemaining
}

func main() {
	//in := formatInput()
	in := "4HBCKKNSWXCY5N7I5UPO2YZ7KHPCJOMML5IOBB22AZJNYCFAPQNFCSEDQ4DU3MY7"
	key, err := base32.StdEncoding.DecodeString(in)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	pwd, remain := oneTimePassword(key, 30) // default interval 30s

	fmt.Printf("%06d (%d second(s) remaining)\n", pwd, remain)
}

func formatInput() string {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "must specify key to use")
		os.Exit(1)
	}

	input := os.Args[1]
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	return strings.ToUpper(inputNoSpaces)

}
