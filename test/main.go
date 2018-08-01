package main

import (

	"fmt"
	"utils"
)

func main() {
	res := utils.GetSign(7, 25)
	fmt.Println(res)
	t:="2018-07-30 07:45:26"
	s,_:=utils.ConvertLocalTimeByString(t)
	fmt.Println(s)
}
