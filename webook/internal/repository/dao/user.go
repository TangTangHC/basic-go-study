package dao

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) Insert(ctx context.Context, user User) error {
	milli := time.Now().UnixMilli()
	user.Ctime = milli
	user.Utime = milli
	err := u.db.WithContext(ctx).Create(&user).Error
	if mE, ok := err.(*mysql.MySQLError); ok {
		const uniqueError uint16 = 1062
		if mE.Number == uniqueError {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (u *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *UserDao) UpdateById(ctx context.Context, user User) error {
	return u.db.WithContext(ctx).Updates(&user).Error
}

func (u *UserDao) FindById(ctx *gin.Context, id int64) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

type User struct {
	Id        int64  `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string
	NikeName  string
	Birthday  string
	Signature string
	Ctime     int64
	Utime     int64
}
