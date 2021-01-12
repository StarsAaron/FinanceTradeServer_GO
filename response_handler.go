/**
* @Author: Aaron
* @Date: 2020/11/23 17:06
 */
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// 返回客户端响应
func response(conn *websocket.Conn, action string, code int32, msg string, data string) error {
	temp := fmt.Sprintf("%s|%d|%s|%s|%s", action, code, msg, data, GetGUID())
	logger.Info.Println("《《《 返回：", temp)
	return conn.WriteMessage(websocket.TextMessage, []byte(temp))
}
