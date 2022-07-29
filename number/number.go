package number

// 为单元测试准备的一维可比较数定义，
// ordered num define for unit test
type Num interface {
	// 测试用临时泛型，一维可比较数字类型
	Uint | Int | Float
}

// 复数泛型
// complex number
type ComplexNumber interface {
	RealNumber | Complex
}
type C interface {
	R | Complex
}

// 实数泛型
// real number
type RealNumber interface {
	RationalNumber
}
type R interface {
	Q
}

// 有理数泛型/可商数泛型
// Rational Number
type RationalNumber interface {
	Integer | Float
}
type Q interface {
	Z | Float
}

// 整数泛型
// Integer
type Integer interface {
	Uint | Int
}
type Z interface {
	Uint | Int
}

// 自然数泛型 N
// natural number
type NaturalNumber interface {
	Uint
}
type N interface {
	Uint
}

// 机器复数类型
// go define of complex
type Complex interface {
	~complex64 | ~complex128
}

// go 浮点数泛型定义
// go define of floating-point number
type Float interface {
	~float32 | ~float64
}

// go 无符号数泛型定义
// go define of unsigned integer
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// go 有符号数泛型定义
// go define of integer
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
