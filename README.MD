# 自游

```
/ 项目结构 # 参考：https://goframe.org/docs/design/project-structure
├── app             # 对外接口
│   ├── gateway     # 网关服务
│   ├── org         # 组织服务
│   └── system      # 系统服务
├── common          # 公共文件
│   ├── berror      # 错误代码
│   ├── enum        # 枚举文件
│   ├── model       # 结构模型
│   └── tools       # 工具文件
│       ├── fvalid  # 自定义验证规则
│       └── string  # 字符串工具
├── config          # 配置文件
│   └── sql         # sql文件
├── hack            # 工具脚本
├── utility         # 工具包
└── go.mod          # 依赖管理
```