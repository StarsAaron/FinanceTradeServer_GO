/**
* @Author: Aaron
* @Date: 2020/11/23 17:17
 */
package main

import (
	"errors"
	"fmt"
	"strings"
)

type LoginAction struct {
}

func (action *LoginAction) HandlerRequest(client *Client, req *BaseReq) (e error) {
	dd := req.ResqData
	ds := strings.Split(dd, ",")
	if len(ds) == 2 {
		username := ds[0]
		pwd := ds[1]
		// 查数据库用户是否正确
		if action.checkPwd(username, pwd) {
			e = response(client.Conn, "LoginRspAction", SUCCEED, "成功登录", "成功登录")
		} else {
			e = response(client.Conn, "LoginRspAction", LOGIN_FAILED, "登录失败", "登录失败")
		}
	} else {
		return errors.New(fmt.Sprintf("数据不合法:%s", dd))
	}
	return
}

func (action *LoginAction) checkPwd(user string, pwd string) bool {
	return user == "aaron" && pwd == "123456"
}
