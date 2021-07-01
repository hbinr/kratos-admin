# interfaces


## 为什么叫 interfaces?
该层相当于 DDD 中的用户接口层

并没有使用Kratos中的 http 组件：
- kratos 可以通过 `.proto` 文件生成 `xx_http.pb.go` 文件
- 入参处理、出参处理、路由定义、中间件(revovery, log)定义


## 该层的作用是什么？
处理用户输入的入参，然后调用应用层，返回响应

- 类似传统三层架构中的controller