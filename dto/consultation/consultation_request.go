package consultationdto

import "time"

type CreateConsultationRequest struct {
	Fullname  	string    		`gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Phone  	  	string    		`gorm:"type: varchar(255)" json:"phone" validate:"required"`
	BornDate  	time.Time		`json:"bornDate" validate:"required"`
	Age		  	int				`json:"age" gorm:"type: int" validate:"required"`
	Height	  	int				`json:"height" gorm:"type: int" validate:"required"`
	Weight	  	int				`json:"weight" gorm:"type: int" validate:"required"`
	Gender	  	string			`json:"gender" gorm:"type: varchar(255)" validate:"required"`
	Subject	  	string			`json:"subject" gorm:"type: varchar(255)" validate:"required"`
	LiveConsul  time.Time		`json:"liveConsul" validate:"required"`
	Desc      	string 	    	`json:"desc" gorm:"type:text" form:"desc" validate:"required"`
	UserID	  	int				`json:"user_id" form:"user_id"`
	Status		string			`json:"status" gorm:"type: varchar(255)"`
	Reply 		string			`json:"reply" gorm:"type:text" form:"reply" `
	LinkLive	string			`json:"linkLive" gorm:"type: varchar(255)"`
}

type UpdateConsultationRequest struct {
	Fullname  	string    		`gorm:"type: varchar(255)" json:"fullname"`
	Phone  	  	string    		`gorm:"type: varchar(255)" json:"phone"`
	BornDate  	time.Time		`json:"bornDate"`
	Age		  	int				`json:"age" gorm:"type: int"`
	Height	  	int				`json:"height" gorm:"type: int"`
	Weight	  	int				`json:"weight" gorm:"type: int"`
	Gender	  	string			`json:"gender" gorm:"type: varchar(255)"`
	Subject	  	string			`json:"subject" gorm:"type: varchar(255)"`
	LiveConsul  time.Time		`json:"liveConsul"`
	Desc      	string 	    	`json:"desc" gorm:"type:text" form:"desc"`
	Status		string			`json:"status" gorm:"type: varchar(255)"`
	Reply 		string			`json:"reply" gorm:"type:text" form:"reply" `
	LinkLive	string			`json:"linkLive" gorm:"type: varchar(255)"`
}