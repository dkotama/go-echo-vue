package models

import "database/sql"

type Task struct {
	ID		int		`json:"id"`
	Name 	string 	`json:"name"`
}


type TaskCollection struct {
	Tasks []Task		`json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection  {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	// Error checking
	if err != nil {
		panic(err)
	}

	// cleanup
	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)

		// error check
		if err2 != nil {
			panic(err2)
		}

		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PutTask(db *sql.DB, name string) (int64, error)  {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// create sql statement
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	// cleanup if panic
	defer stmt.Close()

	// Replace '?' with 'name'
	result, err2 := stmt.Exec(name)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// Prepare statement
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
