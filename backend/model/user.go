package model

import (
	// "time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Name     string `json:"name"`
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
}
type CreateFeedBack struct {
     gorm.Model
     Title  string `json: "title"`
     Category string `json:"category" gorm:"type:enum('features','bug','enhancement')" gorm:"not null"`
     Details string `json:"details"`
}
type Feedback struct { 
      ID  int`json:id`
      Title  string `json: "title"`
     Category string `json:"category"`
     Details string `json:"details"`
}

type LikeFeedback struct {
    gorm.Model
    UserID  uint 
    FeedbackID uint 
}

// type PostComment struct{
//      gorm.Model
//      UserID  uint 
//     Comment string `json: "comment"`
//     PostId  uint  `json:"post_id"`
//     CreatedAt time.Time `json:"created_at"`
//     UpdatedAt time.Time `json:"updated_at"`

// }