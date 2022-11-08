package app

import (
	"errors"
	"flag"
	"github.com/jamesluo111/core_web/framework"
	"github.com/jamesluo111/core_web/framework/util"
	"path/filepath"
)

type HadeApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
}

// NewHadeApp 初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &HadeApp{baseFolder: baseFolder, container: container}, nil
}

//获取基础地址
func (h *HadeApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}
	//如果没有则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}
	return util.GetExecDirectory()
}

//获取日志存放目录
func (h *HadeApp) LogFolder() string {
	return filepath.Join(h.baseFolder, "log")
}

func (h *HadeApp) StorageFolder() string {
	return filepath.Join(h.baseFolder, "storage")
}

func (h *HadeApp) Version() string {
	return "0.0.1"
}

func (h *HadeApp) ConfigFolder() string {
	return filepath.Join(h.baseFolder, "config")
}

func (h *HadeApp) HttpFolder() string {
	return filepath.Join(h.baseFolder, "http")
}

func (h *HadeApp) ProviderFolder() string {
	return filepath.Join(h.baseFolder, "provider")
}

func (h *HadeApp) MiddlewareFolder() string {
	return filepath.Join(h.baseFolder, "middleware")
}

func (h *HadeApp) CommandFolder() string {
	return filepath.Join(h.baseFolder, "command")
}

func (h *HadeApp) RuntimeFolder() string {
	return filepath.Join(h.baseFolder, "runtime")
}

func (h *HadeApp) TestFolder() string {
	return filepath.Join(h.baseFolder, "test")
}
