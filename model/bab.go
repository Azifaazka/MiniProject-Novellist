package model

type Bab struct {
	ID		 	int 		`json:"id" gorm:"primaryKey"`
	Tanggal 	timestamp  	`json:"tanggal"`
	Judul_bab 	string 		`json:"judulbab"`
	konten	 	string  	`json:"konten"`
	ID_Cerita 	Cerita 		`json:"id_cerita" gorm:"foreignkey:"id.cerita"`
}