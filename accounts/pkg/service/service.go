package service

import (
	"context"
	"errors"

	db "github.com/faith/Accounts2/accounts/pkg/db"
	io "github.com/faith/Accounts2/accounts/pkg/io"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AccountsService describes the service.
type AccountsService interface {
	// Add your methods here
	CreateUser(ctx context.Context, user io.User) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

type basicAccountsService struct{}

func (b *basicAccountsService) CreateUser(ctx context.Context, user io.User) (s1 string, e0 error) {
	//logger := log.With(b.logger, "method", "CreateUser")
	uuid, _ := uuid.NewV4()
	user.ID = uuid.String()

	db, err := db.GetSession()
	if err != nil {
		return "Failed", err
	}
	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return "Failed", errors.New("email or password can't be empty")
	}

	// As the validations are successful
	// Hash/encrypt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "Failed", errors.New("Error by encrypting the password")
	}

	user.Password = string(hash) //assigning the hash to password after converting back to string

	_, err2 := db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err2 != nil {
		return "Failed", err2
	}

	//logger.Log("create user", user.ID)

	return user.ID, nil
}

func (b *basicAccountsService) GetUser(ctx context.Context, id string) (s0 string, e1 error) {

	db, err := db.GetSession()
	if err != nil {
		return "", errors.New("error could not get session")
	}

	var email string
	err3 := db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err3 != nil {
		return "", errors.New("error while doing get query using id")
	}

	//logger.Log("Get user", id)

	return email, nil
}

// NewBasicAccountsService returns a naive, stateless implementation of AccountsService.
func NewBasicAccountsService() AccountsService {
	return &basicAccountsService{}
}

// New returns a AccountsService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountsService {
	var svc AccountsService = NewBasicAccountsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
