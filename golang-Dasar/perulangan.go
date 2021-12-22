package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Angka", i)
	// }
	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// }

	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "ddate"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"
	// names[4] = "ez" // baris kode ini menghasilkan error
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)
	// var fruits [4]string

	// // cara horizontal
	// fruits = [4]string{"apple", "grape", "banana", "melon"}

	// // cara vertikal
	// fruits = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }
	// var numbers = [...]int{2, 3, 2, 4, 3}

	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))

	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) // [apple manggo]

}
