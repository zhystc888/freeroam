package boot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/config/nacos/v2"
	rnacos "github.com/gogf/gf/contrib/registry/nacos/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"os"
)

var (
	configPath      = gfile.Pwd() + "/config/config.yaml"
	localConfigPath = gfile.Pwd() + "/app/user/manifest/config/config.yaml"
	env             = os.Getenv("APP_ENV")
)

func init() {
	fmt.Println("全局配置文件:", configPath)
	fmt.Println("本地配置文件:", localConfigPath)
	// 环境变量决定使用本地配置还是 Nacos
	if env == "local" {
		// 本地环境：使用文件配置
		localAdapter, err := gcfg.NewAdapterFile(localConfigPath)
		if err != nil {
			panic(err)
		}
		g.Cfg().SetAdapter(localAdapter)
	} else {
		ctx := gctx.GetInitCtx()
		// 配置中心
		adapter, err := getNacosAdapter(ctx, configPath)
		g.Cfg().SetAdapter(adapter)
		if err != nil {
			panic(err)
		}
		// 注册中心
		getNacosRegister(ctx, configPath)
	}

	// 异步日志
	g.Log().SetAsync(true)
}

func getNacosClientConfig(ctx context.Context, configPath string) constant.ClientConfig {
	return constant.ClientConfig{
		CacheDir:            g.Cfg(configPath).MustGet(ctx, "nacos.cacheDir").String(),
		LogDir:              g.Cfg(configPath).MustGet(ctx, "nacos.logDir").String(),
		NamespaceId:         g.Cfg(configPath).MustGet(ctx, "nacos.namespaceId").String(),
		Username:            g.Cfg(configPath).MustGet(ctx, "nacos.username").String(),
		Password:            g.Cfg(configPath).MustGet(ctx, "nacos.password").String(),
		NotLoadCacheAtStart: true,
		DisableUseSnapShot:  true,
		TimeoutMs:           15000, // 增加超时到 15 秒
	}
}

func getNacosRegister(ctx context.Context, configPath string) {
	addr := g.Cfg(configPath).MustGet(ctx, "nacos.address").String()
	port := g.Cfg(configPath).MustGet(ctx, "nacos.port").Int()
	grpcx.Resolver.Register(rnacos.New(fmt.Sprintf("%s:%d", addr, port), func(config *constant.ClientConfig) {
		*config = getNacosClientConfig(ctx, configPath) // 覆盖 config 指向的结构体内容
	}).SetGroupName(env))
}

func getNacosAdapter(ctx context.Context, configPath string) (adapter gcfg.Adapter, err error) {
	var (
		serverConfig = constant.ServerConfig{
			IpAddr: g.Cfg(configPath).MustGet(ctx, "nacos.address").String(),
			Port:   g.Cfg(configPath).MustGet(ctx, "nacos.port").Uint64(),
		}
		clientConfig = getNacosClientConfig(ctx, configPath)
		configParam  = vo.ConfigParam{
			DataId: "user-config.yaml",
			Group:  env,
		}
	)

	// 在 getNacosAdapter 中添加
	if err := os.MkdirAll(clientConfig.LogDir, 0755); err != nil {
		g.Log().Fatalf(ctx, "Failed to create nocos log dir %s: %v", clientConfig.LogDir, err)
	}
	if err := os.MkdirAll(clientConfig.CacheDir, 0755); err != nil {
		g.Log().Fatalf(ctx, "Failed to create nocos cache dir %s: %v", clientConfig.CacheDir, err)
	}

	iClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{serverConfig},
		"clientConfig":  clientConfig,
	})

	config, _ := iClient.GetConfig(configParam)
	fmt.Println(config)

	// Create anacosClient that implements gcfg.Adapter.
	client, err := nacos.New(ctx, nacos.Config{
		ServerConfigs: []constant.ServerConfig{serverConfig},
		ClientConfig:  clientConfig,
		ConfigParam:   configParam,
	})

	return client, err
}
