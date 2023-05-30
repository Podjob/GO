package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {

	scalarProduct()

}

func scalarProduct() {
	f, err := os.Open("368345.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var v0 []float64 // массив из файла
	var n int64      // число элементов в v0
	for {
		var flt float64
		err := binary.Read(f, binary.LittleEndian, &flt)
		if err == io.EOF {
			break
		}
		v0 = append(v0, flt)
		n++
	}
	var Fv1 []float64 // первый вектор
	var Fv2 []float64 // второй вектор
	// проверка числа элем. массива на четность на четность
	if len(v0)%2 != 0 {
		return
	}
	for i := 0; i < len(v0)/2; i++ {
		Fv1 = append(Fv1, v0[i])
		Fv2 = append(Fv2, v0[len(v0)/2+i])
	}
	//1________//////////////////////////////////////////////////////////////
	fmt.Println("                        1й способ сумм-ия") //(((massiv[0] + massiv[1]) + massiv[2]) + massiv[3])
	var scalProdFl1 float64 = 0                              // 1ое скал произв
	var n1 int = len(Fv1)*len(Fv1) - 1
	for i := 0; i < len(Fv1); i++ {
		scalProdFl1 = scalProdFl1 + Fv1[i]*Fv2[i]
	}
	fmt.Println("Скаляр призвед 1:", scalProdFl1)
	fmt.Println("Число операций:", n1)
	//2________//////////////////////////////////////////////////////////////
	fmt.Println("                        2й способ сумм-ия") //(massiv[0] + massiv[1]) + (massiv[2] + massiv[3])
	var scalProdFl2 float64 = 0                              // 2ое скал произв
	var n2 int
	// проверка числа опреаций для 2ого сопособа
	// log(n)?????????????
	if len(Fv1) == 2 {
		n2 = 1
	} else if len(Fv1) == 1 {
		n2 = 0
	} else {
		n2 = len(Fv1) - 1
	}
	// проверка на четность числа элементов массива
	var ProdFv []float64

	if len(Fv1)%2 == 0 {
		for i := 0; i < len(Fv1); i++ {
			ProdFv = append(ProdFv, Fv1[i]*Fv2[i])
		}
	} else {
		Fv1_copy := Fv1                // копия массива Fv1
		Fv1_copy = append(Fv1_copy, 0) // добавляем в конец массива четный элемент, равный нулю
		Fv2_copy := Fv2                // копия массива Fv2
		Fv2_copy = append(Fv2_copy, 0) // добавляем в конец массива четный элемент, равный нулю
		for i := 0; i < len(Fv1_copy); i++ {
			ProdFv = append(ProdFv, Fv1_copy[i]*Fv2_copy[i])
		}
	}
	var vecZero []float64
	for {
		if len(ProdFv) != 1 {
			vec := vecZero
			for i := 0; i < len(ProdFv); i = i + 2 {
				a := ProdFv[i] + ProdFv[i+1]
				vec = append(vec, a)
			}
			ProdFv = vec
		} else {
			scalProdFl2 = ProdFv[0]
			break
		}
		if len(ProdFv)%2 != 0 && len(ProdFv) != 1 {
			ProdFv = append(ProdFv, 0)
		}
	}

	fmt.Println("Скаляр призвед 2:", scalProdFl2)
	fmt.Println("Число операций:", n2)
	//3________//////////////////////////////////////////////////////////////
	fmt.Println("                        3й способ сумм-ия") //суммирование вставками
	var scalProdFl3 float64 = 0                              // 3е скал произв
	var n3 int = len(Fv1) - 1
	var masFv []float64
	for i := 0; i < len(Fv1); i++ {
		a := Fv1[i] * Fv2[i]
		masFv = append(masFv, a)
	}
	// функиция для сортировки массива (элементы берутся по модулю)
	sort.SliceStable(masFv, func(i, j int) bool {
		return math.Abs(masFv[i]) < math.Abs(masFv[j])
	})
	for i := len(Fv1) - 1; i >= 0; i-- {
		if i != 0 {
			scalProdFl3 = masFv[i] + masFv[i-1]
			masFv = masFv[:len(masFv)-2]       // удаляем последние 2 элеметна массива
			masFv = append(masFv, scalProdFl3) // добавляем в конец сумму 2х удаленных элементов
			sort.SliceStable(masFv, func(i, j int) bool {
				return math.Abs(masFv[i]) < math.Abs(masFv[j])
			})
		} else if i == 0 {
			scalProdFl3 = masFv[i]
		}
	}
	fmt.Println("Скаляр призвед 3:", scalProdFl3)
	fmt.Println("Число операций:", n3)
	////////////////////////////////////////////////////////////////
	y := ulp(1.0) / 2
	var N1 float64 = float64(n1)
	var N2 float64 = float64(n2)
	var N3 float64 = float64(n3)
	Gn1 := (N1 * y) / (1 - N1*y)
	Gn2 := (N2 * y) / (1 - N2*y)
	Gn3 := (N3 * y) / (1 - N3*y)

	var sumAbs float64 = 0.0
	var scalProdSum float64 = 0.0

	for i := 1; i < len(Fv1); i++ {
		sumAbs = sumAbs + Fv1[i]*Fv2[i]
		scalProdSum = scalProdSum + math.Abs(Fv1[i]*Fv2[i])
	}
	//scalProdSum1 := scalProdSum * Gn1
	//scalProdSum2 := scalProdSum * Gn2
	//scalProdSum3 := scalProdSum * Gn3

	// бегущая ошибка
	fmt.Println()

	fmt.Println(scalProdSum * Gn1)
	fmt.Println(scalProdSum * Gn2)
	fmt.Println(scalProdSum * Gn3)

	// 3 априорные ошибки
	scalProdApr1 := scalProdFl1 * Gn1
	scalProdApr2 := scalProdFl2 * Gn2
	scalProdApr3 := scalProdFl3 * Gn3
	fmt.Println()

	fmt.Println(Gn1)
	fmt.Println(Gn2)
	fmt.Println(Gn3)

	fmt.Println()

	fmt.Println(N1)
	fmt.Println(N2)
	fmt.Println(N3)

	fmt.Println()

	fmt.Println(scalProdApr1)
	fmt.Println(scalProdApr2)
	fmt.Println(scalProdApr3)

	fmt.Println()
	fmt.Println()
}

///////////////////////////////////////////////////////////////////////////

func ulp(x float64) (fot float64) {
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
	}
}
