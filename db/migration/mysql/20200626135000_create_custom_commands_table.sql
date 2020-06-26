
-- +migrate Up
CREATE TABLE custom_commands (
    id INT NOT NULL AUTO_INCREMENT,
    group_id INT UNSIGNED NOT NULL,
    command VARCHAR(100) NOT NULL,
    message VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL, 
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE custom_commands;