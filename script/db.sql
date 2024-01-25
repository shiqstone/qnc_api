CREATE TABLE `qnc_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `avatar_url` varchar(256) NOT NULL DEFAULT '' COMMENT '用户头像url',
  `devid` varchar(64) NOT NULL DEFAULT '' COMMENT '设备id',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `salt` varchar(64) NOT NULL DEFAULT '' COMMENT '盐码',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `invite_code` varchar(6) NOT NULL DEFAULT '' COMMENT '邀请码',
  `coin` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '硬币',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:黑名单（强制退出 禁止登录） 3:已注销 ',
  `user_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '类型 0 游客 1 普通 2 机器人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_user_email` (`email`) USING BTREE
) ENGINE=InnoDB  AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COMMENT='用户记录';

CREATE TABLE `qnc_user_reginfo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `tdate` int(11) NOT NULL DEFAULT '0' COMMENT '日期',
  `ip` varchar(32) DEFAULT NULL COMMENT '用户ip',
  `location` varchar(128) DEFAULT NULL COMMENT '地区',
  `ch` varchar(64) DEFAULT '' COMMENT '渠道号',
  `vname` int DEFAULT '0'  COMMENT 'build版本号',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_user_reginfo` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户注册记录';

CREATE TABLE `qnc_coin_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '类型',
  `event_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '对应事件id',
  `incoming` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '增加',
  `expend` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '扣减',
  `balance` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '余额',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  `remark` varchar(512) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='硬币明细表';

CREATE TABLE `quc_deposit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `deposit_id` varchar(32) NOT NULL DEFAULT '' COMMENT '充值订单号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `amount` decimal(20,2) NOT NULL DEFAULT '0' COMMENT '充值金额',
  `currency` varchar(16) DEFAULT '' COMMENT '币种',
  `pay_mode` varchar(20) NOT NULL DEFAULT '' COMMENT '充值方式',
  `pay_bank` varchar(20) NOT NULL DEFAULT '' COMMENT '充值机构',
  `ip` varchar(32) DEFAULT NULL COMMENT '请求ip',
  `bank_trade_no` varchar(64) DEFAULT NULL COMMENT '渠道单号',
  `pay_channel` int(11) NOT NULL DEFAULT '0' COMMENT '充值渠道',
  `finish_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '完成时间',
  `total_refund` decimal(20,2) DEFAULT '0' COMMENT '退款金额',
  `ext` varchar(255) DEFAULT NULL COMMENT '扩展信息',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `deposit_id` (`deposit_id`) USING BTREE,
  KEY `idx_deposit_ctime` (`create_time`) USING BTREE,
  KEY `idx_status_finish_time` (`status`,`finish_time`) USING BTREE,
  KEY `idx_deposit_qid` (`uid`,`deposit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户充值记录记录';

CREATE TABLE `qnc_kv` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(64) NOT NULL DEFAULT '' COMMENT '名称',
  `value` text NOT NULL COMMENT '数据',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  `status` char(1) NOT NULL DEFAULT 'N' COMMENT '数据状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_kv_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `qnc_invite_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `invite_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '被邀请用户id',
  `invite_devid` varchar(32) NOT NULL DEFAULT '' COMMENT '被邀请设备号',
  `ip` varchar(32) DEFAULT NULL COMMENT '被邀请人注册ip',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户邀请记录';

CREATE TABLE `quc_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `prod_name` varchar(50) NOT NULL COMMENT '商品名',
  `prod_id` int(11) NOT NULL COMMENT '商品id',
  `real_cost` decimal(20,2) NOT NULL DEFAULT '0' COMMENT '实际消耗',
  `base_cost` decimal(20,2) NOT NULL  COMMENT '标价',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `ip` varchar(32) DEFAULT NULL COMMENT '请求ip',
  `remark` varchar(512) DEFAULT NULL COMMENT '备注',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `inx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单记录';