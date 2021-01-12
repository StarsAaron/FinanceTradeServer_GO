package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// websocket.Upgrader 配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

// 连接映射
var pool = ConnectPool{}

var orderQueue = NewOrderQueue()

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	//判断请求是否为websocket升级请求。
	if websocket.IsWebSocketUpgrade(r) {
		// 升级协议
		conn, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			logger.Error.Println("Upgrade failed!!!!!")
			return
		}

		// 初始化连接设置
		raddr := conn.RemoteAddr().String()
		logger.Warn.Println("----> Client connect:", raddr)

		if pool.HasConnectObj(raddr) {
			tt := pool.GetConnectObj(raddr).Token
			_ = response(conn, "ConnectRspAction", CONNECT_EXIT, "连接已存在", tt)
			logger.Warn.Println("Existed client!!!!! ip:", raddr)
			return
		}

		//conn.SetCloseHandler(func(code int, text string) error {
		//	//rst :=websocket.FormatCloseMessage(code,text)
		//	//Warning.Println(string(rst[:]))
		//	logger.Info.Println(fmt.Sprintf("Connect closed.code:%d text:%s", code, text))
		//	pool.unregister <- raddr
		//	return nil
		//})

		client := NewClient(raddr, conn, GetGUID())
		pool.AddConnectObj(client)

		err2 := response(conn, "ConnectRspAction", SUCCEED, "连接成功", client.Token)
		if err2 != nil {
			logger.Error.Println(err2)
			return
		}
		go readLoop(client)
	} else {
		//处理普通请求
		logger.Info.Println("普通请求")
	}
}

// 处理请求跟响应
func readLoop(cli *Client) {
	hd := RequestHandlerFactory{
		client: cli,
	}
	defer func() {
		logger.Warn.Printf("<停止监听 %s>", cli.Addr)
		pool.RemoveConnectObj(cli.Addr)
	}()

	for {
		select {
		case <-ctx.Done(): // 通知关闭
			return
		default:
			_, c, e := cli.Conn.ReadMessage()
			if e != nil {
				//logger.Error.Println("ReadMessage failed!!!!! Error:", e.Error())
				return
			} else {
				// 打印内容
				logger.Info.Println("》》》 收到信息：", string(c))
				err := hd.DoAction(string(c))
				if err != nil {
					logger.Error.Println(err)
				}
			}
		}
	}

}

/// 启动websocket服务
func StartServerAndWait(host string, port int) {
	logger.Info.Println("Startting server...")
	//httpAddr := fmt.Sprintf("%s:%d", host, port)
	//http.HandleFunc("/work", serveHTTP)
	//go http.ListenAndServe(httpAddr, nil)

	// 初始化连接管理器
	go orderQueue.loop()
	// 初始化http服务
	mux := http.NewServeMux()
	mux.HandleFunc("/trade", serveHTTP)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}
	go server.ListenAndServe()
	logger.Info.Println("Listenning...")

	// 等待中断信号以优雅地关闭服务器
	waitsignal := make(chan os.Signal)
	signal.Notify(waitsignal, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-waitsignal:
		cancelAllConn()
		server.Shutdown(ctx)
		pool.Release()
		logger.Warn.Printf("XXXXXXXXXXXX 程序终止信号 XXXXXXXXXXXX")
	}
}
