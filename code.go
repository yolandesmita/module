package module

import "fmt"

type Filter func(string) string

type Diskon func(string) bool

type OperasiHitung interface {
	Perkalian() int
}

type Customer struct {
	penerima  string
	usia      int
	pekerjaan string
}

type Nilai struct {
	a, b, c int
}

// return value
func SayHello() string {
	return "Selamat Datang"
}

// multiple return value
func NamaPengunjung() (string, string, string) {
	return "Agus", "Bambang", "Joko"
}

// as a parameter
func DiskonMakan(nama string, filter Filter) {
	filterNama := filter(nama)
	fmt.Println("Selamat", filterNama)
}

func DiskonFilter(nama string) string {
	if nama == "Agus" {
		return "Selamat! Anda Mendapatkan Diskon 50%"
	} else {
		return "Mohon Maaf, Silahkan Coba Lagi di Kesempatan Lain"
	}
}

// anonymous function
func DiscountUsed(name string, diskon Diskon) {
	if diskon(name) {
		fmt.Println("Diskon ini telah digunakan oleh", name)
	} else {
		fmt.Println("Diskon ini belum digunakan oleh", name)
	}
}

// struct and interface

func (Hasil Nilai) Perkalian() int {
	return Hasil.a * Hasil.b * Hasil.c
}
