package repository

import (
	"authentication/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

func (u userRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	dataset := u.db.Insert("users").
		Cols("full_name", "phone", "username", "password").
		Vals(goqu.Vals{
			user.FullName,
			user.Phone,
			user.Username,
			user.Password,
		})

	sql, _, err := dataset.ToSQL() // Generate the SQL statement
	if err != nil {
		return domain.User{}, err
	}

	_, err = u.db.ExecContext(ctx, sql) // Execute the SQL statement
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u userRepository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"username": username,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) GetLastID(ctx context.Context) (int64, error) {
	var result struct {
		ID int64 `db:"id"`
	}

	dataset := u.db.From("users").Order(goqu.C("id").Desc()).Limit(1)

	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}
