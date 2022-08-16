package models

import "gorm.io/gorm"

type User struct {
	FirstName  string `gorm:"column:first_name;type:varchar(100);not null"`
	LastName   string `gorm:"column:last_name;type:varchar(100);not null"`
	UserName   string `gorm:"column:user_name;type:varchar(255);not null;unique"`
	Email      string `gorm:"column:email;type:varchar(255);not null;unique"`
	Password   string `gorm:"column:password;type:varchar(255);not null"`
	Picture    string `gorm:"column:picture;type:varchar(255);not null;default:''"`
	Cover      string `gorm:"column:cover;type:varchar(255);not null;default:''"`
	BirthYear  int    `gorm:"column:birth_year;type:int;not null;default:0"`
	BirthDay   int    `gorm:"column:birth_day;type:int;not null;default:0"`
	BirthMonth int    `gorm:"column:birth_month;type:int;not null;default:0"`
	Verified   bool   `gorm:"column:verified;type:boolean;not null;default:false"`
	Gender     string `gorm:"column:gender;type:varchar(10);not null"`

	Follower  []User `gorm:"many2many:user_followers;"`
	Following []User `gorm:"many2many:user_following;"`
	Search    []User `gorm:"many2many:user_search;"`
	Friends   []User `gorm:"many2many:user_friends;"`
	Detail    Detail `gorm:"embedded"`
	Post      []Post `gorm:"foreignKey:AuthorID;"`
	gorm.Model
}

func (User) TableName() string {
	return "users"
}

type Detail struct {
	Bio          string `gorm:"column:bio;type:varchar(255);not null;default:''"`
	OtherName    string `gorm:"column:other_name;type:varchar(255);not null;default:''"`
	Job          string `gorm:"column:job;type:varchar(255);not null;default:''"`
	Workplace    string `gorm:"column:workplace;type:varchar(255);not null;default:''"`
	Highschool   string `gorm:"column:highschool;type:varchar(255);not null;default:''"`
	College      string `gorm:"column:college;type:varchar(255);not null;default:''"`
	Currentcity  string `gorm:"column:current_city;type:varchar(255);not null;default:''"`
	Hometown     string `gorm:"column:hometown;type:varchar(255);not null;default:''"`
	Relationship string `gorm:"column:relationship;type:varchar(255);not null;default:''"`
	Instagram    string `gorm:"column:instagram;type:varchar(255);not null;default:''"`
}
