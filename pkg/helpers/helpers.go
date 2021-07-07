package helpers

import "github.com/ethereum/go-ethereum/common"

func AddressArrToStringArr(data []common.Address) (result []string) {
	for _, item := range data {
		result = append(result, item.String())
	}
	return
}
