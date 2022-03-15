package module

import "fmt"

type Filter func(string) string

// type Blacklist func(string) bool

// type OperasiHitung interface {
// 	Perkalian() int
// }

// type Nilai struct {
// 	a, b, c int
// }

// func Pengunjung(resto string, blacklist Blacklist) {
// 	if blacklist(resto) {
// 		fmt.Println("Kamu telah mendapat diskon dari", resto)
// 	} else {
// 		fmt.Println("Kamu Punya Kesempatan Untuk Dapat Diskon")
// 	}
// }

// reutrn value
func SayHello() string {
	return "Selamat Datang"
}

// multiple return value
func NamaPengunjung() (string, string, string) {
	return "Agus", "Bambang", "Joko"
}

// // as a parameter
func DiskonMakan(nama string, filter Filter) {
	filterNama := filter(nama)
	fmt.Println("Selamat", filterNama)
}

func DiskonFilter(nama string) string {
	if nama == "Agus" {
		return nama
	} else {
		return "Mohon Maaf"
	}
}

// func (n Nilai) Perkalian() int {
// 	return n.a * n.b * n.c
// }

// // func main()
// // blacklist := func(resto string) bool {
// // 	return resto == "Boga"
// // }

// // Pengunjung("Boga", blacklist)

// // Pengunjung("Agus", func(name string) bool {
// // 	return name == "root"
// // })
