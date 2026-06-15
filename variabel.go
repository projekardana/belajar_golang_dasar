package main

import "fmt"

// kode dibawah untuk mendeklarasikan variabel dengan tipe data string dan mengubah nilainya. Selain itu, kode ini juga menunjukkan cara menggunakan short variable declaration untuk mendeklarasikan variabel dengan tipe data yang berbeda, serta cara mendeklarasikan banyak variabel sekaligus menggunakan blok var.

func variabel() {
    // Deklarasi dengan var
    var nama string = "Golang"
    fmt.Println("Nama awal :", nama)

    // Mengubah nilai variable
    nama = "Go Programming Language"
    fmt.Println("Nama baru :", nama)

    // Short variable declaration
    umur := 10
    tinggi := 170.5
    aktif := true

    fmt.Println("Umur   :", umur)
    fmt.Println("Tinggi :", tinggi)
    fmt.Println("Aktif  :", aktif)

    // Deklarasi banyak variable
    var (
        kota    string = "Jakarta"
        provinsi string = "DKI Jakarta"
    )
    fmt.Println("Kota     :", kota)
    fmt.Println("Provinsi :", provinsi)
}

// Kode di atas mendeklarasikan sebuah variabel bernama `name` dengan tipe data `string`. Variabel ini kemudian
// package main

// import "fmt"

// func main() {
// 	var name string

// 	name = "John Doe"
// 	fmt.Println("Hello, my name is", name)

// 	name = "Jhon Smith"
// 	fmt.Println("Now, my name is", name)
// }