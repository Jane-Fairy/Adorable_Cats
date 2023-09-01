CREATE TABLE `Product`
(
    `id`          VARCHAR(255) PRIMARY KEY COMMENT '商品ID，唯一标识符',
    `product_name`        VARCHAR(255)   NOT NULL COMMENT '商品名称或标题',
    `product_description` TEXT COMMENT '商品详细描述',
    `price`              DECIMAL(10, 2) NOT NULL COMMENT '商品价格',
    `stock_quantity`      INT            NOT NULL COMMENT '商品当前库存数量',
    `manufacturer`       VARCHAR(255) COMMENT '制造商或生产公司',
    `product_category`    VARCHAR(255) COMMENT '商品所属类别或分类',
    `image_url`           VARCHAR(255) COMMENT '商品图片链接，用于展示商品图片',
    `created_date`        DATE COMMENT '商品创建日期',
    `updated_date`        DATE COMMENT '商品最后更新日期',
    `delete_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `del_state`     tinyint  NOT NULL DEFAULT '0' COMMENT '删除状态',
    `version`       bigint   NOT NULL DEFAULT '0' COMMENT '版本号',
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='宠物信息基础表';
SET FOREIGN_KEY_CHECKS = 1;



