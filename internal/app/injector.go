package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/gin-admin/gin-admin-tpl/pkg/auth"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine *gin.Engine
	Auth   auth.Auther
}
