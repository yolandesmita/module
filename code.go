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
	fmt.Println("Selamat", filterNama)
}

func DiskonFilter(nama string) string {
	if nama == "Agus" {
		return nama
	} else {
		return "Mohon Maaf"
	}
}
