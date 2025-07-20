package sqlite_adapter

import (
	"database/sql"

	"github.com/SHshzik/genesys_helper/adapters/queries"
	"github.com/SHshzik/genesys_helper/domain"
	"github.com/SHshzik/genesys_helper/pkg/logger"
)

type SqliteAdapter struct {
	db *sql.DB
	l  logger.Interface
}

func NewSqliteAdapter(db *sql.DB, l logger.Interface) *SqliteAdapter {
	return &SqliteAdapter{db: db, l: l}
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
		s.l.Error("GetCharacterByUserID", "error", err)
		return domain.Character{}, err
	}

	row := s.db.QueryRow(query, args...)

	character := domain.Character{UserID: id}
	err = row.Scan(&character.ID, &character.Name)
	if err != nil {
		s.l.Error("GetCharacterByUserID", "error", err)
		return domain.Character{}, err
	}

	return character, nil
}

func (s *SqliteAdapter) CreateCharacter(character *domain.Character) error {
	query, args, err := queries.CreateCharacterSQL(character.UserID)
	if err != nil {
		s.l.Error("CreateCharacter", "error", err)
		return err
	}

	row := s.db.QueryRow(query, args...)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		s.l.Error("CreateCharacter", "error", err)
		return err
	}

	character.ID = id

	return nil
}

func (s *SqliteAdapter) UpdateCharacter(character *domain.Character) error {
	query, args, err := queries.UpdateCharacterByIDSQL(character.ID, character.Name)
	if err != nil {
		s.l.Error("UpdateCharacter", "error", err)
		return err
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		s.l.Error("UpdateCharacter", "error", err)
		return err
	}

	return nil
}
