-- +goose Up
-- +goose StatementBegin
CREATE TABLE `appointments` (
    `id` varchar(36) NOT NULL,
    `service_id` varchar(36) NOT NULL,
    `customized_package_id` varchar(36) NOT NULL,
    `nursing_id` varchar(36) NOT NULL,
    `patient_id` varchar(36) NOT NULL,
    `appointment_date` date NOT NULL,
    `est_start_time` time NOT NULL,
    `est_total_time` time NOT NULL,
    `status` enum('success', 'nurse_duplicated', 'nurse_reassigned', 'customer_refused') DEFAULT 'success',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `appointments_serviceid_fk` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`) ON UPDATE CASCADE,
    CONSTRAINT `appointments_customizedpackageid_fk` FOREIGN KEY (`customized_package_id`) REFERENCES `customized_packages` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `appointments`;
-- +goose StatementEnd
