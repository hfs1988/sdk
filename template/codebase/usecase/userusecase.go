package usecase

import (
	"database/sql"

	"github.com/hfs1988/sdk/adapters/db"
	"github.com/hfs1988/sdk/client"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int32  `json:"age"`
	Status bool   `json:"status"`
}

type UserUsecase struct {
	db    *sql.DB
	sqlDB client.SQLDB
}

func GetUserUsecase(db *sql.DB, sqlDB client.SQLDB) *UserUsecase {
	return &UserUsecase{
		db:    db,
		sqlDB: sqlDB,
	}
}

func (u *UserUsecase) CreateUser(user User) string {
	err := u.sqlDB.Save(u.db, db.SQLEntity{
		Table: "users",
		ColsVals: db.SQLColsVals{
			Cols:   []string{"name", "email", "age", "status"},
			Values: []any{user.Name, user.Email, user.Age, true},
		},
	})
	if err != nil {
		panic(err)
	}
	return err.Error()
}

func (u *UserUsecase) UpdateUser(user User) string {
	err := u.sqlDB.Update(u.db, db.SQLEntity{
		Table: "users",
		ColsVals: db.SQLColsVals{
			Cols:   []string{"name", "email", "age", "status"},
			Values: []any{user.Name, user.Email, user.Age, user.Status},
		},
		Filters: db.SQLColsVals{
			Cols:   []string{"id"},
			Values: []any{user.ID},
		},
	})
	if err != nil {
		panic(err)
	}
	return err.Error()
}

func (u *UserUsecase) GetAllUser() []User {
	results := u.sqlDB.Get(u.db, db.SQLEntity{
		Table: "users",
		ColsVals: db.SQLColsVals{
			Cols: []string{"name", "email", "age", "status"},
		},
		Filters: db.SQLColsVals{
			Cols:   []string{"status"},
			Values: []any{true},
		},
	})

	var users []User
	for _, v := range results {
		user := User{
			Name:  v["name"].(string),
			Email: v["email"].(string),
			Age:   v["age"].(int32),
		}
		users = append(users, user)
	}
	return users
}

func (u *UserUsecase) GetUserID(userID int) []User {
	results := u.sqlDB.Get(u.db, db.SQLEntity{
		Table: "users",
		ColsVals: db.SQLColsVals{
			Cols: []string{"name", "email", "age", "status"},
		},
		Filters: db.SQLColsVals{
			Cols:   []string{"status", "id"},
			Values: []any{true, userID},
		},
	})

	var users []User
	for _, v := range results {
		user := User{
			Name:  v["name"].(string),
			Email: v["email"].(string),
			Age:   v["age"].(int32),
		}
		users = append(users, user)
	}
	return users
}

func (u *UserUsecase) DeleteUser(user User) string {
	err := u.sqlDB.Update(u.db, db.SQLEntity{
		Table: "users",
		ColsVals: db.SQLColsVals{
			Cols:   []string{"status"},
			Values: []any{false},
		},
		Filters: db.SQLColsVals{
			Cols:   []string{"id"},
			Values: []any{user.ID},
		},
	})
	if err != nil {
		panic(err)
	}
	return err.Error()
}
