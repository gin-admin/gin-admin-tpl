package demo

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/gin-admin/gin-admin-tpl/internal/app/dao/util"
	"github.com/gin-admin/gin-admin-tpl/internal/app/schema"
	"github.com/gin-admin/gin-admin-tpl/pkg/errors"
)

// DemoSet Injection wire
var DemoSet = wire.NewSet(wire.Struct(new(DemoRepo), "*"))

type DemoRepo struct {
	DB *gorm.DB
}

func (a *DemoRepo) getQueryOption(opts ...schema.DemoQueryOptions) schema.DemoQueryOptions {
	var opt schema.DemoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *DemoRepo) Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetDemoDB(ctx, a.DB)
	// TODO: 查询条件

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(util.ParseOrder(opt.OrderFields))

	var list Demos
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.DemoQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDemos(),
	}

	return qr, nil
}

func (a *DemoRepo) Get(ctx context.Context, id uint64, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {
	var item Demo
	ok, err := util.FindOne(ctx, GetDemoDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDemo(), nil
}

func (a *DemoRepo) Create(ctx context.Context, item schema.Demo) error {
	eitem := SchemaDemo(item).ToDemo()
	result := GetDemoDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *DemoRepo) Update(ctx context.Context, id uint64, item schema.Demo) error {
	eitem := SchemaDemo(item).ToDemo()
	result := GetDemoDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *DemoRepo) Delete(ctx context.Context, id uint64) error {
	result := GetDemoDB(ctx, a.DB).Where("id=?", id).Delete(Demo{})
	return errors.WithStack(result.Error)
}

func (a *DemoRepo) UpdateStatus(ctx context.Context, id uint64, status int) error {
	result := GetDemoDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}
