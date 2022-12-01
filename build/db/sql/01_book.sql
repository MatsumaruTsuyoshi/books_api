CREATE TABLE IF NOT EXISTS book (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title          VARCHAR(40) NOT NULL,
    price       VARCHAR(100) NOT NULL,
    author       VARCHAR(100) NOT NULL,
    release_date       VARCHAR(100) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
)