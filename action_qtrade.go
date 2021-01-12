/**
* @Author: Aaron
* @Date: 2020/11/23 19:37
 */
package main

type QTradeAction struct {
}

func (action *QTradeAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	e = response(client.Conn, "QTradeAction", SUCCEED, "成功登录", "成功登录")
	return
}
