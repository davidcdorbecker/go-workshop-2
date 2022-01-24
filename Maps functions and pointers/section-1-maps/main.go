package main

import "fmt"

func main() {
	//var mexicanEquivalence map[string]float64
	mexicanEquivalence := make(map[string]float64)
	mexicanEquivalence = map[string]float64{
		"DOLAR":                 0.050212365768807,
		"EURO":                  0.042607,
		"PESO CHILENO":          39.341643582641,
		"PESO COLOMBIANO":       191.92342342342,
		"UNDEFINED EQUIVALENCE": 0,
	}

	fmt.Println(mexicanEquivalence)

	mexicanEquivalence["YEN"] = 5.5211869897629

	fmt.Println(mexicanEquivalence)
	fmt.Println(mexicanEquivalence["UNDEFINED EQUIVALENCE"])

	//mexicanEquivalence["YEN"] = 5.5211869897629
	//delete(mexicanEquivalence, "PESO CHILENO")
	//fmt.Println(mexicanEquivalence)
	//
	eq := mexicanEquivalence["?"]
	fmt.Println(eq)

	//copy := mexicanEquivalence
	//copy["DOLAR"] = 0.06
	//fmt.Println(copy, mexicanEquivalence)

	mapOfMaps := map[string]map[string]int {
		"a": {"b": 1},
	}

	fmt.Println(mapOfMaps)
}
