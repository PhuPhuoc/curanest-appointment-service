-- +goose Up
-- +goose StatementBegin
CREATE TABLE `service_tasks_custom` (
    `id` varchar(36) NOT NULL,
    `service_plan_custom_id` varchar(36) NOT NULL,
    `order` smallint NOT NULL,
    `name` varchar(255) NOT NULL,
    `client_note` longtext,
    `staff_advice` longtext,
    `est_duration` int NOT NULL,
    `total_cost` decimal(15,2) NOT NULL,
    `unit` enum('quantity','time') DEFAULT 'quantity',
    `total_unit` int NOT NULL,
    `est_date` date NOT NULL,
    `act_date` date NOT NULL,
    `status` enum('available','unavailable') DEFAULT 'available',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_serviceplancustom_servicetaskcustom` FOREIGN KEY (`service_plan_custom_id`) REFERENCES `service_plans_custom` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `service_tasks_custom`;
-- +goose StatementEnd
