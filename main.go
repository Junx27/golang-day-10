package main

import (
	"fmt"
	"golang-day-10/buku"
	"golang-day-10/diskon"
	"golang-day-10/pointer"
	"golang-day-10/salary"
)

type Buku struct {
	Nama  string
	Harga int
}

func (b Buku) viewBook() {
	fmt.Printf("nama buku %s dengan harga %d\n", b.Nama, b.Harga)
}
func (b *Buku) changePrice(p int) {
	b.Harga = p
	fmt.Printf("harga sekarang %d\n", b.Harga)
}

func main() {

	// soal nomor 1
	nama1 := "Tri Saptono"
	nama2 := "Tri Saptono"
	pointer.MenggunakanPointer(&nama1)
	fmt.Println(nama1)
	pointer.TidakMenggunakanPointer(nama2)
	fmt.Println(nama2)

	//soal nomor 2
	//memanggil fungsi
	diskon.GetDiscount()

	//soal nomor 3
	fmt.Println("tidak menggunakan package")
	//menggunakan struct
	buku1 := Buku{
		Nama:  "biologi",
		Harga: 20000,
	}
	//memanggil method
	buku1.viewBook()
	buku1.changePrice(30000)
	buku1.viewBook()

	// soal nomor 4
	fmt.Println("menggunakan package")
	//menggunakan struct
	buku2 := buku.Buku{
		Nama:  "biologi",
		Harga: 20000,
	}
	//memanggil method
	buku2.ViewBook()
	buku2.ChangePrice(30000)
	buku2.ViewBook()

	//soal nomor 5
	//menggunakan struct
	anton := salary.FullTimeEmployee{
		Name:          "anton",
		MonthlySalary: 4000000,
	}
	andi := salary.PartTimeEmployee{
		Name:        "andi",
		HourlyRate:  80000,
		HoursWorked: 20,
	}
	//menggunakan interface
	salary.ViewSalary(anton)
	salary.ViewSalary(andi)

}
