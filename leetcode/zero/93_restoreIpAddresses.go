/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/

package zero

func restoreIPAddresses(s string) []string {
	return rbacktrace(s, "", 1)
}

// 缺点，判断太多，影响了效率。可以考虑讲递归换成循环
func rbacktrace(s, ds string, depth int) []string {
	res := []string{}
	if depth == 5 {
		if len(s) != 0 {
			return res
		}
		return []string{ds[1:]}
	}
	if len(s) == 0 {
		return res
	}

	if s[0] == '0' || len(s) >= 1 {
		res = append(res, rbacktrace(s[1:], ds+"."+s[:1], depth+1)...)
	}
	if len(s) >= 2 && s[0] != '0' {
		res = append(res, rbacktrace(s[2:], ds+"."+s[:2], depth+1)...)
	}
	if len(s) >= 3 && s[:3] <= "255" && s[0] != '0' {
		res = append(res, rbacktrace(s[3:], ds+"."+s[:3], depth+1)...)
	}
	return res
}
