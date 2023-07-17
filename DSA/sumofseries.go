// Write a program to find the sum of the given series 1+2+3+ . . . . . .(N terms) 
/*
Input:
N = 5
Output: 15
Explanation: For n = 5, sum will be 15
+ 2 + 3 + 4 + 5 = 15.
*/

package main

import ("fmt")

func seriesSum(n int) int{
	sum := 0
	for i := 0; i <= n ; i++ {
		sum = sum + i
	}
	return sum
}

func main() {
	fmt.Println(seriesSum(5))
}