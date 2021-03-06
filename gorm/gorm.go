package gorm

import (
	"github.com/goxiaoy/go-saas/common"
	"gorm.io/gorm"
)

type DialectFunc func(s string) gorm.Dialector

// Config 配置参数
type Config struct {
	Debug bool
	//Dialect
	Dialect      DialectFunc
	Cfg          *gorm.Config
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

//struct for multi-tenancy entity
type MultiTenancy struct {
	TenantId HasTenant `gorm:"index"`
}

func BuildPage(db *gorm.DB, p common.Pagination) *gorm.DB {
	return db.Offset(p.Offset).Limit(p.Limit)
}
