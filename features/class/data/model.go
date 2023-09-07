package data

import (
	"kelompok1/immersive-dash/features/class"

	"gorm.io/gorm"
)

// struct Class gorm model
type Class struct {
	gorm.Model
	UserId    uint   `gorm:"column:user_id;not null"`
	ClassName string `gorm:"column:class_name;not null"`
}

func CoreToModel(dataCore class.CoreClass) Class {
	return Class{
		Model:     gorm.Model{},
		UserId:    dataCore.UserId,
		ClassName: dataCore.ClassName,
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Class) class.CoreClass {
	return class.CoreClass{
		ID:        dataModel.ID,
		UserId:    dataModel.UserId,
		ClassName: dataModel.ClassName,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		DeletedAt: dataModel.DeletedAt.Time,
	}
}

func ListModelToCore(dataModel []Class) []class.CoreClass {
	var result []class.CoreClass
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
