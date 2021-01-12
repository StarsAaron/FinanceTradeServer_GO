
## 基础请求
- req
    - action|<reqdata>|guid|sign
    - 请求的第一位跟回调都是 action
    - token 是隐藏字段，跟服务器连接成功由服务器返回
- bak
    - action|code|msgs|<respdata>|guid


## 接口
- 连接
    - bak
        1. action: ConnectRspAction
        2. code: 200 成功
        3. msgs: 提示
        4. respdata: token (用于数据签名)
        5. guid: 随机生成

- 登录 login
    - req
        1. action: LoginAction
        2. reqdata: username,password
        3. guid: 随机生成
        4. sign = md5(req+guid+token)
    
    - bak
        1. action: LoginRspAction
        2. code: 200 成功
        3. msgs: 提示
        4. respdata: 空
        5. guid: 随机生成
        
            

- 登出 logout
    - req
        1. action: LogoutAction
        2. reqdata: 空
        3. guid: 随机生成
        4. sign = md5(req+guid+token)
        
    - bak
        1. action: LogoutRspAction
        2. code: 200 成功
        3. msgs: 提示
        4. respdata: 空
        5. guid: 随机生成

- 下单
    - req
        1. action: OrderInsertAction
        2. reqdata: 
            - otype
               - 0: 市价
               - 1：限价
               - 2：止盈
               - 3：止损 
        3. guid: 随机生成
        4. sign = md5(req+guid+token)
        
    - bak
        1. action: OrderInsertRspAction
        2. code: 200 成功
        3. msgs: 操作提示信息
        4. respdata: 内容
        5. guid: 随机生成

- 查询历史委托
    - req
        1. action: QOrderAction
        2. reqdata: 空
        3. guid: 随机生成
        4. sign = md5(req+guid+token)
        
    - bak
        1. action: QOrderRspAction
        2. code: 200 成功
        3. msgs: 操作提示信息
        4. respdata: 内容
        5. guid: 随机生成
        
- 查询历史成交
    - req
        1. action: QTradeAction
        2. reqdata: 空
        3. sign = md5(req+guid+token)
        
    - bak
        1. action: QTradeRspAction
        2. code: 200 成功
        3. msgs: 操作提示信息
        4. respdata: 内容
        5. guid: 随机生成
        
- 持仓查询
    - req
        1. action: QPositionAction
        2. reqdata: 空
        3. guid: 随机生成
        4. sign = md5(req+guid+token)
        
    - bak
        1. action: QPositionRspAction
        2. code: 200 成功
        3. msgs: 操作提示信息
        4. respdata: 内容
        5. guid: 随机生成
        
- 账户查询
    - req
        1. action: QAccountAction
        2. reqdata: 空
        3. guid: 随机生成
        4. sign = md5(req+guid+token)
        
    - bak
        1. action: QAccountRspAction
        2. code: 200 成功
        3. msgs: 操作提示信息
        4. respdata: 内容
        5. guid: 随机生成