package queries

import sq "github.com/Masterminds/squirrel"

func GetUserByIDSQL(id int64) (string, []interface{}, error) {
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
