package model

import "log"

type Todo struct {
	Id    int64
	Title string
}

func InsertTodo(title string) (todo Todo, err error) {
	insertStatement := "INSERT INTO TODO (TITLE) VALUES ( ? )"
	insertStmt, err := db.Prepare(insertStatement)
	if err != nil {
		log.Println("prepare failure: ", err)
		return
	}
	defer insertStmt.Close()

	res, err := insertStmt.Exec(title)
	if err != nil {
		log.Println("exec failure: ", err)
		return
	}
	todo.Id, err = res.LastInsertId()
	todo.Title = title
	return
}

func GetTodoList() (todos []Todo, err error) {
	rows, err := db.Query("SELECT ID, TITLE from TODO")
	if err != nil {
		log.Println("query failure: ", err)
		return
	}
	defer rows.Close()

	todos = []Todo{}
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Title)
		if err != nil {
			log.Println("scan failure: ", err)
			return
		}
		todos = append(todos, todo)
	}
	return
}

func GetTodoById(id int) (todo Todo, err error) {
	stmt, err := db.Prepare("SELECT ID, TITLE from TODO WHERE ID = ? ")
	if err != nil {
		log.Println("prepare failure: ", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&todo.Id, &todo.Title)
	return
}

func DeleteTodoById(id int) (err error) {
	stmt, err := db.Prepare("DELETE FROM TODO WHERE ID = ? ")
	if err != nil {
		log.Println("prepare failure: ", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return
}
