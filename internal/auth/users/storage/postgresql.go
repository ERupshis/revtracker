package storage

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/logger"
	"gopkg.in/reform.v1"
)

// usersStoragePostgres storageManager implementation for PostgreSQL. Consist of database and QueriesHandler.
// Request to database are synchronized by sync.RWMutex. All requests are done on united transaction. Multi insert/update/delete is not supported at the moment.
type usersStoragePostgres struct {
	db *reform.DB

	log logger.BaseLogger
}

// Create creates usersStoragePostgres implementation. Supports migrations and check connection to database.
func Create(dbConn *reform.DB, log logger.BaseLogger) BaseUsersStorage {
	return &usersStoragePostgres{
		db:  dbConn,
		log: log,
	}
}

func (p *usersStoragePostgres) AddUser(ctx context.Context, user *data.User) (int64, error) {
	// TODO: add impl.
	return -1, nil
}

func (p *usersStoragePostgres) GetUser(ctx context.Context, login string) (*data.User, error) {
	// TODO: add impl.
	return nil, nil
}

func (p *usersStoragePostgres) GetUserID(ctx context.Context, login string) (int64, error) {
	// TODO: add impl.
	return -1, nil
}

func (p *usersStoragePostgres) GetUserRole(ctx context.Context, userID int64) (int, error) {
	user, err := p.getUser(ctx, map[string]interface{}{"id": userID})
	if err != nil {
		return -1, fmt.Errorf("get user role: %w", err)
	}

	if user == nil {
		return -1, nil
	}

	return user.Role, nil
}

func (p *usersStoragePostgres) getUser(ctx context.Context, filters map[string]interface{}) (*data.User, error) {
	// TODO: add impl.
	return nil, nil
}

// TODO: hash password in db.
// password := "user_password" // Replace with the actual password provided by the user
//
// // Hash and salt the password
// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// if err != nil {
// log.Fatal(err)
// }
//
// // Store 'hashedPassword' in the database for the user
//
// // User login: Verify password
// providedPassword := "user_password" // Replace with the password provided during login
//
// // Verify the provided password with the stored hashed password
// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(providedPassword))
// if err == nil {
// fmt.Println("Password is correct!")
// } else if err == bcrypt.ErrMismatchedHashAndPassword {
// fmt.Println("Password is incorrect.")
// } else {
// log.Fatal(err)
// }
