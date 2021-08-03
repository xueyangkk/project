#### go_gin_api.program 
计划

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(128) |  | YES |  |  |
| 3 | slogan | 标语 | varchar(128) |  | YES |  |  |
| 4 | cover_url | 封面图 | varchar(512) |  | YES |  |  |
| 5 | performer_count | 参与人数 | int(11) |  | YES |  |  |
| 6 | enroll_price_in_coin | 报名费用（早币） | int(11) |  | YES |  |  |
| 7 | created_at |  | datetime |  | YES |  |  |
| 8 | type | 计划类型   1肉体计划；2灵魂计划 | smallint(6) |  | YES |  |  |
| 9 | order | 顺序 | int(11) |  | YES |  |  |
| 10 | official | 是否官网创建 | tinyint(1) |  | YES |  |  |
| 11 | is_deleted |  | tinyint(1) |  | YES |  |  |
| 12 | tags |  | varchar(1000) |  | YES |  |  |
| 13 | marking | 默认为1    2 为已经标记的 | tinyint(4) |  | YES |  | 1 |
| 14 | updated_at |  | datetime |  | YES |  |  |
