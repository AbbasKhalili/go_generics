package main

import (
	"fmt"
	"strings"
)

func main() {

	intDic := map[string]int64{
		"sunday":   100,
		"tuesday":  300,
		"thursday": 500,
	}

	floatDic := map[string]float64{
		"saturday":  12.5,
		"monday":    14.5,
		"wednesday": 16.5,
	}

	var intVal, intTitle = SumInts(&intDic)
	var floatVal, floatTitle = SumFloats(&floatDic)

	fmt.Println(intDic, "sum int :", intVal, intTitle)
	fmt.Println(floatDic, "sum float :", floatVal, floatTitle)

	var intByGen = SumIntsOrFloats(intDic)
	fmt.Println("sum int by generics :", intByGen)

	var floatByGen = SumIntsOrFloats(floatDic)
	fmt.Println("sum float by generics :", floatByGen)

	fmt.Println("multiple by generics :", multiple(8))
}

func SumIntsOrFloats[K comparable, V int64 | float64](dic map[K]V) V {
	var num V
	for _, v := range dic {
		num += v
	}
	return num
}

func multiple[K int | int8 | int16 | int32 | int64 | float32 | float64](val K) K {
	return val * val
}

func SumInts(dic *map[string]int64) (int64, string) {
	var num int64
	var title strings.Builder
	for key, val := range *dic {
		num += val
		title.WriteString(key + "|")
	}
	return num, title.String()
}

func SumFloats(dic *map[string]float64) (float64, string) {
	var num float64
	var title strings.Builder
	for key, val := range *dic {
		num += val
		title.WriteString(key + "|")
	}
	return num, title.String()
}
