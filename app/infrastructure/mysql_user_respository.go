package infrastructure

import (
	"context"
	"database/sql"

	"github.com/bayamasa/go-boilerplate/app/domain/user"
	"github.com/bayamasa/go-boilerplate/app/infrastructure/client/db"
)

type MySQLUserRepository struct {
	db *db.DB
}

func NewMySQLUsersRepository(db *db.DB) user.UserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

func (r *MySQLUserRepository) FindBy(ctx context.Context, id int64) (*user.User, error) {
	query := `
	SELECT 
		id, 
		email,
		lastName,
		firstName,
		gender,
		dateOfBirth,
		location
	FROM 
		users 
	WHERE 
		id = ?`
	
	var dbo UserDBO
	err := r.db.Read.QueryRowContext(ctx, query, id).Scan(&dbo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dbo.ToDomainModel(), nil
}
