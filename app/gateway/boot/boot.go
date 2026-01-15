package boot

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

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
)

var (
	configPath      = gfile.Pwd() + "/config/config.yaml"
	localConfigPath = gfile.Pwd() + "/app/gateway/manifest/config/config.yaml"
	env             = os.Getenv("APP_ENV")
)

// envFirstString 优先从环境变量读取（非空白），若不存在则从 configPath 对应的配置文件读取。
func envFirstString(ctx context.Context, envKey, configPath, configKey string) string {
	if v, ok := os.LookupEnv(envKey); ok {
		v = strings.TrimSpace(v)
		if v != "" {
			return v
		}
	}
	return g.Cfg(configPath).MustGet(ctx, configKey).String()
}

// envFirstInt 优先从环境变量读取（非空白），若不存在则从 configPath 对应的配置文件读取。
func envFirstInt(ctx context.Context, envKey, configPath, configKey string) int {
	if v, ok := os.LookupEnv(envKey); ok {
		v = strings.TrimSpace(v)
		if v != "" {
			i, err := strconv.Atoi(v)
			if err == nil {
				return i
			}
		}
	}
	return g.Cfg(configPath).MustGet(ctx, configKey).Int()
}

// envFirstUint64 优先从环境变量读取（非空白），若不存在则从 configPath 对应的配置文件读取。
func envFirstUint64(ctx context.Context, envKey, configPath, configKey string) uint64 {
	if v, ok := os.LookupEnv(envKey); ok {
		v = strings.TrimSpace(v)
		if v != "" {
			u, err := strconv.ParseUint(v, 10, 64)
			if err == nil {
				return u
			}
		}
	}
	return g.Cfg(configPath).MustGet(ctx, configKey).Uint64()
}

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
		CacheDir:            envFirstString(ctx, "NACOS_CACHE_DIR", configPath, "nacos.cacheDir"),
		LogDir:              envFirstString(ctx, "NACOS_LOG_DIR", configPath, "nacos.logDir"),
		NamespaceId:         envFirstString(ctx, "NACOS_NAMESPACE_ID", configPath, "nacos.namespaceId"),
		Username:            envFirstString(ctx, "NACOS_USERNAME", configPath, "nacos.username"),
		Password:            envFirstString(ctx, "NACOS_PASSWORD", configPath, "nacos.password"),
		NotLoadCacheAtStart: true,
		DisableUseSnapShot:  true,
		TimeoutMs:           15000, // 增加超时到 15 秒
	}
}

func getNacosRegister(ctx context.Context, configPath string) {
	addr := envFirstString(ctx, "NACOS_ADDRESS", configPath, "nacos.address")
	port := envFirstInt(ctx, "NACOS_PORT", configPath, "nacos.port")
	grpcx.Resolver.Register(rnacos.New(fmt.Sprintf("%s:%d", addr, port), func(config *constant.ClientConfig) {
		*config = getNacosClientConfig(ctx, configPath) // 覆盖 config 指向的结构体内容
	}).SetGroupName(env))
}

func getNacosAdapter(ctx context.Context, configPath string) (adapter gcfg.Adapter, err error) {
	var (
		serverConfig = constant.ServerConfig{
			IpAddr: envFirstString(ctx, "NACOS_ADDRESS", configPath, "nacos.address"),
			Port:   envFirstUint64(ctx, "NACOS_PORT", configPath, "nacos.port"),
		}
		clientConfig = getNacosClientConfig(ctx, configPath)
		configParam  = vo.ConfigParam{
			DataId: "gateway-config.yaml",
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
	fmt.Println("配置：", config)

	// Create anacosClient that implements gcfg.Adapter.
	client, err := nacos.New(ctx, nacos.Config{
		ServerConfigs: []constant.ServerConfig{serverConfig},
		ClientConfig:  clientConfig,
		ConfigParam:   configParam,
	})

	return client, err
}
