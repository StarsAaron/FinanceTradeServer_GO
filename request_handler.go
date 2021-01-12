/**
* @Author: Aaron
* @Date: 2020/11/6 15:15
 */
package main

import (
	"errors"
	"github.com/StarsAaron/rjtools/injecttools"
)

type Handler interface {
	HandlerRequest(client *Client, req *BaseReq) error
}

type RequestHandlerFactory struct {
	client *Client
}

// 校验数据是否被篡改
func (factory *RequestHandlerFactory) CheckData(req *BaseReq) bool {
	localSign := createSign(factory.client.Token, req.ResqData, req.Guid)

	logger.Info.Println("签名：", localSign, "请求签名：", req.Sign)
	if localSign != req.Sign {
		return false
	}
	return true
}

func (factory *RequestHandlerFactory) CreateAction(opType string) Handler {
	switch opType {
	case "LoginAction":
		return &LoginAction{}
	case "LogoutAction":
		return &LogoutAction{}
	case "OrderInsertAction":
		return &OrderInsertAction{}
	case "QOrderAction":
		return &QOrderAction{}
	case "QTradeAction":
		return &QTradeAction{}
	case "QAccountAction":
		return &QAccountAction{}
	case "QPositionAction":
		return &QPositionAction{}
	}
	return nil
}

func (factory *RequestHandlerFactory) DoAction(msg string) error {
	// TradeRspAction OrderStatusRspAction
	req := BaseReq{}
	err := injecttools.String2Struct(msg, &req)
	if err != nil {
		return err
	}
	if !factory.CheckData(&req) {
		return errors.New("数据校验不符合签名信息！！！客户端：" + factory.client.Addr)
	}

	action := factory.CreateAction(req.Action)
	if action == nil {
		return errors.New("不存在的响应类型，无法选择操作")
	}
	return action.HandlerRequest(factory.client, &req)
}
