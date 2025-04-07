-- +goose Up
-- +goose StatementBegin
CREATE TABLE `medical_records` (
    `id` varchar(36) NOT NULL,
    `nursing_id` varchar(36),
    `customized_package_id` varchar(36) NOT NULL,
    `nursing_report` text,
    `staff_confirmation` text,
    `status` enum('not_done','done') DEFAULT 'not_done',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `medicalrecords_customizedpackageid_fk` FOREIGN KEY (`customized_package_id`) REFERENCES `customized_packages` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `medical_records`;
-- +goose StatementEnd
