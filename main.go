package main

func main() {

}

func SumInts(dic map[string]int64) int64 {
	var num int64
	for _, val := range dic {
		num += val
	}
	return num
}

func SumFloats(dic map[string]float64) float64 {
	var num float64
	for _, val := range dic {
		num += val
	}
	return num
}
