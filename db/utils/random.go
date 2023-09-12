package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generate random currency type
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "JPY", "GBP", "INR", "NPR", "CNY"}
	return currencies[rand.Intn(len(currencies))]
}
