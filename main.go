package main

import (
	."./routine"
	"./conf"

	"fmt"
	"log"
)


func main(){
	conf.Init("./test.toml")
	count:=0
	for i:=0;i<conf.App.C;i++  {
		go func(n int) {
			log.Println(fmt.Sprintf("start:%d",n))
			for j:=0;j<conf.App.N;j++ {
				EleRoutine(1,n)
			}
			count++
		}(i)
	}
	log.Println("finish...")
	//线程已全部开启,发出开始信号
	//close(sign)
	select {

	}
}
