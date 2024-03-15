package main
import "fmt"
func main(){
	//变量以_或字母开头，区分大小写
	a := 2.25//:=短变量声明[:=声明||=赋值] || float64 or float32
	str1 := "666"//len 
	str2 := "777"
	const PI = 3.14//定义常量
	// str2 := "666"//如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误[短变量声明]
	str_compare := str1 == str2 //boolean class[Can't be define other things ,it can only define True or False!!!]
	//var 定义变量[可省略Type(类型)或者=表达式中的任意一个，不可同时省略]
	var b  = complex(9,5) //complex128[actually is float64] or complex64[actually is float32]
	var c int //删除=表达式会为该变量分配默认值，没有未初始化的变量的概念！！！
	// var aa,bb,cc int = 11,22,33//变量也可如以下两种方式定义[相同类型]，可能报未使用的变量错误
	// var aaa,bbb,ccc = 11,"22",False//[不同类型，短变量声明同理]，可能报未使用的变量错误
	// var i,j=os.Open(name)//返回多个值的调用函数，可初始化多个变量[短变量声明同理]，可能报未使用的变量错误
// 	package main
// import "fmt"

// var x, y int
// var (  // 这种因式分解关键字的写法一般用于声明全局变量
//     a int
//     b bool
// )

// var c, d int = 1, 2
// var e, f = 123, "hello"

// //这种不带声明格式的只能在函数体中出现
// //g, h := 123, "hello"

// func main(){
//     g, h := 123, "hello"
//     fmt.Println(x, y, a, b, c, d, e, f, g, h)
// }
	fmt.Printf("int:=%T\n",c)
	fmt.Println(c)
	//end
	fmt.Printf("float:=%T\n",a)//Can use %T[defination] 默认不换行
	fmt.Println(a)//换行
	fmt.Printf("complex:=%T\n",b)
	fmt.Println(b)
	fmt.Printf("bool:=%T\n",str_compare)
	fmt.Println(str_compare)
	fmt.Printf("string_length:=%d",len(str1))
	fmt.Printf("\nstring_contain:=%s",str1)
	fmt.Printf("\nstring:=%T",str1)

}
