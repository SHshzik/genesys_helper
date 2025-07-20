package sqlite_adapter

import (
	"database/sql"

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
	query, args, err := queries.FindUserByIDSQL(id)
	if err != nil {
		return domain.User{}, err
	}

	row := s.db.QueryRow(query, args...)

	var user domain.User
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *SqliteAdapter) CreateUser(user *domain.User) error {
	query, args, err := queries.CreateUserSQL(user.ID, user.FirstName, user.LastName, user.UserName)
	if err != nil {
		return err
	}

	row := s.db.QueryRow(query, args...)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	user.ID = id

	return err
}

func (s *SqliteAdapter) GetCharacterByUserID(id int64) (domain.Character, error) {
	query, args, err := queries.FindCharacterByUserIDSQL(id)
	if err != nil {
		return domain.Character{}, err
	}

	row := s.db.QueryRow(query, args...)

	var character domain.Character
	err = row.Scan(&character.ID, &character.Name)
	if err != nil {
		return domain.Character{}, err
	}

	return character, nil
}
