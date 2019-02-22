package repositoriy

import (
	"time"

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

func (db *ChannelRepository) FindChannelDailyLogByName(name string) []*model.ChannelDailyLog {
	logs := make([]*model.ChannelDailyLog, 0)
	db.Where("name = ?", name).Find(&logs)

	return logs
}

func (db *ChannelRepository) FindChannelLogsByNameAndLogTime(name string, logTime time.Time) []*model.ChannelLog {
	start := logTime.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	logs := make([]*model.ChannelLog, 0)
	db.Where("name = ? and log_time => ? and log_time < ?", name, start, end).Find(&logs)

	return logs
}
