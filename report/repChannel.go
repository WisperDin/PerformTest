package report

import (
	"log"
	"time"
)

var RoutineChan chan RoutineResult

//使用通道接收来自各个用户线程的测试结果
func ListenRoutineResult(reporter *Reporter, timeoutSecond int) {
	timeout := make(chan bool, 1)
	count := 0
	for {
		select {
		case <-timeout:
			log.Println("ListenRoutineResult timeout!")
			return
		case routineRes := <-RoutineChan:
			//解除超时触发器
			log.Println(routineRes)
			go func() {
				<-timeout
			}()
			log.Println("in normal......")
			if !routineRes.IsFinish { //并没有完成的用户线程
				reporter.FailUserNum++
			}
			reporter.RoutineResult = append(reporter.RoutineResult, routineRes)
			count++
			//超时检测
			go func() {
				time.Sleep(time.Second * time.Duration(timeoutSecond))
				timeout <- true
			}()
		}
	}

}
