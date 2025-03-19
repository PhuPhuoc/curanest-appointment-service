-- +goose Up
-- +goose StatementBegin
CREATE TABLE `service_packages` (
    `id` varchar(36) NOT NULL,
    `service_id` varchar(36) DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `description` longtext,
    `combo_days` int,
    `discount` int,
    `time_interval` int,
    `status` enum('available','unavailable') DEFAULT 'available',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `servicepackages_serviceid_fk` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `service_packages`;
-- +goose StatementEnd
