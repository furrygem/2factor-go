package store

import (
	"database/sql"
	"fmt"
)

// Execute query with map of arguments
func (s *Store) QueryStatementFromMap(prefix string, m map[string]interface{}, suffix string) (*sql.Stmt, error) {
	initial_query_template := "%s users WHERE %s %s"

	// i := 1
	// for key := range m {
	// 	if i == len(m) {
	// 		initialQuery += fmt.Sprintf("%s=$%d", key, i)
	// 		i++
	// 		continue
	// 	}
	// 	initialQuery += fmt.Sprintf("%s=$%d AND ", key, i)
	// 	i++
	// }
	var query string
	query, _ = basicqueryfrommap(m, " AND", 1)
	complete_query := fmt.Sprintf(initial_query_template, prefix, query, suffix)
	stmt, err := s.db.Prepare(complete_query)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func (s *Store) UpdateStatementFromMap(set map[string]interface{}, where map[string]interface{}, suffix string) (*sql.Stmt, error) {
	initial_query_template := "UPDATE users SET %s WHERE %s %s"
	set_query := ""
	where_query := ""
	// i := 1
	// for key := range set {
	// 	if i == len(set) {
	// 		set_query += fmt.Sprintf("%s = $%d", key, i)
	// 		i++
	// 		continue
	// 	}
	// 	set_query += fmt.Sprintf("%s = $%d, ", key, i)
	// 	i++
	// }
	// i = 1
	// for key := range where {
	// 	if i == len(where) {
	// 		where_query += fmt.Sprintf("%s = $%d", key, i)
	// 		i++
	// 	}
	// 	where_query += fmt.Sprintf("%s = $%d AND ", key, i)
	// }
	set_query, i := basicqueryfrommap(set, ",", 1)
	where_query, _ = basicqueryfrommap(where, " AND", i)
	query := fmt.Sprintf(initial_query_template, set_query, where_query, suffix)
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func basicqueryfrommap(m map[string]interface{}, delimiter string, i int) (string, int) {
	query := ""
	counter := 1
	for key := range m {
		query += fmt.Sprintf("%s=$%d", key, i)
		if counter != len(m) {
			query += delimiter
		}
		i++
		// if i == lenofm {
		// 	query += fmt.Sprintf("%s = $%d", key, i)
		// 	i++
		// 	break
		// }
		// query += fmt.Sprintf("%s = $%d%s ", key, i, delimiter)
	}
	return query, i
}

// Find difference between two maps and return it as map.
func (s *Store) ModelDiff(target map[string]interface{}, updated map[string]interface{}) map[string]interface{} {
	var difference map[string]interface{} = make(map[string]interface{})
	for key, val := range updated {
		if updated[key] != val {
			difference[key] = updated[key]
		}
	}
	return difference

}
