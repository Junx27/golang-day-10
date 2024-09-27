package buku

import "fmt"

// membuat struct buku
type Buku struct {
	Nama  string
	Harga int
}

// membuat method menampilkan buku
func (b Buku) ViewBook() {
	fmt.Printf("nama buku %s dengan harga %d\n", b.Nama, b.Harga)
}

// membuat method merubah harga buku
func (b *Buku) ChangePrice(p int) {
	b.Harga = p
	fmt.Printf("harga sekarang %d\n", b.Harga)
}
