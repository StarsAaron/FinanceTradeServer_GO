package main

import (
	"context"
	_ "net/http/pprof"
	"runtime"
)

// 全局变量
var logger = NewLog("trade")
var ctx, cancelAllConn = context.WithCancel(context.Background()) /// ctx 主context，控制所有的goroutine退出

func main() {
	//queue := safequeue.NewOneSideQueue()
	//queue.Create("1","2","3","4","5")
	//queue.Print()
	//queue.Pop()
	//queue.Pop()
	//queue.Push("85")
	//queue.Push("84")
	//queue.Print()

	host := "127.0.0.1"
	port := 9002
	logger.Info.Printf("****************************************\n")
	logger.Info.Printf("核心：%d host:%s port:%d \n", runtime.GOMAXPROCS(runtime.NumCPU()), host, port)
	logger.Info.Printf("****************************************")
	StartServerAndWait(host, port)
	logger.close()
}
