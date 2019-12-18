package main

import (
	"fmt"
	 "github.com/HogenYuan/poster-go/poster"
)

func main(){
	_,err :=poster_go.TestPoster()
	if err != nil{
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("生成海报成功！")
	}
}