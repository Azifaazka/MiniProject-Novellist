package model

type User struct {
	ID		 	int 	`json:"id" gorm:"primaryKey"`
	Name 		string 	`json:"name"`
	Email 		string 	`json:"email"`
	Username 	string  `json:"Username"`
	Password 	string 	`json:"password"`
	
}