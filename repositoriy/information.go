package repositoriy

import (
	"github.com/jinzhu/gorm"

	"peercast-yayp/model"
)

type InformationRepository struct {
	*gorm.DB
}

func NewInformationRepository(db *gorm.DB) *InformationRepository {
	return &InformationRepository{db}
}

func (r *InformationRepository) Find() []*model.Information {
	info := make([]*model.Information, 0)
	r.DB.Find(&info)
	return info
}
