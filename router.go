package main

import "jamesluo1/framework"

func registerRouter(core *framework.Core) {
	core.Get("/foo", FooControllerHandler)
}
