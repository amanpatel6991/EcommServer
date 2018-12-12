package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Address struct {
	Id           uint                   `json:"id,omitempty"`
	PhoneNumber  string                 `json:"phone_number,omitempty"`
	AddressLine1 string                 `json:"address_line_1,omitempty"`
	AddressLine2 string                 `json:"address_line_2,omitempty"`
	UserId       uint                   `json:"user_id,omitempty"`
	GoogleUserId uint                   `json:"google_user_id,omitempty"`
	CreatedAt    time.Time              `json:"created_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at,omitempty"`
}

func (Address) TableName() string {
	return "e_address"
}

//Interacting with DB

func GetAddressById(db *gorm.DB, id int) (Address, string) {
	var address Address
	var responseMsg string
	rowsAffected := db.Debug().First(&address, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetAddressById Query Failure !"
	} else {
		responseMsg = "GetAddressById Query Success !"
	}
	return address, responseMsg
}

func GetAllAddresses(db *gorm.DB) ([]Address, string) {
	var addresses []Address
	var responseMsg string

	rowsAffected := db.Debug().Find(&addresses).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetAllAddresses Query Failure !"
	} else {
		responseMsg = "GetAllAddresses Query Success !"
	}
	return addresses, responseMsg
}

func CreateAddress(db *gorm.DB, data Address) (Address, string) {
	var addresses Address
	var responseMsg string
	rowsAffected := db.Debug().Create(&data).RowsAffected
	db.Debug().Preload("Addresses").Last(&addresses)
	if (rowsAffected == 0 ) {
		responseMsg = "CreateAddress Query Failure !"
	} else {
		responseMsg = "CreateAddress Query Success !"
	}
	return addresses, responseMsg
}

//UPDATING OF CHILD IS NOT ALLOWED WHILE UDATING PARENT , CALL CHILD'S UPDATE METHOD EXCLUSIVELY
func UpdateAddressById(db *gorm.DB, data Address, id int) (Address, string) {
	var address Address
	var responseMsg string
	rowsAffected := db.Debug().Model(&Address{}).Where("id=?", id).Update(&data).RowsAffected
	db.Debug().Find(&address, id)
	if (rowsAffected == 0 ) {
		responseMsg = "UpdateAddressById Query Failure !"
	} else {
		responseMsg = "UpdateAddressById Query Success !"
	}
	return address, responseMsg
}

//Deletes without warning even if parent of the CHILD in question exists
func DeleteAddressById(db *gorm.DB, id int) (Address, string) {
	var address Address
	var responseMsg string
	db.Debug().First(&address, id)
	rowsAffected := db.Debug().Delete(&address, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "DeleteAddressById Query Failure !"
	} else {
		responseMsg = "DeleteAddressById Query Success !"
	}
	return address, responseMsg
}