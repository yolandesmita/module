package module

// reutrn value
func SayHello() string {
	return "Selamat Datang!"
}

// multiple return value
func NamaLengkap() (string, string, string) {
	return "Agus", "Budiadi", "Hartono"
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
