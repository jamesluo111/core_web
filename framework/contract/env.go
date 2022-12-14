package contract

const (
	// EnvProduction 代表生产环境
	EnvProduction = "production"
	// EnvTesting 代表测试环境
	EnvTesting = "testing"
	// EnvDevelopment 代表开发环境
	EnvDevelopment = "development"
	// EnvKey 是环境变量服务字符串凭证
	EnvKey = "hade:env"
)

// Env 定义环境变量服务
type Env interface {
	// AppEnv 获取当前环境
	AppEnv() string
	// IsExist 判断一个环境变量是否有被设置
	IsExist(key string) bool
	// Get 获取某个环境变量, 如果没有返回
	Get(key string) string
	// All 获取所有环境变量 .env 和运行环境变量融合后结果
	All() map[string]string
}
