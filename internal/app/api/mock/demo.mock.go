package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var DemoSet = wire.NewSet(wire.Struct(new(DemoMock), "*"))

// DemoMock DemoManage
type DemoMock struct{}

// Query 查询数据
// @Tags DemoManage
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param status query int false "状态(1:启用 2:禁用)"
// @Success 200 {object} schema.ListResult{list=[]schema.Demo} "查询结果"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos [get]
func (a *DemoMock) Query(c *gin.Context) {
}

// Get 查询指定数据
// @Tags DemoManage
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.Demo
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos/{id} [get]
func (a *DemoMock) Get(c *gin.Context) {
}

// Create 创建数据
// @Tags DemoManage
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.Demo true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos [post]
func (a *DemoMock) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags DemoManage
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.Demo true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos/{id} [put]
func (a *DemoMock) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags DemoManage
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos/{id} [delete]
func (a *DemoMock) Delete(c *gin.Context) {
}

// Enable 启用数据
// @Tags DemoManage
// @Summary 启用数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos/{id}/enable [patch]
func (a *DemoMock) Enable(c *gin.Context) {
}

// Disable 禁用数据
// @Tags DemoManage
// @Summary 禁用数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos/{id}/disable [patch]
func (a *DemoMock) Disable(c *gin.Context) {
}
