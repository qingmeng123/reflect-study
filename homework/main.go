/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date:2022/7/20
 */

package main

import (
	"fmt"
	"reflect"
)

type TSlice struct {
	i []interface{}
}

//入队列
func (s *TSlice) push(t interface{}) {
	if s.IsEmpty(){
		s.i = append(s.i, t)
		return
	}
	top := s.top()
	topT:=reflect.TypeOf(top)
	if reflect.TypeOf(t)!=topT{
		fmt.Println("只能放入的类型:",topT)
		return
	}

	s.i=append(s.i,t)

}

//出队列
func (s *TSlice) pop() interface{} {
	item := s.i[0] // 先进先出
	s.i = s.i[1:len(s.i)]
	return item
}

//获取顶部元素
func (s *TSlice) top()interface{} {
	if s.IsEmpty(){
		return nil
	}
	return s.i[0]
}

func (s *TSlice) IsEmpty() bool {
	return len(s.i) == 0
}

//长度
func (s *TSlice) length() int {
	return len(s.i)
}

func test1() {
	fmt.Println("hello")
}

func test2(a int,b string ) {
	fmt.Println(a,b)
}

type teacher struct {
	name string
}

type stu struct {
	name string
}



func main() {
	t:=TSlice{}
	//t.push(test1)

	t.push(test2)
	fmt.Println("长度",t.length())
	//te:=t.pop()
	//teV:=reflect.ValueOf(te)
	//teV.Call(nil)
	//fmt.Println("顶部元素",t.top())
	teparm:=t.pop()
	var params []interface{}
	params=append(params,1)
	params=append(params,"j")
	call(teparm,params...)

}

func call(fun interface{},params ...interface{})  {
	teparmV:=reflect.ValueOf(fun)
	num:=teparmV.Type().NumIn()
	if num!=len(params){
		fmt.Println("参数数量不对")
	}
	in:=make([]reflect.Value,num)
	for i, param := range params {
		in[i]=reflect.ValueOf(param)
	}
	teparmV.Call(in)
}