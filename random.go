package main

import (
	"math/rand"
	"time"
)

const (
	letterIdxBits = 6                                                                // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1                                             // All 1-bits, as many as letterIdxBits
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" //"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	step          = 63 / letterIdxBits                                               // # of letter indices fitting in 63 bits
)

var (
	src    = rand.NewSource(time.Now().UnixNano())
	length = len(letters)
)

/*
// String retrun a random string of size
func String(size int) string {
	return string(randStringBytesMaskImprSrc(size))
}
*/

func randStringBytesMaskImprSrc(n int) (b []byte, err error) {
	/*
		var mu sync.Mutex
		mu.Lock()
		defer mu.Unlock()
	*/
	b = make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), step; i >= 0; cache >>= letterIdxBits {
		if remain == 0 {
			cache, remain = src.Int63(), step
		}
		idx := int(cache&letterIdxMask) % length
		b[i] = letters[idx]
		i--
		remain--
	}
	return b, nil
}
