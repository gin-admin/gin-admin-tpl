package demo

import (
	"context"

	"gorm.io/gorm"

	"github.com/gin-admin/gin-admin-tpl/internal/app/dao/util"
	"github.com/gin-admin/gin-admin-tpl/internal/app/schema"
	"github.com/gin-admin/gin-admin-tpl/pkg/util/structure"
)

// GetDemoDB Get Demo db model
func GetDemoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Demo))
}

// SchemaDemo Demo schema
type SchemaDemo schema.Demo

// ToDemo Convert to Demo entity
func (a SchemaDemo) ToDemo() *Demo {
	item := new(Demo)
	structure.Copy(a, item)
	return item
}

// Demo Demo entity
type Demo struct {
	util.Model
	Code    string `gorm:"size:50;index;"`                // 编号
	Name    string `gorm:"size:50;index;"`                // 名称
	Memo    string `gorm:"size:1024;"`                    // 备注
	Status  int    `gorm:"type:tinyint;index;default:0;"` // 状态(1:启用 2:停用)
	Creator uint64 `gorm:""`                              // 创建者

}

// ToSchemaDemo Convert to Demo schema
func (a Demo) ToSchemaDemo() *schema.Demo {
	item := new(schema.Demo)
	structure.Copy(a, item)
	return item
}

// Demos Demo entity list
type Demos []*Demo

// ToSchemaDemos Convert to Demo schema list
func (a Demos) ToSchemaDemos() []*schema.Demo {
	list := make([]*schema.Demo, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDemo()
	}
	return list
}
