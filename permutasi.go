package main

import "fmt"

func faktorial(n int) int {
	/*fungsi untuk menghitung n faktorial*/
	var hasil, i int
	hasil = 1

	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}

	return hasil
}

func permutasi(a, b int) int {
	/*fungsi untuk menghitung a permutasi b atau b permutasi a*/
	var hasil int
	if a > b {
		hasil = faktorial(a) / faktorial(a-b)
	} else {
		hasil = faktorial(b) / faktorial(b-a)
	}
	return hasil
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(faktorial(x), faktorial(y),permutasi(x, y))
}
