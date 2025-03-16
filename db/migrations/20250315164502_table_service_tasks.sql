-- +goose Up
-- +goose StatementBegin
CREATE TABLE `service_tasks` (
    `id` varchar(36) NOT NULL,
    `service_package_id` varchar(36) NOT NULL,
    `is_must_have` bool NOT NULL,
    `order` smallint NOT NULL,
    `name` varchar(255) NOT NULL,
    `description` longtext,
    `staff_advice` longtext,
    `est_duration` int NOT NULL,
    `cost` decimal(15,2) NOT NULL,
    `additional_cost` decimal(15,2) NOT NULL,
    `additional_cost_desc` text,
    `unit` enum('quantity','time') DEFAULT 'quantity',
    `price_of_step` int NOT NULL,
    `status` enum('available','unavailable') DEFAULT 'available',
    PRIMARY KEY (`id`),
    CONSTRAINT `servicetasks_servicepackageid_fk` FOREIGN KEY (`service_package_id`) REFERENCES `service_packages` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `service_tasks`;
-- +goose StatementEnd
