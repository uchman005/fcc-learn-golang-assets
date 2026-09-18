package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
func lengthOfLongestSubstring(s string) int {
	arrays := breakIntoArrays(s)
	greatestCount := 0
	for _, arr := range arrays {
		str := convertArrToStr(arr)
		fmt.Printf("str: %d %s\n", greatestCount, str)
		if greatestCount > len(str) {
			fmt.Printf("greatestCount: %v\n", greatestCount)
			return greatestCount
		} else {
			newCount := checkCount(str, greatestCount)
			if newCount > greatestCount {
				greatestCount = newCount
			}
		}
	}
	return greatestCount
}

func checkSubStrExits(str string, sub rune, count int) bool {
	result := false
	for _, e := range str {
		if count > len(str) {
			result = true
			break
		}
		if e == sub {
			result = true
			break
		}
	}
	return result
}
func breakIntoArrays(str string) [][]string {
	if len(str) > 100 {
		str = str[0:100]
	}
	arrays := make([][]string, len(str))
	for i := range arrays {
		currentStr := str[i:]
		currentArray := make([]string, len(currentStr))
		for j, elem := range currentStr {
			currentArray[j] = string(elem)
		}
		arrays[i] = currentArray
	}
	return arrays
}

func convertArrToStr(arr []string) string {
	result := ""
	for _, item := range arr {
		result += item
	}
	return result
}

func checkCount(s string, currentCount int) int {
	result1 := string(s[0])
	if currentCount >= len(s) {
		return currentCount
	}
	count := 1
	for i := 1; i < len(s); i++ {
		nextStr := rune(s[i])
		if checkSubStrExits(result1, nextStr, count) {
			result1 = string(nextStr)
		} else {
			result1 += string(nextStr)
			if len(result1) > count {
				count = len(result1)
			}
		}
	}
	return count
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	// test([]emailBill{
	// 	{45},
	// 	{32},
	// 	{43},
	// 	{12},
	// 	{34},
	// 	{54},
	// })

	// test([]emailBill{
	// 	{12},
	// 	{12},
	// 	{976},
	// 	{12},
	// 	{543},
	// })
	fmt.Printf("lengthOfLongestSubstring(\"large\"): %v\n", lengthOfLongestSubstring("bcccabbbaaacbabaaabccbaabacaabbaacaaaccaccaabbacbacbcabcacacaacbabcbccabaaacbbbaaacabaabcbbbccbccabccbacbaaacbcbbbcaaccaabcaaccbaaccabbbccabcbbcabbcababbcaccccaaccbabbccaabacacaaabbaaaabbbbbcacabcbaaaccbaccbcaabcccaabbccacabbbccbaabbbbcabbaabacababbbbbbcbaabaabaababcaaacbbbbaaacbacacbcabbccabcbbcaabbaacbccabcacaabbcaaccccbccaabbcbcbacbbccbababaacbccccccabbaaabcbbbbcacabccabacaaaaacccbbcbcaaccaabaabbaabbccacacabbbaccbbbcabbabacbcbbabbbaababaabbcaaccabbcccbabccacbabbbcbbaccccbaccccbccabbcabbbabaacacccbbbaccbbcccbccccaabbaababcaccabacbbccabcacaabacccacaacbaaccccaaabbbcabacbcaacbbbbbbabccbbbaabbcbcacbbacaacabcbbaacaababaabbbcaaaaaaaacacbbcbbbccaacbcbcbaacbcccccaaccaccacabbbacbcbabacbacbbcabbaccbaabcbbcbcccaaccbcaacacaabaabbbbaabcaabbabbbcaaacbcbacbbbccabcbbcbcacbaababaccaabccbbabaccbaacacbcbbabbbbbabbbbacaabbacaaabcbbabbcbcccaacbbabcbacbbbbccbbaccacccabbcbbccccbabcbaccbbacabcbacababacccccaababbcacaccbbbbcacbbccbbcbbccabaaaaaacaaccbbabbcbaabcccbcccccaabbbbabbbacbacabcbbabcacbbcbaacbbaabacbacccababbaacabbbbbcaaabbbbbaaabcacabcabaabcbabcbccabbbabbbbabbcbabbcaacbcabacaaaccabbabbaccccaabbacbaccbaaaacabbccaabbcaccbabcabcccacbcabaaccabaabcaaaccacbcabcbbcbccbabbcaabcaccbacbabccabcbcbabbcacbacbcabbcccccbacbcaacbabbabcaacabaabcbaabbabbcbabbccbabbaacacaaccaabcccbaaabbaacabccbbaabaaabbaaabacacabbcbbaccacbaaabccbbbcacabcaaaaabcbabbaaacaaccccbaccccccbcccbbabacaabbbabaabbccaaabbaaaaccaabbbcbcccaacaaaaacbbcbbaaaaaacbbbabacbbabcbabaaacbccbaabccccccbabbacaaabbacaaccbaccbcaccbaccbcabcababbccbaaabcbcbabbbcbbbccacacabccabccaabbcbcaccabbbaacccbcacaaaabcbbbbbacaabbbbbacaabaaaccbaaccbbbacaabbbccbcbcccbcbaabaabcaaaacbacbbabbbcccabcaccaacbacbcbbbccbcbabbbcbbcbbbbbbccaaccbcababbaaaabbabcbbcbaacaccbcbcb"))
	fmt.Printf("lengthOfLongestSubstring(\"1R1T7\"): %v\n", lengthOfLongestSubstring("1R1T7"))
	fmt.Printf("lengthOfLongestSubstring(\"pwwkew\"): %v\n", lengthOfLongestSubstring("pwwkew"))
	test([]emailBill{
		{743},
		{13},
		{8},
	})
}
