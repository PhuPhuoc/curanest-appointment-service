-- +goose Up
-- +goose StatementBegin
CREATE TABLE `invoices` (
    `id` varchar(36) NOT NULL,
    `customized_package_id` varchar(36) NOT NULL,
    `total_fee` decimal(15,2) NOT NULL,
    `payment_status` enum('unpaid', 'paid') DEFAULT 'unpaid',
    `note` text,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `invoices_customizedpackageid_fk` FOREIGN KEY (`customized_package_id`) REFERENCES `customized_packages` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `invoices`;
-- +goose StatementEnd
