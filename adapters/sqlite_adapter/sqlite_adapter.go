package sqlite_adapter

import (
	"database/sql"
	"fmt"

	"github.com/SHshzik/genesys_helper/adapters/queries"
	"github.com/SHshzik/genesys_helper/domain"
)

type SqliteAdapter struct {
	db *sql.DB
}

func NewSqliteAdapter(db *sql.DB) *SqliteAdapter {
	return &SqliteAdapter{db: db}
}

func (s *SqliteAdapter) GetUserByID(id int64) (domain.User, error) {
	query, args, err := queries.GetUserByIDSQL(id)
	if err != nil {
		fmt.Println("!!!!")
		fmt.Println(err)
		return domain.User{}, err
	}

	row := s.db.QueryRow(query, args...)

	var user domain.User
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName)
	if err != nil {
		fmt.Println("!!!!")
		fmt.Println(err)
		return domain.User{}, err
	}

	return user, nil
}

func (s *SqliteAdapter) CreateUser(user *domain.User) error {
	query, args, err := queries.CreateUserSQL(user.ID, user.FirstName, user.LastName, user.UserName)
	if err != nil {
		fmt.Println("!!!!")
		fmt.Println(err)
		return err
	}

	row := s.db.QueryRow(query, args...)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("!!!!")
		fmt.Println(err)
		return err
	}

	user.ID = id

	return err
}
