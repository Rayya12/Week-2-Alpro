package main

import "fmt"

func fibonacci(x int) int {
	var min, max, y, i, jumlah int

	// buat angka satu dan angka 2
	min = 0
	max = 1

	// buat nilai y untuk menyimpan fibonacci
	y = 0

	// buat variable jumlah untuk menjumlahkan
	jumlah = 0

	// buat tanpa sama dengan
	for i = 1; i < x; i++ {
		// buat fibonacci
		y = min + max

		// masukkan nilai max ke min
		min = max
		// masukkan fibonacci ke max
		max = y

		// kenapa yang dijumlah yang min? karena y dan max akan menjumlah 2 digit terakhir
		jumlah = min + jumlah
	}

	return jumlah
}

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println(fibonacci(x))
}
