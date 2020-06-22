package degree

import "math"

func Calculate(floor, cup int, innerDetect func(int, int) int) (degree int) {
	// 1.0 If eggs are enough then the binary tree is a non-hollow tree
	log2Floor := math.Log2(float64(floor))
	if float64(cup) >= log2Floor {
		degree = int(math.Ceil(log2Floor))
		return
	}

	return innerDetect(floor, cup)
}

func InnerCalculateA(floor, cup int) (degree int) {
	// 1.1 Prepare the init list (all 1) and the sum list (1,2,3...)
	list := make([]int, cup)
	for c := 0; c < cup; c++ {
		list[c] = 1
	}

	// 1.2 Calculate the list
	sum := 1
	for degree = 1; sum < floor; degree++ {
		calList := make([]int, cup)
		calList[0] = 1
		for c := 1; c < cup; c++ {
			calList[c] = list[c] + list[c-1]
		}
		list = calList
		sum += calList[cup-1]
	}
	return
}

// ∑n=1~degree(∑m=0~cup(C(n,m)))
func InnerCalculateB(floor, cup int) (degree int) {
	// 1.1 Initialize sum
	sum := int(math.Pow(2, float64(cup))) - 1

	// 1.2 Further calculate sum
	for degree = cup; sum < floor; degree++ {
		sum += sumCompose(degree, cup)
	}
	return
}

// ∑m=0~cup(C(n, m))
func sumCompose(n, cup int) int {
	sum := 0
	half := (n + 1) / 2
	if cup <= half {
		for i := 0; i < cup; i++ {
			sum += compose(n, i)
		}
	} else {
		floatHalf := float64(n+1) / 2
		mirror := n + 1 - cup
		for i := 0; float64(i) < floatHalf; i++ {
			if i < mirror || i == half {
				sum += compose(n, i)
			} else {
				sum += compose(n, i) * 2
			}
		}
	}
	return sum
}

// C(n, m)
func compose(n, m int) int {
	// FIXME: Comment out below 2 blocks because they are checked outside of the function
	//// filter out invalid values
	//if m < 0 || n < 0 || m > n {
	//	return 0
	//}
	//
	//// prepare
	//if m > n/2 {
	//	m = n - m
	//}

	// calculate
	result := 1
	j := n
	for i := m; i > 0; i-- {
		result *= j
		j--
	}
	for i := m; i > 1; i-- {
		result /= i
	}
	return result
}
