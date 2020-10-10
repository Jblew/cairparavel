package eventinputtypes

import (
	"strconv"
	"log"
)

func parseIntOrZero(s string, base int, bitSize int) int64 {
	out, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		log.Printf("Error while parsing str to int %v", err)
		return 0
	}
	return out
}
