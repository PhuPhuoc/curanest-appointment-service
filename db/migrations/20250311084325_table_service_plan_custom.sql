-- +goose Up
-- +goose StatementBegin
CREATE TABLE `service_plans_custom` (
    `id` varchar(36) NOT NULL,
    `service_id` varchar(36) DEFAULT NULL,
    `appointment_id` varchar(36) DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `begin_date` datetime NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_appointment_id` (`appointment_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `service_plans_custom`;
-- +goose StatementEnd
