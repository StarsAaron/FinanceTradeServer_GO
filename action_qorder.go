/**
* @Author: Aaron
* @Date: 2020/11/23 19:36
 */
package main

type QOrderAction struct {
}

func (action *QOrderAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	e = response(client.Conn, "QOrderAction", SUCCEED, "成功登录", "成功登录")
	return
}
