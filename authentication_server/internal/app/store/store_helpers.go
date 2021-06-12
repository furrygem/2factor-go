package store

import (
	"database/sql"
	"fmt"
)

// Execute query with map of arguments
func (s *Store) QueryStatementFromMap(m map[string]interface{}) (*sql.Stmt, error) {
	initialQuery := "select * from users where "
	i := 1
	for key := range m {
		initialQuery += fmt.Sprintf("%s=%d ", key, i)
		i++
	}
	stmt, err := s.db.Prepare(initialQuery)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}
