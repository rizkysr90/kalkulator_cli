package main

import (
	"fmt"
)

func inputNumber() []int {
	var arr []int
	y := "y"
	ans := true
	index := 0
	var currentValue int

	for ans {
		fmt.Print("Masukkan angka ke ", index+1, " : ")
		_, err := fmt.Scanln(&currentValue)
		if err != nil {
			fmt.Println("Hanya boleh memasukkan bilangan")
			break

		}
		arr = append(arr, currentValue)

		fmt.Print("Ingin memasukkan angka lagi ? Y / N --> ")
		fmt.Scanln(&y)
		if y != "y" && y != "Y" {
			ans = false
		}
		index++
	}

	return arr
}

func sumCalc(arr []int) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}
func multiply(arr []int) int {
	total := arr[0]
	for i := 1; i < len(arr); i++ {
		total *= arr[i]
	}
	return total
}
func divide(arr []int) int {
	total := arr[0]
	for i := 1; i < len(arr); i++ {
		total /= arr[i]
	}
	return total
}
func subtract(arr []int) int {
	total := arr[0]
	for i := 1; i < len(arr); i++ {
		total -= arr[i]
	}
	return total
}
func main() {
	var choosenMenu int8
	fmt.Println("Kalkulator V.0.1")
	fmt.Println("=================")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")
	fmt.Println("=================")
	fmt.Println("Pilih menu dari angka diatas : ")
	fmt.Scanln(&choosenMenu)

	if choosenMenu == 1 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Penjumlahan :", sumCalc(getNumbers))
	} else if choosenMenu == 2 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Penjumlahan :", subtract(getNumbers))
	} else if choosenMenu == 3 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Pembagian :", divide(getNumbers))
	} else if choosenMenu == 4 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Pembagian :", multiply(getNumbers))
	} else {
		fmt.Println("END")
	}

}
