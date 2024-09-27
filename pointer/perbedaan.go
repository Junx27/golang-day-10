package pointer

// menggunakan pointer
func MenggunakanPointer(r *string) {
	*r = "junx"
}

// tidak menggunakan pointer
func TidakMenggunakanPointer(r string) {
	r = "junx"
}
