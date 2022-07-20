/*******
* @Author:qingmeng
* @Description:反射修改值,反射数据类型变化
* @File:main
* @Date:2022/7/20
 */

package main

import (
	"fmt"
	"reflect"
)

//修改值
func reflect01(b interface{}) {
	rVal:=reflect.ValueOf(b)
	//
	fmt.Printf("rVal kind=%v\n",rVal.Kind())
	//修改值
	rVal.Elem().SetInt(20)
}

//基本类型之间的变化
func reflectTest01(b interface{})  {
	//获取类型
	rTyp:=reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)

	//获取值
	rVal:=reflect.ValueOf(b)
	fmt.Printf("rVal=%v type=%T\n",rVal,rVal)
	//转为int
	n2:=2+rVal.Int()
	fmt.Println("n2=",n2)

	//转回interface{}
	iV:=rVal.Interface()
	//断言
	num2:=iV.(int)
	fmt.Println("num2+2=",num2+2)
}

//结构体变化
func reflectTest02(b interface{})  {
	rTyp:=reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)
	rVal:=reflect.ValueOf(b)

	//获取kind
	kind1:=rVal.Kind()
	kind2:=rTyp.Kind()
	fmt.Printf("kind =%v kind =%v\n",kind1,kind2)

	//转为interface{}
	iv:=rVal.Interface()
	fmt.Printf("iv =%v iv type=%T\n",iv,iv)

	//断言判断类型(b和iv一样)
	switch iv.(type){
	case Student:
		fmt.Println("student.name=",iv.(Student).name)
	}
	stu,ok:=iv.(Student)
	if ok{
		fmt.Println(stu.name)
	}
}

type Student struct {
	name string
	Age int
}

func test()  {
	fmt.Println("hello")
}

func main() {
	typeOf := reflect.TypeOf(test)
	valueOf := reflect.ValueOf(test)
	fmt.Println(typeOf)
	fmt.Println(valueOf.Call(nil))

	fmt.Println("***********")
	var num =1
	//传地址
	reflect01(&num)
	fmt.Println("num=",num)
	reflectTest01(num)

	//结构体反射
	stu:=Student{
		name: "tom",
		Age:  20,
	}

	reflectTest02(stu)
}