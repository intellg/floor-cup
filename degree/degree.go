package degree

import "math"

func Calculate(floor, cup int, innerCalculate func(int, int) int) (degree int) {
	// 1.0 If eggs are enough then the binary tree is a non-hollow tree
	log2Floor := math.Log2(float64(floor))
	if float64(cup) > log2Floor {
		degree = int(log2Floor) + 1
		return
	}

	return innerCalculate(floor, cup)
}
