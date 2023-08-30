DROP TABLE IF EXISTS `pet_info`;
CREATE TABLE pet_info
(
    `id`            VARCHAR(255) PRIMARY KEY COMMENT '宠物ID，主键',
    `name`          VARCHAR(255) COMMENT '宠物名称',
    `type`          VARCHAR(50) COMMENT '宠物类型',
    `breed`         VARCHAR(50) COMMENT '宠物品种',
    `date_of_birth` DATE COMMENT '出生日期',
    `gender`        VARCHAR(10) COMMENT '性别',
    `owner_name`    VARCHAR(255) COMMENT '主人姓名',
    `owner_contact` VARCHAR(255) COMMENT '主人联系方式',
    `create_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `del_state`     tinyint  NOT NULL DEFAULT '0' COMMENT '删除状态',
    `version`       bigint   NOT NULL DEFAULT '0' COMMENT '版本号',
    `note`          TEXT COMMENT '备注'

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='宠物信息基础表';
SET FOREIGN_KEY_CHECKS = 1;

