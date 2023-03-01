// Solution for leetcode problem #2529 (Count max of positive and negative integers in a slice)
//
//  Link: https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer
//
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maximumCountNoBranches(sample))
	fmt.Println(maximumCountDumb(sample))
}

var sample = []int{-2, 0, -3, -1, 1, 2, 0}

func maximumCountNoBranches(nx []int) int {
	// count negatives and positives
	neg, pos := 0, 0
	for _, x := range nx {
		neg += int(0x1 & bits.RotateLeft(uint(x), 1))
		pos += int(0x1 & bits.RotateLeft(uint(-1*x), 1))
	}

	// figure out which is larger
	if neg < pos {
		return pos
	}
	return neg
}

func maximumCountDumb(nums []int) int {
	positive := 0
	negative := 0
	for _, number := range nums {
		if number > 0 {
			positive++
		}
		if number < 0 {
			negative++
		}
	}
	if positive < negative {
		return negative
	}
	return positive
}
