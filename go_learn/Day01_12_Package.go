package main

import (
	"package_test"
)

// package 导入有2中方法
// 1. GOPATH 是之前用的比较多的管理方式，只要把package的地址添加到GOPATH中即可找到该包
// 2. gomod 是现在比较推荐的 如果在一个file下有一个go.mod文件，则声明了这个file 是一个package
func main() {
	package_test.AddTwoNumbers(1, 2)

}
