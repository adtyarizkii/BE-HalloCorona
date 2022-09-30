package models

import "time"

type Consultation struct {
	ID        	int       		`json:"id" gorm:"primary_key:auto_increment"`
	Fullname  	string    		`json:"fullname" gorm:"type: varchar(255)"`
	Phone  	  	string    		`json:"phone" gorm:"type: varchar(255)"`
	BornDate  	time.Time		`json:"bornDate"`
	Age		  	int				`json:"age" gorm:"type: int"`
	Height	  	int				`json:"height" gorm:"type: int"`
	Weight	  	int				`json:"weight" gorm:"type: int"`
	Gender	  	string			`json:"gender" gorm:"type: varchar(255)"`
	Subject	  	string			`json:"subject" gorm:"type: varchar(255)"`
	LiveConsul  time.Time		`json:"liveConsul"`
	Desc      	string 	    	`json:"desc" gorm:"type:text" form:"desc"`
	UserID	  	int				`json:"user_id" form:"user_id"`
	User	  	UserResponse 	`json:"user"`
	Status		string			`json:"status" gorm:"type: varchar(255)"`
	Reply 		string			`json:"reply" gorm:"type:text" form:"reply"`
	LinkLive	string			`json:"linkLive" gorm:"type: varchar(255)"`
	CreatedAt 	time.Time 		`json:"-"`
	UpdatedAt 	time.Time 		`json:"-"`
}

type ConsultationResponse struct {
	ID        	int       		`json:"id" gorm:"primary_key:auto_increment"`
	Fullname  	string    		`json:"fullname" gorm:"type: varchar(255)"`
	Phone  	  	string    		`json:"phone" gorm:"type: varchar(255)"`
	BornDate  	time.Time		`json:"bornDate"`
	Age		  	int				`json:"age" gorm:"type: int"`
	Height	  	int				`json:"height" gorm:"type: int"`
	Weight	  	int				`json:"weight" gorm:"type: int"`
	Gender	  	string			`json:"gender" gorm:"type: varchar(255)"`
	Subject	  	string			`json:"subject" gorm:"type: varchar(255)"`
	LiveConsul  time.Time		`json:"liveConsul"`
	Desc      	string 	    	`json:"desc" gorm:"type:text" form:"desc"`
	UserID	  	int				`json:"user_id" form:"user_id"`
	User	  	UserResponse 	`json:"user"`
	Status		string			`json:"status" gorm:"type: varchar(255)"`
	Reply 		string			`json:"reply" gorm:"type:text" form:"reply"`
	LinkLive	string			`json:"linkLive" gorm:"type: varchar(255)"`
	CreatedAt 	time.Time 		`json:"createdAt"`
	UpdatedAt 	time.Time 		`json:"updatedAt"`
}

func (ConsultationResponse) TableName() string {
	return "users"
}