package zero

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理
*/

// eg: "123"*"456"
// 转换成如下字符串加法
// "00738" + "06150" + "49200"
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	ss := make([]string, l2)
	for i, c := range num2 {
		bs := make([]byte, l1+l2)
		offset, carry := l2-1-i, 0
		for n := 0; n < l1+l2; n++ {
			idx := l1 - 1 - (n - offset)
			if n < offset || idx < 0 {
				bs[n] = byte(carry) + '0'
				carry = 0 // 进位清零
				continue
			}
			sum := int(c-'0')*int(num1[idx]-'0') + carry
			bs[n], carry = byte(sum%10)+'0', sum/10
		}
		// [002940 051600 837000]
		ss[i] = string(bs)
	}
	bs, carry := make([]byte, l1+l2), 0
	for i := 0; i < l1+l2; i++ {
		sum := carry
		for _, s := range ss {
			sum += int(s[i] - '0')
		}
		bs[l1+l2-1-i], carry = byte(sum%10)+'0', sum/10
	}
	if carry != 0 {
		bs[0] = byte(carry + '0')
	}
	if bs[0] == '0' {
		return string(bs[1:])
	}
	return string(bs)
}

/* 一段示例代码
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	b1, b2 := []byte(num1), []byte(num2)
	temp := make([]int, len(b1)+len(b2))
	for i := 0; i < len(b1); i++ {
		for j := 0; j < len(b2); j++ {
			temp[i+j+1] += int(b1[i]-'0') * int(b2[j]-'0')
		}
	}
    for i:=len(temp)-1;i>0;i--{
		temp[i-1]+=temp[i]/10
		temp[i]=temp[i]%10
	}

	if temp[0] == 0 {
		temp = temp[1:]
	}
	result := make([]byte, len(temp))
	for i, v := range temp {
		result[i] = byte(v + '0')
	}
	return string(result)

}
*/
