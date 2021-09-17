package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/gin-admin/gin-admin-tpl/internal/app/contextx"
	"github.com/gin-admin/gin-admin-tpl/internal/app/ginx"
	"github.com/gin-admin/gin-admin-tpl/internal/app/schema"
	"github.com/gin-admin/gin-admin-tpl/internal/app/service"
)

var DemoSet = wire.NewSet(wire.Struct(new(DemoAPI), "*"))

type DemoAPI struct {
	DemoSrv *service.DemoSrv
}

func (a *DemoAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DemoQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.DemoSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *DemoAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DemoSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *DemoAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Demo
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = contextx.FromUserID(ctx)
	result, err := a.DemoSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *DemoAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Demo
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.DemoSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *DemoAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DemoSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *DemoAPI) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DemoSrv.UpdateStatus(ctx, ginx.ParseParamID(c, "id"), 1)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *DemoAPI) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DemoSrv.UpdateStatus(ctx, ginx.ParseParamID(c, "id"), 2)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
