DROP TABLE IF EXISTS `goods`;
CREATE TABLE goods (
                       goods_id int(11) UNSIGNED NOT NULL COMMENT '商品id',
                       goods_name varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
                       selling_point varchar(500) NOT NULL DEFAULT '' COMMENT '商品卖点',
                       category_id int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商品分类id',
                       deduct_stock_type tinyint(3) UNSIGNED NOT NULL DEFAULT '20' COMMENT '库存计算方式(10下单减库存 20付款减库存)',
                       content longtext NOT NULL COMMENT '商品详情',
                       sales_initial int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '初始销量',
                       sales_actual int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '实际销量',
                       goods_sort int(11) UNSIGNED NOT NULL DEFAULT '100' COMMENT '商品排序(数字越小越靠前)',
                       goods_status tinyint(3) UNSIGNED NOT NULL DEFAULT '10' COMMENT '商品状态(10上架 20下架)',
                       is_delete tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除',
                       create_time int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                       update_time int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品记录表';

DROP TABLE IF EXISTS `goods_category`;
CREATE TABLE `goods_category` (
                                  `category_id` int(11) UNSIGNED NOT NULL COMMENT '商品分类id',
                                  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '分类名称',
                                  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '上级分类id',
                                  `image_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类图片id',
                                  `sort` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序方式(数字越小越靠前)',
                                  `create_time` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                                  `update_time` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品分类表';

DROP TABLE IF EXISTS `goods_upload_file`;
CREATE TABLE goods_upload_file (
                                   file_id int(11) UNSIGNED NOT NULL COMMENT '文件id',
                                   storage varchar(20) NOT NULL DEFAULT '' COMMENT '存储方式',
                                   file_url varchar(255) NOT NULL DEFAULT '' COMMENT '存储域名',
                                   file_name varchar(255) NOT NULL DEFAULT '' COMMENT '文件路径',
                                   file_size int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文件大小(字节)',
                                   file_type varchar(20) NOT NULL DEFAULT '' COMMENT '文件类型',
                                   extension varchar(20) NOT NULL DEFAULT '' COMMENT '文件扩展名',
                                   is_user int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否为c端用户上传',
                                   is_recycle tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已回收',
                                   is_delete tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '软删除',
                                   create_time int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片记录表';

DROP TABLE IF EXISTS `goods_image`;
CREATE TABLE goods_image (
                             id int(11) UNSIGNED NOT NULL COMMENT '主键id',
                             goods_id int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商品id',
                             image_id int(11) NOT NULL COMMENT '图片id(关联图片记录表)',
                             create_time int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品图片记录表';

DROP TABLE IF EXISTS `goods_spec`;
CREATE TABLE goods_spec (
                            spec_id int(11) UNSIGNED NOT NULL COMMENT '规格组id',
                            spec_name varchar(255) NOT   NULL DEFAULT '' COMMENT '规格组名称',
                            create_time int(11) NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品规格组记录表';


DROP TABLE IF EXISTS `goods_sku`;
CREATE TABLE goods_sku (
                           goods_sku_id int(11) UNSIGNED NOT NULL COMMENT '商品规格id',
                           goods_id int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商品id',
                           spec_sku_id varchar(255) NOT NULL DEFAULT '0' COMMENT '商品sku记录索引 (由规格id组成)',
                           image_id int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '图片id',
                           goods_no varchar(100) NOT NULL DEFAULT '' COMMENT '商品编码',
                           goods_price decimal(10,2) UNSIGNED NOT NULL DEFAULT '0.00' COMMENT '商品价格',
                           line_price decimal(10,2) UNSIGNED NOT NULL DEFAULT '0.00' COMMENT '商品划线价',
                           stock_num int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '当前库存数量',
                           goods_sales int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商品销量',
                           goods_weight double UNSIGNED NOT NULL DEFAULT '0' COMMENT '商品重量(Kg)',
                           create_time int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                           update_time int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品规格表';\

drop table IF exists `spec_value`;
CREATE TABLE spec_value (
                            spec_value_id int(11) UNSIGNED NOT NULL COMMENT '规格值id',
                            spec_value varchar(255) NOT NULL COMMENT '规格值',
                            spec_id int(11) NOT NULL COMMENT '规格组id',
                            create_time int(11) NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品规格值记录表';