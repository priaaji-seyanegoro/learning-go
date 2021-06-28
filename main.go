package main

import "fmt"

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Persegi struct {
	Sisi int
}

// group types together based on method
type Luas interface {
	CariLuas() int
}

func (num PersegiPanjang) CariLuas() int {
	luas := num.Panjang * num.Lebar
	return luas
}

func (num Persegi) CariLuas() int {
	luas := num.Sisi * num.Sisi
	return luas
}

func PrintHasil(name string, luas Luas) {
	fmt.Printf("Luas %v : %v \n", name, luas.CariLuas())
}

func main() {
	luasPersegi := Persegi{
		Sisi: 10,
	}

	luasPersegiPanjang := PersegiPanjang{
		Panjang: 10,
		Lebar:   5,
	}

	PrintHasil("Persegi", luasPersegi)
	PrintHasil("Persegi Panjang", luasPersegiPanjang)
}
