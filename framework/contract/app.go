package contract

// AppKey定义字符串凭证
const AppKey = "hade:app"

type App interface {
	// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
	AppID() string
	//Version 定义当前版本
	Version() string
	//BaseFolder 定义项目基础地址
	BaseFolder() string
	//ConfigFolder 定义了配置文件地址
	ConfigFolder() string
	//LogFolder 定义了日志所在路径
	LogFolder() string
	//ProviderFolder 定义服务提供者所在路径
	ProviderFolder() string
	//MiddlewareFolder 定义中间件
	MiddlewareFolder() string
	//CommandFolder 定义业务定义的命令
	CommandFolder() string
	//RuntimeFolder 定义业务的运行中间态
	RuntimeFolder() string
	//TestFolder 存放测试所需要的信息
	TestFolder() string
	// HttpFolder 存放的文件
	HttpFolder() string
}
