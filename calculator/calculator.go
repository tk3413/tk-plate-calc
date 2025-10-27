package calculator

import (
	sort "sort"

	gen_code "github.com/tk3413/tk-weight-calc/server_gen"
)

func intPtr(i int) *int {
	return &i
}

func CalculateWeights(weight float64) gen_code.PlateCount {
	weightRemaining := weight - 45.0 // Subtract bar weight
	plates := map[float64]int{
		45:  0,
		35:  0,
		25:  0,
		10:  0,
		5:   0,
		2.5: 0,
	}

	plateSizes := make([]float64, 0, len(plates))
	for p := range plates {
		plateSizes = append(plateSizes, p)
	}
	// sort descending so we try largest plates first
	sort.Slice(plateSizes, func(i, j int) bool { return plateSizes[i] > plateSizes[j] })

	for _, plate := range plateSizes {
		count := int(weightRemaining / (plate * 2))
		if count > 0 {
			plates[plate] = count
			weightRemaining -= float64(count) * plate * 2
		}
	}
	return gen_code.PlateCount{
		N45:   intPtr(plates[45]),
		N25:   intPtr(plates[25]),
		N10:   intPtr(plates[10]),
		N5:    intPtr(plates[5]),
		N2pt5: intPtr(plates[2.5]),
	}
}
