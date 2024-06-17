package binarysearchtree

import (
	"crypto/rand"
	"math/big"
)

func SliceToTree(slice []int) *Node {
	root := &Node{Val: slice[0]}
	slice = slice[1:]
	for i := range slice {
		Insert(root, slice[i])
	}

	return root
}

func RandomSlice(quanitity int) []int {
	bigInt := big.NewInt(int64(quanitity))
	result := make([]int, quanitity)

	for i := 0; i < quanitity; i++ {
		number, _ := rand.Int(rand.Reader, bigInt)
		result[i] = int(number.Int64())
	}

	return result
}

func FindGreatestSlice(nums []int) int {
	var greatest int

	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}

	return greatest
}
