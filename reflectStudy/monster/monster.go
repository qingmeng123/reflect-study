/*******
* @Author:qingmeng
* @Description:
* @File:monster
* @Date:2022/7/20
 */

package monster

import "fmt"

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32
	Sex string
}

func (s Monster) Print(){
	fmt.Println("-------start------")
	fmt.Println(s)
	fmt.Println("-------end------")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1+n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name=name
	s.Age=age
	s.Sex=sex
	s.Score=score
}
