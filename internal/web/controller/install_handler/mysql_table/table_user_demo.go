package mysql_table

//CREATE TABLE `user_demo` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`user_name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
//`nick_name` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',
//`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户Demo表';

func CreateUserDemoTableSql() (sql string) {
	sql = "CREATE TABLE `user_demo` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`user_name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',"
	sql += "`nick_name` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',"
	sql += "`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "PRIMARY KEY (`id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户Demo表';"

	return
}
