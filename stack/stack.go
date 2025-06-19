package stack

// https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {
	// ([])
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	brakets := make([]byte, 0, len(s))
	for _, c := range []byte(s) {
		if len(brakets) > 0 && mp[c] == brakets[len(brakets)-1] {
			brakets = brakets[:len(brakets)-1]
		} else {
			brakets = append(brakets, c)
		}
	}
	return len(brakets) == 0
}

// https://leetcode.cn/problems/basic-calculator-ii/
// multiplication to addition
func calculate(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0') //continuous number
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

// https://leetcode.cn/problems/next-greater-element-i/description/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// nums1 = [4,1,2], nums2 = [1,3,4,2].

	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums2); j++ {

		}
	}
}
