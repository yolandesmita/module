package module

import "fmt"

// reutrn value
func SayHello() string {
	return "Selamat Datang!"
}

// multiple return value
func NamaPengunjung() (string, string, string) {
	return "Agus", "Bambang", "Joko"
}

// as a parameter
func Diskon(nama string, filter func(string) string) {
	fmt.Println("Pengunjung dengan nama", filter(nama), "akan mendapatkan diskon")
}

//interface struct

type OperasiHitung interface {
	Perkalian() int
	// Keliling() int
	// Luas() int
}

type Nilai struct {
	a, b, c int
}

func (n Nilai) Perkalian() int {
	return n.a * n.b * n.c
}

// func (in InputPersegi) Luas() int {
// 	return in.a * in.b
// }
