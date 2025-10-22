package main

import "fmt"

func Function2() {
	var num1 int
	fmt.Println("Введи")
	fmt.Scan(&num1)

	var fl bool = true
	var num = num1

	if num1 <= 12307 {
		for num < 12307 {
			if num%9 == 0 && num%13 == 0 && num != 0 {
				fl = false
				break
			} else if num < 0 {
				num *= -1
			} else if num%7 == 0 && num != 0 {
				num *= 39
			} else if num%9 == 0 && num != 0 {
				num = num*13 + 1
			} else {
				num = (num + 2) * 3
			}
		}

		if !fl {
			fmt.Println("server error")
		} else {
			fmt.Println(num)
		}

	} else {
		fmt.Println("Число", num1, "уже больше числа 12307")
	}
}
