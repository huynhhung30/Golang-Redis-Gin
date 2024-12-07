package models

import (
	"time"
	"trinity-app/config"
	"trinity-app/utils/constants"
	"trinity-app/utils/functions"
)

type CouponModel struct {
	Id            int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
	UserId        int    `json:"user_id" gorm:"user_id"`
	Discount      string    `json:"discount" gorm:"discount"`
	IsValid      bool    `json:"is_valid" gorm:"is_valid"`
	Tittle 		  string      `json:"tittle" gorm:"tittle"`
	CountRegistrations int64 `json:"count_registrations" gorm:"count_registrations"`
	DueDate       time.Time `json:"due_date" gorm:"due_date"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
}
type QueryCoupon struct {
	Id            int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
}

func CreateCoupon(body *CouponModel) (rs CouponModel, err error) {
	functions.ShowLog("bodybodybodybodybodyerr", body)
	body.CreatedAt = functions.CurrentTime()
	body.UpdatedAt = functions.CurrentTime()
	body.IsValid = true
	body.CountRegistrations=0
	err = config.DB.Debug().Create(&body).Error

	return rs, err
}


func FindCouponId(id int) (rs CouponModel) {

	config.DB.Debug().
		Table("coupon_models").
		Where("id = ?", id).
		Take(&rs)
		functions.ShowLog("rs",rs)
	return rs
}

func UpdateCountRegistrations(info CouponModel) (rs CouponModel ) {
   config.DB.Debug().Table("coupon_models").Where("id = ?", info.Id).Updates(map[string]interface{}{
	   "count_registrations": info.CountRegistrations - 1,
   })
   couponCurrent := CouponModel{}
   config.DB.Debug().Table("coupon_models").Where("id = ?", info.Id).Take(&couponCurrent)
   functions.ShowLog("=-=-=-=-=-=-cbepppppcheckkcouponCurrentcouponCurrentsk",couponCurrent)
   if couponCurrent.CountRegistrations == 0{
	config.DB.Debug().Table("coupon_models").Where("id = ?", info.Id).Updates(map[string]interface{}{
		"is_valid": false,
	})
   }
   return rs
}


func FindCouponList(params PageLimitQueryModel) (rs []CouponModel, totalCount int64) {
	sortParams := "id asc"
	keywordParams := ""
	if params.Keyword != "" {
		// keywordParams = ""
		keywordParams = "title LIKE" + "'%" + params.Keyword 

	}
	if params.Sort == constants.SORT_PARAMS_DESC {
		sortParams = "id desc"
	}
	config.DB.Table("coupon_models").
		Debug().
		Order(sortParams).
		Where("is_valid=1").
		Where(keywordParams).
		Count(&totalCount).
		Limit(params.Limit).
		Offset((params.Page - 1) * params.Limit).
		Find(&rs)
	return rs, totalCount
}