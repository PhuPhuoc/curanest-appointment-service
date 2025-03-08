-- +goose Up
-- +goose StatementBegin
CREATE TABLE `services` (
    `id` varchar(36) NOT NULL,
    `category_id` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `description` longtext,
    `est_duration` varchar(255) NOT NULL,
    `status` enum('available','unavailable') DEFAULT 'available',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`),
    CONSTRAINT `fk_services_categories` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `services` DROP FOREIGN KEY `fk_services_categories`;
DROP INDEX `idx_category_id` ON `services`;
DROP TABLE `services`;
-- +goose StatementEnd
