package module

import "fmt"

type Blacklist func(string) bool

func Pengunjung(resto string, blacklist Blacklist) {
	if blacklist(resto) {
		fmt.Println("Kamu telah mendapat diskon dari", resto)
	} else {
		fmt.Println("Kamu Punya Kesempatan Untuk Dapat Diskon")
	}
}

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
	filterNama := filter(nama)
	fmt.Println("Selamat", filterNama)
}

func DiskonFilter(nama string) string {
	if nama == "Agus" {
		return nama
	} else {
		return "Mohon Maaf"
	}

	blacklist := func(resto string) bool {
		return resto == "Boga"
	}

	Pengunjung("Boga", blacklist)

	Pengunjung("Agus", func(name string) bool {
		return name == "root"
	})
}
