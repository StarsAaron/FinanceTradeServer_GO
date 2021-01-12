package main

import (
	list2 "container/list"
	"github.com/gorilla/websocket"
	"sync"
	"sync/atomic"
)

///////// 订单队列
type Order struct {
	Clientid string

	Price  float64
	Volume float64
}

type OrderQueue struct {
	reqCh chan *Order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{
		reqCh: make(chan *Order, 1000),
	}
}

func (orderQueue *OrderQueue) order(req *Order) {
	orderQueue.reqCh <- req
}

func (orderQueue *OrderQueue) loop() {
	defer logger.Info.Println("<OrderQueue exit loop>")
	for {
		select {
		case <-ctx.Done(): // 通知关闭
			return
		case req, ok := <-orderQueue.reqCh:
			if ok {
				logger.Info.Println(req)
			} else {
				return
			}
		}
	}
	//for req := range orderQueue.reqCh {
	//	logger.Info.Println(req)
	//}
}

///////// 客户端连接实例
type Client struct {
	Addr  string
	Conn  *websocket.Conn
	Token string
}

// NewClient 创建Client实例的简洁方式
func NewClient(addr string, conn *websocket.Conn, token string) *Client {
	cl := new(Client)
	cl.Addr = addr
	cl.Conn = conn
	cl.Token = token
	return cl
}

///////// 连接池

//////// 加锁是没有意义的，里面的实例还是可以被修改的，这里只是用来获取内容，不能修改，否则并发会有问题
type ConnectPool struct {
	clients sync.Map
	length  uint64
}

func (pool *ConnectPool) AddConnectObj(client *Client) {
	pool.clients.Store(client.Addr, client)
	atomic.AddUint64(&(pool.length), 1)
	logger.Info.Printf("新客户端加入:%s", client.Addr)
}

func (pool *ConnectPool) RemoveConnectObj(key string) {
	if v, ok := pool.clients.Load(key); ok {
		client := v.(*Client)
		err := client.Conn.Close()
		if err != nil {
			logger.Error.Println("关闭连接出错：", err)
		}
		pool.clients.Delete(key)
		atomic.AddUint64(&(pool.length), ^uint64(1-1))
		logger.Info.Printf("客户端移除:%s 当前客户端数量：%d", key, atomic.LoadUint64(&(pool.length)))
	}
}

// GetConnectObj 获取连接实例
func (pool *ConnectPool) GetConnectObj(id string) *Client {
	if v, ok := pool.clients.Load(id); ok {
		return v.(*Client)
	} else {
		return nil
	}
}

// HasConnectObj 判断是否存在相同名称的实例
func (pool *ConnectPool) HasConnectObj(key string) bool {
	_, ok := pool.clients.Load(key)
	return ok
}

// HasConnectObj 判断是否存在相同名称的实例
func (pool *ConnectPool) Release() {
	list := list2.New()
	pool.clients.Range(func(key, value interface{}) bool {
		client := value.(*Client)
		client.Conn.Close()
		list.PushBack(key)
		logger.Info.Println("移除连接：", key)
		return true
	})
	//l := list.New()
	for i := list.Front(); i != nil; i = i.Next() {
		pool.clients.Delete(i)
	}
}

///////// 其它
// 请求的格式： action|<reqdata>|guid|sign
type BaseReq struct {
	Action   string
	ResqData string
	Guid     string
	Sign     string
}
