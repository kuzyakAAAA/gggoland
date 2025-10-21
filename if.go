package main

import "fmt"

func Function2() {
	var num int
	fmt.Println("Введи")
	fmt.Scan(&num)
	var fl bool
	for num < 12307 {
		if num%9==0 && num%13==0 {
			fl=false
			break
		} else if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num&9 == 0 {
			num = num*13 + 1
		} else {
			num = (num + 2) * 3
		}
	}
	if !fl{
			fmt.Println("server error")
	} else {
			fmt.Println(num)
	}
}
