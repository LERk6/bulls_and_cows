package main

import (
	"fmt"
	"time"
)

const (
	digits_in_number = 4
)

func main() {
	var (
		str string
		//num uint16
	)
	input(&str)

}

func input(str *string) {
	fmt.Print("Ваше число: ")
	for fmt.Scan(str); len(*str) != digits_in_number; fmt.Scan(str) {
		fmt.Println(*str, "len = ", len(*str))
		fmt.Println("Некорректный ввод!", *str)
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Ваше число: ")
	}
	for i := 0; i < digits_in_number; i++ {

	}
}
