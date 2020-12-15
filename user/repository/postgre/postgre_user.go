package postgre

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"foodmarket/domain"
)

//PsqlUserRepository struct
type PsqlUserRepository struct {
	DB *gorm.DB
}

//NewPsqlUserRepository psql
func NewPsqlUserRepository(Coon *gorm.DB) domain.UserRepository {
	return &PsqlUserRepository{Coon}
}

func (u *PsqlUserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User

	if result := u.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *PsqlUserRepository) CreateUser(ctx context.Context, usr *domain.User) (user *domain.User, err error) {
	panic("coming soon")
}

func (u *PsqlUserRepository) Attempt(ctx context.Context, credential *domain.Credential) (user *domain.User, err error) {

	panic("coming soon")
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *PsqlUserRepository) Update(ctx context.Context, usr *domain.User) (user *domain.User, err error) {
	panic("coming soon")

	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *PsqlUserRepository) Find(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	panic("coming soon")
	//
	//user = new(domain.User)
	//err = u.DB.Model(user).Where("id = ? ", id).First()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return user, nil
}

func (u *PsqlUserRepository) FindBy(ctx context.Context, key, value string) (user *domain.User, err error) {
	panic("implement me")
}
