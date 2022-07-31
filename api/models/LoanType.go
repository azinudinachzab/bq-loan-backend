package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type LoanType struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name  string `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (lt *LoanType) Prepare() {
	lt.ID = 0
	lt.Name = html.EscapeString(strings.TrimSpace(lt.Name))
	lt.CreatedAt = time.Now()
	lt.UpdatedAt = time.Now()
}

func (lt *LoanType) SaveLoanType(db *gorm.DB) (*LoanType, error) {
	var err error
	err = db.Debug().Create(&lt).Error
	if err != nil {
		return &LoanType{}, err
	}
	return lt, nil
}

func (lt *LoanType) FindAllLoanTypes(db *gorm.DB) (*[]LoanType, error) {
	var err error
	loanTypes := make([]LoanType, 0)
	err = db.Debug().Model(&LoanType{}).Limit(paginationLimit).Find(&loanTypes).Error
	if err != nil {
		return &[]LoanType{}, err
	}
	return &loanTypes, err
}

func (lt *LoanType) FindLoanTypeByID(db *gorm.DB, uid uint32) (*LoanType, error) {
	var err error
	err = db.Debug().Model(LoanType{}).Where("id = ?", uid).Take(&lt).Error
	if err != nil {
		return &LoanType{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &LoanType{}, errors.New("loan type not Found")
	}
	return lt, err
}

func (lt *LoanType) UpdateALoanType(db *gorm.DB, uid uint32) (*LoanType, error) {
	db = db.Debug().Model(&LoanType{}).Where("id = ?", uid).Take(&LoanType{}).UpdateColumns(
		map[string]interface{}{
			"name":      lt.Name,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &LoanType{}, db.Error
	}
	// This is the display the updated loan type
	err := db.Debug().Model(&LoanType{}).Where("id = ?", uid).Take(&lt).Error
	if err != nil {
		return &LoanType{}, err
	}
	return lt, nil
}

func (lt *LoanType) DeleteALoanType(db *gorm.DB, uid uint32) error {
	db = db.Debug().Model(&LoanType{}).Where("id = ?", uid).Take(&LoanType{}).Delete(&LoanType{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
