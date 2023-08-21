package customer

type Customer struct {
	ID    int64  `gorm:"primary_key;auto_increment;not_null"`
	Name  string `gorm:"not null" json:"name,omitempty"`
	Email string `gorm:"not null" json:"email,omitempty"`
}

