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
	// in := "4HBCKKNSWXCY5N7I5UPO2YZ7KHPCJOMML5IOBB22AZJNYCFAPQNFCSEDQ4DU3MY7"
	// key, err := formatInput(64)

	key, err := randStringBytesMaskImprSrc(64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("The key is", string(key))
	for {
		pwd, remain := oneTimePassword(key, 30) // default interval is 30 second
		for remain > 0 {
			/* clear the srceen not suitable
			fmt.Printf("\033[H\033[2J")
			fmt.Printf("%06d (%d second(s) remaining)\n", pwd, remain)
			time.Sleep(1 * time.Second)
			*/
			fmt.Printf("%06d (%d second(s) remaining)\r", pwd, remain)
			time.Sleep(1 * time.Second)
			remain--

		}
	}
}

func formatInput() (b []byte, err error) {
	if len(os.Args) < 2 {
		err = fmt.Errorf("must specify key to use")
		return
	}
	input := os.Args[1]
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	in := strings.ToUpper(inputNoSpaces)
	if len(in) != 64 {
		err = fmt.Errorf("need 64 size key, got %d", len(in))
		return
	}
	b, err = base32.StdEncoding.DecodeString(in)
	return
}
