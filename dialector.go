package duckgorm

import (
	_ "github.com/marcboeker/go-duckdb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

const (
	driverName = "duckdb"
)

var _ gorm.Dialector = (*Dialector)(nil)

type Dialector struct {
	*Config
}

type Config struct {
	DriverName string
	DSN        string
	Conn       gorm.ConnPool
}

func (d *Dialector) Name() string {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) Initialize(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) Migrator(db *gorm.DB) gorm.Migrator {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) DataTypeOf(field *schema.Field) string {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) DefaultValueOf(field *schema.Field) clause.Expression {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) QuoteTo(writer clause.Writer, s string) {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) Explain(sql string, vars ...interface{}) string {
	//TODO implement me
	panic("implement me")
}
