package helpers

import (
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	"regexp"
	"time"
	"unicode/utf8"
)

func AddressArrToStringArr(data []common.Address) (result []string) {
	for _, item := range data {
		result = append(result, item.String())
	}
	return
}

func IsValidAddress(v string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(v)
}

func TrimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func RandomNonce() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
