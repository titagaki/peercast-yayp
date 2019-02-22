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

func (db *ChannelRepository) SaveOrCreate(channel *model.Channel) {
	if db.NewRecord(channel) {
		db.Create(channel)
	} else {
		db.Save(channel)
	}
}

func (db *ChannelRepository) FindPlayingChannels() []*model.Channel {
	channel := make([]*model.Channel, 0)
	db.Where("is_playing = ?", true).Find(&channel)

	return channel
}
