package postgre

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"foodmarket/domain"
	"github.com/go-pg/pg/v10"

)

//PsqlUserRepository struct
type PsqlUserRepository struct {
	DB *pg.DB
}

//NewPsqlUserRepository psql
func NewPsqlUserRepository(db *pg.DB) domain.UserRepository {
	return PsqlUserRepository{DB: db}
}

func (u PsqlUserRepository) CreateUser(ctx context.Context, usr *domain.User) (user *domain.User, err error) {
	_, err = u.DB.Model(usr).Insert()
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u PsqlUserRepository) Attempt(ctx context.Context, credential *domain.Credential) (user *domain.User, err error) {
	user = new(domain.User)
	err = u.DB.Model(user).Where("email = ?", credential.Email).Select()
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u PsqlUserRepository) Update(ctx context.Context, usr *domain.User) (user *domain.User, err error) {
	_, err = u.DB.Model(usr).Update()

	if err != nil {
		return nil, err
	}

	return usr, nil}

func (u PsqlUserRepository) Find(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	user = new(domain.User)
	err = u.DB.Model(user).Where("id = ? ", id).First()
	if err != nil {
		return nil, err
	}

	return user, nil}

func (u PsqlUserRepository) FindBy(ctx context.Context, key, value string) (user *domain.User, err error) {
	panic("implement me")
}