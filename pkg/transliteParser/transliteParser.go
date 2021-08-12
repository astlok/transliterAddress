package transliteParser

import (
	"fmt"
	"transliterAdress/constants"
)

func TransliteAddress(name string, street string, city string, region string, index string, country string) string {
	address := name + street +  city + region + index + country
	var resultAddress string
	for _, char := range address {
		if char > 1103 {
			continue
		}
		kek := constants.MatchLetters[string((char))]
		fmt.Println(kek)
		resultAddress += constants.MatchLetters[string(char)]
	}
	return resultAddress
}
