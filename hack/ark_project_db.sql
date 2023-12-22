/*
 Navicat Premium Data Transfer

 Source Server         : himi_prod
 Source Server Type    : MySQL
 Source Server Version : 80013 (8.0.13)
 Source Host           : pc-t4n0woj13om5jbza5.rwlb.singapore.rds.aliyuncs.com:3306
 Source Schema         : himi_db

 Target Server Type    : MySQL
 Target Server Version : 80013 (8.0.13)
 File Encoding         : 65001

 Date: 21/12/2023 20:36:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for active_person_statistics
-- ----------------------------
DROP TABLE IF EXISTS `active_person_statistics`;
CREATE TABLE `active_person_statistics` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `begin_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `app_channel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道',
  `app_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app版本',
  `active_type` int(11) NOT NULL DEFAULT '1' COMMENT '类型1日活，2周活，3月活',
  `active_users_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '活跃用户',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1862 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for activity_channel_config
-- ----------------------------
DROP TABLE IF EXISTS `activity_channel_config`;
CREATE TABLE `activity_channel_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ac_id` int(11) NOT NULL DEFAULT '0' COMMENT '活动Id activity_config 自增id',
  `show_channel` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'ALL' COMMENT '渠道：ALL全部渠道、指定渠道',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=212 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for activity_config
-- ----------------------------
DROP TABLE IF EXISTS `activity_config`;
CREATE TABLE `activity_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '活动标题',
  `res_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '资源图 活动横图',
  `link_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '活动地址',
  `start_time` bigint(11) NOT NULL DEFAULT '1451577600' COMMENT '活动开始时间',
  `end_time` bigint(11) NOT NULL DEFAULT '1451577600' COMMENT '活动结束时间',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '状态 0下架 1上架',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序，顺序排列',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for activity_page_config
-- ----------------------------
DROP TABLE IF EXISTS `activity_page_config`;
CREATE TABLE `activity_page_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ac_id` int(11) NOT NULL DEFAULT '0' COMMENT '活动Id activity_config 自增id',
  `res_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '资源地址',
  `duration` int(11) NOT NULL DEFAULT '0' COMMENT '持续时间 单位秒 >0才倒计时',
  `link_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `show_style` tinyint(8) NOT NULL DEFAULT '0' COMMENT '0无配置 1全屏 2半屏 后续其他扩展',
  `page_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '页面类别：启动页(start_up)、首页(home_page)、直播间(room_live)、后续扩展',
  `page_location` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '页面位置：启动页：只有一个位置不用配置，首页：弹窗（pop_up）、浮窗（float）；直播间：弹窗（pop_up）、浮窗（float）、底部图标（bottom_icon）、直播间-礼物（gift_top）；',
  `start_time` bigint(11) NOT NULL DEFAULT '0' COMMENT '活动开始时间',
  `end_time` bigint(11) NOT NULL DEFAULT '0' COMMENT '活动结束时间',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for activity_plan
-- ----------------------------
DROP TABLE IF EXISTS `activity_plan`;
CREATE TABLE `activity_plan` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `activity_type` int(11) NOT NULL DEFAULT '0' COMMENT '活动类别,榜单活动',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '活动标题',
  `icon` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `explain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '活动说明',
  `start_time` bigint(11) NOT NULL DEFAULT '1451577600' COMMENT '活动开始时间',
  `end_time` bigint(11) NOT NULL DEFAULT '1451577600' COMMENT '活动结束时间',
  `time_level` int(11) NOT NULL DEFAULT '0' COMMENT '活动数据类型,（自然周、自然天、时）也可以自定义时间类型（最小颗粒度：小时）、周期内（all）：0',
  `reward_cycle` int(11) NOT NULL DEFAULT '0' COMMENT '奖励发放周期配置：小时、天、周、默认0不奖励、统一指定时间奖励',
  `object` varchar(50) NOT NULL DEFAULT '0' COMMENT '奖励对象：主播/公会长/用户 或者多个',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖励数据源：rank_diamonds、rank_coin、rank_gift(根据礼物id判断)……,以此类推；根据key标记调用对应的封装接口 不配置默认为不调用数据源（特殊活动/非模板活动）',
  `app_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '能参加活动的渠道',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '活动状态是否有效，1有效，2无效',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for activity_reward_for_guild
-- ----------------------------
DROP TABLE IF EXISTS `activity_reward_for_guild`;
CREATE TABLE `activity_reward_for_guild` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户Id 主播id',
  `guild_id` int(11) NOT NULL COMMENT '公会Id 关联guild_info自增',
  `activity_type` tinyint(8) NOT NULL DEFAULT '0' COMMENT '活动类型  0默认 1首播活动 2四周活动 后续扩展',
  `effect_live_time` bigint(11) NOT NULL DEFAULT '0' COMMENT '有效直播时长（s） 首播、四周活动存储热门直播时长',
  `effect_day` int(11) NOT NULL DEFAULT '0' COMMENT '有效直播天',
  `is_reward` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否已经奖励 0未发放 1已发放',
  `reward` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励、存储印尼盾分，后续其他使用可以存储别的 根据不同的活动类型',
  `reward_date` bigint(10) NOT NULL DEFAULT '0' COMMENT '发放日期',
  `created_time` bigint(10) NOT NULL DEFAULT '1451577600',
  `updated_time` bigint(10) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_id_activity_type` (`user_id`,`activity_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6155 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='公会活动奖励记录';

-- ----------------------------
-- Table structure for activity_reward_for_pull_guild
-- ----------------------------
DROP TABLE IF EXISTS `activity_reward_for_pull_guild`;
CREATE TABLE `activity_reward_for_pull_guild` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户Id 主播id',
  `guild_id` int(11) NOT NULL COMMENT '公会Id 关联guild_info自增',
  `guild_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户所属公会长Id （这里是B公会长ID）',
  `puller_guild_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'A公会拉的B公会 这里存储A公会长的ID',
  `live_time` bigint(11) NOT NULL DEFAULT '0' COMMENT '直播时长 单位秒',
  `is_reward` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否已经奖励 0未发放 1已发放',
  `reward` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励、存储印尼盾分 奖励给B公会长的',
  `reward_date` bigint(10) NOT NULL DEFAULT '0' COMMENT '发放日期',
  `created_time` bigint(10) NOT NULL DEFAULT '1451577600',
  `updated_time` bigint(10) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='公会拉公会奖励记录';

-- ----------------------------
-- Table structure for activity_user_config
-- ----------------------------
DROP TABLE IF EXISTS `activity_user_config`;
CREATE TABLE `activity_user_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ac_id` int(11) NOT NULL DEFAULT '0' COMMENT '活动Id activity_config 自增id',
  `show_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'ALL' COMMENT '展示对象：ALL 全部、UNPAID未充值用户 、PAID已经充值用户 、HOST主播公会 后续扩展',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for adjust_message
-- ----------------------------
DROP TABLE IF EXISTS `adjust_message`;
CREATE TABLE `adjust_message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `install_time` bigint(20) NOT NULL DEFAULT '0' COMMENT 'install_time',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `order_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `event_token` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `ga_id` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '订单id',
  `device_id` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '设备id',
  `app_name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `app_id` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `tracker` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `network` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `adjust_id` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `currency` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `revenue` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_install_time` (`install_time`),
  KEY `idx_app_name` (`app_name`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_revenue` (`revenue`),
  KEY `idx_ga_id` (`ga_id`)
) ENGINE=InnoDB AUTO_INCREMENT=212020354 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
 
-- ----------------------------
-- Table structure for agent_balance_info
-- ----------------------------
DROP TABLE IF EXISTS `agent_balance_info`;
CREATE TABLE `agent_balance_info` (
  `user_id` bigint(20) NOT NULL,
  `total_fee` bigint(20) NOT NULL COMMENT '总充值 印尼盾 单位：分',
  `total_sale_fee` bigint(20) NOT NULL COMMENT '总销售 印尼盾 单位：分',
  `total_buy_diamonds` bigint(20) NOT NULL COMMENT '总购买钻石',
  `total_sale_diamonds` bigint(20) NOT NULL COMMENT '总销售钻石',
  `remain_diamonds` bigint(20) NOT NULL COMMENT '剩余钻石数',
  `pay_password` varchar(255) NOT NULL COMMENT '支付密码',
  `cost_price` decimal(10,2) NOT NULL COMMENT '成本单价',
  `is_lock` int(11) NOT NULL DEFAULT '0' COMMENT '是否锁定',
  `agent_states` int(11) NOT NULL DEFAULT '1' COMMENT '代理状态 1正常',
  `agent_rank` int(11) NOT NULL DEFAULT '3' COMMENT '代理等级 1 2 3',
  `agent_ban_time` bigint(20) NOT NULL COMMENT '禁止转账截止时间',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`user_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for agent_balance_record
-- ----------------------------
DROP TABLE IF EXISTS `agent_balance_record`;
CREATE TABLE `agent_balance_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `target_user_id` bigint(20) NOT NULL,
  `record_type` int(11) NOT NULL COMMENT '消费类型 1支出 2收入',
  `one_price` decimal(10,2) NOT NULL COMMENT '售出单价',
  `cost_price` decimal(10,2) NOT NULL COMMENT '成本单价',
  `diamonds` int(11) NOT NULL COMMENT '钻石数额',
  `before_change` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变更之前余额',
  `after_change` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变更之后余额',
  `link_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '引起余额变更的id',
  `certificate_url` varchar(255) NOT NULL COMMENT '凭证url',
  `payment_id` int(11) NOT NULL COMMENT '支付id',
  `payment_name` varchar(255) NOT NULL COMMENT '支付名称',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_record_type` (`record_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16930 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for agent_detail_record
-- ----------------------------
DROP TABLE IF EXISTS `agent_detail_record`;
CREATE TABLE `agent_detail_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `operate_user_id` int(11) NOT NULL COMMENT '操作用户id',
  `change_type` int(11) NOT NULL COMMENT '触发类型 1 后台 2 系统',
  `update_type` int(11) NOT NULL COMMENT '更新类型 1 代理等级 2 支付密码 3 禁止转账截止时间 4 是否锁定 5 代理状态',
  `before_change` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '变更之前值',
  `after_change` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '变更之后值',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '更新备注',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for agent_giveaway_props
-- ----------------------------
DROP TABLE IF EXISTS `agent_giveaway_props`;
CREATE TABLE `agent_giveaway_props` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '直播间Id',
  `user_id` bigint(20) NOT NULL COMMENT '主播ID',
  `props_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '道具名称',
  `props_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '道具icon',
  `props_day` int(11) NOT NULL COMMENT '道具天数 30',
  `props_states` int(11) NOT NULL COMMENT '道具状态 1 待使用 2 已使用',
  `props_type` int(11) NOT NULL COMMENT '道具类型 1 礼物 2 vip 3 装扮',
  `props_link` int(11) NOT NULL COMMENT '道具关联id',
  `balance_record_id` bigint(20) NOT NULL COMMENT '流水记录id',
  `use_time` bigint(20) NOT NULL COMMENT '使用时间',
  `target_user_id` int(11) NOT NULL COMMENT '使用目标id',
  `expiration_time` bigint(20) NOT NULL COMMENT '过期时间 -1 不过期',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=449 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for agent_opera_record
-- ----------------------------
DROP TABLE IF EXISTS `agent_opera_record`;
CREATE TABLE `agent_opera_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `opera_ip` varchar(255) NOT NULL COMMENT '操作ip',
  `link_id` varchar(255) NOT NULL COMMENT '关联id',
  `user_agent` varchar(255) NOT NULL COMMENT '系统类型',
  `opera_type` int(11) NOT NULL COMMENT '操作类型 1 登录 2 转账 3 充值',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18063 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for anchor_day_statistics
-- ----------------------------
DROP TABLE IF EXISTS `anchor_day_statistics`;
CREATE TABLE `anchor_day_statistics` (
  `statistics_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `statistics_time` date NOT NULL COMMENT '统计时间',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '工会id',
  `all_income` bigint(20) NOT NULL COMMENT '总收益',
  `normal_gift_income` bigint(20) NOT NULL COMMENT '普通礼物收益',
  `lucky_gift_income` bigint(20) NOT NULL COMMENT '幸运礼物收益',
  `live_time` int(11) NOT NULL COMMENT '直播时长（s）',
  `effect_live_time` int(11) NOT NULL DEFAULT '0' COMMENT '有效直播时长（s） 大于30分钟的直播才算有效直播累计',
  `is_effect_day` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否有效天 0非 1是',
  `pk_count` int(11) NOT NULL COMMENT '总pk次数',
  `effective_pk_count` int(20) NOT NULL COMMENT '有效pk次数',
  `watch_member_count` int(11) NOT NULL DEFAULT '0' COMMENT '观看人数',
  `new_fans_count` int(11) NOT NULL DEFAULT '0' COMMENT '新增粉丝数',
  `send_gift_person` int(11) NOT NULL DEFAULT '0' COMMENT '新增送礼人数',
  `is_first_live` int(2) NOT NULL COMMENT '是否首播',
  `role` int(2) NOT NULL COMMENT '角色',
  `reward_day_for_guild` bigint(20) NOT NULL DEFAULT '0' COMMENT '当天对公会的奖励（印尼盾 单位分），比如主播当天首播满1小时给公会的奖励',
  `reward_day_for_guild_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动奖励说明',
  `sum_hot_award` int(11) NOT NULL COMMENT '热门钻石奖励',
  `hot_effect_live` bigint(20) NOT NULL DEFAULT '0' COMMENT '热门开播',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `day_water_diamond` bigint(20) NOT NULL DEFAULT '0' COMMENT '直播天流水',
  PRIMARY KEY (`statistics_id`) USING BTREE,
  UNIQUE KEY `statistics_UNIQUE` (`statistics_time`,`user_id`) USING BTREE,
  KEY `idx_time` (`statistics_time`) USING BTREE,
  KEY `idx_guildid_userid_effectlivetime_statisticstime` (`guild_id`,`user_id`,`effect_live_time`,`statistics_time`),
  KEY `idx_userid_statisticstime_effectlivetime_iseffectday` (`user_id`,`statistics_time`,`effect_live_time`,`is_effect_day`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_guildid_statisticstime` (`guild_id`,`statistics_time`),
  KEY `idx_day_water_diamond` (`day_water_diamond`)
) ENGINE=InnoDB AUTO_INCREMENT=2077775 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='主播收益统计表';

-- ----------------------------
-- Table structure for anchor_game_log
-- ----------------------------
DROP TABLE IF EXISTS `anchor_game_log`;
CREATE TABLE `anchor_game_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `anchor_id` bigint(20) NOT NULL COMMENT '主播ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `room_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '第三方直播间ID全局唯一',
  `game_id` int(10) NOT NULL DEFAULT '0' COMMENT '直播间配置的游戏id',
  `diamonds` bigint(20) unsigned NOT NULL COMMENT '钻石数',
  `anchor_rate` int(11) NOT NULL DEFAULT '0' COMMENT '主播分成  *10000的整数 最小为万分之一',
  `before_diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '当前钻石数',
  `log_type` int(11) NOT NULL COMMENT '流水类型 1 增加 2 扣除',
  `after_diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '增加之后的钻石数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_anchorId` (`anchor_id`) USING BTREE,
  KEY `idx_room_id` (`room_id`) USING BTREE,
  KEY `idx_diamonds` (`diamonds`)
) ENGINE=InnoDB AUTO_INCREMENT=8577444 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='主播游戏流水表';

-- ----------------------------
-- Table structure for anchor_label_apply
-- ----------------------------
DROP TABLE IF EXISTS `anchor_label_apply`;
CREATE TABLE `anchor_label_apply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签',
  `states` int(11) NOT NULL COMMENT '状态 1 待审核 2 已通过 3 已驳回',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28499 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for anchor_level
-- ----------------------------
DROP TABLE IF EXISTS `anchor_level`;
CREATE TABLE `anchor_level` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `level` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '等级',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '等级',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '状态1有效，2无效',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_level` (`level`) USING BTREE,
  KEY `idx_state` (`state`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10079 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for anchor_live_time_rank
-- ----------------------------
DROP TABLE IF EXISTS `anchor_live_time_rank`;
CREATE TABLE `anchor_live_time_rank` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `rank_date_time` date NOT NULL COMMENT '排行日期',
  `effect_live_seconds` int(11) NOT NULL COMMENT '有效直播时长',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id_time` (`user_id`,`rank_date_time`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_rank_date_time` (`rank_date_time`)
) ENGINE=InnoDB AUTO_INCREMENT=90800 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for anchor_rank_rule
-- ----------------------------
DROP TABLE IF EXISTS `anchor_rank_rule`;
CREATE TABLE `anchor_rank_rule` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL DEFAULT '0' COMMENT '等级从0开始, 0为等级1',
  `min` bigint(20) NOT NULL DEFAULT '0' COMMENT '最小值',
  `max` bigint(20) NOT NULL DEFAULT '0' COMMENT '最大值',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8mb4 COMMENT='主播等级规则';

-- ----------------------------
-- Table structure for anchor_warning_record
-- ----------------------------
DROP TABLE IF EXISTS `anchor_warning_record`;
CREATE TABLE `anchor_warning_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `icon` varchar(500) NOT NULL DEFAULT '' COMMENT '警告图片证据',
  `appeal` varchar(500) NOT NULL DEFAULT '' COMMENT '申诉',
  `state` int(6) NOT NULL DEFAULT '0' COMMENT '0待申诉，1申诉中，2成功，3失败',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播id',
  `room_id` varchar(500) NOT NULL DEFAULT '' COMMENT '直播间id',
  `warn_type` int(6) NOT NULL DEFAULT '0' COMMENT '警告类型',
  `reason` varchar(500) NOT NULL DEFAULT '' COMMENT '警告原因',
  `warn_big_type` int(11) NOT NULL DEFAULT '0',
  `admin_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '提醒大类',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3328 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for anchor_week_settlement
-- ----------------------------
DROP TABLE IF EXISTS `anchor_week_settlement`;
CREATE TABLE `anchor_week_settlement` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `week` int(11) NOT NULL DEFAULT '0' COMMENT '周，一年中的第几周 例如:202230 2022年的第30周',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会ID',
  `cut_rate` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '公会抽成 设置范围0-100',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `total_income` bigint(20) NOT NULL DEFAULT '0' COMMENT '总收益 单位分',
  `live_time` int(11) NOT NULL DEFAULT '0' COMMENT '直播时长（s）',
  `effect_live_day` tinyint(8) NOT NULL DEFAULT '0' COMMENT '有效出勤天数',
  `star_level` tinyint(8) NOT NULL DEFAULT '0' COMMENT '星级',
  `is_first_live` tinyint(4) NOT NULL DEFAULT '0' COMMENT '本周是否首播 0非首播；1首播',
  `settle_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '结算状态：\r\n1可以结算  达到结算要求的标星主播或有公会的主播\r\n0：兑换为钻石（不可以结算）；',
  `surplus_income` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '不结算时兑换完钻石后，剩余金币，结算时这里存储0',
  `exchange_diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '结算时兑换钻石数，不结算的话这里存储0',
  `income_gear` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '金币档位 例如：0-2500000',
  `gear_condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '档位条件；例如：周1至周3',
  `reward_formula_for_host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT ' 例如：coin*0.08*50% ',
  `host_rate` int(11) NOT NULL DEFAULT '0' COMMENT '对应中reward_formula_for_host的最后奖励比例，比如：50% *100的整数',
  `reward_money_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播金币档位奖励给主播 印尼盾 分',
  `gift_income` bigint(30) unsigned NOT NULL DEFAULT '0' COMMENT '礼物收入=reward_formula_for_host 印尼盾 分',
  `guild_rate` int(11) NOT NULL DEFAULT '0' COMMENT '公会礼物提成比列  *100的整数',
  `guild_gift_income` bigint(30) unsigned NOT NULL DEFAULT '0' COMMENT '公会礼物提成收入 印尼盾 分',
  `week_rank` int(11) NOT NULL DEFAULT '0' COMMENT '周排行榜 标星主播按照金币又大到小排名 ',
  `week_rank_reward_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '周榜主播奖励（印尼盾）同时需要按照例如：周1至周3工作几天时长多久来发放，如果排行榜靠前，时长或出勤天未达到同样不发放奖励  印尼盾 分',
  `week_rank_reward_for_guild_rate` int(11) NOT NULL DEFAULT '0' COMMENT '周榜公会获得提成比率 *100的整数',
  `week_rank_reward_for_guild` bigint(20) NOT NULL DEFAULT '0' COMMENT '周榜公会获得提成收入 印尼盾 分',
  `host_total_income_cut_before` bigint(30) unsigned NOT NULL DEFAULT '0' COMMENT '公会抽成前主播总收入 印尼盾 分',
  `host_total_income_cut` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会从主播收入中抽成  印尼盾 分',
  `pay_commission` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '支付手续费、（is_pay_to_guild此数据关联）给公会该值为0，单独打给主播需要扣除手续 印尼盾分',
  `host_total_income` bigint(30) unsigned NOT NULL DEFAULT '0' COMMENT '主播总收入 印尼盾 分',
  `activity_reward_for_guild` bigint(20) NOT NULL DEFAULT '0' COMMENT '活动奖励给公会 印尼盾 分 比如分四周奖励  \r\n该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会',
  `live_reward_for_guild` bigint(20) NOT NULL DEFAULT '0' COMMENT '直播奖励给公会 印尼盾 分 比如首播1小时\r\n该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会',
  `guild_total_income` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会总收入 印尼盾 分',
  `is_pay_to_guild` tinyint(8) NOT NULL DEFAULT '0' COMMENT '工资是否发放给公会 0非 1是',
  `settlement_bank_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发放账户名',
  `settlement_bank_accounts` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发放银行账号',
  `settlement_bank_person` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发放银行卡持有人',
  `is_settle` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否已经结算过 0未结算  1结算过',
  `classify_week_rank` int(11) NOT NULL DEFAULT '0' COMMENT '类型周排行',
  `classify_week_rank_type` int(11) NOT NULL DEFAULT '0' COMMENT '类型周排行类型 1 新档位 2 中档位 3 高档位',
  `water_rank` int(11) NOT NULL DEFAULT '0' COMMENT '周流水排名',
  `water_diamond` int(11) NOT NULL DEFAULT '0' COMMENT '周流水钻石',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_record` (`week`,`user_id`) USING BTREE COMMENT '聚合唯一约束',
  KEY `idx_query_host` (`user_id`) USING BTREE,
  KEY `idx_week` (`week`,`star_level`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68373 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for app_config
-- ----------------------------
DROP TABLE IF EXISTS `app_config`;
CREATE TABLE `app_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `system` int(11) NOT NULL DEFAULT '0' COMMENT '所属系统；0所有,1安卓 2ios',
  `version` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '版本',
  `channel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道',
  `key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'key',
  `value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '值',
  `remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_system_ver_chan_key` (`system`,`version`,`channel`,`key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for apply_battle
-- ----------------------------
DROP TABLE IF EXISTS `apply_battle`;
CREATE TABLE `apply_battle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `guild_id` int(11) NOT NULL DEFAULT '0' COMMENT '公会Id 关联guild_info自增',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播ID',
  `admin_operate_name` varchar(100) DEFAULT '' COMMENT '审核人id',
  `status` tinyint(11) NOT NULL DEFAULT '0' COMMENT '状态：0待审核 1通过 2不通过 ',
  `apply_type` tinyint(11) NOT NULL DEFAULT '0' COMMENT '报名类型，1公会，2 主播',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=206 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for apply_evaluation_score
-- ----------------------------
DROP TABLE IF EXISTS `apply_evaluation_score`;
CREATE TABLE `apply_evaluation_score` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '弹窗类型 1:真 2:假',
  `operate` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '用户操作 1:确认 2:取消',
  `score` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '用户评分',
  `action` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '评分触发动作 1:举报拉黑 2:直播卡顿 3:频繁切换直接间 4:送礼物 5:主播关播',
  `device_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_device_id` (`device_id`)
) ENGINE=InnoDB AUTO_INCREMENT=388 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='苹果评价分数';

-- ----------------------------
-- Table structure for area_info
-- ----------------------------
DROP TABLE IF EXISTS `area_info`;
CREATE TABLE `area_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `area_code` int(11) NOT NULL DEFAULT '1' COMMENT '地区',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `title_cn` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `currency` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '货币编码',
  `currency_code` int(11) NOT NULL DEFAULT '0' COMMENT '国家编码',
  `sort` int(11) NOT NULL DEFAULT '0',
  `is_rich` tinyint(1) DEFAULT '0' COMMENT '是否富有，0普通，1富有',
  `is_show` tinyint(1) NOT NULL DEFAULT '0',
  `created_time` bigint(20) NOT NULL,
  `updated_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uq_title_cn` (`title_cn`) USING BTREE,
  KEY `idx_area_ac_sort_isshow` (`area_code`,`is_show`,`sort`) USING BTREE,
  KEY `idx_currency` (`currency`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1467681760023400457 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for balance_info
-- ----------------------------
DROP TABLE IF EXISTS `balance_info`;
CREATE TABLE `balance_info` (
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `total_fee` bigint(20) unsigned NOT NULL COMMENT '总花费、只有付费购买才累计到这里 印尼盾 单位：分',
  `total_payout_diamonds` bigint(20) unsigned NOT NULL COMMENT '总流水，凡是用户花费的钻石都记录到这里 累计',
  `total_pay_diamonds` bigint(20) unsigned NOT NULL COMMENT '总购买钻石 只累计购买的',
  `remain_diamonds` bigint(20) unsigned NOT NULL COMMENT '剩余钻石数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`user_id`),
  KEY `balances_total` (`total_fee`) USING BTREE,
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for balance_record
-- ----------------------------
DROP TABLE IF EXISTS `balance_record`;
CREATE TABLE `balance_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `source_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户Id',
  `target_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '主播Id',
  `depletion_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '消费类型',
  `diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '钻石变动数额',
  `before_change` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变更之前余额',
  `after_change` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变更之后余额',
  `link_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '引起余额变更的id',
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '消费类型 1支出 2收入',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_source_user_id` (`source_user_id`),
  KEY `idx_target_user_id` (`target_user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_depletion_type` (`depletion_type`),
  KEY `idx_link_id` (`link_id`),
  KEY `idx_sourceuserid_depletiontype_type` (`source_user_id`,`depletion_type`,`type`)
) ENGINE=InnoDB AUTO_INCREMENT=235352061490565336 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


-- ----------------------------
-- Table structure for balance_statistics
-- ----------------------------
DROP TABLE IF EXISTS `balance_statistics`;
CREATE TABLE `balance_statistics` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `depletion_type` int(10) unsigned NOT NULL COMMENT '消费类型',
  `diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '钻石变动数额',
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `time_level` int(11) NOT NULL COMMENT '时间维度：1.每小时，2.每天',
  `big_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '消费类型 1支出 2收入',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=128663 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for balance_statistics_user
-- ----------------------------
DROP TABLE IF EXISTS `balance_statistics_user`;
CREATE TABLE `balance_statistics_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `date` char(10) NOT NULL COMMENT '日期',
  `depletion_type` tinyint(3) unsigned NOT NULL COMMENT '消费类型',
  `type` tinyint(3) unsigned NOT NULL COMMENT '消费类型 1支出 2收入',
  `diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '钻石变动数额',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `diamonds_IDX` (`user_id`,`date`,`depletion_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=168630 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户维度钻石收支统计';

-- ----------------------------
-- Table structure for ban_user_info
-- ----------------------------
DROP TABLE IF EXISTS `ban_user_info`;
CREATE TABLE `ban_user_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '封禁用户id',
  `admin_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `link_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联警告表主键id',
  `type` int(6) NOT NULL DEFAULT '0' COMMENT '封禁类型，1禁发私信，2禁发公屏聊天，3禁开播，4禁止登录，5封禁设备',
  `ban_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '封禁时间',
  `device_id` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '1' COMMENT '封禁设备',
  `ban_reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '封禁原因',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_device_id` (`device_id`),
  KEY `idx_ban_time` (`ban_time`),
  KEY `idx_ban_time_device_id` (`ban_time`,`device_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1515 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='banner_info';

-- ----------------------------
-- Table structure for banner_info
-- ----------------------------
DROP TABLE IF EXISTS `banner_info`;
CREATE TABLE `banner_info` (
  `bid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'bid',
  `ranking` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `client_os` int(11) NOT NULL DEFAULT '0' COMMENT '0 全部；1 android；2 iOS',
  `app_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'app名称',
  `cover` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面URL',
  `target` int(11) NOT NULL DEFAULT '0' COMMENT '跳转类型，0 网页，1 app内页',
  `link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转URL',
  `show_type` int(11) NOT NULL DEFAULT '0' COMMENT '展示类型，0 全部，1 未充值用户 2 充值用户 3 主播',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  `state` int(6) NOT NULL DEFAULT '0' COMMENT '状态1有效、2失效',
  PRIMARY KEY (`bid`) USING BTREE,
  KEY `idx_app_name` (`app_name`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='banner_info';

-- ----------------------------
-- Table structure for battle_match_record
-- ----------------------------
DROP TABLE IF EXISTS `battle_match_record`;
CREATE TABLE `battle_match_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `battle_id` bigint(20) NOT NULL COMMENT 'pk id',
  `win_user_id` bigint(20) NOT NULL COMMENT '赢方用户id',
  `hang_up_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '挂断人id',
  `match_season` int(11) NOT NULL COMMENT '比赛赛季 递增',
  `win_match_score` int(11) NOT NULL COMMENT '赢方分数',
  `lose_match_score` int(11) NOT NULL COMMENT '输方分数',
  `user_interrupt_score` int(11) NOT NULL COMMENT '用户中断分数',
  `win_streak_score` int(11) NOT NULL COMMENT '用户连胜分数',
  `win_segment_diff_score` int(11) NOT NULL COMMENT '赢方用户段位差分数',
  `double_score` int(11) NOT NULL COMMENT '翻倍卡分数',
  `fatigued_score` int(11) NOT NULL COMMENT '疲劳扣除分数',
  `lose_segment_diff_score` int(11) NOT NULL COMMENT '输方用户段位差分数',
  `lose_score_guard` int(11) NOT NULL COMMENT '积分保护',
  `streak_count_guard` int(11) NOT NULL COMMENT '连胜保护',
  `rank_guard` int(11) NOT NULL COMMENT '段位保护',
  `win_curr_match_score` bigint(20) NOT NULL COMMENT '赢方用户当前分数',
  `lose_curr_match_score` bigint(20) NOT NULL COMMENT '输方用户当前分数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for beautiful_num
-- ----------------------------
DROP TABLE IF EXISTS `beautiful_num`;
CREATE TABLE `beautiful_num` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `super_id` bigint(20) NOT NULL COMMENT '靓号',
  `classify_id` bigint(20) NOT NULL COMMENT '分类id',
  `icon_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'icon地址',
  `shop_status` int(11) NOT NULL DEFAULT '2' COMMENT '状态 1可用 2不可用',
  `sort_weight` int(11) NOT NULL COMMENT '排序权重',
  `sku_list_price` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '[]' COMMENT '规格列表价格',
  `user_level_limit` int(11) NOT NULL COMMENT '用户等级限制',
  `price` int(11) NOT NULL COMMENT '单价',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `is_wear` int(11) NOT NULL DEFAULT '0' COMMENT '是否佩戴',
  `experience` bigint(20) NOT NULL DEFAULT '0' COMMENT '经验值',
  `expiration_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
  `reclaim_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '回收时间',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_super_id` (`super_id`) USING BTREE,
  KEY `idx_classify_id` (`classify_id`) USING BTREE,
  KEY `idx_expiration_time` (`expiration_time`) USING BTREE,
  KEY `idx_shop_status` (`shop_status`) USING BTREE,
  KEY `idx_is_wear_user_id` (`user_id`,`is_wear`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for beautiful_num_classify
-- ----------------------------
DROP TABLE IF EXISTS `beautiful_num_classify`;
CREATE TABLE `beautiful_num_classify` (
  `classify_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `classify_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `translate_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '翻译key',
  `is_show` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '是否显示',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`classify_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for black_list
-- ----------------------------
DROP TABLE IF EXISTS `black_list`;
CREATE TABLE `black_list` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `operator_id` bigint(20) NOT NULL COMMENT '操作者ID',
  `blocked_Id` bigint(20) NOT NULL COMMENT '被拉黑的用户ID',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '拉黑结束时间,为0就是永久拉黑',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`blocked_Id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28097 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='黑名单';

-- ----------------------------
-- Table structure for charm_anchor_day_statistics
-- ----------------------------
DROP TABLE IF EXISTS `charm_anchor_day_statistics`;
CREATE TABLE `charm_anchor_day_statistics` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `statistics_time` date NOT NULL COMMENT '统计时间',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '工会id',
  `live_time` int(11) NOT NULL COMMENT '直播时长（s）',
  `effect_live_time` int(11) NOT NULL DEFAULT '0' COMMENT '有效直播时长（s） 大于30分钟的直播才算有效直播累计',
  `is_effect_day` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否有效天 0非 1是',
  `new_fans_count` int(11) NOT NULL COMMENT '新增粉丝数',
  `new_send_gift_person` int(11) NOT NULL COMMENT '新增送礼人数',
  `all_send_gift_person` int(11) NOT NULL COMMENT '总送礼人数',
  `all_active_person_count` int(11) NOT NULL COMMENT '直播间活跃人数，发送过消息',
  `join_room_person_count` int(11) NOT NULL COMMENT '进入房间总人数',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `charm_anchor_day_statistics_un` (`statistics_time`,`user_id`),
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_guild_id` (`guild_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=225297 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for commodity_free_gift
-- ----------------------------
DROP TABLE IF EXISTS `commodity_free_gift`;
CREATE TABLE `commodity_free_gift` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cid` bigint(20) NOT NULL,
  `state` int(11) NOT NULL COMMENT '状态 1 启用',
  `gift_type` int(11) NOT NULL COMMENT '赠送类型	1 装扮  2 经验',
  `gift_link` bigint(20) NOT NULL COMMENT '赠送关联	商品id',
  `gift_val` int(11) NOT NULL COMMENT '赠送数	装扮为天 经验为经验值',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=165 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for commodity_info
-- ----------------------------
DROP TABLE IF EXISTS `commodity_info`;
CREATE TABLE `commodity_info` (
  `cid` bigint(20) unsigned NOT NULL COMMENT '商品Id 唯一约束',
  `payment_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '关联payment_info表主键',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称 比如80钻',
  `price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '价格 分',
  `value` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '钻石数',
  `bonus` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '赠送钻石数',
  `currency` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '价格对应的单位 比如：IDR、MYR 前台支付',
  `totalFee_cent` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '对应印尼盾价格 单位分 统计使用',
  `is_first_topup` tinyint(10) NOT NULL DEFAULT '0' COMMENT '是否首次充值优惠档位，0默认，1首充优惠档位、2非首充档位',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序、顺序',
  `state` tinyint(255) NOT NULL DEFAULT '0' COMMENT '状态；0下架；1上架',
  `is_default` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付',
  `origin_bonus` int(11) NOT NULL COMMENT '折扣前赠送',
  `discount_type` int(11) NOT NULL COMMENT '折扣类型，1.首充折扣',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`cid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for complex_prize_config
-- ----------------------------
DROP TABLE IF EXISTS `complex_prize_config`;
CREATE TABLE `complex_prize_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `activity_type` int(11) NOT NULL COMMENT '抽奖活动类型 1 圣诞 2 元旦',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图标',
  `draw_type` int(11) NOT NULL COMMENT '抽奖类型 1装扮 2钻石 3优惠券 4 实物',
  `link_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '关联商品id  1,2  逗号分割',
  `prize_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '奖品名称',
  `num` int(11) NOT NULL COMMENT '数量',
  `is_first_pool` int(11) NOT NULL COMMENT '是否第一个奖池出现',
  `rate` int(11) NOT NULL COMMENT '概率万分比  *10000的整数',
  `status` int(11) NOT NULL COMMENT '配置状态 0失效 1生效',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for complex_prize_record
-- ----------------------------
DROP TABLE IF EXISTS `complex_prize_record`;
CREATE TABLE `complex_prize_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `c_id` int(11) NOT NULL COMMENT '抽奖配置id',
  `activity_type` int(11) NOT NULL COMMENT '抽奖活动类型 1 圣诞 2 元旦',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图标',
  `draw_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '抽奖类型 1装扮 2钻石 3优惠券 4 实物',
  `props_log_id` bigint(20) NOT NULL COMMENT '扣除道具流水id',
  `link_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '关联商品id  1,2  逗号分割',
  `prize_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '奖品名称',
  `prize_count` int(11) NOT NULL COMMENT '奖池期数',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=114832 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for coupon_info
-- ----------------------------
DROP TABLE IF EXISTS `coupon_info`;
CREATE TABLE `coupon_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `coupon_state` int(11) NOT NULL COMMENT '状态 1未使用 2已使用',
  `use_time` bigint(20) NOT NULL COMMENT '使用时间',
  `coupon_type` int(11) NOT NULL COMMENT '优惠券类型 1 钻石赠送券',
  `coupon_ratio` int(11) NOT NULL COMMENT '赠送比例 30 则赠送订单总额的30%',
  `min_shift` int(11) NOT NULL COMMENT '最小支持充值档位',
  `max_shift` int(11) NOT NULL COMMENT '最大支持充值档位',
  `link_id` bigint(20) NOT NULL COMMENT '优惠券来源',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22554 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for diamond_turntable_config
-- ----------------------------
DROP TABLE IF EXISTS `diamond_turntable_config`;
CREATE TABLE `diamond_turntable_config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `gid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品表id',
  `gtype` tinyint(4) NOT NULL DEFAULT '0' COMMENT '商品类型 1装扮 2钻石',
  `num` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '数量',
  `rate` int(11) NOT NULL DEFAULT '0' COMMENT '概率百分比  *100的整数',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '商品状态 0失效 1生效',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_gid_num` (`gid`,`num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='钻石转盘商品配置表';

-- ----------------------------
-- Table structure for diamond_turntable_record
-- ----------------------------
DROP TABLE IF EXISTS `diamond_turntable_record`;
CREATE TABLE `diamond_turntable_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `cid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '配置表id',
  `gid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品表id',
  `num` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '数量',
  `turntable_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '奖池期数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_query` (`cid`,`gid`,`num`) USING BTREE,
  KEY `idx_turntable_num` (`turntable_count`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23683 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='钻石转盘中奖记录表';

-- ----------------------------
-- Table structure for diamond_turntable_statistics
-- ----------------------------
DROP TABLE IF EXISTS `diamond_turntable_statistics`;
CREATE TABLE `diamond_turntable_statistics` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `watch_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '在直播间的有效观看时间 单位s ',
  `raffle_num` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '获得的抽奖次数',
  `use_num` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '使用的抽奖次数',
  `gid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品表id',
  `num` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '数量',
  `turntable_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '奖池期数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `cid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖配置id',
  `prize_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '奖品名称',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_query` (`gid`,`num`) USING BTREE,
  KEY `idx_turntable_num` (`turntable_count`) USING BTREE,
  KEY `idx_createtime_rafflenum` (`create_time`,`raffle_num`)
) ENGINE=InnoDB AUTO_INCREMENT=877184 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='钻石转盘数据统计表';

-- ----------------------------
-- Table structure for feed_comment
-- ----------------------------
DROP TABLE IF EXISTS `feed_comment`;
CREATE TABLE `feed_comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '评论用户Id',
  `feed_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '动态Id',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父级评论ID，\r\n0标识对一级动态回复、\r>0标识对评论回复只存储一级回复的Id（自增id）',
  `to_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'A回复了B，这里存储B的用户ID ，一级回复id这里存储0',
  `comment_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '该条评论针对那个评论的回复，一级回复存储0，针对评论的的回复存储回复评论的ID，该数据只是存储',
  `reply_context` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论/回复内容',
  `reply_like_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '评论点赞数',
  `state` tinyint(6) NOT NULL DEFAULT '1' COMMENT '0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `feed_comment_user_id_IDX` (`user_id`) USING BTREE,
  KEY `idx_feedid_state` (`feed_id`,`state`)
) ENGINE=InnoDB AUTO_INCREMENT=34028 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for feed_comment_mark
-- ----------------------------
DROP TABLE IF EXISTS `feed_comment_mark`;
CREATE TABLE `feed_comment_mark` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `feed_comment_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '评论Id',
  `mark` tinyint(8) NOT NULL DEFAULT '0' COMMENT '标记类型：1喜欢、 后续扩展',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_feed_comment_id` (`feed_comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7031 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for feed_info
-- ----------------------------
DROP TABLE IF EXISTS `feed_info`;
CREATE TABLE `feed_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `text` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '动态内容',
  `images` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '动态图片组，采用数组存储,目前规定数组长度为9，即最多9张图片',
  `like_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `state` tinyint(6) NOT NULL DEFAULT '1' COMMENT '0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除',
  `score` int(6) NOT NULL DEFAULT '0' COMMENT '动态排序权重',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `video` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '动态视频',
  `video_picture` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '视频封面',
  `type` int(6) NOT NULL DEFAULT '0' COMMENT '动态类型，1图片、2视频',
  `upload_video_width` int(11) NOT NULL DEFAULT '0',
  `upload_video_height` int(11) NOT NULL DEFAULT '0',
  `is_vote` int(2) NOT NULL DEFAULT '0' COMMENT '是否pk动态',
  `option_one` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'pk选项1',
  `option_two` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'pk选项2',
  `system_os` int(2) NOT NULL DEFAULT '0' COMMENT '发动态的设备类型',
  `app_channel` varchar(255) NOT NULL COMMENT '渠道包',
  PRIMARY KEY (`id`),
  KEY `idx_userid` (`user_id`),
  KEY `idx_system_os` (`system_os`) USING BTREE,
  KEY `idx_app_channel` (`app_channel`) USING BTREE,
  KEY `idx_score` (`score`) USING BTREE,
  KEY `feed_info_user_id_IDX` (`user_id`,`score`,`create_time`,`app_channel`,`system_os`) USING BTREE,
  KEY `idx_state` (`state`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=49134 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for feed_mark
-- ----------------------------
DROP TABLE IF EXISTS `feed_mark`;
CREATE TABLE `feed_mark` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `feed_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '动态Id',
  `mark` tinyint(8) NOT NULL DEFAULT '0' COMMENT '标记类型：1喜欢、2不感兴趣 后续扩展',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE_user_id_feed_id_mark` (`user_id`,`feed_id`,`mark`) USING BTREE,
  KEY `INDEX_mark` (`mark`) USING BTREE,
  KEY `INDEX_feed_id` (`feed_id`) USING BTREE,
  KEY `INDEX_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=140047 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for feed_topic
-- ----------------------------
DROP TABLE IF EXISTS `feed_topic`;
CREATE TABLE `feed_topic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `topic_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '话题Id',
  `feed_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '动态Id',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_topic_id_feed_id` (`topic_id`,`feed_id`) USING BTREE,
  KEY `idx_topic_id` (`topic_id`) USING BTREE,
  KEY `idx_feed_id` (`feed_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21172 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='动态话题关联表';

-- ----------------------------
-- Table structure for first_settlement_chance
-- ----------------------------
DROP TABLE IF EXISTS `first_settlement_chance`;
CREATE TABLE `first_settlement_chance` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `week` int(11) NOT NULL DEFAULT '0' COMMENT '周，一年中的第几周 例如:202230 2022年的第30周',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_unique` (`user_id`),
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14006 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for game_config
-- ----------------------------
DROP TABLE IF EXISTS `game_config`;
CREATE TABLE `game_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `game_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '游戏ID',
  `game_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方业务唯一标识',
  `game_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '游戏名称',
  `icon_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图片地址',
  `mini_icon_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '小图标',
  `game_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '游戏地址',
  `game_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '游戏类型 0自研 1joyplay 2cocos',
  `is_mini` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否半屏 0全屏 1半屏',
  `is_hot` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否热门 0不是 1是热门',
  `system_rate` int(11) NOT NULL DEFAULT '0' COMMENT '抽水比例  *10000的整数 最小为万分之一',
  `anchor_rate` int(11) NOT NULL DEFAULT '0' COMMENT '主播分成  *10000的整数 最小为万分之一',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `game_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '游戏版本',
  `md5_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'MD5版本',
  `cocos_engine_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'cocos引擎版本',
  `show_system` int(11) NOT NULL DEFAULT '0' COMMENT '显示系统 0 全部 1 android 2 ios',
  `win_ratio` decimal(10,2) NOT NULL COMMENT '游戏宽高比',
  `entrypoint` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Ws接入点',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1 正常 2 下架',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `orientation` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏方向 横竖屏',
  `is_full_win` int(11) NOT NULL COMMENT '是否全屏 1全屏',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_game_id` (`game_id`) USING BTREE,
  KEY `idx_is_mini` (`is_mini`) USING BTREE,
  KEY `idx_is_hot` (`is_hot`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10030 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='游戏配置表';

-- ----------------------------
-- Table structure for game_room_statistics
-- ----------------------------
DROP TABLE IF EXISTS `game_room_statistics`;
CREATE TABLE `game_room_statistics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `time_level` int(11) NOT NULL COMMENT '时间维度：1.每小时，2.每天',
  `room_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '直播间id',
  `diamonds` bigint(20) DEFAULT NULL COMMENT '总钻石',
  `game_id` bigint(20) DEFAULT NULL COMMENT '游戏id',
  `anchor_reward` bigint(20) DEFAULT NULL COMMENT '主播收益',
  `play_num` bigint(20) DEFAULT NULL COMMENT '操作数',
  `play_person` bigint(20) DEFAULT NULL COMMENT '游戏人数',
  `anchor_id` bigint(20) DEFAULT NULL COMMENT '主播id',
  `master_id` bigint(20) DEFAULT NULL COMMENT '公会长id',
  `server_id` bigint(20) DEFAULT NULL COMMENT '运营号id',
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_time_level` (`time_level`) USING BTREE,
  KEY `idx_room_id` (`room_id`) USING BTREE,
  KEY `idx_game_id` (`game_id`) USING BTREE,
  KEY `idx_server_id` (`server_id`) USING BTREE,
  KEY `idx_master_id` (`master_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=53035 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for game_statistics
-- ----------------------------
DROP TABLE IF EXISTS `game_statistics`;
CREATE TABLE `game_statistics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'app渠道',
  `game_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '游戏key',
  `game_income` bigint(20) NOT NULL COMMENT '游戏收益',
  `game_outlay` bigint(20) NOT NULL COMMENT '游戏支出',
  `play_user_count` bigint(20) NOT NULL COMMENT '游戏人数',
  `room_count` bigint(20) NOT NULL COMMENT '房间数',
  `play_count` bigint(20) NOT NULL COMMENT '操作数',
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `time_level` int(11) NOT NULL COMMENT '时间维度：1.每小时，2.每天',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=146813 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for gift_info
-- ----------------------------
DROP TABLE IF EXISTS `gift_info`;
CREATE TABLE `gift_info` (
  `gid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '礼物Id 自增',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '礼物名称',
  `diamonds` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '价格-钻石',
  `original_diamonds` int(11) NOT NULL DEFAULT '0' COMMENT '原价',
  `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '礼物大类 0普通礼物  1幸运礼物',
  `gift_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '0普通小礼物（无特效）1特效礼物  2 房间红包雨 3 全站红包雨 4 全站通知礼物',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '礼物图标',
  `anim_effect_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '特效礼物url',
  `file_md5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '礼物md5值',
  `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0下架  1上架',
  `sort` int(10) unsigned NOT NULL COMMENT '排序 顺序',
  `res_voice_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '声音',
  `res_type` int(11) NOT NULL DEFAULT '0' COMMENT '1 icon/2 gif/3 文件资源；4为3d；5为面部识别；6为svga；',
  `is_batter` int(11) NOT NULL DEFAULT '1' COMMENT '是否连送',
  `min_level` int(11) NOT NULL COMMENT '最小等级',
  `luck_pool_type` int(11) NOT NULL DEFAULT '0' COMMENT '奖池类型 0 90% 1 95%',
  `badge` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '礼物徽章',
  `start_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `end_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `created_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `updated_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`gid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=260 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for gift_log
-- ----------------------------
DROP TABLE IF EXISTS `gift_log`;
CREATE TABLE `gift_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sender_id` bigint(20) unsigned NOT NULL COMMENT '送礼者uid',
  `receiver_id` bigint(20) unsigned NOT NULL COMMENT '收礼者uid',
  `gift_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '礼物id',
  `gift_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物类型',
  `gift_num` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '礼物个数',
  `room_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '房间id',
  `diamonds` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物单价',
  `total_diamonds` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物总价',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_sender_id` (`sender_id`),
  KEY `idx_receiver_id` (`receiver_id`),
  KEY `idx_roomid` (`room_id`),
  KEY `idx_gift_type` (`gift_type`),
  KEY `idx_gift_id` (`gift_id`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_receiverid_createtime_senderid` (`receiver_id`,`create_time`,`sender_id`)
) ENGINE=InnoDB AUTO_INCREMENT=630062 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gift_lucky_log
-- ----------------------------
DROP TABLE IF EXISTS `gift_lucky_log`;
CREATE TABLE `gift_lucky_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sender_id` bigint(20) unsigned NOT NULL COMMENT '送礼者uid',
  `receiver_id` bigint(20) unsigned NOT NULL COMMENT '收礼者uid',
  `gift_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '礼物id',
  `gift_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物类型',
  `gift_num` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '礼物个数',
  `room_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '房间id',
  `diamonds` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物单价',
  `total_diamonds` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物总价',
  `lucky_multiple` bigint(20) NOT NULL COMMENT '中奖倍速',
  `lucky_diamonds` bigint(20) NOT NULL COMMENT '中奖钻石',
  `create_time` bigint(20) DEFAULT '0',
  `update_time` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_sender_id` (`sender_id`),
  KEY `idx_roomid` (`room_id`),
  KEY `idx_receiver_id` (`receiver_id`) USING BTREE,
  KEY `idx_gift_type` (`gift_type`) USING BTREE,
  KEY `idx_create_time` (`create_time`) USING BTREE,
  KEY `idx_gift_id` (`gift_id`)
) ENGINE=InnoDB AUTO_INCREMENT=91874107655958383 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gift_statistics
-- ----------------------------
DROP TABLE IF EXISTS `gift_statistics`;
CREATE TABLE `gift_statistics` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '礼物id',
  `gift_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '礼物个数',
  `total_diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '礼物总价',
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `time_level` int(11) NOT NULL COMMENT '时间维度：1.每小时，2.每天',
  `lucky_diamonds` bigint(20) NOT NULL COMMENT '中奖钻石',
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `user_count` bigint(20) DEFAULT NULL COMMENT '送礼人数',
  `room_count` bigint(20) DEFAULT NULL COMMENT '房间数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=86558 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for google_order
-- ----------------------------
DROP TABLE IF EXISTS `google_order`;
CREATE TABLE `google_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `receipt` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'inAppPurchaseData 回调',
  `signature` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'google签名',
  `internal_ip` bigint(20) unsigned NOT NULL COMMENT '总购买钻石 只累计购买的',
  `product_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '商品Id',
  `gp_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'google订单idGPA.0000',
  `purchase_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `app_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app名称',
  `package` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '包名 com.aaa.bb',
  `order_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '系统订单ID',
  `state` tinyint(10) NOT NULL DEFAULT '0' COMMENT '状态0待处理 1处理完成',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_gp_order_id` (`gp_order_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for guild_anchor
-- ----------------------------
DROP TABLE IF EXISTS `guild_anchor`;
CREATE TABLE `guild_anchor` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统自增',
  `guild_id` int(11) NOT NULL COMMENT '公会Id 关联guild_info自增',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `guild_role` int(11) NOT NULL DEFAULT '2' COMMENT '1公会长 2公会成员',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1 正常；2 失效',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `join_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '加入时间',
  `star_level` int(11) NOT NULL DEFAULT '0' COMMENT '主播星级',
  `label_type` int(11) NOT NULL DEFAULT '0' COMMENT '标签类型',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '主播标签',
  `score` int(11) NOT NULL DEFAULT '3000' COMMENT '主播基础分数',
  `live_permissions` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1钻石房，2密码房，3钻石房和密码房',
  `create_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  `cloud_recording` tinyint(4) NOT NULL DEFAULT '0' COMMENT '录制状态 0不录制，1录制',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `labour_union_to_anchors_lu_type` (`status`) USING BTREE,
  KEY `idx_guild_role` (`guild_role`),
  KEY `idx_join_time` (`join_time`),
  KEY `idx_guild_id` (`guild_id`),
  KEY `userId` (`user_id`) USING BTREE,
  KEY `idx_status_starlevel_jointime_userid_guildid` (`status`,`star_level`,`join_time`,`user_id`,`guild_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28593 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for guild_anchor_record
-- ----------------------------
DROP TABLE IF EXISTS `guild_anchor_record`;
CREATE TABLE `guild_anchor_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `admin_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '管理员id',
  `expiration_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
  `background_adjustment` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否后台调整',
  `update_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '调整类型，1星级，2公会，3标签',
  `is_hot` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否在热门列表',
  `old_content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '修改前内容',
  `new_content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '修改后内容',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22616 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='新用户热门记录表';

-- ----------------------------
-- Table structure for guild_apply
-- ----------------------------
DROP TABLE IF EXISTS `guild_apply`;
CREATE TABLE `guild_apply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint(20) NOT NULL COMMENT '申请人用户Id',
  `guild_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公会名称',
  `anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '主播数量',
  `whatsapp` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系WhatsApp',
  `state` tinyint(11) NOT NULL COMMENT '状态：0待审核 1通过 2不通过 ',
  `operator` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '审核人',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `userId` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5299 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for guild_info
-- ----------------------------
DROP TABLE IF EXISTS `guild_info`;
CREATE TABLE `guild_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '公会Id',
  `user_id` bigint(20) NOT NULL COMMENT '公会长Id',
  `guild_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公会名称',
  `gcode` int(10) NOT NULL DEFAULT '0' COMMENT '公会邀请码',
  `member_count` int(10) NOT NULL DEFAULT '0' COMMENT '成员数量',
  `portrait` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公会头像',
  `status` int(10) NOT NULL DEFAULT '0' COMMENT '状态',
  `notice` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公告',
  `cut_rate` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '公会抽成 设置范围0-100',
  `is_pay_to_guild` tinyint(8) NOT NULL DEFAULT '0' COMMENT '工资是否发放给公会 0非 1是',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `server_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  `source_type` tinyint(8) NOT NULL DEFAULT '0' COMMENT '绑定来源0无，1运营号直接邀请，2工会邀请，3后台绑定',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_guild_name` (`guild_name`),
  UNIQUE KEY `idx_ gcode` (`gcode`),
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4537 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for guild_pull_guild_info
-- ----------------------------
DROP TABLE IF EXISTS `guild_pull_guild_info`;
CREATE TABLE `guild_pull_guild_info` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `puller_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'A拉B，这里存放A公会长ID',
  `pulled_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'A拉B，这里存放B公会长ID',
  `manager_id` int(8) NOT NULL DEFAULT '0' COMMENT '操作管理员ID',
  `created_time` bigint(10) NOT NULL DEFAULT '1451577600',
  `updated_time` bigint(10) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_pulled` (`pulled_user_id`) USING BTREE,
  KEY `idx_puller` (`puller_user_id`) USING BTREE,
  KEY `idx_pulled` (`pulled_user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='公会拉公会绑定记录';

-- ----------------------------
-- Table structure for guild_settlement_reward
-- ----------------------------
DROP TABLE IF EXISTS `guild_settlement_reward`;
CREATE TABLE `guild_settlement_reward` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `week` int(11) NOT NULL DEFAULT '0' COMMENT '周，一年中的第几周 例如:202230 2022年的第30周',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会id',
  `reward_type` int(11) NOT NULL DEFAULT '0' COMMENT '奖励类型，  比如type  1：为3周2次结算  2……  后续扩展',
  `reward` bigint(20) NOT NULL DEFAULT '0' COMMENT '活动奖励给公会 印尼盾 分 \r\n该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9571 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for guild_week_settlement
-- ----------------------------
DROP TABLE IF EXISTS `guild_week_settlement`;
CREATE TABLE `guild_week_settlement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `settlement_week` int(11) NOT NULL COMMENT '周 example:202230',
  `guild_id` bigint(20) NOT NULL COMMENT '公会ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `live_time` int(11) NOT NULL DEFAULT '0' COMMENT '直播时长（s）',
  `pay_commission` bigint(20) NOT NULL COMMENT '支付手续费',
  `total_income` bigint(20) NOT NULL COMMENT '工会所有主播 工会总收益金币收益 单位分',
  `effect_live_day` int(11) NOT NULL COMMENT '工会所有主播 总有效出勤天数',
  `gift_income` bigint(20) NOT NULL COMMENT '工会所有主播 礼物收入 印尼盾 分',
  `guild_gift_income` bigint(20) NOT NULL COMMENT '工会所有主播 公会礼物提成收入 印尼盾 分',
  `week_rank_reward_for_guild` bigint(20) NOT NULL COMMENT '工会所有主播 周榜公会获得提成收入 印尼盾 分',
  `host_total_income_cut` bigint(20) NOT NULL COMMENT '工会所有主播 公会从主播收入中抽成  印尼盾 分',
  `host_total_income` bigint(20) NOT NULL COMMENT '工会所有主播 主播总收入 印尼盾 分',
  `activity_reward_for_guild` bigint(20) NOT NULL COMMENT '工会所有主播 活动奖励给公会',
  `live_reward_for_guild` bigint(20) NOT NULL COMMENT '工会所有主播 直播奖励给公会 印尼盾 分',
  `guild_total_income` bigint(20) NOT NULL COMMENT '公会总收入 印尼盾 分',
  `settlement_bank_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发放账户名',
  `settlement_bank_accounts` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发放银行账号',
  `settlement_bank_person` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发放银行卡持有人',
  `is_settle` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否已经结算过 0未结算  1结算过',
  `up_standard_host_count` bigint(20) NOT NULL COMMENT '工会所有主播 达标结算主播数',
  `not_standard_host_count` bigint(20) NOT NULL COMMENT '工会所有主播 未达标结算主播数',
  `guild_cut_rate` bigint(20) NOT NULL COMMENT '工会抽成',
  `host_total_income_before` bigint(20) NOT NULL COMMENT '工会所有主播 未扣除主播总收入 印尼盾 分',
  `is_pay_to_guild` tinyint(8) NOT NULL DEFAULT '0' COMMENT '工资是否发放给公会 0非 1是',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5452 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for guild_week_statistics
-- ----------------------------
DROP TABLE IF EXISTS `guild_week_statistics`;
CREATE TABLE `guild_week_statistics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `server_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  `week` int(11) NOT NULL DEFAULT '0' COMMENT '周，一年中的第几周 例如:202230 2022年的第30周',
  `guild_id` int(11) NOT NULL DEFAULT '0' COMMENT '公会Id 关联guild_info自增',
  `total_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '总主播数',
  `total_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '总直播时长(秒)',
  `total_live_person` int(11) NOT NULL DEFAULT '0' COMMENT '总开播人数',
  `new_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '新增主播人数',
  `new_anchor_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '新增主播直播时长',
  `new_anchor_live_person` int(11) NOT NULL DEFAULT '0' COMMENT '新增主播开播人数',
  `star_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '总标星主播数',
  `star_anchor_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '总标星主播总开播时长',
  `star_anchor_live_person` int(11) NOT NULL DEFAULT '0' COMMENT '总标星主播开播人数',
  `new_settlement_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '新增达标结算主播数',
  `settlement_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '达标结算主播数',
  `new_star_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '新增标星主播人数',
  `new_star_anchor_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '新增标星主播总开播时长',
  `new_star_anchor_live_person` int(11) NOT NULL DEFAULT '0' COMMENT '新增标星主播开播人数',
  `settlement_anchor_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '达标结算主播总直播时长',
  `settlement_income` bigint(20) NOT NULL DEFAULT '0' COMMENT '达标结算收益',
  `new_lived_unsettlement_person` int(11) NOT NULL DEFAULT '0' COMMENT '新增开播未达标结算主播数',
  `lived_unsettlement_person` int(11) NOT NULL DEFAULT '0' COMMENT '已开播未结算主播数',
  `lived_unsettlement_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '已开播未结算主播总开播时长',
  `week_award_diamonds` int(11) NOT NULL DEFAULT '0' COMMENT '本周奖励钻石数',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '奖励钻石的备注',
  `invite_person_num` int(11) NOT NULL DEFAULT '0' COMMENT '邀新总人数',
  `total_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '拉新总冲值',
  `invite_normal_user_num` int(11) NOT NULL DEFAULT '0' COMMENT '邀新普通用户数',
  `invite_normal_user_total_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀新普通用户充值',
  `invite_anchor_num` int(11) NOT NULL DEFAULT '0' COMMENT '邀新主播数',
  `invite_anchor_total_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀新主播充值',
  `invite_master_num` int(11) NOT NULL DEFAULT '0' COMMENT '邀新公会数',
  `invite_master_total_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀新公会长充值',
  `effective_user_num` int(11) NOT NULL DEFAULT '0' COMMENT '有效用户数',
  `effective_user_total_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '有效用户金币',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `guild_week_statistics_un` (`week`,`guild_id`)
) ENGINE=InnoDB AUTO_INCREMENT=56490 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for income_info
-- ----------------------------
DROP TABLE IF EXISTS `income_info`;
CREATE TABLE `income_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) NOT NULL DEFAULT '-1' COMMENT '用户id',
  `total_income` bigint(20) NOT NULL DEFAULT '-1' COMMENT '总收益 单位分',
  `balances_income` bigint(20) NOT NULL DEFAULT '-1' COMMENT '收益余额 单位分',
  `total_commission_diamonds` bigint(20) NOT NULL DEFAULT '0',
  `balances_commission_diamonds` bigint(20) NOT NULL DEFAULT '0',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '总佣金收入 扩大10000',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '佣金收入 扩大10000',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=481256 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='收益表';

-- ----------------------------
-- Table structure for income_record
-- ----------------------------
DROP TABLE IF EXISTS `income_record`;
CREATE TABLE `income_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `source_id` bigint(20) NOT NULL COMMENT '收益来源用户的id',
  `user_record_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '对应用户扣费记录id',
  `income_type` int(11) NOT NULL DEFAULT '0' COMMENT '收益类型',
  `link_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '引起变更的id',
  `income` bigint(20) NOT NULL DEFAULT '0' COMMENT '收益变动数额',
  `before_change` bigint(20) NOT NULL DEFAULT '0' COMMENT '变更之前余额',
  `after_change` bigint(20) NOT NULL DEFAULT '0' COMMENT '变更之后余额',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'createdAt',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'updatedAt',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_user_record_id` (`user_record_id`) USING BTREE,
  KEY `idx_link_id` (`link_id`) USING BTREE,
  KEY `idx_source_id` (`source_id`) USING BTREE,
  KEY `idx_created_at` (`create_time`) USING BTREE,
  KEY `idx_userid_createtime` (`user_id`,`create_time`) COMMENT 'create by DAS-d74e5782-ef82-4a15-a3c3-14e48bf0df94'
) ENGINE=InnoDB AUTO_INCREMENT=235352058256095404 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='收益明细表';

-- ----------------------------
-- Table structure for inviter_info
-- ----------------------------
DROP TABLE IF EXISTS `inviter_info`;
CREATE TABLE `inviter_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `stat` int(11) NOT NULL,
  `invite_user_count` int(11) DEFAULT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_id_stat` (`user_id`,`stat`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_stat` (`stat`),
  KEY `idx_stat_userid_createtime` (`stat`,`user_id`,`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=29352 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for live_classify
-- ----------------------------
DROP TABLE IF EXISTS `live_classify`;
CREATE TABLE `live_classify` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级id 0 为顶级',
  `classify_type` int(11) NOT NULL COMMENT '1 推荐 2 最新 3 热门 4 pk 5 密码房 6 钻石房 7 标签主播 8 无标签主播 9 新秀 10 夜播',
  `classify_default_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '默认分类名',
  `classify_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类备注',
  `classify_status` int(11) NOT NULL DEFAULT '1' COMMENT '1 正常',
  `classify_rank` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for live_classify_lang
-- ----------------------------
DROP TABLE IF EXISTS `live_classify_lang`;
CREATE TABLE `live_classify_lang` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `classify_id` int(11) NOT NULL COMMENT '分类id',
  `classify_lang` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类语言',
  `classify_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `classify_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '分类图标',
  `classify_pick_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for live_cloud_record
-- ----------------------------
DROP TABLE IF EXISTS `live_cloud_record`;
CREATE TABLE `live_cloud_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统自增',
  `task_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '录制任务 ID，长度固定为 16 个字节的字符串。',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会id',
  `room_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '录制房间 ID',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '录制流对应的推流用户 ID（混流时，为 mix_output_stream_id）。',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '录制状态，1开始录制，2录制结束',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_room_id` (`room_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for live_cloud_video
-- ----------------------------
DROP TABLE IF EXISTS `live_cloud_video`;
CREATE TABLE `live_cloud_video` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统自增',
  `room_id` varchar(255) NOT NULL DEFAULT '' COMMENT '录制房间 ID。',
  `duration` bigint(20) NOT NULL DEFAULT '0' COMMENT '文件时长，单位：ms。',
  `begin_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '直播开始时间',
  `file_url` varchar(255) NOT NULL DEFAULT '' COMMENT '文件访问 URL。第三方存储为七牛云或阿里云 Vod 时不返回。',
  `file_size` bigint(20) NOT NULL DEFAULT '0' COMMENT '文件大小，单位：字节。',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_room_id` (`room_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6122 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for login_record
-- ----------------------------
DROP TABLE IF EXISTS `login_record`;
CREATE TABLE `login_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `app_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'app名称',
  `app_channel` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'app渠道',
  `app_version` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'app版本',
  `device_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '设备号',
  `ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登陆ip',
  `language` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '语言',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `app_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'app登录token',
  `login_type` int(10) NOT NULL COMMENT '1 gg 2 fb 3 apple',
  `app_phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'app手机型号',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_app_name` (`app_name`),
  KEY `idx_app_channel` (`app_channel`),
  KEY `idx_device_id` (`device_id`),
  KEY `idx_created_time` (`create_time`),
  KEY `idx_login_type` (`login_type`),
  KEY `idx_userid` (`user_id`) COMMENT 'create by DAS-fcf0af5b-59bd-4fcd-815f-95e8f766b817',
  KEY `idx_ip` (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=2392150 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for manager_message_record
-- ----------------------------
DROP TABLE IF EXISTS `manager_message_record`;
CREATE TABLE `manager_message_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  `admin_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '管理员id',
  `message` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '发送的消息',
  `whats_app` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'whatsapp号码',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for medal_app_config
-- ----------------------------
DROP TABLE IF EXISTS `medal_app_config`;
CREATE TABLE `medal_app_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `medal_id` bigint(20) NOT NULL COMMENT '勋章id',
  `app_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'app名称',
  `medal_name_config` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '{}' COMMENT '勋章名称配置json',
  `medal_desc_config` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '{}' COMMENT '勋章描述配置json',
  PRIMARY KEY (`id`),
  UNIQUE KEY `config_unique` (`medal_id`,`app_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='勋章app配置表';

-- ----------------------------
-- Table structure for medal_info
-- ----------------------------
DROP TABLE IF EXISTS `medal_info`;
CREATE TABLE `medal_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `medal_icon` varchar(255) NOT NULL COMMENT '勋章图片',
  `medal_grey_icon` varchar(255) NOT NULL COMMENT '勋章灰色图片',
  `medal_name_config` varchar(255) NOT NULL DEFAULT '{}' COMMENT '勋章名称配置json',
  `medal_desc_config` varchar(500) NOT NULL DEFAULT '{}' COMMENT '勋章描述配置',
  `conditions_type` int(11) NOT NULL DEFAULT '0' COMMENT '获取类型',
  `begin_time` bigint(20) NOT NULL COMMENT '获取勋章开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '获取勋章结束时间',
  `medal_states` int(11) NOT NULL COMMENT '勋章状态 1',
  `medal_sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=112 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for medal_user
-- ----------------------------
DROP TABLE IF EXISTS `medal_user`;
CREATE TABLE `medal_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `medal_id` bigint(20) NOT NULL COMMENT '勋章id',
  `access_count` int(11) NOT NULL COMMENT '获得次数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_medal_id_user_id` (`user_id`,`medal_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_medal_id` (`medal_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24651 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for message_info
-- ----------------------------
DROP TABLE IF EXISTS `message_info`;
CREATE TABLE `message_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `sender_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送用户Id',
  `receiver_id` varchar(100) NOT NULL DEFAULT '0' COMMENT '接收人',
  `conversation_type` int(6) NOT NULL DEFAULT '0' COMMENT '会话类型（1直播间消息、2私信消息）',
  `message_info` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `message_type` int(6) NOT NULL DEFAULT '0' COMMENT '消息类型',
  `message_source` varchar(100) NOT NULL DEFAULT '0' COMMENT '消息来源',
  `send_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送时间',
  PRIMARY KEY (`id`),
  KEY `idx_sendtime_messageinfo` (`send_time`,`message_info`(5)) COMMENT 'create by DAS-68134cde-4ff7-4f72-aca5-75cfb7564f1b',
  KEY `idx_senderid_sendtime` (`sender_id`,`send_time`) COMMENT 'create by DAS-b4bacb16-5bf3-4069-a55e-279253a341a8',
  KEY `idx_conversationtype_sendtime_messageinfo` (`conversation_type`,`send_time`,`message_info`(5)) COMMENT 'create by DAS-9b68bc6d-f45f-4fdf-945e-cdfbd0d2fc27',
  KEY `idx_conversationtype_messageinfo` (`conversation_type`,`message_info`(5)),
  KEY `idx_receiver_id` (`receiver_id`),
  KEY `idx_messagetype` (`message_type`),
  KEY `idx_messagetype_sendtime` (`message_type`,`send_time`),
  KEY `idx_conversationtype_messagetype` (`conversation_type`,`message_type`)
) ENGINE=InnoDB AUTO_INCREMENT=39890203 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for noble_config
-- ----------------------------
DROP TABLE IF EXISTS `noble_config`;
CREATE TABLE `noble_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `noble_id` bigint(20) NOT NULL COMMENT '贵族id',
  `privilege_id` bigint(20) NOT NULL COMMENT '特权id',
  `privilege_link` bigint(20) NOT NULL COMMENT '特权值 类型为加成类型时val为倍数、商品时为商品id',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_noble_id` (`noble_id`) USING BTREE,
  KEY `idx_privilege_id` (`privilege_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='COLUMNAR=1 贵族配置列表';

-- ----------------------------
-- Table structure for noble_info
-- ----------------------------
DROP TABLE IF EXISTS `noble_info`;
CREATE TABLE `noble_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `noble_name` varchar(255) NOT NULL COMMENT '贵族名称',
  `noble_states` int(11) NOT NULL COMMENT '贵族状态 1 上架',
  `noble_icon` varchar(255) NOT NULL COMMENT '贵族图标',
  `current_price` int(11) NOT NULL COMMENT '现价',
  `original_price` int(11) NOT NULL COMMENT '原价',
  `noble_type` int(11) NOT NULL COMMENT '贵族类型 1 vip',
  `noble_level` int(11) NOT NULL COMMENT '贵族等级',
  `noble_day` int(11) NOT NULL COMMENT '贵族天数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_noble_type` (`noble_type`),
  KEY `idx_ noble_level` (`noble_level`),
  KEY `idx_noble_states` (`noble_states`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='COLUMNAR=1 贵族';

-- ----------------------------
-- Table structure for noble_privilege
-- ----------------------------
DROP TABLE IF EXISTS `noble_privilege`;
CREATE TABLE `noble_privilege` (
  `privilege_id` int(11) NOT NULL AUTO_INCREMENT,
  `privilege_name` varchar(255) NOT NULL COMMENT '特权名称',
  `privilege_icon` varchar(255) NOT NULL COMMENT '特权图片',
  `privilege_states` int(11) NOT NULL COMMENT '特权状态 1 正常',
  `is_privilege` int(11) NOT NULL COMMENT '是否特权 1特权 2商品',
  `privilege_type` int(11) NOT NULL COMMENT '特权类型',
  `privilege_desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '描述 多语言配置 {en:xxx}',
  `noble_type` int(11) NOT NULL COMMENT '贵族类型 1 vip',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`privilege_id`) USING BTREE,
  KEY `idx_privilege_states` (`privilege_states`),
  KEY `idx_noble_type` (`noble_type`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='COLUMNAR=1';

-- ----------------------------
-- Table structure for noble_user
-- ----------------------------
DROP TABLE IF EXISTS `noble_user`;
CREATE TABLE `noble_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `noble_id` bigint(20) NOT NULL COMMENT '贵族id',
  `expiration_time` bigint(20) NOT NULL COMMENT '过期时间',
  `noble_states` int(11) NOT NULL COMMENT '贵族状态 1 有效',
  `paid_then_diamonds` bigint(20) DEFAULT '0' COMMENT '购买vip时的付费钻石数',
  `paid_product_day` bigint(20) DEFAULT NULL COMMENT '购买的规格天',
  `paid_product_diamond` bigint(20) DEFAULT NULL COMMENT '购买的钻石数',
  `renew_inspection_time` bigint(20) DEFAULT NULL COMMENT '自动续费检查时间',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_noble_id` (`noble_id`) USING BTREE,
  KEY `idx_expiration_time` (`expiration_time`),
  KEY `idx_noble_states` (`noble_states`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='COLUMNAR=1';

-- ----------------------------
-- Table structure for oauth_token
-- ----------------------------
DROP TABLE IF EXISTS `oauth_token`;
CREATE TABLE `oauth_token` (
  `user_id` bigint(20) unsigned NOT NULL,
  `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `login_type` int(11) NOT NULL,
  `create_time` bigint(10) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(10) NOT NULL DEFAULT '1451577600',
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE,
  UNIQUE KEY `idx_open_id-login_type` (`login_type`,`open_id`) USING BTREE,
  KEY `idx_open_id` (`open_id`),
  KEY `idx_login_type` (`login_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info` (
  `oid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '票据Id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户Id',
  `product` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品Id 对应commodity_info主键',
  `product_value` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品钻石',
  `product_bonus` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品赠送钻石',
  `product_description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品描述',
  `total_fee` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段',
  `currency` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '价格对应的单位 比如：IDR、MYR 前台支付\r\n对应commodity_info的currency字段',
  `currency_price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '价格 分 对应商品手机支付当地货币价格',
  `is_first_topup` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否首次充值 0非首充 1首充\r\n对应商品表的is_first_topup字段',
  `state` tinyint(11) NOT NULL DEFAULT '0' COMMENT '状态 0为待支付 1为已支付 ',
  `system_os` int(11) NOT NULL DEFAULT '1' COMMENT '所属系统；0未知,1安卓 2ios',
  `app_channel` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '所属渠道',
  `app_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '所属app名称',
  `pay_type` int(11) NOT NULL DEFAULT '0' COMMENT '支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay',
  `channel_code` int(11) NOT NULL DEFAULT '-1' COMMENT '支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定',
  `remarks` varchar(140) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注说明',
  `coupon_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '优惠券id',
  `referrer_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '推荐人id',
  `certificate_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '代理充值凭证',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`oid`) USING BTREE,
  KEY `idx_user_id` (`user_id`),
  KEY `idx_product_value` (`product_value`),
  KEY `idx_product` (`product`),
  KEY `idx_app_channel` (`app_channel`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_appchannel_state_updatetime` (`app_channel`,`state`,`update_time`),
  KEY `idx_state_updatetime` (`state`,`update_time`),
  KEY `idx_referrer_id` (`referrer_id`) USING BTREE,
  KEY `idx_paytype` (`pay_type`),
  KEY `idx_states` (`state`),
  KEY `idx_product_updatetime` (`product`,`update_time`),
  KEY `idx_userid_state_productvalue` (`user_id`,`state`,`product_value`)
) ENGINE=InnoDB AUTO_INCREMENT=1288469 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for order_statistics
-- ----------------------------
DROP TABLE IF EXISTS `order_statistics`;
CREATE TABLE `order_statistics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `begin_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间\r',
  `time_level` int(11) NOT NULL DEFAULT '0' COMMENT '时间维度：1.每小时，2.每天\r',
  `app_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app渠道\r',
  `pay_type` int(11) NOT NULL DEFAULT '0' COMMENT '支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay',
  `channel_code` int(11) NOT NULL DEFAULT '-1' COMMENT '支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定',
  `total_fee` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段',
  `all_order_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '总订单数',
  `success_order_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '付费订单数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unidx_olny_one` (`begin_time`,`time_level`,`app_channel`,`pay_type`,`channel_code`),
  KEY `idx_begin_time` (`begin_time`),
  KEY `idx_time_level` (`time_level`),
  KEY `idx_app_channel` (`app_channel`),
  KEY `idx_pay_type` (`pay_type`)
) ENGINE=InnoDB AUTO_INCREMENT=72621 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for payment_info
-- ----------------------------
DROP TABLE IF EXISTS `payment_info`;
CREATE TABLE `payment_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'app渠道',
  `pay_type` int(11) NOT NULL DEFAULT '0' COMMENT '支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay',
  `payment_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '支付渠道：比如Google、OVO、DANA等等',
  `currency_code` int(11) NOT NULL DEFAULT '0' COMMENT '国家数字码，比如印尼：360；马来：458  0为默认',
  `pay_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '支付图标',
  `channel_code` int(11) NOT NULL DEFAULT '0' COMMENT '支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定',
  `outside` tinyint(11) NOT NULL DEFAULT '0' COMMENT '是否外部网页打开：0不需要，1需要',
  `state` tinyint(11) NOT NULL DEFAULT '0' COMMENT '状态：0下架、1上架',
  `system_os` tinyint(11) NOT NULL DEFAULT '0' COMMENT '所属系统；0所有,1安卓 2ios',
  `is_default` tinyint(11) NOT NULL DEFAULT '0' COMMENT '是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序，顺序排列',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '说明备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for payout_info
-- ----------------------------
DROP TABLE IF EXISTS `payout_info`;
CREATE TABLE `payout_info` (
  `payout_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '代付id',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id 自增',
  `payout_type` int(11) NOT NULL COMMENT '转账类型  1 主播转账 2 公会长提现  3公会长主播提现',
  `order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '系统订单号',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '交易订单号',
  `operate_user_id` bigint(20) NOT NULL COMMENT '操作员用户id',
  `currency_code` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '付款币种编码',
  `currency_fee` bigint(20) NOT NULL COMMENT '付款币种总额，单位分',
  `paid_fee` bigint(20) DEFAULT NULL COMMENT '实际交易金额，单位分',
  `process_fee` bigint(20) DEFAULT NULL COMMENT '手续费，单位分',
  `bank_code` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '银行代码',
  `payee_account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收款人银行账号',
  `payee_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收款人姓名',
  `payee_email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收款人邮箱',
  `payee_phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收款人电话',
  `transfer_date` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '放款时间',
  `transfer_purpose` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用途',
  `pay_type` tinyint(4) NOT NULL COMMENT '支付类型',
  `transfer_status` tinyint(4) NOT NULL COMMENT '转账状态，0.未转账，1.转账中，2.已转账，3.转账失败，4.已退款，5.取消转账',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '修改时间',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`payout_id`) USING BTREE,
  UNIQUE KEY `order_no_UNIQUE` (`order_no`) USING BTREE,
  KEY `idx_trade_no` (`trade_no`) USING BTREE,
  KEY `idx_operate_user_id` (`operate_user_id`) USING BTREE,
  KEY `idx_pay_type` (`pay_type`) USING BTREE,
  KEY `idx_transfer_status` (`transfer_status`) USING BTREE,
  KEY `idx_create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='代付体现订单';

-- ----------------------------
-- Table structure for payout_record
-- ----------------------------
DROP TABLE IF EXISTS `payout_record`;
CREATE TABLE `payout_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `payout_id` bigint(20) NOT NULL COMMENT '代付id',
  `order_no` varchar(255) NOT NULL COMMENT '系统订单号',
  `trade_no` varchar(255) NOT NULL COMMENT '交易订单号',
  `status` varchar(255) NOT NULL COMMENT '状态',
  `failure_code` varchar(255) NOT NULL COMMENT '失败状态',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for prize_cashback_config
-- ----------------------------
DROP TABLE IF EXISTS `prize_cashback_config`;
CREATE TABLE `prize_cashback_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `cashback_ratio` int(11) NOT NULL COMMENT '返现比例',
  `five_times_number` int(11) NOT NULL COMMENT '5倍数量',
  `ten_times_number` int(11) NOT NULL COMMENT '10倍数量',
  `fifty_times_number` int(11) NOT NULL COMMENT '50倍数量',
  `one_hundred_times_number` int(11) NOT NULL COMMENT '100倍数量',
  `five_hundred_times_number` int(11) NOT NULL DEFAULT '0' COMMENT '500倍',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='奖池返现配置';

-- ----------------------------
-- Table structure for prize_pool_config
-- ----------------------------
DROP TABLE IF EXISTS `prize_pool_config`;
CREATE TABLE `prize_pool_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `gid` bigint(20) NOT NULL DEFAULT '0' COMMENT '幸运礼物id',
  `start_time` time NOT NULL DEFAULT '00:00:00' COMMENT '配置生效时间',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '配置类型 1:默认 2:其他',
  `cashback_ratio_id` bigint(20) NOT NULL COMMENT '奖池返现配置id',
  `state` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1:生效 2:失效',
  `use_count` int(11) NOT NULL DEFAULT '0' COMMENT '使用次数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='奖池配置';

-- ----------------------------
-- Table structure for quest_config
-- ----------------------------
DROP TABLE IF EXISTS `quest_config`;
CREATE TABLE `quest_config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `quest_name_json` varchar(255) NOT NULL DEFAULT '' COMMENT '任务名称',
  `quest_icon` varchar(255) NOT NULL DEFAULT '' COMMENT '任务图标',
  `quest_start_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务开始时间',
  `quest_end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务结束时间',
  `quest_key` varchar(100) NOT NULL DEFAULT '' COMMENT '任务唯一标识 key唯一则表示套娃任务 区分等级',
  `quest_parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务上级id -1 为一级任务',
  `quest_rank` int(11) NOT NULL DEFAULT '0' COMMENT '任务等级 套娃任务等级递增',
  `quest_cycle` int(11) NOT NULL DEFAULT '0' COMMENT '任务周期 1 每日 2 每周',
  `quest_theme` int(11) NOT NULL DEFAULT '0' COMMENT '任务主题 1 日常任务 2 主播任务 3 用户钻石任务 4 用户活跃宝箱',
  `quest_user_scope` int(11) NOT NULL DEFAULT '0' COMMENT '任务用户范围 1 所有用户 2 星级主播 3 公会长',
  `quest_type` int(11) NOT NULL DEFAULT '0' COMMENT '任务类型 common.QuestTypeSignIn',
  `quest_achieved` int(11) NOT NULL DEFAULT '0' COMMENT '任务达成值 1 次数任务代表1次 计时为1分钟',
  `is_reward` int(11) NOT NULL DEFAULT '0' COMMENT '任务是否有奖励 1 有',
  `quest_is_must` int(11) NOT NULL DEFAULT '0' COMMENT '任务是否必须 1 是',
  `quest_sort` int(11) NOT NULL DEFAULT '0' COMMENT '任务排序',
  `quest_states` int(11) NOT NULL DEFAULT '0' COMMENT '任务状态',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'createdAt',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'updatedAt',
  `automatic_release` int(11) NOT NULL COMMENT '自动发放 1 是',
  PRIMARY KEY (`id`),
  KEY `idx_quest_start_time` (`quest_start_time`),
  KEY `idx_quest_end_time` (`quest_end_time`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='任务配置表';

-- ----------------------------
-- Table structure for quest_prize
-- ----------------------------
DROP TABLE IF EXISTS `quest_prize`;
CREATE TABLE `quest_prize` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `quest_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务id',
  `prize_name` varchar(100) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `prize_icon` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品图标',
  `prize_type` int(11) NOT NULL DEFAULT '0' COMMENT '奖品类型 1装扮 2钻石 3优惠券 4贝壳 5 活跃值 6 用户经验 7 主播经验 8 主播收益',
  `link_id` varchar(100) NOT NULL DEFAULT '' COMMENT '关联商品id  1,2  逗号分割',
  `prize_num` int(11) NOT NULL DEFAULT '0' COMMENT '奖品数量',
  `status` bigint(20) NOT NULL DEFAULT '0' COMMENT '配置状态 0失效 1生效',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'createdAt',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'updatedAt',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='任务奖励';

-- ----------------------------
-- Table structure for quest_user_detail
-- ----------------------------
DROP TABLE IF EXISTS `quest_user_detail`;
CREATE TABLE `quest_user_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `quest_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务id',
  `quest_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '参与用户id',
  `expiration_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
  `quest_achieved` int(11) NOT NULL DEFAULT '0' COMMENT '任务达成值 1 次数任务代表1次 计时为1分钟',
  `is_complete` int(11) NOT NULL DEFAULT '0' COMMENT '是否完成',
  `complete_time` int(11) NOT NULL DEFAULT '0' COMMENT '完成时间',
  `quest_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '任务唯一标识 key唯一则表示套娃任务 区分等级',
  `quest_rank` int(11) NOT NULL DEFAULT '0' COMMENT '任务等级 套娃任务等级递增',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'createdAt',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT 'updatedAt',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `quest_user_detail_quest_id_quest_user_id_expiration_time_uindex` (`quest_id`,`quest_user_id`,`expiration_time`) COMMENT '任务进度唯一约束',
  KEY `idx_quest_id` (`quest_id`),
  KEY `idx_user_id` (`quest_user_id`),
  KEY `idx_ expiration_time` (`expiration_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1698855 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='任务用户参与详情';

-- ----------------------------
-- Table structure for recharge_points_record
-- ----------------------------
DROP TABLE IF EXISTS `recharge_points_record`;
CREATE TABLE `recharge_points_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `points` int(11) NOT NULL COMMENT '积分',
  `points_type` int(11) NOT NULL COMMENT '积分类型',
  `before_points` int(11) NOT NULL COMMENT '之前积分',
  `after_points` int(11) NOT NULL COMMENT '之后积分',
  `before_rank` int(11) NOT NULL COMMENT '之前排名',
  `after_rank` int(11) NOT NULL COMMENT '之后排名',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_points_type` (`points_type`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=5933 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='充值积分记录';

-- ----------------------------
-- Table structure for red_envelope_info
-- ----------------------------
DROP TABLE IF EXISTS `red_envelope_info`;
CREATE TABLE `red_envelope_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `anchor_user_id` bigint(20) NOT NULL COMMENT '主播id',
  `room_id` varchar(255) NOT NULL COMMENT '房间id',
  `red_envelope_type` int(11) NOT NULL COMMENT '红包类型 1 房间红包 2 全站红包',
  `red_envelope_count` int(11) NOT NULL COMMENT '红包个数',
  `red_envelope_diamond` int(11) NOT NULL COMMENT '红包钻石',
  `spent_diamond` int(11) NOT NULL COMMENT '已抢钻石',
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7998 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for red_envelope_record
-- ----------------------------
DROP TABLE IF EXISTS `red_envelope_record`;
CREATE TABLE `red_envelope_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `r_id` bigint(20) NOT NULL COMMENT '红包id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `room_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '房间id',
  `diamond` int(11) NOT NULL COMMENT '红包钻石',
  `before_diamond` int(11) NOT NULL COMMENT '红包之前余额',
  `after_diamond` int(11) NOT NULL COMMENT '红包之后余额',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=182523 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for report_record
-- ----------------------------
DROP TABLE IF EXISTS `report_record`;
CREATE TABLE `report_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `report_type` tinyint(8) NOT NULL DEFAULT '0' COMMENT '举报类型：1直播间举报、2用户、3动态、4、评论，后续扩展',
  `report_object` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '举报对象 1直播间举报存储直播间Id，动态举报存储动态Id',
  `content_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '举报内容类型，1侮辱谩骂，2色情低俗，3政治谣言，4其他',
  `describe` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '举报用户描述',
  `state` tinyint(8) NOT NULL DEFAULT '0' COMMENT '状态：0待处理 1已处理 后续扩展',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23615 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for revenue_statistics
-- ----------------------------
DROP TABLE IF EXISTS `revenue_statistics`;
CREATE TABLE `revenue_statistics` (
  `statistics_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `begin_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `time_level` int(11) NOT NULL DEFAULT '0' COMMENT '时间维度：1.每小时，2.每天',
  `total_expend_diamonds` bigint(20) NOT NULL DEFAULT '0' COMMENT '总消费（流水时间在时间范围内的支出类型的钻石数量）',
  `total_revenue_diamonds` bigint(20) NOT NULL DEFAULT '0' COMMENT '总收入（流水时间在时间范围内的支出类型的钻石数量）',
  `all_user_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '总用户量',
  `new_user_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '每天日活--认证时间在今天',
  `active_users_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '每天日活--认证时间在今天',
  `active_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播日活',
  `auth_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '认证主播数',
  `auth_star_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '认证标星主播数',
  `live_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播开播数',
  `live_star_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '标星主播开播数',
  `effective_live_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '有效开播数',
  `new_effective_live_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '新增有效开播数',
  `reach_effective_live_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '达标开播数',
  `total_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '总收入(印尼盾)',
  `incremental_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '充值用户数',
  `all_order_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '总订单数',
  `success_order_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '付费订单数',
  `total_paid_user_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '总支付人数',
  `incremental_paid_user_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '新用户支付人数',
  `yesterday_new_user_count` int(11) NOT NULL DEFAULT '0' COMMENT '昨天新增用户',
  `yesterday_user_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '昨天注册用户在今天登录数',
  `week_new_user_count` int(11) NOT NULL DEFAULT '0' COMMENT '上周新增用户',
  `week_user_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '上周注册用户在今天登录数',
  `month_new_user_count` int(11) NOT NULL DEFAULT '0' COMMENT '上月新增用户',
  `month_user_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '上月注册用户在今天登录数',
  `yesterday_new_anchor_count` int(11) NOT NULL DEFAULT '0' COMMENT '昨天新增认证主播',
  `yesterday_anchor_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '昨天注册主播在今天登录数',
  `week_new_anchor_count` int(11) NOT NULL DEFAULT '0' COMMENT '上周新增认证主播',
  `week_anchor_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '上周注册主播在今天登录数',
  `month_new_anchor_count` int(11) NOT NULL DEFAULT '0' COMMENT '上月新增认证主播',
  `month_anchor_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '上月注册主播在今天登录数',
  `yesterday_new_pay_user_count` int(11) NOT NULL DEFAULT '0' COMMENT '昨天新增付费用户',
  `yesterday_pay_user_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '昨天付费在今天登录数',
  `week_new_pay_user_count` int(11) NOT NULL DEFAULT '0' COMMENT '上周新增付费用户',
  `week_pay_user_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '上周付费在今天登录数',
  `month_new_pay_user_count` int(11) NOT NULL DEFAULT '0' COMMENT '上月新增付费用户',
  `month_pay_user_today_count` int(11) NOT NULL DEFAULT '0' COMMENT '上月付费在今天登录数',
  `remain_diamonds` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '剩余钻石数',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `app_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app名称',
  `app_channel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app渠道',
  `advertising_costs` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '广告费',
  PRIMARY KEY (`statistics_id`) USING BTREE,
  UNIQUE KEY `idx_begin_timeand_app_channel` (`begin_time`,`time_level`,`app_channel`) USING BTREE,
  KEY `idx_begin_time` (`begin_time`) USING BTREE,
  KEY `idx_app_channel` (`app_channel`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=60115 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for reward_record
-- ----------------------------
DROP TABLE IF EXISTS `reward_record`;
CREATE TABLE `reward_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会id',
  `reward_type` int(11) NOT NULL DEFAULT '0' COMMENT '奖励类型，  比如type  1：为3周2次结算  2……  后续扩展',
  `reward` bigint(20) NOT NULL DEFAULT '0' COMMENT '活动奖励给公会 印尼盾 分 \r\n该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9774 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for room_activity_banner
-- ----------------------------
DROP TABLE IF EXISTS `room_activity_banner`;
CREATE TABLE `room_activity_banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `icon` varchar(255) NOT NULL COMMENT '图片',
  `activity_type` int(11) NOT NULL COMMENT '活动类型 1 首冲 2 h5',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `state` int(11) NOT NULL DEFAULT '1' COMMENT '状态1启用',
  `show_type` int(11) NOT NULL DEFAULT '0' COMMENT '展示类型，0 全部，1 未充值用户 2 充值用户 3 主播',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for room_admin
-- ----------------------------
DROP TABLE IF EXISTS `room_admin`;
CREATE TABLE `room_admin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID 存储主播的id',
  `manager_id` bigint(20) NOT NULL COMMENT '管理员ID 主播指定的管理员 ',
  `valid` tinyint(8) NOT NULL DEFAULT '0' COMMENT '是否有效，0无效 1有效',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_user_id_manager_id` (`user_id`,`manager_id`) USING BTREE,
  KEY `uid` (`user_id`,`manager_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5825 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='直播间管理员';

-- ----------------------------
-- Table structure for room_live
-- ----------------------------
DROP TABLE IF EXISTS `room_live`;
CREATE TABLE `room_live` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '直播间Id',
  `user_id` bigint(20) NOT NULL COMMENT '主播ID',
  `room_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '直播间名称',
  `room_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '第三方直播间ID全局唯一',
  `stream_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '推拉流ID全局唯一',
  `image_path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '图片路径',
  `score` int(10) NOT NULL DEFAULT '0' COMMENT '排序分值',
  `member_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '累计观看人数',
  `send_gift_person_count` int(10) NOT NULL DEFAULT '0' COMMENT '送礼人数',
  `send_lucky_gift_person_count` int(10) NOT NULL DEFAULT '0' COMMENT '送幸运礼人数',
  `status` int(10) NOT NULL DEFAULT '0' COMMENT '-1 //所有状态\r\n0  //预开播\r\n1  //开播中\r\n2  //直播正常结束\r\n3  //由巡管或运营管理员将直播间关闭\r 4  //由服务器轮询出异常退出房间\r\n5  //预开播时检测到有未关播的房间，将其退出房间\r\n6  //其他异常退出房间\r\n7  //正常关闭预开播\r 8  //系统检测关闭预开播\r 9  //暂时离开',
  `coin_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '打赏金币数量 放大100倍',
  `begin_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '直播开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '直播结束时间',
  `last_heartbeat_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次心跳时间',
  `focus_count` int(10) NOT NULL DEFAULT '0' COMMENT '本次直播涨粉数量',
  `role` tinyint(4) NOT NULL DEFAULT '0' COMMENT '主播类型 0 用户 1 主播',
  `battle_status` int(10) NOT NULL DEFAULT '0' COMMENT 'pk状态 0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK',
  `private_status` int(2) NOT NULL DEFAULT '0' COMMENT '0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间',
  `game_id` int(10) NOT NULL DEFAULT '0' COMMENT '直播间配置的游戏id',
  `unlock_price` int(11) NOT NULL DEFAULT '0' COMMENT '钻石房解锁价格',
  `password` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码房密码',
  `wish_gift_id` int(11) NOT NULL DEFAULT '0' COMMENT '心愿礼物id',
  `wish_gift_count` int(11) NOT NULL DEFAULT '0' COMMENT '心愿礼物个数',
  `room_type` int(11) NOT NULL DEFAULT '0' COMMENT '房间类型 0 普通直播间  2 多人直播间',
  `max_multiple` int(11) NOT NULL DEFAULT '9' COMMENT '最大值窗口数',
  `layout_set` int(11) NOT NULL DEFAULT '1' COMMENT '布局设置  1 突出 2 平铺',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_userId_roomId` (`user_id`,`room_id`) USING BTREE,
  KEY `idx_room_id` (`room_id`),
  KEY `idx_room_name` (`room_name`) USING BTREE,
  KEY `idx_userid_updatetime_status` (`user_id`,`update_time`,`status`),
  KEY `idx_status_roomid` (`status`,`room_id`),
  KEY `idx_begin_time` (`begin_time`) USING BTREE,
  KEY `idx_create_time` (`create_time`),
  KEY `idx_end_time` (`end_time`),
  KEY `idx_update_time` (`update_time`),
  KEY `idx_status_lastheartbeattime_begintime` (`status`,`last_heartbeat_time`,`begin_time`),
  KEY `idx_lastheartbeattime_userid` (`last_heartbeat_time`,`user_id`),
  KEY `idx_streamid` (`stream_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2634457 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='COLUMNAR=1 直播间记录表';

-- ----------------------------
-- Table structure for room_live_battle
-- ----------------------------
DROP TABLE IF EXISTS `room_live_battle`;
CREATE TABLE `room_live_battle` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '直播间Id',
  `user_id` bigint(20) NOT NULL COMMENT '主播ID',
  `battle_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'pk类型 0 随机 1 指定 2后台拉起',
  `stream_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '推拉流ID全局唯一',
  `room_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '房间ID',
  `diamonds_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '打赏钻石数量',
  `score_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '打赏积分数量',
  `target_user_id` bigint(20) NOT NULL COMMENT '对手主播ID',
  `target_stream_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '对手stream_id',
  `target_diamonds_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '对手打赏钻石数量',
  `target_score_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '对手打赏积分数量',
  `target_room_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '房间ID',
  `battle_status` int(10) NOT NULL DEFAULT '0' COMMENT '0 匹配中 1 pk中 2 邀请人pk取消 3 对方拒绝 4 惩罚中 5 pk正常结束 6 pk异常结束',
  `end_type` int(11) NOT NULL DEFAULT '-1' COMMENT '结束原因 0 正常结束 1 超时结束 2 手动结束 3 系统结束',
  `begin_time` bigint(20) NOT NULL DEFAULT '0' COMMENT 'pk开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT 'pk结束时间',
  `user_answer_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户响应时间',
  `target_answer_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '对方响应时间',
  `create_time` bigint(20) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`),
  KEY `idx_battle_status` (`battle_status`),
  KEY `idx_end_time` (`end_time`),
  KEY `idx_room_id` (`room_id`),
  KEY `idx_target_room_id` (`target_room_id`),
  KEY `idx_target_user_id` (`target_user_id`),
  KEY `idx_room_live_battle` (`battle_type`),
  KEY `idx_userid_battlestatus_updatetime` (`user_id`,`battle_status`,`update_time`),
  KEY `idx_targetuserid_battlestatus_updatetime` (`target_user_id`,`battle_status`,`update_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1144371 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='pk记录表';

-- ----------------------------
-- Table structure for room_live_call
-- ----------------------------
DROP TABLE IF EXISTS `room_live_call`;
CREATE TABLE `room_live_call` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `anchor_user_id` bigint(20) NOT NULL COMMENT '主播用户id',
  `room_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '房间id',
  `call_states` int(11) NOT NULL COMMENT '连麦状态 0 排队中 1连麦中 2 已结束',
  `stream_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '连麦流id',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=94857 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for room_multiple_user
-- ----------------------------
DROP TABLE IF EXISTS `room_multiple_user`;
CREATE TABLE `room_multiple_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `portrait` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '头像',
  `room_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '房间id',
  `anchor_user_id` bigint(20) NOT NULL COMMENT '主播用户id',
  `live_control` int(11) NOT NULL DEFAULT '1' COMMENT '直播控制 1 正常 2 禁视频 ',
  `sound_control` int(11) NOT NULL DEFAULT '1' COMMENT '直播控制 1 正常 2 静音',
  `live_states` int(11) NOT NULL COMMENT '多人房状态 0 申请中 1 上麦中 2 已拒绝 3 已结束 4 忽略',
  `stream_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '流id 房主的为流混流拉流id',
  `multiple_user_role` int(11) NOT NULL COMMENT '多人房身份 1 多人房主播 2 上麦者',
  `main_status` int(11) NOT NULL DEFAULT '0' COMMENT '主麦状态',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40725 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for room_remove_list
-- ----------------------------
DROP TABLE IF EXISTS `room_remove_list`;
CREATE TABLE `room_remove_list` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT 'UID',
  `operator_id` bigint(20) NOT NULL COMMENT '操作者ID',
  `room_id` bigint(20) NOT NULL COMMENT '房间ID',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '移出结束时间,为0就是永久拉黑',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`user_id`,`room_id`,`end_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='移出房间记录';

-- ----------------------------
-- Table structure for room_review_record
-- ----------------------------
DROP TABLE IF EXISTS `room_review_record`;
CREATE TABLE `room_review_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room_id` varchar(100) NOT NULL COMMENT '房间id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `stat_code` int(11) NOT NULL COMMENT '审核状态。0：审核中1：审核结束',
  `risk_level` int(11) NOT NULL COMMENT '风险级别  0：pass，正常内容， 1：review，可疑内容 2：reject，违规内容',
  `content_type` int(11) NOT NULL COMMENT '用来区分音频和图片回调。 1：该回调为图片回调 2：该回调为音频回调',
  `stream_id` varchar(150) NOT NULL COMMENT '开始视频流审核时输入的流 ID。',
  `process_start_time` bigint(20) NOT NULL COMMENT '开始处理时间戳（毫秒）。',
  `process_end_time` bigint(20) NOT NULL COMMENT '检测完成时间戳（毫秒）',
  `risk_type` int(11) NOT NULL COMMENT '标识风险类型。 0：正常 100：涉政 200：色情 210：性感 300：广告 310：二维码 320：水印 400：暴恐 500：违规 510：不良场景 530：行业违规',
  `risk_source` int(11) NOT NULL COMMENT '风险来源。 1000：无风险。 1001：文字有风险（OCR 图片识别文字）。 1002：截帧有风险。',
  `matched_detail` varchar(100) NOT NULL COMMENT '命中所有敏感词名单的详情。 调用 开始视频流审核 时，如果 ImageType[] 传值包括 5（广告识别）、10（OCR），会触发敏感词识别。',
  `matched_item` varchar(100) NOT NULL COMMENT '命中的具体敏感词',
  `matched_list` varchar(100) NOT NULL COMMENT '命中敏感词所在的名单名称',
  `description` varchar(100) NOT NULL COMMENT '策略规则风险原因描述。',
  `detect_type` int(11) NOT NULL COMMENT '用来区分截帧图片是否过了检测。 1：截帧图片过了检测 2：截帧图片没过检测',
  `similarity` decimal(10,2) NOT NULL COMMENT '与上一张截帧图片的相似概率值 取值范围 [0-1]，数值越接近 1 越相似。',
  `still_time` int(11) NOT NULL COMMENT '展示静止画面时间（秒）',
  `image_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '截帧图片 url 地址。',
  `image_time` bigint(20) NOT NULL COMMENT '视频流截帧图片违规发生的时间戳（毫秒）。',
  `image_text` varchar(100) NOT NULL COMMENT '视频中画面识别出的文字内容',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_room_id` (`room_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_risk_level` (`risk_level`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_risklevel_createtime` (`risk_level`,`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=6427682 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for room_score_record
-- ----------------------------
DROP TABLE IF EXISTS `room_score_record`;
CREATE TABLE `room_score_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `room_id` varchar(255) NOT NULL,
  `operation_type` int(11) NOT NULL COMMENT '操作类型',
  `operation_id` bigint(20) NOT NULL COMMENT '操作id',
  `score` int(11) NOT NULL COMMENT '排序分值',
  `before_score` int(11) NOT NULL COMMENT 'before排序分值',
  `after_score` int(11) NOT NULL COMMENT 'after排序分值',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  KEY `idx_roomid_createtime` (`room_id`,`create_time`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_room_id` (`room_id`)
) ENGINE=InnoDB AUTO_INCREMENT=313350 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='房间分数记录';

-- ----------------------------
-- Table structure for room_unlock_record
-- ----------------------------
DROP TABLE IF EXISTS `room_unlock_record`;
CREATE TABLE `room_unlock_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `room_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `anchor_user_id` bigint(20) NOT NULL COMMENT '主播id',
  `is_unlock` int(255) NOT NULL DEFAULT '0' COMMENT '是否解锁',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `unlock_price` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4377 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for room_user_record
-- ----------------------------
DROP TABLE IF EXISTS `room_user_record`;
CREATE TABLE `room_user_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `room_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '房间ID',
  `exit_type` int(10) NOT NULL DEFAULT '0' COMMENT '用户退出房间原因，0 自己退出；1 房间管理员；2 主播提出房间 3 超级管理踢出房间；4 主播拉黑；5 异常退出房间',
  `extend` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '额外信息',
  `is_chat` int(2) NOT NULL DEFAULT '0' COMMENT '是否有聊天',
  `is_wait_long` int(2) NOT NULL DEFAULT '0' COMMENT '是否待满5分钟',
  `diamonds` bigint(20) NOT NULL DEFAULT '0' COMMENT '在此房间消费的钻石流水',
  `begin_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '进入房间时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '退出房间时间',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`user_id`) USING BTREE,
  KEY `idx_userid_roomid` (`user_id`,`room_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7400202 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户在直播间操作记录表';

-- ----------------------------
-- Table structure for safe_info
-- ----------------------------
DROP TABLE IF EXISTS `safe_info`;
CREATE TABLE `safe_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `content` varchar(500) NOT NULL DEFAULT '' COMMENT '内容',
  `safe_type` int(6) NOT NULL DEFAULT '0' COMMENT '类型，1头像，2昵称，3签名，后续扩展',
  `state` int(6) NOT NULL DEFAULT '0' COMMENT '状态，0待审核，1审核通过，2审核不通过',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for server_guild_live_statistics
-- ----------------------------
DROP TABLE IF EXISTS `server_guild_live_statistics`;
CREATE TABLE `server_guild_live_statistics` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `begin_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '工会id',
  `guild_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '工会名称',
  `server_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  `live_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播开播数',
  `live_star_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '标星主播开播数',
  `active_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播在线数',
  `star_active_anchor_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '标星主播在线数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=411141 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for server_notice
-- ----------------------------
DROP TABLE IF EXISTS `server_notice`;
CREATE TABLE `server_notice` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统自增',
  `server_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `create_admin_id` bigint(10) NOT NULL DEFAULT '0' COMMENT '通知创建者Id  0系统创建',
  `notice_type` int(11) NOT NULL DEFAULT '2' COMMENT '通知类型：1系统星级调整 2人工星级调整  后续其他扩展',
  `before_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '之前内容',
  `after_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '之后内容',
  `system_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '系统备注说明',
  `message_type` int(11) NOT NULL DEFAULT '1' COMMENT '消息说明类别：系统星级下：0系统星级保持、1系统星级升级、2系统星级降级  \r\n人工星级调整\r\n1人工星级降级 2人工星级降级  后续扩展',
  `state` int(11) NOT NULL DEFAULT '3000' COMMENT '状态 0未处理 1已知晓 2已处理',
  `operator` bigint(10) NOT NULL DEFAULT '0' COMMENT '最后操作人',
  `server_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '运营号备注',
  `create_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `labour_union_to_anchors_lu_type` (`message_type`) USING BTREE,
  KEY `idx_server_id` (`server_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_state` (`state`) USING BTREE,
  KEY `idx_notice_type` (`notice_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=60647 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for server_week_statistics
-- ----------------------------
DROP TABLE IF EXISTS `server_week_statistics`;
CREATE TABLE `server_week_statistics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  `week` int(11) NOT NULL DEFAULT '0' COMMENT '周，一年中的第几周 例如:202230 2022年的第30周',
  `guild_num` int(11) NOT NULL DEFAULT '0' COMMENT '公会数量',
  `effective_guild_num` int(11) NOT NULL DEFAULT '0' COMMENT '有效公会数量',
  `new_guild_num` int(11) NOT NULL DEFAULT '0' COMMENT '新增公会数量',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=994 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for share_new_user_info
-- ----------------------------
DROP TABLE IF EXISTS `share_new_user_info`;
CREATE TABLE `share_new_user_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `share_id` int(11) NOT NULL COMMENT '分享ID',
  `owner_user_id` bigint(20) NOT NULL COMMENT '分享者uid',
  `watch_live_time` int(11) NOT NULL COMMENT '观看直播时长',
  `register_time` bigint(20) DEFAULT NULL COMMENT '注册时间',
  `ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登陆地IP',
  `smid` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备ID',
  `device_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备型号',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '0:初始化；1:待加币；2:加币成功',
  `watch_live_rewards` int(11) NOT NULL DEFAULT '1' COMMENT '1:看播奖励有效；2:看播奖励无效',
  `is_old_user` int(11) NOT NULL DEFAULT '0' COMMENT '1:老用户；0:新用户',
  `full_reason_code` int(11) NOT NULL DEFAULT '0' COMMENT '失败原因 status -1时有效',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_owner_uid` (`owner_user_id`),
  KEY `idx_share_id` (`share_id`),
  KEY `idx_status` (`status`),
  KEY `uniq_uid` (`user_id`) USING BTREE,
  KEY `idx_owneruserid_userid_status` (`owner_user_id`,`user_id`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=61034 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for share_new_user_record
-- ----------------------------
DROP TABLE IF EXISTS `share_new_user_record`;
CREATE TABLE `share_new_user_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `owner_uid` bigint(20) NOT NULL COMMENT '分享者uid',
  `type` int(11) NOT NULL COMMENT '1:新用户注册；2:新用户充值',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '0:初始化；1:待加币；2:加币成功',
  `owner_order_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分享者加币订单号',
  `owner_diamond_num` int(11) DEFAULT NULL COMMENT '分享者获得的钻石数',
  `new_user_order_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '新用户加币订单',
  `new_user_diamond_num` int(11) DEFAULT NULL COMMENT '新用户获得的钻石数',
  `code_id` int(11) NOT NULL COMMENT '分享ID',
  `register_time` bigint(20) DEFAULT NULL COMMENT '注册时间',
  `recharge_num` int(11) NOT NULL COMMENT '充值的虚拟币数量',
  `watch_live_time` int(11) NOT NULL COMMENT '观看直播时长',
  `ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登陆地IP',
  `smid` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备ID',
  `device_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备型号',
  `update_time` bigint(20) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for shop_goods
-- ----------------------------
DROP TABLE IF EXISTS `shop_goods`;
CREATE TABLE `shop_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `category_id` int(11) unsigned NOT NULL COMMENT '类型',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '商品状态 0下架 1上架',
  `icon_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标地址',
  `res_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '资源地址',
  `icon_size` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标尺寸',
  `price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '单价 钻石数',
  `forever_price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '永久售卖价值 后续使用对于有永久售卖的商品',
  `desc` varchar(521) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品描述',
  `discount_price` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '[]' COMMENT '折扣 json 配置不同天数的价格 ',
  `user_level_limit` int(11) NOT NULL COMMENT '用户等级限制',
  `shelf_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '上架时间',
  `is_renewable` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可续费 该字段基本是没用',
  `sort_weight` int(11) NOT NULL DEFAULT '0' COMMENT '排序权重 升序',
  `sell_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '售卖类型 1-售卖商品；2-活动商品',
  `quality` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '品质',
  `extra` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '{}' COMMENT '扩展属性',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `number` bigint(20) NOT NULL DEFAULT '0' COMMENT '号码',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_query` (`category_id`,`status`,`sell_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=182 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='商场商品表';

-- ----------------------------
-- Table structure for shop_order
-- ----------------------------
DROP TABLE IF EXISTS `shop_order`;
CREATE TABLE `shop_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `shop_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '购买数量（月）',
  `buy_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买人uid',
  `receive_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '接收人',
  `coin_number` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '购买金额 钻石数（不扩大）',
  `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单描述',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1617 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商场订单表';

-- ----------------------------
-- Table structure for short_urls
-- ----------------------------
DROP TABLE IF EXISTS `short_urls`;
CREATE TABLE `short_urls` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `raw_url` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `short_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `create_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shortCode` (`short_code`)
) ENGINE=InnoDB AUTO_INCREMENT=33899 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='短链映射表';

-- ----------------------------
-- Table structure for star_level_reset_config
-- ----------------------------
DROP TABLE IF EXISTS `star_level_reset_config`;
CREATE TABLE `star_level_reset_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `star_level` int(11) NOT NULL DEFAULT '0' COMMENT '主播星级',
  `rule_type` tinyint(8) NOT NULL DEFAULT '0' COMMENT '类别 0保级 1升级 2降级',
  `income_min` bigint(20) NOT NULL DEFAULT '0' COMMENT '金币min 不用扩大100',
  `send_gift_person_min` int(11) NOT NULL DEFAULT '0' COMMENT '送礼人数min',
  `self_send_rate_max` tinyint(8) NOT NULL DEFAULT '0' COMMENT '自己送自己金币占比max 0.4 存储40 扩大100倍',
  `create_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(10) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for sync_task
-- ----------------------------
DROP TABLE IF EXISTS `sync_task`;
CREATE TABLE `sync_task` (
  `id` bigint(20) unsigned NOT NULL COMMENT '任务id',
  `src_table` varchar(255) NOT NULL DEFAULT '' COMMENT '数据来源表',
  `dst_table` varchar(255) NOT NULL DEFAULT '' COMMENT '数据去向表',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `sync_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '已同步/统计的数据来源表的最大记录id',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sys_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `nick_name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  `status` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `sys_admin_un` (`user_name`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员表';

-- ----------------------------
-- Table structure for sys_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_role`;
CREATE TABLE `sys_admin_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `server_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=372 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户与角色关系表';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '上级菜单ID',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单标题',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '组件名称',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '组件',
  `menu_sort` int(11) NOT NULL COMMENT '排序',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接地址',
  `create_time` bigint(20) NOT NULL COMMENT '创建日期',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `titleCn` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '中文菜单标题',
  PRIMARY KEY (`menu_id`) USING BTREE,
  KEY `inx_pid` (`pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=157 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统菜单';

-- ----------------------------
-- Table structure for sys_operate_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_operate_log`;
CREATE TABLE `sys_operate_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `begin_time` bigint(20) NOT NULL COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求地址',
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'ip',
  `params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求参数',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'HTTP方法',
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '返回信息',
  `host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求主机',
  `brower` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '浏览器',
  PRIMARY KEY (`id`),
  KEY `idx_begin_time` (`begin_time`),
  KEY `idx_admin_id` (`admin_id`),
  KEY `idx_uri` (`uri`)
) ENGINE=InnoDB AUTO_INCREMENT=108578 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `role_cn_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色中文名称',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色表';

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11318 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色与权限关系表';

-- ----------------------------
-- Table structure for sys_server_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_server_role`;
CREATE TABLE `sys_server_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联admin主键',
  `role` tinyint(4) NOT NULL DEFAULT '0' COMMENT '角色0普通，1管理',
  `group` tinyint(4) NOT NULL DEFAULT '0' COMMENT '管理组，1雅加达，2三宝垄',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建日期',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_server_role_un` (`admin_id`),
  KEY `idx_role` (`role`) USING BTREE,
  KEY `idx_admin_id` (`admin_id`) USING BTREE,
  KEY `idx_group` (`group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for third_party_busi_game_order
-- ----------------------------
DROP TABLE IF EXISTS `third_party_busi_game_order`;
CREATE TABLE `third_party_busi_game_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `order_type` int(32) NOT NULL DEFAULT '0' COMMENT '1   : 增加钻石, 2:扣减钻石',
  `busi_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方业务唯一标识',
  `himi_uid` int(11) NOT NULL COMMENT 'uid',
  `coin_num` int(11) NOT NULL DEFAULT '0' COMMENT '订单虚拟币额度',
  `busi_uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方业务用户唯一标识',
  `order_status` int(11) NOT NULL DEFAULT '0' COMMENT '订单状态 0-下单,1-执行成功',
  `room_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '房间id',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_busi_order` (`busi_id`) USING BTREE,
  KEY `idx_uid_busi_id` (`himi_uid`,`busi_id`) USING BTREE,
  KEY `idx_order_type` (`order_type`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_coinnum` (`coin_num`) COMMENT 'create by DAS-1e536130-fb4f-42de-9413-a0986bbcfebb'
) ENGINE=InnoDB AUTO_INCREMENT=183027773 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='第三方游戏业务订单信息表';

-- ----------------------------
-- Table structure for topic_follow
-- ----------------------------
DROP TABLE IF EXISTS `topic_follow`;
CREATE TABLE `topic_follow` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `topic_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '话题Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_topic_id_user_id` (`topic_id`,`user_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_topic_id` (`topic_id`) USING BTREE,
  KEY `idx_create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2577 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='话题关注表';

-- ----------------------------
-- Table structure for topic_info
-- ----------------------------
DROP TABLE IF EXISTS `topic_info`;
CREATE TABLE `topic_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '话题标题',
  `cover_img` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面图',
  `score` bigint(20) NOT NULL DEFAULT '0' COMMENT '值越大 热度越高',
  `topic_type` tinyint(11) NOT NULL DEFAULT '0' COMMENT '0普通话题 1官方 2审核模式话题',
  `recommend_score` int(11) NOT NULL DEFAULT '0' COMMENT '推荐 默认是0  排序倒序 值越大越靠前',
  `state` tinyint(11) NOT NULL DEFAULT '0' COMMENT '状态：0待审核 1审核通过 2审核不通过',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_title` (`title`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_score` (`score`) USING BTREE,
  KEY `idx_create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=690 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='话题信息表';

-- ----------------------------
-- Table structure for trades_info
-- ----------------------------
DROP TABLE IF EXISTS `trades_info`;
CREATE TABLE `trades_info` (
  `orderId` bigint(20) unsigned NOT NULL COMMENT '订单Id',
  `tid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '票据Id',
  `pay_type` int(11) NOT NULL DEFAULT '0' COMMENT '支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay',
  `total_fee` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '支付费用，三方支付通知有，没有则填充订单的对应货币金额 单位分',
  `currency` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '实际支付货币，三方通知有说明，比如：IDR，没有则填充对应货币单位',
  `state` tinyint(11) NOT NULL DEFAULT '-1' COMMENT '''SUCCESS'': 1,   // 交易成功\r    ''UNKNOWN'': 0,   // 未知的交易状态\r    ''REFUNDED'': 2,  // 退款\r    ''RENEW'': 3,     // 续订\r    ''WAITING'': 4,   // 等待支付',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`orderId`) USING BTREE,
  UNIQUE KEY `trades_unique_tid_paytype` (`tid`,`pay_type`) USING BTREE,
  KEY `index_tid` (`tid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for update_record
-- ----------------------------
DROP TABLE IF EXISTS `update_record`;
CREATE TABLE `update_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `admin_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '管理员id',
  `update_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '调整类型，1星级，2公会，3标签，4备注',
  `before_body` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '修改前内容',
  `after_body` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '修改后内容',
  `remark` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '修改后内容',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=206 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='新用户热门记录表';

-- ----------------------------
-- Table structure for user_bank
-- ----------------------------
DROP TABLE IF EXISTS `user_bank`;
CREATE TABLE `user_bank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `id_card_num` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '身份证',
  `id_card_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '身份证姓名',
  `id_card_front` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '身份证正面',
  `id_card_back` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '身份证背面',
  `bank_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账户名',
  `bank_accounts` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行账号',
  `bank_person` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行卡持有人',
  `created_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `updated_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11848 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for user_base_info
-- ----------------------------
DROP TABLE IF EXISTS `user_base_info`;
CREATE TABLE `user_base_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户Id',
  `first_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '首播日期',
  `first_auth_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次认证日期',
  `last_auth_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后认证日期',
  `live_call_switch` int(11) NOT NULL DEFAULT '0' COMMENT '连麦开关 1打开',
  `live_call_type` int(11) NOT NULL DEFAULT '1' COMMENT '连麦类型 1 所有人 2 关注我的人 3我的好友',
  `review_up_user` int(11) NOT NULL DEFAULT '1' COMMENT '多人房上麦审核',
  `multiple_user_type` int(11) NOT NULL DEFAULT '1' COMMENT '多人房上麦权限  1 所有人 2 关注我的人 3我的好友',
  `created_time` bigint(10) NOT NULL DEFAULT '1451577600',
  `updated_time` bigint(10) NOT NULL DEFAULT '1451577600',
  `server_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运营号id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  KEY `idx_first_auth_time` (`first_auth_time`)
) ENGINE=InnoDB AUTO_INCREMENT=184353 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_battle_match
-- ----------------------------
DROP TABLE IF EXISTS `user_battle_match`;
CREATE TABLE `user_battle_match` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `match_season` int(11) NOT NULL COMMENT '比赛赛季 递增',
  `match_score` int(11) NOT NULL COMMENT '赛季分数',
  `match_max_score` int(11) NOT NULL COMMENT '赛季最高分数',
  `win_battle_count` int(11) NOT NULL COMMENT '赢场次',
  `dog_fall_battle_count` int(11) NOT NULL COMMENT '平局场次',
  `all_battle_count` int(11) NOT NULL COMMENT '所有pk赛场次',
  `win_streak_count` int(11) NOT NULL COMMENT '连胜场次（中断归零）',
  `grading_is_complete` int(11) NOT NULL COMMENT '是否完成定级',
  `grading_count` int(11) NOT NULL COMMENT '已打定级赛数量',
  `win_grading_count` int(11) NOT NULL DEFAULT '0' COMMENT '已赢定级赛数量',
  `grading_score` int(11) NOT NULL COMMENT '定级赛分数',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3216 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for user_check_first
-- ----------------------------
DROP TABLE IF EXISTS `user_check_first`;
CREATE TABLE `user_check_first` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `check_type` int(11) NOT NULL COMMENT '检查类型',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_userid_checktype` (`user_id`,`check_type`)
) ENGINE=InnoDB AUTO_INCREMENT=2080186 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for user_dress_up
-- ----------------------------
DROP TABLE IF EXISTS `user_dress_up`;
CREATE TABLE `user_dress_up` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `shop_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `category_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类ID',
  `start_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间 -1标识永久配置',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0-未佩戴 1-佩戴 ',
  `effect_days` int(11) NOT NULL DEFAULT '0' COMMENT '有效期 （单位：天）-1标识永久',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uid_shop_id` (`user_id`,`shop_id`) USING BTREE COMMENT '根据用户UID和商品ID确认一条用户装扮'
) ENGINE=InnoDB AUTO_INCREMENT=137097 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户装扮表';

-- ----------------------------
-- Table structure for user_dress_up_log
-- ----------------------------
DROP TABLE IF EXISTS `user_dress_up_log`;
CREATE TABLE `user_dress_up_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `shop_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `category_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类ID',
  `dress_type` tinyint(8) NOT NULL DEFAULT '0' COMMENT '0：卸载装扮 1：穿戴装扮',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=121268 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户装扮表';

-- ----------------------------
-- Table structure for user_experience
-- ----------------------------
DROP TABLE IF EXISTS `user_experience`;
CREATE TABLE `user_experience` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `user_exp` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'user经验值',
  `anchor_exp` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'anchor经验值',
  `consume` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '消费',
  `income` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '收入',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`user_id`),
  KEY `income` (`income`),
  KEY `consume` (`consume`)
) ENGINE=InnoDB AUTO_INCREMENT=273949361 DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin COMMENT='用户个人统计信息';

-- ----------------------------
-- Table structure for user_focus
-- ----------------------------
DROP TABLE IF EXISTS `user_focus`;
CREATE TABLE `user_focus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '关注ID',
  `like_id` bigint(20) NOT NULL COMMENT '被关注ID',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0 相互未关注 1 单方关注 2互相关注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`),
  KEY `idx_like_id` (`like_id`),
  KEY `idx_user_id-like_id` (`user_id`,`like_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2319466 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户关注主播列表';

-- ----------------------------
-- Table structure for user_focus_feed_record
-- ----------------------------
DROP TABLE IF EXISTS `user_focus_feed_record`;
CREATE TABLE `user_focus_feed_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '关注ID',
  `like_id` bigint(20) NOT NULL COMMENT '被关注ID',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_user_id-like_id` (`user_id`,`like_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_like_id` (`like_id`) USING BTREE,
  KEY `idx_user_id-like_id` (`user_id`,`like_id`) USING BTREE,
  KEY `idx_create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=728798 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户查看关注主播最后一次时间';

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户Id 自增',
  `super_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '靓号 默认为0',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码  空字符串说明没有密码',
  `gender` int(11) NOT NULL DEFAULT '1' COMMENT '性别 1男 2女 0未知',
  `birthday` bigint(20) unsigned NOT NULL COMMENT '生日 时间戳',
  `height` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '身高 cm',
  `education` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '学历，枚举 小学 高中 大学 本科 扩展',
  `industry` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '行业，枚举 学生、文化/艺术 扩展',
  `intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '自我介绍 签名',
  `portrait` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '头像',
  `live_state` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0未开播 1开播',
  `whats_app` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'whatsapp号码',
  `role` int(11) NOT NULL DEFAULT '0' COMMENT '0普通用户 1主播 2管理员 后续扩展',
  `ban_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '封禁时间，0代表未封禁',
  `last_heartbeat_at` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '用户心跳',
  `follow_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '我关注别人的数量',
  `followed_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '我的粉丝数量（别人关注我）',
  `is_new_user` tinyint(6) unsigned NOT NULL DEFAULT '0' COMMENT '0为老用户 1为新用户',
  `device_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `app_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app名称',
  `app_channel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道',
  `app_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'app版本',
  `phone_mode` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机型号',
  `phone_system_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机系统版本',
  `last_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '最后一次登录ip',
  `system` int(10) NOT NULL DEFAULT '0' COMMENT '1为安卓 2为ios',
  `user_language` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户语言',
  `currency_code` int(11) NOT NULL DEFAULT '0' COMMENT '国家数字码，比如印尼：360；马来：458  0为默认',
  `remarks` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `anchor_high_quality` int(11) DEFAULT '0' COMMENT '是否高质量主播 1高质量',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  KEY `idx_ nickname` (`nickname`),
  KEY `idx_live_state` (`live_state`),
  KEY `idx_app_name` (`app_name`),
  KEY `idx_last_heartbeat_at` (`last_heartbeat_at`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_app_channel` (`app_channel`),
  KEY `idx_role` (`role`)
) ENGINE=InnoDB AUTO_INCREMENT=2595001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_props
-- ----------------------------
DROP TABLE IF EXISTS `user_props`;
CREATE TABLE `user_props` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `props_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '道具名称',
  `props_type` int(11) NOT NULL COMMENT ' 道具类型 1 礼物卡',
  `props_link` bigint(20) NOT NULL COMMENT '道具关联id',
  `props_states` int(11) NOT NULL COMMENT '道具状态 1待使用 2已使用',
  `use_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '使用时间',
  `target_user_id` bigint(20) NOT NULL COMMENT '使用目标id',
  `expiration_time` bigint(20) NOT NULL COMMENT '过期时间 -1 不过期',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`),
  KEY `idx_props_type` (`props_type`),
  KEY `idx_props_states` (`props_states`),
  KEY `idx_props_link` (`props_link`),
  KEY `idx_expiration_time` (`expiration_time`),
  KEY `idx_createtime` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=811986 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_props_log
-- ----------------------------
DROP TABLE IF EXISTS `user_props_log`;
CREATE TABLE `user_props_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `target_user_id` bigint(20) NOT NULL COMMENT '使用目标id',
  `from_type` int(11) NOT NULL COMMENT '来源类型 common.PropsLogFromTypeGift',
  `num` int(11) NOT NULL COMMENT '数量',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`),
  KEY `idx_target_user_id` (`target_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=687214 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_rank_rule
-- ----------------------------
DROP TABLE IF EXISTS `user_rank_rule`;
CREATE TABLE `user_rank_rule` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL DEFAULT '0' COMMENT '等级从0开始, 0为等级1',
  `min` bigint(20) NOT NULL DEFAULT '0' COMMENT '最小值',
  `max` bigint(20) NOT NULL DEFAULT '0' COMMENT '最大值',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=161 DEFAULT CHARSET=utf8mb4 COMMENT='用户等级规则';

-- ----------------------------
-- Table structure for week_settlement_reward
-- ----------------------------
DROP TABLE IF EXISTS `week_settlement_reward`;
CREATE TABLE `week_settlement_reward` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `settlement_week` int(11) NOT NULL DEFAULT '0' COMMENT '结算周',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播Id',
  `reward_role` tinyint(10) NOT NULL DEFAULT '0' COMMENT '奖励给谁 1 主播  2公会 ',
  `reward_money` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励金额 印尼盾 ',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_week-user_id-reward_role` (`settlement_week`,`user_id`,`reward_role`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1635 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='周结算奖励表';

-- ----------------------------
-- Table structure for week_settlement_special_anchor
-- ----------------------------
DROP TABLE IF EXISTS `week_settlement_special_anchor`;
CREATE TABLE `week_settlement_special_anchor` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `week` int(11) NOT NULL DEFAULT '0' COMMENT '周，一年中的第几周 例如:202230 2022年的第30周',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `guild_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '公会ID',
  `live_time` int(11) NOT NULL DEFAULT '0' COMMENT '直播时长（s）-优质主播统计',
  `effect_live_day` tinyint(8) NOT NULL DEFAULT '0' COMMENT '有效出勤天数-优质主播统计',
  `level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '优质等级 A/AA等',
  `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否达标优质结算\n0未达标  1达标 ',
  `state_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '状态说明',
  `base_gear_money_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '基础金币档位底薪 印尼盾 单位分',
  `special_gear_money_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '特殊主播底薪 印尼盾 单位分 比如A类、AA类主播',
  `real_gear_money_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '真实主播底薪 印尼盾 单位分 比如A类、AA类主播\n真实底薪（如果基础底薪小于优质底薪 这里存储最高的，结算表anchor_week_settlement中的RewardMoneyForHost与该字段保持一致）',
  `surplus_money_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '底薪差额 RealGearMoneyForHost-baseGearMoneyForHost ',
  `base_guild_rate` int(11) NOT NULL DEFAULT '0' COMMENT '基础公会礼物提成比列  *100的整数 基础公会提成比列',
  `special_guild_rate` int(11) NOT NULL DEFAULT '0' COMMENT '特殊公会礼物提成比列  *100的整数 比如优质主播公会提成比列',
  `real_guild_rate` int(11) NOT NULL DEFAULT '0' COMMENT '真实公会礼物提成比列  *100的整数 与结算表anchor_week_settlement中的guild_rate与该字段保持一致',
  `reward_money_for_guild` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励给公会 印尼盾 分 这里只是记录，真正的数据累加进结算表的 activity_reward_for_guild字段中',
  `reward_money_for_host` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励给主播 印尼盾 分 这里只是记录，真正的数据累加进结算表的 HostTotalIncomeCutBefore字段中',
  `server_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '运营号备注',
  `create_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '1451577600' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_userId_week` (`week`,`user_id`) USING BTREE COMMENT '聚合唯一约束',
  KEY `idx_userId` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1542 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
