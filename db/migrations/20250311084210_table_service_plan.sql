-- +goose Up
-- +goose StatementBegin
CREATE TABLE `service_plans` (
    `id` varchar(36) NOT NULL,
    `service_id` varchar(36) DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `description` longtext,
    `combo_days` int,
    `discount` int,
    `status` enum('available','unavailable') DEFAULT 'available',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_service_id` (`service_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `service_plans`;
-- +goose StatementEnd
