package main

import "fmt"

/**
http://wiki.jikexueyuan.com/project/kmp-algorithm/define.html
*/
func kmp(s, p string) int {
	i, j := 0, 0
	pLen := len(p)
	sLen := len(s)
	next := getNext(p)
	fmt.Println(next)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] { //s[i]!=s[0]=>j=next[0]=-1,第0位不匹配所以i++，j++;j=0
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}

/**
next [j] = k，代表 j 之前的字符串中有最大长度为 k 的相同前缀后缀。
此也意味着在某个字符失配时，该字符对应的 next 值会告诉你下一步匹配中，模式串应该跳到哪个位置（跳到next [j] 的位置）。如果 next [j] 等于 0 或 -1，则跳到模式串的开头字符，若 next [j] = k 且 k > 0，代表下次匹配跳到 j 之前的某个字符，而不是跳到开头，且具体跳过了 k 个字符。
*/
func getNext(p string) []int {
	pLen := len(p)
	next := make([]int, pLen, pLen)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < pLen-1 { //因为next[pLen-1]由s[i] == s[pLen-2]算出
		if i == -1 || p[i] == p[j] { //-1代表了起始位不匹配，i=0,s[0]!=s[j]=>i=next[0]=-1
			i++
			j++
			next[j] = i
		} else {
			i = next[i]
		}
	}
	return next
}

func main() {
	ret := kmp("abbaabbbabba", "babba")
	fmt.Println(ret)
}
