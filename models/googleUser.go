package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type GoogleUser struct {
	Id             uint                   `json:"id,omitempty"`
	Token          string                 `json:"token,omitempty" gorm:"type:varchar(2000)"`
	Name           string                 `json:"name,omitempty"`
	Email          string                 `json:"email,omitempty"`
	SignedInSource string                 `json:"signed_in_source,omitempty"` //  google (if user created from any other source it wont have this value "maunal" and can be blacklisted)
	Addresses      []Address              `json:"addresses,omitempty" gorm:"ForeignKey:UserId;AssociationForeignKey:Id"` //one to many
	CreatedAt      time.Time              `json:"created_at,omitempty"`
	UpdatedAt      time.Time              `json:"updated_at,omitempty"`
}

func (GoogleUser) TableName() string {
	return "e_google_user"
}

//Interacting with DB

func GetGoogleUserById(db *gorm.DB, id int) (GoogleUser, string) {
	var googleUser GoogleUser
	var responseMsg string
	rowsAffected := db.Debug().First(&googleUser, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetGoogleUserById Query Failure !"
	} else {
		responseMsg = "GetGoogleUserById Query Success !"
	}
	return googleUser, responseMsg
}

func GetGoogleUserWithAssociationsById(db *gorm.DB, id int) (GoogleUser, string) {
	var googleUser GoogleUser
	var responseMsg string

	rowsAffected := db.Debug().Preload("Addresses").First(&googleUser, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetGoogleUserWithAssociationsById Query Failure !"
	} else {
		responseMsg = "GetGoogleUserWithAssociationsById Query Success !"
	}
	return googleUser, responseMsg
}

func GetGoogleAllUsers(db *gorm.DB) ([]GoogleUser, string) {
	var googleUsers []GoogleUser
	var responseMsg string

	rowsAffected := db.Debug().Find(&googleUsers).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetGoogleAllUsers Query Failure !"
	} else {
		responseMsg = "GetGoogleAllUsers Query Success !"
	}
	return googleUsers, responseMsg
}

func GetGoogleAllUserswithAssociations(db *gorm.DB) ([]GoogleUser, string) {
	var googleUsers []GoogleUser
	var responseMsg string

	rowsAffected := db.Debug().Preload("Addresses").Find(&googleUsers).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetGoogleAllUserswithAssociations Query Failure !"
	} else {
		responseMsg = "GetGoogleAllUserswithAssociations Query Success !"
	}
	return googleUsers, responseMsg
}

//CREATES NEW USER ONLY IF USER NOT ALREADY PRESENT IN DB
func CreateGoogleUser(db *gorm.DB, data GoogleUser) (GoogleUser, string) {
	var googleUser GoogleUser
	var responseMsg string
	data.SignedInSource = "google"

	//Check if google user record already present
	db.Where("email = ?" , data.Email).First(&googleUser)
	if googleUser.Id !=0 {
		tempId := googleUser.Id
		rowsAffected := db.Debug().Model(&GoogleUser{}).Where("id=?", tempId).Update(&data).RowsAffected
		db.Debug().Preload("Addresses").Find(&googleUser, tempId)
		if (rowsAffected == 0 ) {
			responseMsg = "CreateGoogleUpdateUser Query Failure !"
		} else {
			responseMsg = "CreateGoogleUpdateUser Query Success !"
		}
		return googleUser , responseMsg
	}

	rowsAffected := db.Debug().Create(&data).RowsAffected
	db.Debug().Preload("Addresses").Last(&googleUser)
	if (rowsAffected == 0 ) {
		responseMsg = "CreateGoogleUser Query Failure !"
	} else {
		responseMsg = "CreateGoogleUser Query Success !"
	}
	return googleUser, responseMsg
}

//UPDATING OF CHILD IS NOT ALLOWED WHILE UDATING PARENT , CALL CHILD'S UPDATE METHOD EXCLUSIVELY
func UpdateGoogleUserById(db *gorm.DB, data GoogleUser, id int) (GoogleUser, string) {
	var googleUser GoogleUser
	var responseMsg string
	rowsAffected := db.Debug().Model(&GoogleUser{}).Where("id=?", id).Update(&data).RowsAffected
	db.Debug().Preload("Addresses").Find(&googleUser, id)
	if (rowsAffected == 0 ) {
		responseMsg = "UpdateGoogleUserById Query Failure !"
	} else {
		responseMsg = "UpdateGoogleUserById Query Success !"
	}
	return googleUser, responseMsg
}

//Deletes without warning even if parent of the CHILD in question exists
func DeleteGoogleUserById(db *gorm.DB, id int) (GoogleUser, string) {
	var googleUser GoogleUser
	var responseMsg string
	db.Debug().Preload("Addresses").First(&googleUser, id)
	rowsAffected := db.Debug().Delete(&googleUser, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "DeleteGoogleUserById Query Failure !"
	} else {
		responseMsg = "DeleteGoogleUserById Query Success !"
	}
	return googleUser, responseMsg
}
