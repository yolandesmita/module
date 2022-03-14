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
	filterNama := filter(nama)
	fmt.Println("Pengunjung dengan nama", filterNama, "akan mendapatkan diskon")
}

func DiskonFilter(nama string) string {
	if nama == "Agus" {
		return "Selamat! Anda mendapatkan diskon 50%"
	} else {
		return "Anda mendapat diskon 5%"
	}
}
