/**
* @Author: Aaron
* @Date: 2020/11/23 19:39
 */
package main

type QPositionAction struct {
}

func (action *QPositionAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	e = response(client.Conn, "QPositionAction", SUCCEED, "成功登录", "成功登录")
	return
}
