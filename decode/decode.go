package decode

import (
	"strconv"
)

// logic function
func logic(ch rune, num int) string {
	intCh, err := strconv.Atoi(string(ch))
	if err != nil {
		return "Z"
	}
	res := intCh + num
	if res >= 0 && res <= 2 {
		return strconv.Itoa(res)
	}
	return "Z"
}

// decode function
func Decode(ch string, l string) string {
	if ch == "" && !(l[len(l)-1] < '0' || l[len(l)-1] > '2') {
		return l
	}

	if len(l) == 0 {
		if ch[0] == 'L' {
			return min(
				Decode(ch[1:], "21"),
				Decode(ch[1:], "20"),
				Decode(ch[1:], "10"),
			)
		} else if ch[0] == 'R' {
			return min(
				Decode(ch[1:], "01"),
				Decode(ch[1:], "02"),
				Decode(ch[1:], "12"),
			)
		} else {
			return min(
				Decode(ch[1:], "00"),
				Decode(ch[1:], "11"),
				Decode(ch[1:], "22"),
			)
		}
	} else {
		if l[len(l)-1] < '0' || l[len(l)-1] > '2' {
			return "ZZZZZZZ"
		} else if ch[0] == 'L' {
			return min(
				Decode(ch[1:], l+logic(rune(l[len(l)-1]), -1)),
				Decode(ch[1:], l+logic(rune(l[len(l)-1]), -2)),
			)
		} else if ch[0] == 'R' {
			return min(
				Decode(ch[1:], l+logic(rune(l[len(l)-1]), 1)),
				Decode(ch[1:], l+logic(rune(l[len(l)-1]), 2)),
			)
		} else if ch[0] == '=' {
			return Decode(ch[1:], l+string(l[len(l)-1]))
		}
	}
	return "ZZZZZZZ"
}

// min function to find the lexicographically smallest string
func min(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	minStr := strs[0]
	for _, str := range strs[1:] {
		if str < minStr {
			minStr = str
		}
	}
	return minStr
}