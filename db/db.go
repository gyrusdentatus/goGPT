package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./conversation_history.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, role TEXT, content TEXT);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	return db
}

func addMessage(db *sql.DB, role, content string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO messages(role, content) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(role, content)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}

func getConversationHistory(db *sql.DB) []openai.ChatCompletionMessage {
	rows, err := db.Query("SELECT role, content FROM messages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var messages []openai.ChatCompletionMessage
	for rows.Next() {
		var role, content string
		err = rows.Scan(&role, &content)
		if err != nil {
			log.Fatal(err)
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: content,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return messages
}

