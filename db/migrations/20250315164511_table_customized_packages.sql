-- +goose Up
-- +goose StatementBegin
CREATE TABLE `customized_packages` (
    `id` varchar(36) NOT NULL,
    `service_package_id` varchar(36) NOT NULL,
    `patient_id` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `customizedpackages_servicepackageid_fk` FOREIGN KEY (`service_package_id`) REFERENCES `service_packages` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `customized_packages`;
-- +goose StatementEnd
