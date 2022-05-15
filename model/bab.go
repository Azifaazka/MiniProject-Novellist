package model

type Bab struct {
	ID		 	int 		`json:"id" gorm:"primaryKey"`
	Tanggal 	string  	`json:"tanggal"`
	Judulbab 	string 		`json:"judulbab"`
	Konten	 	string  	`json:"konten"`
	IDCerita 	int 		`json:"idcerita" gorm:"foreignkey:"idcerita;references:ID"`
}