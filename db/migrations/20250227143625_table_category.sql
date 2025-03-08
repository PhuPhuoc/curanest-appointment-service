-- +goose Up
-- +goose StatementBegin
CREATE TABLE `categories` (
    `id` varchar(36) NOT NULL,
    `staff_id` varchar(36) DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `description` longtext,
    `thumbnail` varchar(255) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_staff_id` (`staff_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX `unique_staff_id` ON `categories`;
DROP TABLE `categories`;
-- +goose StatementEnd
