/*

Print numbers from 1 to N without the help of loops.

Example 1:

Input:
N = 10
Output: 1 2 3 4 5 6 7 8 9 10

*/

package main

import ("fmt")

func PrintNumber(n int){
	if n > 0 {
		PrintNumber(n-1)
		fmt.Printf("%d",n)
	}
}

func main(){
	n := 10
	PrintNumber(n)
}