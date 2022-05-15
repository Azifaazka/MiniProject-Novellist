package model

type Comment struct {
	ID		 	int 		`json:"id" gorm:"primaryKey"`
	Deskcom	 	string  	`json:"deskcom"`
	IDCerita 	int 		`json:"idcerita" gorm:"foreignkey:"idcerita;references:ID"`
	IDUser 		int 		`json:"iduser" gorm:"foreignkey:"iduser;references:ID"`
	
}