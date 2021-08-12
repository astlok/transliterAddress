package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"transliterAdress/pkg/transliteParser"
)

func main() {
	fmt.Println("Выберете что вам нужно сделать\n1. Перевод адреса на английский")
	reader := bufio.NewReader(os.Stdin)
	//scan := bufio.NewScanner(os.Stdin)
	char, err := reader.ReadString('\n')
	char = strings.Replace(char, "\n", "", -1)
	for err != nil {
		fmt.Println("Введите корректное число")
		char, err = reader.ReadString('\n')
	}
	switch char {
	case "1":
		fmt.Println("Введите ваше ФИО в формате \"Иванов И.И.\"")
		name, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Введите вашу улицу и дом корректно в формате \"ул. Космонавтов 35с3\"")
			name, err = reader.ReadString('\n')
		}

		fmt.Println("Введите вашу улицу и дом в формате \"ул. Космонавтов 35с3\"")
		street, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Введите вашу улицу и дом корректно в формате \"ул. Космонавтов 35с3\"")
			street, err = reader.ReadString('\n')
		}
		fmt.Println("Введите ваш город/населенный пункт формате \"Сысерть\"")
		city, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Введите ваш город/населенный пункт корректно формате \"Сысерть\"")
			city, err = reader.ReadString('\n')
		}
		fmt.Println("Введите вашу область/край/республику формате \"Краснодарский край\"")
		region, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Введите вашу область/край/республику корректно в формате \"Краснодарский край\"")
			region, err = reader.ReadString('\n')
		}
		fmt.Println("Введите ваш индекс")
		index, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Введите индекс корректно в формате \"228322\"")
			index, err = reader.ReadString('\n')
		}
		fmt.Println("Введите вашу страну")
		country, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Введите вашу страну корректно")
			country, err = reader.ReadString('\n')
		}
		fmt.Println("\nРезультат:\n")
		fmt.Println(transliteParser.TransliteAddress(name, street, city, region, index, country))
	}
}
