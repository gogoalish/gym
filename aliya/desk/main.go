package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// func main() {
// 	if len(os.Args) <= 1 || len(os.Args) >= 3 {
// 		log.Fatal("Usage: go run change_db.go <argument>")
// 	}
// 	flag := os.Args[1]
// 	db, err := Connect("./db/servicedesk.db")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer db.Close()

// 	switch flag {
// 	case "create":
// 		if err := CreateTable(db); err != nil {
// 			log.Fatalln(err)
// 		}
// 		log.Println("Successful")
// 	case "drop":
// 		if err := DropAllDB(db); err != nil {
// 			log.Fatalln(err)
// 		}
// 		log.Println("Successful")
// 	default:
// 		log.Fatalf("%s: unknown flag. Use: 'create' or 'drop'", flag)
// 	}
// }

func main() {
	
}

func CreateTable(db *sql.DB) error {
	path := "./db/migrations"

	// reading all direcotries
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range dir {
		info, err := file.Info()
		if err != nil {
			return err
		}
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", path, info.Name()))
		if err != nil {
			return err
		}
		if _, err := db.Exec(string(data)); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func DropAllDB(db *sql.DB) error {
	records := `DROP TABLE IF EXISTS`

	tabls, err := SelectAllTable(db)
	if err != nil {
		return err
	}
	for _, table := range tabls {
		_, err := db.Exec(fmt.Sprintf("%s %s", records, table))
		if err != nil {
			return err
		}
	}
	return nil
}

func SelectAllTable(db *sql.DB) ([]string, error) {
	records := `SELECT name FROM sqlite_master WHERE type='table';`

	stmt, err := db.Prepare(records)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var tabls []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		} else if table == "sqlite_sequence" {
			continue
		}

		tabls = append(tabls, table)
	}
	return tabls, nil
}

func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*
CREATE TABLE `users` (
`id` int,
`first_name` TEXT,
`last_name` TEXT,
`description` TEXT,
`email` TEXT,
`password` TEXT,
`location` TEXT,
`phone` TEXT
);

CREATE TABLE `roles` (
`id` int,
`name` TEXT
);

CREATE TABLE `projects` (
`id` int,
`title` TEXT,
`description` TEXT,
`created` date
);

CREATE TABLE `sprints` (
`id` int,
`project_id` int,
`name` TEXT,
`created` date,
`expires` date,
`description` TEXT
);

CREATE TABLE `tasks` (
`id` int,
`title` TEXT,
`description` TEXT,
`author_id` int,
`type_id` int,
`project_id` int,
`sprint_id` int,
`status_id` int,
`priority_id` int
);

CREATE TABLE `task_statuses` (
`id` int,
`name` TEXT
);

CREATE TABLE `user_project_role` (
`id` int,
`user_id` int,
`project_id` int,
`role_id` int
);

CREATE TABLE `assignee` (
`id` int,
`user_id` int,
`task_id` int
);

CREATE TABLE `types` (
`id` int,
`name` TEXT
);

CREATE TABLE `priority` (
`id` int,
`name` int
);

CREATE TABLE `files` (
`id` int,
`task_id` int,
`name` TEXT,
`uploaded_date` date,
`content_type` TEXT,
`size` int
);
ALTER TABLE `sprints` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);
ALTER TABLE `tasks` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);
ALTER TABLE `tasks` ADD FOREIGN KEY (`sprint_id`) REFERENCES `sprints` (`id`);
ALTER TABLE `tasks` ADD FOREIGN KEY (`status_id`) REFERENCES `task_statuses` (`id`);
ALTER TABLE `user_project_role` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_project_role` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);
ALTER TABLE `user_project_role` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
ALTER TABLE `assignee` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `assignee` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);
ALTER TABLE `tasks` ADD FOREIGN KEY (`type_id`) REFERENCES `types` (`id`);
ALTER TABLE `tasks` ADD FOREIGN KEY (`priority_id`) REFERENCES `priority` (`id`);
ALTER TABLE `tasks` ADD FOREIGN KEY (`author_id`) REFERENCES `users` (`id`);
ALTER TABLE `files` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);





COMMIT;
PRAGMA ignore_check_constraints = ON;
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;
PRAGMA synchronous = NORMAL;
*/
