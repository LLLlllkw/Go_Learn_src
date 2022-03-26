package main

func main() {
	// 1. Bool default:false 0被认为是false
	var a1 bool
	println(a1)
	a1 = true
	println(a1)
	// 2. Int
	var b1 int8 = 12   // -2 ^ 7 ~ 2 ^ 7 - 1
	var b2 int16 = 200 // -2 ^ 15 ~ 2 ^ 15 - 1
	var b3 int32 = 99999
	var b4 int64 = 99999999
	println(b1, b2, b3, b4)
	var (
		c1 uint8  = 200
		c2 uint16 = 2999
		c3 uint32 = 299999
		c4 uint64 = 21219999
	)
	println(c1, c2, c3, c4)
	var (
		c5 int  = 20000 // 根据底层平台，如果是32位就是32位int，如果是64位就是64位int
		c6 uint = 99999 // 根据底层平台，如果是32位就是32位uint，如果是64位就是64位uint
	)
	println(c5, c6)

	// 3. float32, float 64, complex
	var (
		d1 float32    = 2.33          // 32 bits
		d2 float64    = 2.99          // 64 bits
		d3 complex64  = 8 + 9i        // 64 bits 复数，实部虚部各占 32 bits
		d4 complex128 = complex(8, 9) // 128 bits 复数，实部虚部各占 64 bits
	)
	println(d1, d2, d3, d4)

	// 4. others
	var (
		e1 byte    = 12 // ~= uint8
		e2 rune    = 13 // ~= int32
		e3 uint    = 99 // 基于平台，32 或 64 bits
		e4 uintptr = 89 // 存放一个指针，所以是uint
	)
	println(e1, e2, e3, e4)

	//  5. 字符串  go 中字符串由字符拼接, 字符严格使用单引号('), 字符串严格使用双引号(")
	var s1 string = "hello"
	println(s1)
	//  6. 类型转换  a = Type(Variable)  要求转型后的变量a 的类型要和转型的Type一致  字符串和其他类型转换可以用strconv
	f1 := string(e1)  // 将e1从uint8用ascii码转成了string
	c4 = uint64(c3)  // 将变量向上转型
	b3 = int32(b4)   // 将变量向下转型 会有溢出的问题
	ff := int(d1)  // float和int直接的转换
	println(f1,ff)
}
