-- +goose Up
-- +goose StatementBegin
CREATE TABLE `customized_tasks` (
    `id` varchar(36) NOT NULL,
    `customized_package_id` varchar(36) NOT NULL,
    `task_order` smallint NOT NULL,
    `name` varchar(255) NOT NULL,
    `client_note` longtext,
    `staff_advice` longtext,
    `est_duration` int NOT NULL,
    `total_cost` decimal(15,2) NOT NULL,
    `unit` enum('quantity','time') DEFAULT 'quantity',
    `total_unit` int NOT NULL,
    `est_date` date NOT NULL,
    `act_date` date NOT NULL,
    `status` enum('not_done','done') DEFAULT 'not_done',
    PRIMARY KEY (`id`),
    CONSTRAINT `customizedtasks_cutomizedpackageid_fk` FOREIGN KEY (`customized_package_id`) REFERENCES `customized_packages` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `customized_tasks`;
-- +goose StatementEnd
