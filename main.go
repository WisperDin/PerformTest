package main

import (
	"./conf"
	. "./routine"
	/*
			"fmt"
			"log"*/
	"./report"
)

/*func run() {
	count := 0
	for i := 0; i < conf.App.C; i++ {
		go func(n int) {
			log.Println(fmt.Sprintf("start:%d", n))
			for j := 0; j < conf.App.N; j++ {
				EleRoutine(1, n)
			}
			count++
		}(i)
	}
	log.Println("finish...")
	select {}
}*/

func main() {
	conf.Init("./test.toml")
	//EleRoutine(1, 1)
	//run()
	rep := report.Reporter{}
	rep.UserNum = 1000
	report.RoutineChan = make(chan report.RoutineResult)
	go report.ListenRoutineResult(&rep, 5)

	go EleRoutine(1, 1)
	for {

	}
}
