package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Participant struct {
	ID                      uint32    `gorm:"primary_key;auto_increment" json:"id"`
	HoursSpentOnSocialDaily uint8     `json:"hoursSpentDailyOnSocialMedia"`
	Age                     uint8     `json:"age"`
	Gender                  string    `gorm:"size:1" json:"gender"`
	CountryCode             string    `gorm:"size:2" json:"countryCode"`
	CreatedAt               time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *Participant) TableName() string {
	return "participants"
}

func (p *Participant) SaveParticipant(db *gorm.DB) (*Participant, error) {
	err := db.Debug().Model(&Participant{}).Create(&p).Error
	if err != nil {
		return &Participant{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Participant{}).Where("id = ?", p.ID).Take(&p.CreatedAt).Error
		if err != nil {
			return &Participant{}, err
		}
	}
	return p, nil
}

func (p *Participant) GetAllParticipants(db *gorm.DB) (*[]Participant, error) {
	participants := []Participant{}
	err := db.Debug().Model(&Participant{}).Limit(10000).Find(&participants).Error
	if err != nil {
		return &[]Participant{}, err
	}
	return &participants, err
}
