package framework

type IGroup interface {
	//实现httpMethod
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)
	//实现嵌套group
	Group(string) IGroup

	// 嵌套中间件
	Use(middlewares ...ControllerHandler)
}

type Group struct {
	core   *Core
	parent *Group
	prefix string

	middlewares []ControllerHandler // 存放中间件
}

//初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandler := append(g.middlewares, handlers...)
	g.core.Get(uri, allHandler...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandler := append(g.middlewares, handlers...)
	g.core.Post(uri, allHandler...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandler := append(g.middlewares, handlers...)
	g.core.Put(uri, allHandler...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandler := append(g.middlewares, handlers...)
	g.core.Delete(uri, allHandler...)
}

func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.prefix + g.prefix
}

func (g *Group) Use(middleware ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middleware...)
}
