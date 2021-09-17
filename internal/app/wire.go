// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"github.com/gin-admin/gin-admin-tpl/internal/app/api"
	"github.com/gin-admin/gin-admin-tpl/internal/app/dao"
	"github.com/gin-admin/gin-admin-tpl/internal/app/router"
	"github.com/gin-admin/gin-admin-tpl/internal/app/service"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
