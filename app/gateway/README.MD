# 网关服务

```
/ 项目结构 # 参考：https://goframe.org/docs/design/project-structure
├── api             # 对外接口
├── boot            # 启动初始化
├── hack            # 工具脚本
├── internal        # 内部逻辑
│   ├── cmd         # 入口指令
│   ├── consts      # 常量定义
│   ├── controller  # 接口处理
│   ├── dao         # 数据访问
│   ├── logic       # 业务逻辑
│   ├── model       # 结构模型
│   |   ├── do      # 领域对象
│   │   └── entity  # 数据模型
│   ├── packed      # 资源管理 参考：https://goframe.org/docs/core/gres-pack-using-cli
│   └── service     # 业务实现
├── manifest        # 交付清单
│   ├── config      # 配置管理
│   ├── deploy      # 部署文件
│   ├── docker      # 镜像文件
│   ├── i18n        # 多语言文件
│   └── protobuf    # 协议文件
├── resource        # 静态资源
├── utility         # 工具包
├── go.mod          # 依赖管理
└── main.go         # 入口文件
```