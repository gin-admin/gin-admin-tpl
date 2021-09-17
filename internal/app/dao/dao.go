package dao

import (
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/gin-admin/gin-admin-tpl/internal/app/config"
	"github.com/gin-admin/gin-admin-tpl/internal/app/dao/demo"
	"github.com/gin-admin/gin-admin-tpl/internal/app/dao/util"
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	util.TransSet,
	demo.DemoSet,
) // end

// Define repo type alias
type (
	TransRepo = util.Trans
	DemoRepo  = demo.DemoRepo
) // end

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(demo.Demo),
	) // end
}
