CREATE TABLE `user_auth` (
                               `uid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                               `user_name` varchar(20) not null unique COMMENT '用户名',
                               `password` varchar(65) not null COMMENT '密码',
                               `follow_count` bigint default 0 COMMENT '关注人数',
                               `follower_count` varchar(20) default 0  COMMENT '粉丝总数',
                               `is_follow` boolean default false COMMENT '是否关注，true-已关注，false-未关注',
                               `avatar` varchar(255) default '' COMMENT '用户头像',
                               `background_image` varchar(255) default ''  COMMENT '用户个人页顶部大图',
                               `signature` varchar(255) default '' COMMENT '个人简介',
                               `total_favorite` varchar(255) default '' COMMENT '获赞数量',
                               `work_count` bigint default 0 COMMENT '作品数',
                               `favorite_count` bigint default 0 COMMENT '喜欢数',
                               `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                               `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户鉴权表';



CREATE TABLE `video` (
                               `vid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                               `uid` bigint not null COMMENT '作者',
                               `title` varchar(255) not null COMMENT '视频标题',
                               `play_url` varchar(255) not null COMMENT '播放地址',
                               `cover_url` varchar(255) not null COMMENT '封面地址',
                               `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                               `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`vid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE `comment` (
                           `cid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                           `vid` bigint not null COMMENT '视频id',
                           `uid` bigint not null COMMENT '用户id',
                           `content` varchar(255) not null COMMENT '评论内容，不能超过255字符',
                           `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

CREATE TABLE `favorite` (
                           `lid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                           `vid` bigint not null COMMENT '视频id',
                           `uid` bigint not null COMMENT '用户id',
                           `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                           `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`lid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='喜欢表';