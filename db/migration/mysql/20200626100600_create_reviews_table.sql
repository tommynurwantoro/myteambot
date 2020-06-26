
-- +migrate Up
CREATE TABLE reviews (
    id INT NOT NULL AUTO_INCREMENT,
    group_id INT UNSIGNED NOT NULL,
    url VARCHAR(255) NOT NULL,
    title VARCHAR(100) NOT NULL,
    is_reviewed BOOLEAN NOT NULL,
    is_tested BOOLEAN NOT NULL,
    users VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE reviews;