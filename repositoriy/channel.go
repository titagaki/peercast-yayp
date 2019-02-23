package repositoriy

import (
	"github.com/jinzhu/gorm"

	"peercast-yayp/model"
)

type ChannelRepository struct {
	*gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{db}
}

func (r *ChannelRepository) SaveOrCreate(channel *model.Channel) {
	if r.DB.NewRecord(channel) {
		r.DB.Create(channel)
	} else {
		r.DB.Save(channel)
	}
}

func (r *ChannelRepository) FindPlayingChannels() model.ChannelList {
	channel := make([]*model.Channel, 0)
	r.DB.Where("is_playing = ?", true).Find(&channel)

	return channel
}
