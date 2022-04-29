package model

type Star struct {
	ID		 	int 		`json:"id" gorm:"primaryKey"`
	deskcom	 	string  	`json:"deskcom"`
	ID_Cerita 	Cerita 		`json:"id_cerita" gorm:"foreignkey:"id.cerita"`
	ID_User 	User 		`json:"id_user" gorm:"foreignkey:"id.user"`
}