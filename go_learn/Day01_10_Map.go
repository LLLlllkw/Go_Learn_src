package main

import "fmt"

func main() {
	// Map 就是字典
	// 1. 新建Map map[KeyDataType] ValDataType
	// map 类型是引用类型，在声明后，如果没有初始化，是没有给其分配空间的。
	// 在go中，基本类型在变量声明时就会有默认初始值赋值，而引用对象是没有的，相当于指向了nil
	//		This variable m is a map of string keys to int values:
	//		var m map[string]int
	//		Map types are reference types, like pointers or slices, and so the value of m above is nil;
	//		it doesn’t point to an initialized map. A nil map behaves like an empty map when reading,
	//			but attempts to write to a nil map will cause a runtime panic; don’t do that.
	//		To initialize a map, use the built in make function:
	//			m = make(map[string]int)
	var Map1 map[int]string
	Map1 = make(map[int]string)
	Map1[1] = "haha"
	Map1[2] = "heihei"
	fmt.Println(Map1)
	// 2. range 遍历 map的key
	for number := range Map1 {
		fmt.Println("number", number, "val", Map1[number])
	}
	// 3. 验证某个key是不是在map里
	number1, ok := Map1[1] // special case  ok-idiom
	fmt.Println(number1, ok)
	// 4. 通过key去删除元素
	delete(Map1, 1)
	fmt.Println(Map1)
	// 5. 用range 遍历 map; 如果for循环中有1个变量，则是key;如果有2个变量，则是key和val https://studygolang.com/articles/20049
	for x := range Map1 {
		fmt.Println(x)
	}
	// 6. Map 不能用 == 号去比较，只能和nil 比较  https://segmentfault.com/q/1010000019940462 回答中说map可以比较是错的，其他的可以看
	// 7. Map 在函数传参的时候, 可以直接传值而不需要使用指针
	//    原因是
	fmt.Println("Map1", Map1)
	AddNewItem(Map1, 3, "huhu")
	fmt.Println("Map1", Map1)
}
func AddNewItem(dict map[int]string, key int, val string) {
	dict[key] = val
}
