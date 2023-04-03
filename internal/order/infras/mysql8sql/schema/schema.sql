CREATE TABLE `xf_order`
(
    `order_id`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `price`        decimal(10, 2) NULL DEFAULT NULL,
    `order_type`   tinyint(0) NULL DEFAULT NULL,
    `order_status` tinyint(0) NULL DEFAULT NULL,
    `delete_is`    tinyint(0) NULL DEFAULT NULL,
    `create_time`  datetime(0) NULL DEFAULT NULL,
    `update_time`  datetime(0) NULL DEFAULT NULL,
    PRIMARY KEY (`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;