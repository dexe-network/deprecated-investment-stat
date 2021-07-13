package helpers

import (
	"github.com/ethereum/go-ethereum/common"
	"regexp"
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