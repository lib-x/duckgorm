package duckgorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/migrator"
	"gorm.io/gorm/schema"
)

var _ gorm.Migrator = (*Migrator)(nil)

type Migrator struct {
	migrator.Migrator
}

func (m Migrator) AutoMigrate(dst ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) CurrentDatabase() string {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) FullDataTypeOf(field *schema.Field) clause.Expr {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) GetTypeAliases(databaseTypeName string) []string {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) CreateTable(dst ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) DropTable(dst ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) HasTable(dst interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) RenameTable(oldName, newName interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) GetTables() (tableList []string, err error) {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) TableType(dst interface{}) (gorm.TableType, error) {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) AddColumn(dst interface{}, field string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) DropColumn(dst interface{}, field string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) AlterColumn(dst interface{}, field string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) MigrateColumn(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) MigrateColumnUnique(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) HasColumn(dst interface{}, field string) bool {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) RenameColumn(dst interface{}, oldName, field string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) ColumnTypes(dst interface{}) ([]gorm.ColumnType, error) {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) CreateView(name string, option gorm.ViewOption) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) DropView(name string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) CreateConstraint(dst interface{}, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) DropConstraint(dst interface{}, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) HasConstraint(dst interface{}, name string) bool {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) CreateIndex(dst interface{}, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) DropIndex(dst interface{}, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) HasIndex(dst interface{}, name string) bool {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) RenameIndex(dst interface{}, oldName, newName string) error {
	//TODO implement me
	panic("implement me")
}

func (m Migrator) GetIndexes(dst interface{}) ([]gorm.Index, error) {
	//TODO implement me
	panic("implement me")
}
