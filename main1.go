package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Task struct to represent a task record
type Task struct {
	AssignedUserID string    `faker:"uuid_hyphenated"`
	CreatedUserID  string    `faker:"uuid_hyphenated"`
	Title          string    `faker:"sentence"`
	Description    string    `faker:"paragraph"`
	Status         int       `faker:"oneof:1,2,3,4"`
	StartDate      time.Time `faker:"time_this_year"`
	EndDate        time.Time `faker:"time_between_dates=now,30d"`
}

func main() {
	// Connect to your MySQL database
	db, err := sql.Open("mysql", "root:Aa@123456@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	numGoroutines := 10

	recordsPerGoroutine := 1200000 / numGoroutines
	var waitGroup sync.WaitGroup
	waitGroup.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go generateAndInsertFakeRecords(db, i, i+recordsPerGoroutine, &waitGroup)
	}
	waitGroup.Wait()

	fmt.Println("Fake records inserted successfully.")
}

// Function to generate and insert fake records
func generateAndInsertFakeRecords(db *sql.DB, start, end int, wg *sync.WaitGroup) {
	fmt.Println("Starting goroutine", start)
	for i := start; i < end; i++ { // Generate 1.2 million records
		fmt.Printf("============ %d ============ \n", i)
		var task Task

		// Check if Title is empty and set a default value
		if task.Title == "" {
			task.Title = fmt.Sprintf("%s%d", "Default Title", i)
		}

		// Check if Description is empty and set a default value
		if task.Description == "" {
			task.Description = fmt.Sprintf("%s%d", "Default Description", i)
		}

		// Check for invalid start_date values and replace with a valid default
		if task.StartDate.IsZero() {
			task.StartDate = time.Now()
		}
		if task.EndDate.IsZero() {
			task.EndDate = time.Now()
		}
		task.CreatedUserID = "509e2b5d-ae95-11ee-9fb6-309c23d934ba"

		// Insert the fake record into the tasks table
		_, err := db.Exec("INSERT INTO tasks (assigned_user_id, created_user_id, title, description, status, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?, ?)",
			task.AssignedUserID, task.CreatedUserID, task.Title, task.Description, task.Status, task.StartDate, task.EndDate)
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Done()
	fmt.Println("Fake records inserted successfully.")
}
