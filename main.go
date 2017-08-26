package main

import (
	."./routine"
	"./conf"
	"log"
	"fmt"
)


func main(){
	conf.Init("./test.toml")
	sign:=make(chan bool)
	count:=0
	for i:=0;i<conf.App.C;i++  {
		go func(n int) {
			//阻塞线程,等待开始信号
			<-sign
			log.Println(fmt.Sprintf("start:%d",n))
			for j:=0;j<conf.App.N;j++ {
				EleRoutine(1,n)
			}
			count++
		}(i)
	}
	log.Println("finish...")
	//线程已全部开启,发出开始信号
	close(sign)
	for {
		if count>=1000{
			log.Println("get 1000")
		}
		if count>=3000{
			log.Println("get 3000")
		}
		if count>=5000{
			log.Println("get 5000")
		}
		if count>=7000{
			log.Println("get 7000")
		}
		if count>=10000{
			log.Println("get 10000")
		}
	}
}
