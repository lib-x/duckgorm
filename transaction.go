package duckgorm

import "gorm.io/gorm"

var _ gorm.SavePointerDialectorInterface = (*Dialector)(nil)

func (d *Dialector) SavePoint(tx *gorm.DB, name string) error {
	//TODO implement me
	panic("implement me")
}

func (d *Dialector) RollbackTo(tx *gorm.DB, name string) error {
	//TODO implement me
	panic("implement me")
}
