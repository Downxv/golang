package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

// 全局变量
var (
	n = 1
)

// go语言中有未引入的包和变量会导致编译不通过
func main() {
	fmt.Println("n = ", n)
	// 转义字符
	fmt.Println("a\tb\nc", "\\\"d\re")
	// 定义变量
	var i int
	i = 10
	fmt.Println("i=", i)
	var num = 20
	fmt.Println("num = ", num)
	sum := i + num
	fmt.Println("sum = ", sum)
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3) //切片 映射 函数 指针变量默认为nil
	// var n1, n2 = 10, "ok"
	// n1, n2 := 10, "ok"
	// +号使用
	fmt.Println(4+5, "hk"+"lo")
	// 整数类型
	// int8 int16 int32 int64 int rune=int32存中文
	// uint8 uint16 uint32 uint64 uint byte=uint8
	var q = 1
	fmt.Printf("q的类型是%T,字节数是%d\n", q, unsafe.Sizeof(q))
	// 浮点类型
	var p float32
	var o float64 //开发中使用
	fmt.Println(p, o)
	// 字符类型
	var c1 byte = 'a'
	var c2 rune = '啊'
	fmt.Printf("c1=%c,c2=%c", c1, c2)
	// 布尔类型
	var bl bool = true
	fmt.Println(bl)
	// 字符串类型, 字符串不可变
	var str string = "hello"
	fmt.Println(str)
	var str1 string = `kl\n"` // 反引号输出原始字符串
	str1 += "ok" +            // +号放在上一行
		"op"
	fmt.Println(str1)
	// 类型转换，需要显式转换
	var ii int = 2
	var j float64 = float64(ii)
	fmt.Println(j)
	// 基本数据类型转字符串
	var bb bool = false
	var strr string = fmt.Sprintf("%t", bb)
	fmt.Printf("strr=%q\n", strr)
	// string转基本类型
	var sss string = "true"
	bbb, _ := strconv.ParseBool(sss) //多重赋值 匿名变量
	fmt.Println(bbb, "bbb的地址是", &bbb, "bbb的类型为", reflect.TypeOf(bbb))
	// 字节缓冲处理大量字符串操作
	var ss1, ss2 string = "qqqq", "wwww"
	var ss3 bytes.Buffer
	ss3.WriteString(ss1)
	ss3.WriteString(ss2)
	fmt.Println(ss3.String())
	index1 := strings.Index(ss3.String(), "q")
	index2 := strings.LastIndex(ss3.String(), "q")
	fmt.Println(ss3.String()[index1:], ss3.String()[index2:])
	// 常量
	const ss4 string = "hello world"
	// 常量组定义
	const (
		aa1 = 1
		aa2
		aa3 = 3
	)
	fmt.Println(ss4, aa1, aa2, aa3)
	// 指针
	var ptr *int
	fmt.Println(ptr, "指针变量ptr的地址为", &ptr)
	var p1 int = 2
	ptr = &p1
	fmt.Println("指针Ptr指向的值为", *ptr)
	*ptr = 4
	fmt.Println(p1)
	// 初始化指针
	var ptr1 *int = new(int)
	fmt.Println(*ptr1)
	// 位运算符
	var aa4 int = 1
	var aa5 int = 0
	var bb1 bool = false
	fmt.Println(aa4&aa5, aa4|aa5, aa4^aa5, aa4<<1, aa4>>1, !bb1, ^aa5)
	// if语句
	if num := runtime.NumCPU(); num >= 1 {
		fmt.Println(num)
	} else if num > 2 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
	// for语句
	for i := 1; i < 5; i++ {
		fmt.Print(i)
	}
	// 带标签break continue类似
