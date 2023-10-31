package dao

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicate = errors.New("数据冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

type UserDao interface {
	Insert(ctx context.Context, user User) error
	FindByEmail(context.Context, string) (User, error)
	UpdateById(context.Context, User) error
	FindById(context.Context, int64) (User, error)
	FindByPhone(context.Context, string) (User, error)
}

type GORMUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GORMUserDao{db: db}
}

func (u *GORMUserDao) Insert(ctx context.Context, user User) error {
	milli := time.Now().UnixMilli()
	user.Ctime = milli
	user.Utime = milli
	err := u.db.WithContext(ctx).Create(&user).Error
	if mE, ok := err.(*mysql.MySQLError); ok {
		const uniqueError uint16 = 1062
		if mE.Number == uniqueError {
			return ErrUserDuplicate
		}
	}
	return err
}

func (u *GORMUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *GORMUserDao) UpdateById(ctx context.Context, user User) error {
	return u.db.WithContext(ctx).Updates(&user).Error
}

func (u *GORMUserDao) FindById(ctx context.Context, id int64) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *GORMUserDao) FindByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	return user, err
}

type User struct {
	Id        int64          `gorm:"primaryKey"`
	Email     sql.NullString `gorm:"unique"`
	Phone     sql.NullString `gorm:"unique"`
	Password  string
	NikeName  string
	Birthday  string
	Signature string
	Ctime     int64
	Utime     int64
}
