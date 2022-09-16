package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	k := len(letters)
	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	const currencies = "USD, EUR, GBP, CHF, JPY, CAD, AUD, NZD"
	currenciesList := strings.Split(currencies, ", ")
	n := len(currenciesList)
	return currenciesList[rand.Intn(n)]
}