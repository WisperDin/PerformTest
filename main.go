package main

import (
	."./routine"
	"./conf"
	//"fmt"
	"fmt"
)


func main(){
	conf.Init("./test.toml")
	for i:=0;i<1000;i++  {
		go func() {
			for  {
				EleRoutine(1,i)
			}
		}()
		fmt.Println(i)
	}

}
