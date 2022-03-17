package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	digits_in_number = 4
)

func main() {
	var (
		str string
	)
	input(&str)
	fmt.Println("ИТОГ - " + str)

}

func input(str *string) {
	exit := 0
	for exit != 1 {
		exit = 0
		fmt.Print("Введите ", digits_in_number, "-значное число(без повтора цифр): ")
		for fmt.Scan(str); len(*str) != digits_in_number; fmt.Scan(str) {
			fmt.Println("val =", *str, "len =", len(*str))
			incorrect_input(str)
			fmt.Print("Повторите попытку: ")
		}
		if i := fmt.Sprint(strings.ToUpper(*str)); i == "EXIT" {
			*str = i
			break
		}
		exit = 2
		for _, v := range *str {
			if v < 48 || v > 57 {
				exit = -1
				break
			}
		}
	Flag:
		switch exit {
		case -1:
			incorrect_input(str)
		case 2:
			for i := 1; i < len(*str); i++ {
				for j := 0; j < i; j++ {
					if (*str)[i] == (*str)[j] {
						exit = 0
						break
					}
				}
				if exit == 0 {
					incorrect_input(str)
					break Flag
				}
			}
			exit = 1
		default:
			fmt.Println("ERROR exit = ", exit)
		}
	}
}

func incorrect_input(str *string) {
	fmt.Print("[")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("]")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Некорректный ввод! (" + *str + ")")
}
