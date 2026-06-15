package main

import "fmt"

func main() {
    var nama string = "Budi Santoso"
    var pesan string = "Belajar Golang"

    fmt.Println("Nama   :", nama)
    fmt.Println("Pesan  :", pesan)

    fmt.Println("Panjang nama   :", len(nama))
    fmt.Println("Karakter ke-0  :", string(nama[0]))
    fmt.Println("Karakter ke-1  :", string(nama[1]))

    // String kosong
    var kosong string = ""
    fmt.Println("String kosong, panjang:", len(kosong))

    // String multiline menggunakan backtick
    multiline := `Baris pertama
Baris kedua
Baris ketiga`
    fmt.Println(multiline)
}