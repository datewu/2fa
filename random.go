package main

import (
	"encoding/binary"
	"time"

	"github.com/sinalpha/security"
)

func randStr() string {
	randBytes := security.NewEncryptionKey()
	return security.ToString(randBytes[:])
}

// Gen2fa ...
func Gen2fa(tag string, interval int64) (uint32, int64) {
	epochSeconds := time.Now().Unix()
	secondsRemaining := interval - (epochSeconds % interval)

	value := make([]byte, 8)
	binary.PutVarint(value, epochSeconds/interval)

	h := security.Hash(tag, value)

	offset := h[len(h)-1] & 0x0F
	bytes := h[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	bytes[0] = bytes[0] & 0x7F

	number := binary.BigEndian.Uint32(bytes)
	return number % 1000000, secondsRemaining
}
