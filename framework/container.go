package framework

import (
	"errors"
	"sync"
)

// Container是一个服务容器,提供绑定和获取服务的功能
type Container interface {
	// Bind绑定一个服务提供者,如果关键字凭证已经存在,会进行替换操作,返回error
	Bind(provider ServiceProvider) error
	//IsBind关键字是否已经绑定服务提供者
	IsBind(key string) bool

	//Make根据一个关键字凭证获取一个服务
	Make(key string) (interface{}, error)
	//MustMake 根据关键字凭证获取服务,如果这个关键字凭证未绑定服务提供者,那么会panic
	//所以使用这个接口的时候必须保证服务容器已经为这个关键字凭证绑定了服务提供者
	MustMake(key string) interface{}
	//MakeNew根据关键字获取一个服务,只是这个服务并不是单例模式的
	//他是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	//这个函数在需要为不同的参数启动不同的实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// HadeContainer是服务容器的具体实现
type HadeContainer struct {
	Container //强制要求HadeContainer实现Container接口
	// providers 储存注册的服务提供者, key为字符串凭证
	providers map[string]ServiceProvider
	// instance 储存具体的实例, key为字符串凭证
	instances map[string]interface{}
	//lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewHadeContainer() *HadeContainer {
	return &HadeContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// 将服务容器和关键字进行绑定
func (hade *HadeContainer) Bind(provider ServiceProvider) error {

	key := provider.Name()
	hade.providers[key] = provider

	//注册就要实例化
	if provider.IsDefer() == false {
		if err := provider.Boot(hade); err != nil {
			return err
		}
		//实例化方法
		params := provider.Params(hade)
		method := provider.Register(hade)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		hade.lock.Lock()
		hade.instances[key] = instance
		hade.lock.Unlock()
	}
	return nil
}

func (hade *HadeContainer) IsBind(key string) bool {
	return hade.findServiceProvider(key) != nil
}

// Make方式调用内部的make实现
func (hade *HadeContainer) Make(key string) (interface{}, error) {
	return hade.make(key, nil, false)
}

func (hade *HadeContainer) MustMake(key string) interface{} {
	seve, err := hade.make(key, nil, true)
	if err != nil {
		panic(err)
	}
	return seve
}

// MakeNew 方式使用的内部的make初始化
func (hade *HadeContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return hade.make(key, params, true)
}

// 真正的实例化一个服务
func (hade *HadeContainer) make(key string, params []interface{}, foreNew bool) (interface{}, error) {
	hade.lock.RLock()
	defer hade.lock.RUnlock()

	//查询是否已注册过这个服务提供者,如果没有注册，则返回错误
	sp := hade.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if ins, ok := hade.instances[key]; ok {
		return ins, nil
	}

	if foreNew {
		return hade.newInstance(sp, params)
	}

	//容器中若还未实例化,则进行实例化
	inst, err := hade.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	hade.instances[key] = inst
	return inst, nil
}

func (hade *HadeContainer) findServiceProvider(key string) ServiceProvider {
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	if sp, ok := hade.providers[key]; ok {
		return sp
	}
	return nil
}

func (hade *HadeContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	//fore new
	if err := sp.Boot(hade); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(hade)
	}
	method := sp.Register(hade)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, nil
}
