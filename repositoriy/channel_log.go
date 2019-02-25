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

func (r *ChannelLogRepository) CreateChannelLogs(logTime time.Time, channels model.ChannelList) {
	for _, c := range channels {
		log := model.ChannelLog{
			LogTime:         logTime,
			ChannelID:       c.ID,
			CID:             c.CID,
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
		r.DB.Create(&log)
	}
}

func (r *ChannelLogRepository) FindChannelLogsByNameAndLogTime(name string, logTime time.Time) []*model.ChannelLog {
	start := time.Date(logTime.Year(), logTime.Month(), logTime.Day(), 0, 0, 0, 0, time.Local)
	end := start.Add(24 * time.Hour)

	logs := make([]*model.ChannelLog, 0)
	r.DB.Where("name = ? and log_time >= ? and log_time < ?", name, start, end).Find(&logs)
	return logs
}

func (r *ChannelLogRepository) FindChannelLogsByLogTime(logTime time.Time) []*model.ChannelLog {
	start := time.Date(logTime.Year(), logTime.Month(), logTime.Day(), 0, 0, 0, 0, time.Local)
	end := start.Add(24 * time.Hour)

	logs := make([]*model.ChannelLog, 0)
	r.DB.Where("log_time >= ? and log_time < ?", start, end).
		Order("name asc, log_time asc").Find(&logs)
	return logs
}
