package repository

import ("gorm.io/gorm"
		. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
)

type UserRepository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
	FindByid(id uint) (*User, error)
	FindAll() ([]User, error)
	Update(user *User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error{
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where("username = ?",username).First(&user).Error; err != nil{
		return nil, err
	}
	return &user,nil
}

func (r *userRepository) FindByid(id uint) (*User, error) {
	var user User
	if err := r.db.Where("id = ?",id).First(&user).Error; err != nil{
		return nil, err
	}
	return &user,nil
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil{
		return nil, err
	}
	return users,nil
}

func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}


