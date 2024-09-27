package diskon

import (
	"fmt"
	"strconv"
)

func GetDiscount() {
	//membuat variabel
	var totalPrice, finalPrice, discount int

	//membuat perulangan input user
	for {
		fmt.Println("masukan total belanja")
		var price string
		fmt.Scan(&price)
		number, err := strconv.Atoi(price)
		if err != nil {
			fmt.Println("masukan angka salah")
			continue
		} else {
			totalPrice = number
		}
		break
	}

	//membuat validasi diskon
	if totalPrice > 1000000 {
		discount = totalPrice * 20 / 100
		finalPrice = totalPrice - discount
	} else {
		discount = 0
		finalPrice = totalPrice
	}
	//menampilkan nilai variabel
	fmt.Printf("total belanja adalah %d jumlah diskon %d jumlah yang harus dibayar %d\n", totalPrice, discount, finalPrice)
}
