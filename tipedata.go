package main

import (
	"fmt"

)

func main() {
	//int && float
	fmt.Println("satu =", 1)
	fmt.Println("dua =", 2)
	fmt.Println("tiga koma lima =", 3.5)

	//bool
	fmt.Println("benar =", true)
	fmt.Println("salah =", false)

	//string
	var x string = "Kawan SAYA"
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[2])

	//variabel
	var name string
	name = "Jhonny English"
	fmt.Println(name)

	var nama = "Jhonny English"
	fmt.Println(nama)

	nam := "Jhonny English"
	fmt.Println(nam)
	nam = "Jhonny China"
	fmt.Println(nam)

	var (
		firstName = "Jhonny"
		lastName  = "English"
	)
	fmt.Printf("nama = %s %s \n", firstName, lastName)

	//Constant
	const namaDepan = "Jhonny"
	const namaBelakang string = "English"

	fmt.Println("nama =", namaDepan, namaBelakang)

	var kata string = "Kawan SAYA"
	fmt.Println(kata[2])
	var w uint8 = kata[2]
	fmt.Println(w)
	var wString = string(w)
	fmt.Println(wString)

	//Declaration
	type NoHp string

	var noSaya NoHp = "083244"

	var noTeman string = "033842"
	var noteman NoHp = NoHp(noTeman)

	fmt.Println(noSaya)
	fmt.Println(noteman)

	//operasi mat
	var a = 20
	var b = 23
	var c = 21

	var e = a - b + c
	fmt.Println(e)

	var f = 21
	f += 5
	fmt.Println(f)

	//operasi bool
	var nilai = 90
	var absen = 12

	var nilaiAkhir bool = nilai > 90
	var absenAkhir bool = absen > 11

	var lulus bool = nilaiAkhir && absenAkhir
	fmt.Println(lulus)

	//Array
	var namateman = [...]string{
		"Jhonny",
		"English",
		"Saputra",
		"Christiany",
	}
	fmt.Println(namateman)
	fmt.Println(namateman[0])

	slice1 := namateman[1:3]

	fmt.Println(slice1)
	//Map
	minggu := map [string]string{
		"hari" : "minggu",
		"tanggal" : "12",
		// "tanggal" : "13",
	}
	fmt.Println(minggu["hari"])
	fmt.Println(minggu["tanggal"])

	//If Else Expresion
	namas := "Joko"

	if namas == "Jhonny"{
		fmt.Println("Hello Jhonny")
	}else if namas == "English"{
		fmt.Println("Hello",namas)

	}else{
		fmt.Println("Elu Sapa??")
	}

	if nilai > 90 && absen > 11   {
		fmt.Println("Kamu Lulus")
	}else{
		fmt.Println("Kamu Tetap Di Sini")
	}

	//Switch
	switch nilai > 80 && absen > 11{
	case true:
		fmt.Println("Kamu Lulus")
	case false:
		fmt.Println("Kamu Tetap Di Sini")

	}
	//For Loops
	for number := 1; number <= 10; number++{
		fmt.Println("Perulangan",number)
	}
	for i :=0; i < 10; i++{
		if i%2 == 0{
		continue
		}
		fmt.Println(i)
	}
}