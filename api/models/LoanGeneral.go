package models

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"log"
	"os"
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
	Status    int        `gorm:"type:int;not null" json:"status"`
	LoanTypeID uint32    `gorm:"type:int;not null" json:"loan_type_id"`
	LoanTypes LoanType   `json:"loan_type"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type LoanGeneralRaw struct {
	LoanTypeName string
	UserName  string
	ID        uint32
	UserID    uint32
	Title     string
	Amount    float64
	Datetime  time.Time
	Tenor     int
	Status    int
	LoanTypeID uint32
	CreatedAt time.Time
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

func (lg *LoanGeneral) FindAllLoanGenerals(db *gorm.DB) (*[]LoanGeneralRaw, error) {
	DbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	rdb, err := sql.Open("mysql", DbURL)
	if err != nil {
		log.Printf("error while opening mysql DB: %v", err)
		return nil, err
	}

	rows, err := rdb.Query(`SELECT lt.name, u.name, lg.id, lg.user_id, lg.title, lg.amount, lg.datetime, lg.tenor, lg.status, 
	lg.loan_type_id, lg.created_at FROM loan_generals lg 
	LEFT JOIN loan_types lt ON lt.id = lg.loan_type_id LEFT JOIN users u ON u.id = lg.user_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lGeneral := make([]LoanGeneralRaw, 0)

	for rows.Next() {
		var (
			typeName sql.NullString
			userName sql.NullString
			id uint32
			userID uint32
			title string
			amount float64
			datetime time.Time
			tenor int
			status int
			loanTypeID uint32
			createdAt time.Time
		)

		if err := rows.Scan(&typeName, &userName, &id, &userID, &title, &amount, &datetime, &tenor, &status, &loanTypeID, &createdAt); err != nil {
			return nil, err
		}

		lgd := LoanGeneralRaw{
			LoanTypeName: typeName.String,
			UserName:     userName.String,
			ID:           id,
			UserID:       userID,
			Title:        title,
			Amount:       amount,
			Datetime:     datetime,
			Tenor:        tenor,
			Status:       status,
			LoanTypeID:   loanTypeID,
			CreatedAt:    createdAt,
		}
		lGeneral = append(lGeneral, lgd)

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
