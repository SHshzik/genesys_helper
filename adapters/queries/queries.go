package queries

import sq "github.com/Masterminds/squirrel"

func FindUserByIDSQL(id int64) (string, []interface{}, error) {
	query := sq.Select("id", "first_name", "last_name", "user_name").
		From("users").
		Where(sq.Eq{"id": id})

	return query.ToSql()
}

func CreateUserSQL(id int64, firstName string, lastName string, userName string) (string, []interface{}, error) {
	query := sq.Insert("users").
		Columns("id", "first_name", "last_name", "user_name").
		Values(id, firstName, lastName, userName).
		Suffix("RETURNING id")

	return query.ToSql()
}

func FindCharacterByUserIDSQL(id int64) (string, []interface{}, error) {
	query := sq.Select("id", "name").
		From("characters").
		Where(sq.Eq{"user_id": id})

	return query.ToSql()
}

func CreateCharacterSQL(user_id int64) (string, []interface{}, error) {
	query := sq.Insert("characters").
		Columns("user_id").
		Values(user_id).
		Suffix("RETURNING id")

	return query.ToSql()
}

func UpdateCharacterByIDSQL(id int64, name string) (string, []interface{}, error) {
	query := sq.Update("characters").
		Set("name", name).
		Where(sq.Eq{"id": id})

	return query.ToSql()
}
