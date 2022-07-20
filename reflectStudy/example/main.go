/*******
* @Author:qingmeng
* @Description:反射实例，主要作用：做框架
* @File:main
* @Date:2022/7/20
 */

package main

import (
	"fmt"
	"reflect"
	"reflect-homework/reflectStudy/monster"
)

func TestStruct(a interface{}) {
	typ:=reflect.TypeOf(a)
	val:=reflect.ValueOf(a)
	kd:=val.Kind()
	if kd!=reflect.Struct{
		fmt.Println("expect struct")
		return
	}
	num:=val.NumField()	//属性数量
	fmt.Printf("struct has %d fields\n",num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n",i,val.Field(i))
		//获取tag标签
		tagVal:=typ.Field(i).Tag.Get("json")
		if tagVal!=""{
			fmt.Printf("Field %d: tag=%v\n",i,tagVal)
		}
	}
	
	//获取结构体方法(就算用typ,也不能统计私有方法数量，也不能调用私有方法)
	numOfMethod:=typ.NumMethod()
	fmt.Printf("struct has %d methods\n",numOfMethod)
	for i := 0; i < numOfMethod; i++ {
		fmt.Println("filed",val.Method(i))
	}
	fmt.Println("numfiled:",num)
	//调用方法，以方法首字母排序
	val.Method(1).Call(nil)
	//调用第一个方法
	var params []reflect.Value
	params=append(params,reflect.ValueOf(10))
	params=append(params,reflect.ValueOf(40))
	res:=val.Method(0).Call(params)
	fmt.Println("res=",res[0].Int())
}

func main() {
 	a:=monster.Monster{
		Name:  "duryun",
		Age:   20,
		Score: 100,
		Sex:   "男",
	}
	TestStruct(a)
}
