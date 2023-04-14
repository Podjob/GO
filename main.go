package main

import (
	"fmt"
	"math"
)

var a float32 = 0.1
var b float32 = 0.2
var c float32 = 0.3

func f1() (a2 string, b2 string, c2 string) {

	var a1 uint32 = math.Float32bits(a)
	//fmt.Println(a1)
	a2 = fmt.Sprintf("%032b", a1)
	//fmt.Println(a2)

	var b1 uint32 = math.Float32bits(b)
	//fmt.Println(b1)
	b2 = fmt.Sprintf("%032b", b1)
	//fmt.Println(b2)

	var c1 uint32 = math.Float32bits(c)
	//fmt.Println(c1)
	c2 = fmt.Sprintf("%032b", c1)
	//fmt.Println(c2)

	return
}

func E() (E1 int64, E2 int64, E3 int64) {

	i, j, k := f1()
	bin_a := i[1:9]
	bin_b := j[1:9]
	bin_c := k[1:9]

	/*	fmt.Println(len(bin_a))
		fmt.Println(len(bin_b))
		fmt.Println(len(bin_c))

		fmt.Println(bin_a)
		fmt.Println(bin_b)
		fmt.Println(bin_c)
	*/
	var s1, s2, s3 int64 = 1, 1, 1

	for i := len(bin_a) - 1; i >= 0; i-- {
		if bin_a[i] == '1' {
			E1 += s1
		}
		s1 *= 2
	}
	//fmt.Println(E1)

	for i := len(bin_b) - 1; i >= 0; i-- {
		if bin_b[i] == '1' {
			E2 += s2
		}
		s2 *= 2
	}
	//fmt.Println(E2)

	for i := len(bin_c) - 1; i >= 0; i-- {
		if bin_c[i] == '1' {
			E3 += s3
		}
		s3 *= 2
	}
	//fmt.Println(E3)

	return
}

func M() (M1 int64, M2 int64, M3 int64) {

	i, j, k := f1()
	bin_a := i[9:]
	bin_b := j[9:]
	bin_c := k[9:]
	/*
		fmt.Println(len(bin_a))
		fmt.Println(len(bin_b))
		fmt.Println(len(bin_c))

		fmt.Println(bin_a)
		fmt.Println(bin_b)
		fmt.Println(bin_c)
	*/
	var s1, s2, s3 int64 = 1, 1, 1

	for i := len(bin_a) - 1; i >= 0; i-- {
		if bin_a[i] == '1' {
			M1 += s1
		}
		s1 *= 2
	}
	//fmt.Println(M1)

	for i := len(bin_b) - 1; i >= 0; i-- {
		if bin_b[i] == '1' {
			M2 += s2
		}
		s2 *= 2
	}
	//fmt.Println(M2)

	for i := len(bin_c) - 1; i >= 0; i-- {
		if bin_c[i] == '1' {
			M3 += s3
		}
		s3 *= 2
	}
	//fmt.Println(M3)

	return

}

func Schet() {
	i, j, k := E()

	e1 := float64(i)
	e2 := float64(j)
	e3 := float64(k)

	x, y, z := M()

	m1 := float64(x)
	m2 := float64(y)
	m3 := float64(z)

	var U1 float64 = math.Pow(-1, 0) * math.Pow(2, (e1-127)) * (1 + (m1 / math.Pow(2, 23)))
	//fmt.Println(U1)

	var U2 float64 = math.Pow(-1, 0) * math.Pow(2, (e2-127)) * (1 + (m2 / math.Pow(2, 23)))
	//fmt.Println(U2)

	var U3 float64 = math.Pow(-1, 0) * math.Pow(2, (e3-127)) * (1 + (m3 / math.Pow(2, 23)))
	//fmt.Println(U3)

	if (U1 + U2) == U3 {
		fmt.Println("U1 + U2 =", U1+U2, "=", U3, "\n Гыыыы")
	} else {
		fmt.Println("U1 + U2 =", U1+U2, "!=", U3, "\n Ну нет так нет")
	}

}

func main() {

	E()
	M()
	Schet()
	// 3 функции
	// 1 переводит число в бинорный вид
	// 2 описание числа в виде символов S E M
	// 3 посчитать пример в фотке
}
