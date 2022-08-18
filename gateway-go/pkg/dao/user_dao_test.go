package dao

import (
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/database"
	"ledungcobra/gateway-go/pkg/models"
	"testing"
)

func TestUserDAO_Find(t *testing.T) {

	users := generateUsers(2)
	config.Init()

	dbConnector := database.NewSQLConnector(config.Cfg.SqlDsn)
	dbConnector.Connect()
	dbConnector.MigrateModels()
	dao := NewUserDao(dbConnector.GetDatabase())
	for i, user := range users {
		err := dao.Save(&user)
		users[i] = user
		if err != nil {
			t.Errorf("Save() error = %v", err)
		}
	}
	defer deleteUsers(dbConnector.GetDatabase(), users)
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		query any
		args  any
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		foundUser *models.User
		wantErr   bool
	}{
		{name: "Test find user should found",
			fields: fields{db: dbConnector.GetDatabase()},
			args: args{
				query: "user_name=?",
				args:  users[0].UserName,
			},
			foundUser: &users[0],
			wantErr:   false,
		},
		{name: "Test find user should not found",
			fields: fields{db: dbConnector.GetDatabase()},
			args: args{
				query: "user_name=?",
				args:  users[0].UserName + fake.UUID(),
			},
			foundUser: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDAO{
				db: tt.fields.db,
			}
			user, err := u.Find(tt.args.query, tt.args.args)
			if diff := cmp.Diff(user, tt.foundUser, cmp.Comparer(comparatorUser)); tt.foundUser != nil && diff != "" {
				t.Errorf("Find() diff = %v", diff)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func comparatorUser(u1, u2 *models.User) bool {
	if u1 == nil && u2 == nil {
		return true
	}
	if u1 == nil || u2 == nil {
		return false
	}
	return u1.ID == u2.ID && u1.Verified == u2.Verified &&
		u1.UserName == u2.UserName && u1.Email == u2.Email &&
		u1.Password == u2.Password && u1.Picture == u2.Picture &&
		u1.BirthYear == u2.BirthYear && u1.BirthDay == u2.BirthDay &&
		u1.BirthMonth == u2.BirthMonth
}

func deleteUsers(getDatabase *gorm.DB, users []models.User) {
	for _, user := range users {
		getDatabase.Delete(&user)
	}
}

func generateUsers(size int) []models.User {
	users := make([]models.User, size)
	for i := 0; i < size; i++ {
		users[i] = models.User{
			FirstName:  fake.FirstName(),
			LastName:   fake.LastName(),
			UserName:   fake.Username() + fake.UUID(),
			Email:      fake.Email(),
			Password:   fake.Password(true, true, true, true, false, 10),
			Picture:    fake.Name(),
			BirthYear:  fake.Year(),
			BirthDay:   fake.Day(),
			BirthMonth: fake.Month(),
			Verified:   false,
			Gender:     "",
			Follower:   nil,
			Following:  nil,
			Search:     nil,
			Friends:    nil,
			Post:       nil,
		}
	}
	return users
}
