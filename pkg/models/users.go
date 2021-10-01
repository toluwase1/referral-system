package models

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID
	Name string
	Email  string
	Password string
	Balance float64
	Referral
}

type Referral struct {
	referralcode uuid.UUID
	referralcount uint64
	dueforbonus bool
}

func (user *User) Register () *User{

	return user
}

//get a user by the referral code
// referral count method
//register
//transaction
	//transfer

//referral method
	//notification based on 2 conditions