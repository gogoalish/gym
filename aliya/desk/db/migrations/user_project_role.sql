CREATE TABLE user_project_role (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
    FOREIGN KEY (project_id) REFERENCES projects(id)
    FOREIGN KEY (role_id) REFERENCES roles(id)
);