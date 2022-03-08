package main

import "fmt"

func main() {
	// 利用或操作+空格将英文字符转换为小写
	lowerCase := string('A' | ' ')
	fmt.Println(lowerCase)

	// 利用与操作+下划线将英文字符转换为大写
	upperCase := string('a' & '_')
	fmt.Println(upperCase)

	// 利用异或操作+空格进行大小写呼唤
	lowerCase1 := string('d' ^ ' ')
	upperCase1 := string('D' ^ ' ')
	fmt.Println(lowerCase1, upperCase1)

	// 判断两个数是否异号
	var x, y int
	x, y = -1, 2
	fmt.Println(x ^ y)
	x, y = 3, 2
	fmt.Println(x ^ y)

	// 按位与消除最后一个1
	fmt.Println(9 & 8)
}
