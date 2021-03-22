package main

import (
	"fmt"
	"math"
)

func main() {
	// a := math.MaxInt32
	// b := math.MinInt32
	c := -10
	c1 := c + math.MaxUint32 + 1
	fmt.Println(c, c1)
}

package problem405

import (
	"math"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	// 负数转为补码
	if num < 0 {
		num = num + math.MaxUint32 + 1
	}
	hexList := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	result := ""
	for num > 0 {
		// 使用0xf (00...01111)取得后四位, 原数字右移动四位
		result = hexList[num&0xf] + result
		num >>= 4
	}
	return result
}

class Solution:
    def toHex(self, num: int) -> str:
        stra = ''
        if num < 0:
            num = (abs(num) ^ ((2 ** 32) - 1)) + 1
        elif num == 0:
            return '0'
        while (num >> 4) > 0 or num > 0:
            a = str(num % 16)
            if a == '10': a = 'a'
            elif a == '11': a = 'b'
            elif a == '12': a = 'c'
            elif a == '13': a = 'd'
            elif a == '14': a = 'e'
            elif a == '15': a = 'f'
            stra += a
            num = num >> 4
        return ''.join(reversed(stra))

作者：CoderSaru
链接：https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/solution/fen-hao-luo-ji-dui-fu-shu-0-zheng-shu-fen-bie-chu-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


作者：Takagi-san
链接：https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/solution/go-can-kao-jie-fa-by-caaalabash/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。