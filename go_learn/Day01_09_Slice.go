package main

import (
	"fmt"
	"reflect"
)

func main() {
	// slice是对于数组Array的更高一层的包装、抽象(Abstraction)
	// Array 本身的长度是无法变化，但Slice 包装后形成了一个长度可以变化的Array
	// Slice 是一个引用对象，指向底层对应的一个数组对象。所以说如果要修改Array或者Slice，直接把这个值取出来修改即可
	// 底层逻辑
	//     struct slice{
	//         *ptr  指向对应底层数组的一个指针
	//         len   当前slice(对应的数组中已经使用的)长度
	//         cap   最大长度，即当前slice指向的数组的最大长度
	//   	}
	// 也就是说，Slice对应指向的底层Array的大小可能和已经被使用的大小是不一样的。
	// 1. 创建Slice
	//    1.1 直接定义初值
	s := []int{1, 2, 3}                        // 和数组创建的区别是不需要在'[]' 中指定大小
	fmt.Println("type of ", reflect.TypeOf(s)) // 可以看到得到的结果是 []int， 如果是数组的话结果是 [5]int
	//    1.2 利用数组生产切片. 将数组进行切片，得到的就是一个切片，不过好像不能直接写 '[5]int{1,2,3}[:]' 会报错，目测有可能和GC有关
	array1 := [5]int{1, 2, 3}
	a := array1[:]
	fmt.Println(reflect.TypeOf(a).Kind())
	//    1.3 利用内置make 生成
	b := make([]int, 5, 10)
	fmt.Println(b)
	// 2. len() & cap
	fmt.Println("len: ", len(b), "Cap: ", cap(b))
	// 3. 由于slice是一个引用对象，如果是空slice，其对象= nil
	var o []int
	fmt.Println(o, len(o), cap(o))
	fmt.Println(o == nil) // true
	// 4. append & copy
	// append 踩坑点: 如果append以后没有超过array的cap, 那底层的array对象就不变
	//             : 如果append以后超过array的cap, 那底层会重新生成一个新的array。把这个slice指向新的array。
	//				 但如果同时有其他slice也指向原本append之前的那个array，那个其他的slice的内容就不会变。因为底层新增了一个array
	first := [2]int{1, 2}
	slice1 := first[:]
	slice2 := slice1
	slice1 = append(slice1, 2, 3, 4, 5)
	fmt.Println("slice1:", slice1, "slice2", slice2) // slice2 还是 [1,2]
	// copy(a,b) 把 b copy 给 a
	var slice3 []int
	fmt.Println("slice3", slice3 == nil)
	var slice4 []int = []int{1, 2}
	copy(slice3, slice4) // 失败，因为copy的工作逻辑是拷贝的个数是以a，b两个slice中较短长度来定的，这里是0，所以没拷贝
	fmt.Println("slice3", slice3, "slice4", slice4)
	slice5 := make([]int, 5, 5)
	copy(slice5, slice4) // 拷贝slice4 中的[1,2] 给 slice5
	fmt.Println("slice4", slice4, "slice5", slice5)

	// 数组也可以直接使用copy,但必须先转成slice
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6}
	copy(arr1[:], arr2[:]) // 直接 copy(arr1,arr2)是错误的
	fmt.Println("arr1", arr1, "arr2", arr2)

	// slice对象在传参时是会复制一份，也就是go里面全是传值
	arr3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr4 := arr3[3:5]
	add2(&arr4, 0)
	fmt.Println(arr4)
	add2(&arr4, 1)
	fmt.Println(arr4)
}
func add2(sli *[]int, num int) {
	*sli = append(*sli, num)
	fmt.Println(*sli)
	l := new(*[]int)
}
