-- +goose Up
-- +goose StatementBegin
CREATE TABLE `appointments` (
    `id` varchar(36) NOT NULL,
    `service_id` varchar(36) NOT NULL,
    `customized_package_id` varchar(36) NOT NULL,
    `nursing_id` varchar(36),
    `patient_id` varchar(36) NOT NULL,
    `patient_address` text NOT NULL,
    `patient_lat_lng` varchar(255) NOT NULL,
    `est_date` datetime NOT NULL,
    `act_date` datetime,
    `total_est_duration` int NOT NULL,
    `status` enum('waiting', 'confirmed', 'upcoming', 'success', 'cancel') DEFAULT 'waiting',
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
