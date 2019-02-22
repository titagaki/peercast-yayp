package repositoriy

import (
	"time"

	"github.com/jinzhu/gorm"

	"peercast-yayp/model"
)

type ChannelLogRepository struct {
	*gorm.DB
}

func NewChannelLogRepository(db *gorm.DB) *ChannelLogRepository {
	return &ChannelLogRepository{db}
}

func (db *ChannelLogRepository) CreateChannelLogs(logTime time.Time, channels model.ChannelList) {
	for _, c := range channels {
		log := model.ChannelLog{
			LogTime:         logTime,
			ChannelID:       c.ID,
			GnuID:           c.GnuID,
			Name:            c.Name,
			Bitrate:         c.Bitrate,
			ContentType:     c.ContentType,
			Listeners:       c.Listeners,
			Relays:          c.Relays,
			Age:             c.Age,
			Genre:           c.Genre,
			Description:     c.Description,
			Url:             c.Url,
			Comment:         c.Comment,
			TrackArtist:     c.TrackArtist,
			TrackTitle:      c.TrackTitle,
			TrackAlbum:      c.TrackAlbum,
			TrackGenre:      c.TrackGenre,
			TrackContact:    c.TrackContact,
			HiddenListeners: c.HiddenListeners,
		}
		db.Create(&log)
	}
}

func (db *ChannelLogRepository) FindChannelDailyLogByName(name string) []*model.ChannelDailyLog {
	logs := make([]*model.ChannelDailyLog, 0)
	db.Where("name = ?", name).Find(&logs)

	return logs
}

func (db *ChannelLogRepository) FindChannelLogsByNameAndLogTime(name string, logTime time.Time) []*model.ChannelLog {
	start := logTime.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	logs := make([]*model.ChannelLog, 0)
	db.Where("name = ? and log_time => ? and log_time < ?", name, start, end).Find(&logs)

	return logs
}
