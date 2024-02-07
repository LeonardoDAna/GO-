// 参数任意数量 类型唯一的函数

package main

import (
	"fmt"
)

func test(args ...int) int {
	fmt.Println(args)
	count := 0
	for i := 0; i < len(args); i++ {
		count += args[i]
	}
	return count
}

// 传入指针 才能得到内存地址 然后修改内存中指定的数值
func countAdd(num *int) {
	*num += 1
}

func main() {
	count := test(1, 2, 3, 4, 5)
	fmt.Println(count)
	countAdd(&count)
	fmt.Println(count)

}
