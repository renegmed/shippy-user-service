package go_micro_srv_auth

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	// if err != nil {
	// 	return err
	// }
	return scope.SetColumn("Id", uuid.String())
}
