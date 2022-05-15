package model

type Cerita struct {
	ID		 	int 		`json:"id" gorm:"primaryKey"`
	Tanggal 	string  	`json:"tanggal"`
	Judul 		string 		`json:"judul"`
	Genre		string 		`json:"genre`
	Deskripsi 	string  	`json:"deskripsi"`
	IsCompleted bool 		`json:"iscompleted"`
	IDUser 		int 		`json:"iduser" gorm:"foreignkey:id_user;references:ID"`	
}