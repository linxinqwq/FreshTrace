package consts

import "time"

const (
	ConfigFilePath = "config/config.yaml"            // 配置文件所在目录
	LogFilePath    = "internal/logs/application.log" // 日志文件存放路径

	// TokenSecret 定义加密的盐,生成token和解析的时候都需要用到，使用同一个
	TokenSecret         = "ChainLogix"   // TokenSecret
	TokenExpireDuration = time.Hour * 24 //  定义token过期时间 一天过期

	EmailCodeExpireDuration = time.Minute * 5     // 邮箱验证码过期时间
	MysqlFilePath           = "internal/database" // mysql生成文件存放路径

	FiscoBcosFilePath = "config/config.toml" // fsico bcos配置文件

)
