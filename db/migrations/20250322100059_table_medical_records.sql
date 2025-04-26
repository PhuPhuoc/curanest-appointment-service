-- +goose Up
-- +goose StatementBegin
CREATE TABLE `medical_records` (
    `id` varchar(36) NOT NULL,
    `nursing_id` varchar(36),
    `appointment_id` varchar(36) NOT NULL,
    `nursing_report` text,
    `staff_confirmation` text,
    `status` enum('not_done','done') DEFAULT 'not_done',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `medicalrecords_appointmentid_fk` FOREIGN KEY (`appointment_id`) REFERENCES `appointments` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `medical_records`;
-- +goose StatementEnd
