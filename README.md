# 介绍
Go基础语法训练，base golang language dev 1.20 version

# 结构体
- 内聚、逻辑上的抽象

 # 微服务 
- 微服务架构下，组件间通信(http/tcp)、治理(注册与发现、节点选择、可用性、可观测性)问题。
- 模块化依然是单体架构、服务化粒度更细更加独立的部署
    
# GRPC
- 使用IDL生成代码的RPC框架，适合多语言通信场景～
- 小型系统直接使用HTTP接口～

# GORM
- GO和数据库交互的框架

# 启动项目
```shell
go run cmd/main.go -config configs/config.toml
```