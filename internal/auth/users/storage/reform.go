package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	"github.com/erupshis/revtracker/internal/logger"
	"gopkg.in/reform.v1"
)

var (
	_ BaseUsersStorage = (*usersReform)(nil)
)

// usersReform storageManager implementation for PostgreSQL. Consist of database and QueriesHandler.
// Request to database are synchronized by sync.RWMutex. All requests are done on united transaction. Multi insert/update/delete is not supported at the moment.
type usersReform struct {
	db *reform.DB

	log logger.BaseLogger
}

// Create creates usersReform implementation. Supports migrations and check connection to database.
func Create(dbConn *reform.DB, log logger.BaseLogger) BaseUsersStorage {
	return &usersReform{
		db:  dbConn,
		log: log,
	}
}

func (r *usersReform) InsertUser(ctx context.Context, user *data.User) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, user)
}

func (r *usersReform) UpdateUser(ctx context.Context, user *data.User) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, user)
}

func (r *usersReform) SelectUserByID(ctx context.Context, ID int64) (*data.User, error) {
	return r.selectUser(ctx, nil, map[string]utils.Argument{"id": utils.CreateArgument(ID)})
}

func (r *usersReform) SelectUserByLogin(ctx context.Context, login string) (*data.User, error) {
	return r.selectUser(ctx, nil, map[string]utils.Argument{"login": utils.CreateArgument(login)})
}

func (r *usersReform) DeleteUserByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, map[string]utils.Argument{"id": utils.CreateArgument(ID)}, data.UserTable)
}

func (r *usersReform) selectUser(ctx context.Context, tx *reform.TX, filters map[string]utils.Argument) (*data.User, error) {
	content, err := requests.SelectOne(ctx, r.db, tx, filters, data.UserTable)

	if content == nil {
		return nil, err
	}
	return content.(*data.User), err
}

// TODO: hash password in db. Need to split storage - should be storage and manager.
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