esc:
	for {
		for {
			if i := 1; i < 5 {
				fmt.Print(i)
			}
			fmt.Println()
			break esc
		}
	}
	// switch语句
	switch 1 + 1 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}
	// 数组
	var array [3]string
	array = [...]string{"l", "k", "j"}
	fmt.Println(array)
	for k, v := range array {
		fmt.Println("下标", k, "值", v)
	}
	// 切片 动态数组
	var sl []int = []int{1, 2, 3, 4}
	fmt.Print(len(sl), cap(sl), sl == nil)
	// make初始化切片
	var sl1 []int
	sl1 = make([]int, 2, 5)
	fmt.Print(len(sl1), cap(sl1), sl1 == nil)
	sl1 = append(sl1, 5) // 添加元素
	sl1 = append(sl1, 8)
	fmt.Println(sl1)
	sl1 = append(sl1[0:1], sl1[2:]...) // 删除元素, ...表示切片展开
	fmt.Println(sl1)
	// 映射
	var mp map[string]int = map[string]int{
		"i": 8,
		"l": 0,
		"k": 9,
	}
	fmt.Println(mp, len(mp))
	// make 初始化映射
	var mp1 map[string]int = make(map[string]int)
	mp1["u"] = 6
	mp1["p"] = 7
	for k, v := range mp1 {
		fmt.Println(k, v)
	}
	delete(mp, "l") // 删除键值对
	fmt.Println(mp)
	fmt.Println("sum=", add(1, 3))
	// 函数变量
	var f1 func(x int, y int) (sum int)
	f1 = add
	f2 := add // 短变量赋值
	fmt.Println(f1(1, 2), f2(3, 4))
	var aa6 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(addall(aa6...))
	// 匿名函数
	func(w string) {
		fmt.Println("hello" + w)
	}("world")
	f := func(w string) {
		fmt.Println("hello" + w)
	}
	f("world")
	// 闭包 (调用了全局变量的匿名函数)
	var ii1 int = 1
	func() {
		ii1++
		fmt.Println(ii1)
	}()
	f4 := f3(12)
	fmt.Println(f4())
	// 延迟执行函数 defer关键字 常用于释放资源
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("not defer")
	// 内置函数 略
	// 包管理 包中可定义init函数 包中大写开头的标识符是公有的，小写开头的是私有的
	// 结构体
	// type St struct {
	// 	id   int
	// 	Name string // 首字母大写是公有属性
	// }
	// 结构体实例化
	var st1 St
	st1.id = 10
	st1.Name = "o"
	fmt.Println(st1)
	st2 := new(St)
	fmt.Println(st2)
	// 结构体初始化
	var st3 St = St{
		id:   10,
		Name: "p",
	}
	fmt.Println(st3)
	// 结构体方法使用
	st3.changeId()
	fmt.Println(st3)
	// 结构体内嵌
	type At struct {
		St
		id int
	}
	At1 := &At{
		St: St{
			id:   12,
			Name: "k",
		},
		id: 23,
	}
	At1.St.id = 44
	fmt.Println(At1)
	// 错误处理和panic recover errors.New()
	pc()
	fmt.Println("出错后")
	// 文件操作
	// 列出目录
	file, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("读取目录失败")
	} else {
		for _, v := range file {
			if v.IsDir() {
				fmt.Println("目录 ", v.Name())
			} else {
				fmt.Println("文件 ", v.Name())
			}
		}
	}
	// 递归遍历目录用Walk
	// 创建目录
	err1 := os.Mkdir("./app", 0777)
	if err1 != nil {
		fmt.Println("目录创建失败", err1)
	} else {
		os.Chmod("./app", 0777)
		fmt.Println("目录创建成功")
	}
	// 创建多级目录
	err2 := os.MkdirAll("./app/qqq/www", 0777)
	if err2 != nil {
		fmt.Println("多级目录创建失败", err2)
	} else {
		fmt.Println("多级目录创建成功")
		os.Chmod("./app/qqq/www", 0777)
	}
	// 删除目录或文件
	err3 := os.Remove("./app/qqq/www")
	if err3 != nil {
		fmt.Println("删除目录失败", err3)
	} else {
		fmt.Println("删除目录成功")
	}
	// 删除非空目录
	err4 := os.RemoveAll("./app")
	if err4 != nil {
		fmt.Println("删除非空目录失败")
	} else {
		fmt.Println("删除非空目录成功")
	}
	// 文件读取 Read / ReadAtu
	ff, err := os.OpenFile("./语法.go", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件 ./语法.go失败")
	} else {
		buf := make([]byte, 1024)
		for {
			le, _ := ff.Read(buf)
			if le != 0 {
				// fmt.Println(string(buf))
			} else {
				break
			}
		}
	}
	ff.Close()
	// 文件写入 Write WriteAt
	ff1, err := os.OpenFile("./a.txt", os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	var by []byte
	by = []byte("数据\r\n")
	ff1.Write(by)
	ff1.Close()
	// 处理json文件
	// 正则表达式
	mt, err := regexp.Match("as*", []byte("dfeassv"))
	fmt.Println(mt)
	// regexp.MatchString()
	re := regexp.MustCompile(`\^\d+`)
	mt1 := re.FindStringIndex("dfo^131sdf")
	fmt.Println(mt1)
}

// 函数
func add(x int, y int) (sum int) {
	sum = x + y
	return sum
}

// 可变参数
func addall(s ...int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return sum
}

// 函数返回闭包
func f3(a int) func() int {
	return func() int {
		a++
		return a
	}
}

type St struct {
	id   int
	Name string // 首字母大写是公有属性
}

func (s *St) changeId() {
	s.id = 30
}

// 错误处理
func pc() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获错误")
		}
	}()
	fmt.Println("错误前")
	panic("出错了")
}
