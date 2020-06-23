package degree

import (
	"math"
)

func Calculate(floor, cup int, innerCalculate func(int, int) int) (degree int) {
	// 1.0 If eggs are enough then the binary tree is a non-hollow tree
	log2Floor := math.Log2(float64(floor))
	if float64(cup) > log2Floor {
		degree = int(log2Floor) + 1
		return
	}

	return innerCalculate(floor, cup)
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
func InnerCalculateB(floor, cup int) (degree int) { // this function relies on the caller function to filter out the log2(Floor)
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

func InnerCalculateC(floor, cup int) (degree int) {
	dp := make([][]int, cup+1)
	for i := 0; i < cup+1; i++ {
		dp[i] = make([]int, floor+1)
	}
	for i := 1; i < floor+1; i++ {
		dp[1][i] = i
	}
	for i := 1; i < cup+1; i++ {
		dp[i][1] = 1
	}

	for i := 2; i < cup+1; i++ {
		for j := 2; j < floor+1; j++ {
			res := math.MaxInt64
			for k := 1; k < j+1; k++ {
				tmp := max(dp[i-1][k-1], dp[i][j-k])
				res = min(tmp, res)
			}
			dp[i][j] = res + 1
		}
	}

	return dp[cup][floor]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
