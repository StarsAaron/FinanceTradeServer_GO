/**
* @Author: Aaron
* @Date: 2020/11/23 19:35
 */
package main

type OrderInsertAction struct {
}

func (action *OrderInsertAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	e = response(client.Conn, "OrderInsertAction", SUCCEED, "成功登录", "成功登录")
	return
}
