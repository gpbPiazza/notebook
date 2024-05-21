package array

// [20,30,40,50]
// 1. 20 + addUntil100([30,40,50]) = 140
// 2. 30 + addUntil100([40,50]) = 120
// 2. 40 + addUntil100([50]) = 90
// 2. 50 + addUntil100([]) = 50

// [99,10,20,55,10]
// 1. 99 + addUntil100([10,20,55,10]) = 194
// 2. 10 + addUntil100([20,55,10]) = 95
// 3. 20 + addUntil100([55,10]) = 85
// 4. 55 + addUntil100([10]) = 65
// 5. 10 + addUntil100([]) = 10

// O(2ˆn)
func addUntil100(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if numbers[0]+addUntil100(numbers[1:]) > 100 {
		return addUntil100(numbers[1:])
	} else {
		return numbers[0] + addUntil100(numbers[1:])
	}
}

func addUntil100_On_implementation(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	sumOfRemaingNums := addUntil100(numbers[1:])

	if numbers[0]+sumOfRemaingNums > 100 {
		return sumOfRemaingNums
	} else {
		return numbers[0] + sumOfRemaingNums
	}
}
