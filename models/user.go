package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id             uint                   `json:"id,omitempty"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Password       string                 `json:"password,omitempty"`
	SignedInSource string                 `json:"signed_in_source,omitempty"` //  manual (if user created from any other source it wont have this value "maunal" and can be blacklisted)
	Addresses      []Address              `json:"addresses,omitempty" gorm:"ForeignKey:UserId;AssociationForeignKey:Id"` //one to many
	CreatedAt      time.Time              `json:"created_at,omitempty"`
	UpdatedAt      time.Time              `json:"updated_at,omitempty"`
}

func (User) TableName() string {
	return "e_users"
}


//Interacting with DB

func GetUserById(db *gorm.DB, id int) (User, string) {                                //tested
	var user User
	var responseMsg string
	rowsAffected := db.Debug().First(&user, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetUserById Query Failure !"
	} else {
		responseMsg = "GetUserById Query Success !"
	}
	return user, responseMsg
}

func GetUserWithAssociationsById(db *gorm.DB, id int) (User, string) {                     //tested
	var user User
	var responseMsg string

	rowsAffected := db.Debug().Preload("Addresses").First(&user, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetUserWithAssociationsById Query Failure !"
	} else {
		responseMsg = "GetUserWithAssociationsById Query Success !"
	}
	return user, responseMsg
}

func GetAllUsers(db *gorm.DB) ([]User, string) {                                       //tested
	var users []User
	var responseMsg string

	rowsAffected := db.Debug().Find(&users).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetAllUsers Query Failure !"
	} else {
		responseMsg = "GetAllUsers Query Success !"
	}
	return users, responseMsg
}

func GetAllUserswithAssociations(db *gorm.DB) ([]User, string) {                        //tested
	var users []User
	var responseMsg string

	rowsAffected := db.Debug().Preload("Addresses").Find(&users).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "GetAllUserswithAssociations Query Failure !"
	} else {
		responseMsg = "GetAllUserswithAssociations Query Success !"
	}
	return users, responseMsg
}

//CRATING OF CHILD IS ALLOWED WHILE CREATING PARENT BUT STILL PREFER CHILD'S CREATE METHOD EXCLUSIVELY
func CreateUser(db *gorm.DB, data User) (User, string) {                                 //tested
	var user User
	var responseMsg string
	data.SignedInSource = "manual"
	rowsAffected := db.Debug().Create(&data).RowsAffected
	db.Debug().Preload("Addresses").Last(&user)
	if (rowsAffected == 0 ) {
		responseMsg = "CreateUserBy Query Failure !"
	} else {
		responseMsg = "CreateUserBy Query Success !"
	}
	return user, responseMsg
}

//UPDATING OF CHILD IS NOT ALLOWED WHILE UDATING PARENT , CALL CHILD'S UPDATE METHOD EXCLUSIVELY
func UpdateUserById(db *gorm.DB, data User, id int) (User, string) {          //tested (new child gets created with foreign key 0)
	var user User
	var responseMsg string
	rowsAffected := db.Debug().Model(&User{}).Where("id=?", id).Update(&data).RowsAffected
	db.Debug().Preload("Addresses").Find(&user, id)
	if (rowsAffected == 0 ) {
		responseMsg = "UpdateUserById Query Failure !"
	} else {
		responseMsg = "UpdateUserById Query Success !"
	}
	return user, responseMsg
}

//Deletes without warning even if parent of the CHILD in question exists
func DeleteUserById(db *gorm.DB, id int) (User, string) {                    //tested (deletes even if child present)
	var user User
	var responseMsg string
	db.Debug().Preload("Addresses").First(&user, id)
	rowsAffected := db.Debug().Delete(&user, id).RowsAffected
	if (rowsAffected == 0 ) {
		responseMsg = "DeleteUserById Query Failure !"
	} else {
		responseMsg = "DeleteUserById Query Success !"
	}
	return user, responseMsg
}
