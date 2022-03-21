package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	Digits_in_number = 4
	Sleeps_times     = 50
)

func main() {
	var (
		str, winning_number string
	)
	input(&str)
	generating_numbers(&winning_number)
	fmt.Printf("\nИТОГ:   загаданное число - %v \n\tваше число - %v\n\n", winning_number, str)
}

func generating_numbers(winning_number *string) {
	temp := int(time.Now().Unix()) * int(time.Now().Unix())
	exit := 0
	t := temp % 10
	*winning_number += fmt.Sprint(t)
	temp /= 10
	for i := 1; i < Digits_in_number; i++ {
		for exit != 1 {
			t = temp % 10
			exit = 1
			for _, v := range *winning_number {
				if int(v) == t+48 {
					exit = 0

					break
				}
			}
			temp /= 10
		}
		exit = 0
		*winning_number += fmt.Sprint(t)
		if temp == 0 {
			temp = int(time.Now().Unix()) * int(time.Now().Unix()) / 2
		}
	}
}

func input(str *string) {
	exit := 0
	for exit != 1 {
		fmt.Printf("\033c")
		exit = 0
		fmt.Print("Введите ", Digits_in_number, "-значное число(без повтора цифр): ")
		for fmt.Scan(str); len(*str) != Digits_in_number; fmt.Scan(str) {
			incorrect_input(str)
			fmt.Print("Введите ", Digits_in_number, "-значное число(без повтора цифр): ")
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
	// fmt.Printf("\033c")
	loading()
	fmt.Println("Некорректный ввод! (" + *str + ")")
	loading()
	fmt.Printf("\033c")
}

func loading() {
	fmt.Print("[")
	for i := 0; i < 10; i++ {
		time.Sleep(Sleeps_times * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(Sleeps_times * time.Millisecond)
	fmt.Println("]")
	time.Sleep(Sleeps_times * time.Millisecond)
}
