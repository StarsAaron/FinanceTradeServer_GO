/**
* @Author: Aaron
* @Date: 2020/11/23 17:17
 */
package main

type LogoutAction struct {
}

func (action *LogoutAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	pool.RemoveConnectObj(client.Addr)
	e = response(client.Conn, "LogoutRspAction", LOGOUT_FAILED, "退出成功", "退出成功")
	//if err != nil {
	//	logger.Error.Printf("退出失败%s", client.Addr)
	//	e = response(client.Conn, "LogoutRspAction", SUCCEED, "退出失败", "退出成功")
	//} else {
	//	e = response(client.Conn, "LogoutRspAction", LOGOUT_FAILED, "退出成功", "退出成功")
	//}
	return
}
