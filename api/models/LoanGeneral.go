package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanGeneral struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32     `gorm:"type:int;not null" json:"user_id"`
	Users     User       `json:"user"`
	Title     string     `gorm:"size:255;not null" json:"title"`
	Amount    float64    `gorm:"type:decimal;not null" json:"amount"`
	Datetime  time.Time  `gorm:"type:timestamp;not null" json:"datetime"`
	Tenor     int        `gorm:"type:int;not null" json:"tenor"`
	LoanTypeID uint32    `gorm:"type:int;not null" json:"loan_type_id"`
	LoanTypes LoanType   `json:"loan_type"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (lg *LoanGeneral) Prepare() {
	lg.ID = 0
	lg.Title = html.EscapeString(strings.TrimSpace(lg.Title))
	lg.CreatedAt = time.Now()
	lg.UpdatedAt = time.Now()
}

//func (lg *LoanGeneral) Validate() error {
//	if p.Title == "" {
//		return errors.New("Required Title")
//	}
//	if p.Content == "" {
//		return errors.New("Required Content")
//	}
//	if p.AuthorID < 1 {
//		return errors.New("Required Author")
//	}
//	return nil
//}

func (lg *LoanGeneral) SaveLoanGeneral(db *gorm.DB) (*LoanGeneral, error) {
	var err error
	err = db.Debug().Model(&LoanGeneral{}).Create(&lg).Error
	if err != nil {
		return &LoanGeneral{}, err
	}
	if lg.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", lg.UserID).Take(&lg.Users).Error
		if err != nil {
			return &LoanGeneral{}, err
		}
		err = db.Debug().Model(&LoanType{}).Where("id = ?", lg.LoanTypeID).Take(&lg.LoanTypes).Error
		if err != nil {
			return &LoanGeneral{}, err
		}
	}
	return lg, nil
}

func (lg *LoanGeneral) FindAllLoanGenerals(db *gorm.DB) (*[]LoanGeneral, error) {
	var err error
	lGeneral := make([]LoanGeneral,0)
	err = db.Debug().Model(&LoanGeneral{}).Limit(paginationLimit).Find(&lGeneral).Error
	if err != nil {
		return &[]LoanGeneral{}, err
	}
	return &lGeneral, nil
}

func (lg *LoanGeneral) FindLoanGeneralByID(db *gorm.DB, lid uint32) (*LoanGeneral, error) {
	var err error
	err = db.Debug().Model(&LoanGeneral{}).Where("id = ?", lid).Take(&lg).Error
	if err != nil {
		return &LoanGeneral{}, err
	}
	if lg.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", lg.UserID).Take(&lg.Users).Error
		if err != nil {
			return &LoanGeneral{}, err
		}
		err = db.Debug().Model(&LoanType{}).Where("id = ?", lg.LoanTypeID).Take(&lg.LoanTypes).Error
		if err != nil {
			return &LoanGeneral{}, err
		}
	}
	return lg, nil
}

//func (lg *LoanGeneral) UpdateALoanGeneral(db *gorm.DB) (*LoanGeneral, error) {
//
//	var err error
//	// db = db.Debug().Model(&LoanGeneral{}).Where("id = ?", pid).Take(&LoanGeneral{}).UpdateColumns(
//	// 	map[string]interface{}{
//	// 		"title":      p.Title,
//	// 		"content":    p.Content,
//	// 		"updated_at": time.Now(),
//	// 	},
//	// )
//	// err = db.Debug().Model(&LoanGeneral{}).Where("id = ?", pid).Take(&p).Error
//	// if err != nil {
//	// 	return &LoanGeneral{}, err
//	// }
//	// if p.ID != 0 {
//	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
//	// 	if err != nil {
//	// 		return &LoanGeneral{}, err
//	// 	}
//	// }
//	err = db.Debug().Model(&LoanGeneral{}).Where("id = ?", p.ID).Updates(LoanGeneral{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
//	if err != nil {
//		return &LoanGeneral{}, err
//	}
//	if p.ID != 0 {
//		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
//		if err != nil {
//			return &LoanGeneral{}, err
//		}
//	}
//	return p, nil
//}

func (lg *LoanGeneral) DeleteALoanGeneral(db *gorm.DB, lid uint32) error {
	db = db.Debug().Model(&LoanGeneral{}).Where("id = ?", lid).Take(&LoanGeneral{}).Delete(&LoanGeneral{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("LoanGeneral not found")
		}
		return db.Error
	}
	return nil
}
