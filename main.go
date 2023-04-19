package main

import (
	"fmt"
	"strings"
	//"runtime/debug"
	"math"
	"math/big"
)

func main() {

	//Rounding()
	// Float_ex()
	// Rat_ex()
	//fmt.Println(ulp())
	//ulp()
	fmt.Println(fid())
	fmt.Println(dif())
	pog()

}

func tuchka() (big.Rat){
	var X_ici [31]big.Rat
	var x0 big.Rat = 4
	var x1 big.Rat = 4.25
	X_ici[0] = x0
	X_ici[1] = x1
	var f big.Rat
	y := x0
	z := x1
	for i:=0; i<30; i++{
	    f:= 
	    X_ici[i+2]:=f
	}
	return X_ici
}
func dif() float64 {
	var d float64 = 0
	for i := 1000000.00; i >= 1; i-- {
		d += 1 / i
	}
	return d
}
func fid() float64 {
	var d float64 = 0
	for i := 1.00; i < 1000000; i++ {
		d += 1 / i
	}
	return d
}

func pog() {
	fmt.Println((ulp() / 2 * 1000000) / (1 - ulp()/2*1000000))
}

// надо написать что-то окр(а+б) = (а+б)(1+f)// где |f|<=u // u= ulp/2
var a float32 = 0.1
var b float32 = 0.2
var c float32 = 0.3

var x float64 = 1.00
var y float64 = 2.00

func ulp() (fot float64) {

	var x2 string
	var x1 uint64 = math.Float64bits(x)
	x2 = fmt.Sprintf("%064b", x1)

	bin_x := x2[1:12]

	var E_x float64 = 0
	var s float64 = 1

	if bin_x == "00000000000" {
		var fot float64
		ULP := math.Pow(2, -1074)
		fot = ULP
		return fot

	} else if bin_x == "11111111111" && !strings.Contains(x2[12:], "1") {
		//1. если после 11 единиц идет 52 нуля, то возвращаяем бескончность
		return math.Inf(1)

	} else if bin_x == "11111111111" && strings.Contains(x2[12:], "1") {
		//2. если после 11 единиц идет хотя-бы 1 единица, то возвращаем NAN
		fot := math.NaN()
		return fot

	} else {
		for i := len(bin_x) - 1; i >= 0; i-- {
			if bin_x[i] == '1' {
				E_x += s
			}
			s *= 2
		}
		fot := math.Pow(2, E_x-1023-52)
		return fot

		//Sravnebie := new(big.Rat).SetFloat64((y - x) / fot)
		//fmt.Println(Sravnebie)
	}
}

func Rounding() {

	var f float64 = 1.00000005960464488641292746251565404236316680908203125
	var f1 float32 = 1.00000005960464488641292746251565404236316680908203125
	var f2 float32 = float32(f)
	fmt.Println(f1, f2)
	var x2 string
	var x1 uint64 = math.Float64bits(x)
	fmt.Println(x1)
	x2 = fmt.Sprintf("%032b", x1)
	fmt.Println(x2)
}

func Float_ex() {

	a1 := new(big.Float)
	b1 := new(big.Float)
	c1 := new(big.Float)
	st1 := fmt.Sprintf("%.27f\n", float64(a))
	st2 := fmt.Sprintf("%.26f\n", float64(b))
	st3 := fmt.Sprintf("%.24f\n", float64(c))

	_, err1 := fmt.Sscan(st1, a1)
	if err1 != nil {
		fmt.Println("Sscan: error_a1")
		return
	}

	_, err2 := fmt.Sscan(st2, b1)
	if err2 != nil {
		fmt.Println("Sscan: error_a2")
		return
	}

	_, err3 := fmt.Sscan(st3, c1)
	if err3 != nil {
		fmt.Println("Sscan: error_a3")
		return
	}

	sum := new(big.Float)
	sum.Add(a1, b1)
	fmt.Println(a1, "+", b1, "=", sum)
	if sum == c1 {
		fmt.Println(sum, "==", c1)
	} else {
		fmt.Println(sum, "!=", c1)
	}

}

func Rat_ex() {

	a1 := new(big.Rat)
	b1 := new(big.Rat)
	c1 := new(big.Rat)

	st1 := fmt.Sprintf("%.27f\n", float64(a))
	st2 := fmt.Sprintf("%.26f\n", float64(b))
	st3 := fmt.Sprintf("%.24f\n", float64(c))

	_, err1 := fmt.Sscan(st1, a1)
	if err1 != nil {
		fmt.Println("Sscan: error_a1")
		return
	}

	_, err2 := fmt.Sscan(st2, b1)
	if err2 != nil {
		fmt.Println("Sscan: error_a2")
		return
	}

	_, err3 := fmt.Sscan(st3, c1)
	if err3 != nil {
		fmt.Println("Sscan: error_a3")
		return
	}

	var sum = &big.Rat{}
	sum.Add(a1, b1)
	fmt.Println(a1, "+", b1, "=", sum)
	if sum == c1 {
		fmt.Println(sum, "==", c1)
	} else {
		fmt.Println(sum, "!=", c1)
	}

}
