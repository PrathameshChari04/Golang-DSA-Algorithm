/*

Time to Words                                                                                                                                                                                             Given a time in the format of hh:mm (12-hour format) 0 < HH < 12, 0 <= MM < 60. The task is to convert it into words.
Note: There are different corner cases appears for simplicity refer this example:

6:00 six o'clock
6:10 ten minutes past six
6:15 quarter past six
6:30 half past six
6:45 quarter to seven
6:47 thirteen minutes to seven

Example 1:

Input: H = 6, M = 0
Output: six o' clock
Explanation: 6H:0M = six o' clock.

*/

package main

import "fmt"

var hours = [...]string{
	"twelve", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten", "eleven",
}

var minutes = [...]string{
	"o' clock", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten", "eleven",
	"twelve", "thirteen", "fourteen", "quarter", "sixteen",
	"seventeen", "eighteen", "nineteen", "twenty", "twenty one",
	"twenty two", "twenty three", "twenty four", "twenty five",
	"twenty six", "twenty seven", "twenty eight", "twenty nine",
	"half",
}

func convertTimeToWords(H, M int) string {
	if M == 0 {
		return hours[H] + " " + minutes[M]
	} else if M == 1 {
		return minutes[M] + " minute past " + hours[H]
	} else if M == 59 {
		return minutes[60-M] + " minute to " + hours[(H+1)%12]
	} else if M == 15 || M == 30 {
		return minutes[M] + " past " + hours[H]
	} else if M == 45 {
		return minutes[60-M] + " to " + hours[(H+1)%12]
	} else if M <= 30 {
		return minutes[M] + " minutes past " + hours[H]
	} else {
		return minutes[60-M] + " minutes to " + hours[(H+1)%12]
	}
}

func main() {
	H := 0
	M := 50
	timeInWords := convertTimeToWords(H, M)
	fmt.Println(timeInWords)
}

