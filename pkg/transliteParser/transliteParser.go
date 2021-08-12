package transliteParser

import (
	"transliterAdress/constants"
)

func TransliteAddress(name string, street string, city string, region string, index string, country string) string {
	address := name + street +  city + region + index
	var resultAddress string
	for _, char := range address {
		resultAddress += constants.MatchLetters[string(char)]
	}
	resultAddress += constants.MatchLetters[country]
	return resultAddress
}
