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
		user_number, winning_number string
	)
	generating_numbers(&winning_number)
	input(&user_number)
	fmt.Printf("\nИТОГ:   загаданное число - %v \n\tваше число - %v\n\n", winning_number, user_number)
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

func input(user_number *string) {
	for exit := 0; exit != 1; {
		fmt.Printf("\033c")
		exit = 0
		fmt.Print("Введите ", Digits_in_number, "-значное число(без повтора цифр): ")
		for fmt.Scan(user_number); len(*user_number) != Digits_in_number; fmt.Scan(user_number) {
			incorrect_input(user_number)
			fmt.Print("Введите ", Digits_in_number, "-значное число(без повтора цифр): ")
		}
		if i := fmt.Sprint(strings.ToUpper(*user_number)); i == "EXIT" {
			*user_number = i
			break
		}
		exit = 2
		for _, v := range *user_number {
			if v < 48 || v > 57 {
				exit = -1
				break
			}
		}
	Flag:
		switch exit {
		case -1:
			incorrect_input(user_number)
		case 2:
			for i := 1; i < len(*user_number); i++ {
				for j := 0; j < i; j++ {
					if (*user_number)[i] == (*user_number)[j] {
						exit = 0
						break
					}
				}
				if exit == 0 {
					incorrect_input(user_number)
					break Flag
				}
			}
			exit = 1
		default:
			fmt.Println("ERROR exit = ", exit)
		}
	}
}

func incorrect_input(user_number *string) {
	// fmt.Printf("\033c")
	loading()
	fmt.Println("Некорректный ввод! (" + *user_number + ")")
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
