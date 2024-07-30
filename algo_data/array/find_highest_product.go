package array

// [5,-10,-6,9,4]
// Vou iterar por cada item no array
// Multiplicar o item anterior com o item atual
// Se a mutiplicação atual for maior que a multiplicação anterior
// Essa é a melhor multiplicação

// var melhorMulti = 60
// 5x-10 = -50
// -10*-6 = 60
// -6*9=-54
// 9*4=36

func FindHighestProduct(nums []int) int {
	highestProduct := 0

	before := nums[0]
	for i := 1; i < len(nums); i++ {
		current := nums[i]

		if (before * current) > highestProduct {
			highestProduct = (before * current)
		}
		before = current
	}

	return highestProduct
}
func FindHighestProduct_BookImplementation(nums []int) int {
	firstHighest, secondHighest := nums[0], nums[0]
	firstLowest, secondLowest := 0, 0

	for _, num := range nums {
		if num >= firstHighest {
			secondHighest = firstHighest
			firstHighest = num
		} else if num > secondHighest {
			secondHighest = num
		}

		if num <= firstLowest {
			secondLowest = firstLowest
			firstLowest = num
		} else if num < secondLowest {
			secondLowest = num
		}
	}

	highestProduct := firstHighest * secondHighest
	lowestProduct := firstLowest * secondLowest

	if highestProduct > lowestProduct {
		return highestProduct
	}
	return lowestProduct
}
