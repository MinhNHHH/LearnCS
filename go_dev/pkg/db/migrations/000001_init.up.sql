CREATE TABLE Users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL
);

CREATE TABLE Documents (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL, -- Add user_id column
    is_public BOOLEAN DEFAULT False,
    html_content JSON DEFAULT '{}'::json, -- Ensure JSON default is properly cast
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE -- Remove trailing comma
);


INSERT INTO Users (id, first_name, last_name, password, created_at, updated_at)
VALUES (1, 'Admin', 'User', '123456', '2022-08-19 00:00:00', '2022-08-19 00:00:00');

