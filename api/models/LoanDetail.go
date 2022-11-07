package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type LoanDetail struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	LoanGeneralID uint32 `gorm:"type:int;not null" json:"loan_general_id"`
	General LoanGeneral
	Amount float64       `gorm:"type:decimal;not null" json:"amount"`
	Datetime  time.Time  `gorm:"type:timestamp;not null" json:"datetime"`
	Status    int        `gorm:"type:int;not null" json:"status"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ld *LoanDetail) Prepare() {
	ld.ID = 0
	ld.CreatedAt = time.Now()
	ld.UpdatedAt = time.Now()
}

func (ld *LoanDetail) SaveLoanDetail(db *gorm.DB) (*LoanDetail, error) {
	var err error
	err = db.Debug().Model(&LoanDetail{}).Create(&ld).Error
	if err != nil {
		return &LoanDetail{}, err
	}
	if ld.ID != 0 {
		err = db.Debug().Model(&LoanGeneral{}).Where("id = ?", ld.LoanGeneralID).Take(&ld.General).Error
		if err != nil {
			return &LoanDetail{}, err
		}
	}
	return ld, nil
}

func (ld *LoanDetail) FindAllLoanDetails(db *gorm.DB) (*[]LoanDetail, error) {
	var err error
	loanDetails := make([]LoanDetail, 0)
	err = db.Debug().Model(&LoanDetail{}).Limit(paginationLimit).Find(&loanDetails).Error
	if err != nil {
		return &[]LoanDetail{}, err
	}
	return &loanDetails, err
}

func (ld *LoanDetail) FindLoanDetailByID(db *gorm.DB, lid uint32) (*LoanDetail, error) {
	var err error
	err = db.Debug().Model(&LoanDetail{}).Where("id = ?", lid).Take(&ld).Error
	if err != nil {
		return &LoanDetail{}, err
	}
	if ld.ID != 0 {
		err = db.Debug().Model(&LoanGeneral{}).Where("id = ?", ld.LoanGeneralID).Take(&ld.General).Error
		if err != nil {
			return &LoanDetail{}, err
		}
	}
	return ld, nil
}

func (ld *LoanDetail) FindLoanDetailByLoanGeneralID(db *gorm.DB, lid uint32) (*[]LoanDetail, error) {
	var err error
	loanDetails := make([]LoanDetail, 0)
	err = db.Debug().Model(&LoanDetail{}).Where("loan_general_id = ?", lid).Find(&loanDetails).Error
	if err != nil {
		return &[]LoanDetail{}, err
	}
	return &loanDetails, err
}

func (ld *LoanDetail) UpdateALoanDetail(db *gorm.DB, uid uint32) (*LoanDetail, error) {
	db = db.Debug().Model(&LoanDetail{}).Where("id = ?", uid).Take(&LoanDetail{}).UpdateColumns(
		map[string]interface{}{
			"amount":      ld.Amount,
			"datetime":    ld.Datetime,
			"status":      ld.Status,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &LoanDetail{}, db.Error
	}
	// This is the display the updated data
	err := db.Debug().Model(&LoanDetail{}).Where("id = ?", uid).Take(&ld).Error
	if err != nil {
		return &LoanDetail{}, err
	}
	return ld, nil
}

func (ld *LoanDetail) DeleteALoanDetail(db *gorm.DB, lid uint32) error {
	db = db.Debug().Model(&LoanDetail{}).Where("id = ?", lid).Take(&LoanDetail{}).Delete(&LoanDetail{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("LoanDetail not found")
		}
		return db.Error
	}
	return nil
}

func (ld *LoanDetail) BulkSaveLoanDetail(db *gorm.DB, ald []LoanDetail) error {
	for _, val := range ald {
		err := db.Debug().Model(&LoanDetail{}).Create(&val).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (ld *LoanDetail) UpdateStatus(db *gorm.DB, uid uint32, status int) error {
	err := db.Debug().Model(&LoanDetail{}).Where("id = ?", uid).Take(&LoanDetail{}).UpdateColumns(
		map[string]interface{}{
			"status": status,
			"updated_at": time.Now(),
		},
	).Error

	return err
}
