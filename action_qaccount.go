/**
* @Author: Aaron
* @Date: 2020/11/23 19:37
 */
package main

type QAccountAction struct {
}

func (action *QAccountAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	e = response(client.Conn, "QAccountAction", SUCCEED, "成功登录", "成功登录")
	return
}
