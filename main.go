package main

import (
	"fmt"

	Op "github.com/rizkysr90/kalkulator_cli/helper"
)

func inputNumber() []float32 {
	var arr []float32
	y := "y"
	ans := true
	index := 0
	var currentValue float32

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
		fmt.Println("Hasil Penjumlahan :", Op.SumCalc(getNumbers))
	} else if choosenMenu == 2 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Penjumlahan :", Op.Subtract(getNumbers))
	} else if choosenMenu == 3 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Pembagian :", Op.Divide(getNumbers))
	} else if choosenMenu == 4 {
		getNumbers := inputNumber()
		fmt.Println("Hasil Pembagian :", Op.Multiply(getNumbers))
	} else {
		fmt.Println("END")
	}

}
