CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR (64) NOT NULL,
    description VARCHAR (225),
    author_id INTEGER NOT NULL,
    type_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    sprint_id INTEGER,
    status_id INTEGER NOT NULL,
    priority_id id INTEGER NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (type_id) REFERENCES task_types(id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (sprint_id) REFERENCES sprints(id),
    FOREIGN KEY (status_id) REFERENCES task_statuses(id),
    FOREIGN KEY (priority_id) REFERENCES priority(id)
);