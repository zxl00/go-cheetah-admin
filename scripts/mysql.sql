INSERT INTO `go-easy-admin`.sys_menu (created_at,updated_at,deleted_at,create_by,name,name_code,is_show,icon,`path`,sort,parent_id,component,menu_type) VALUES
('2025-04-21 11:33:44.587','2025-08-04 15:07:55.984',NULL,'admin','首页','home',1,'house','/welcome',1,0,'Welcome',2),
('2025-04-21 11:43:08.876','2025-08-04 15:14:22.232',NULL,'admin','系统管理','sys',1,'Setting','/sys',2,0,'Sys',2),
('2025-04-21 14:27:19.709','2025-08-01 14:07:08.846',NULL,'admin','用户管理','user',1,'user','/user',1,2,'User',2),
('2025-04-21 14:28:04.319','2025-04-21 14:28:04.319',NULL,'admin','菜单管理','menu',1,'menu','/menu',2,2,'Menu',2),
('2025-04-21 14:28:27.591','2025-04-21 14:28:27.591',NULL,'admin','角色管理','role',1,'crop','/role',3,2,'Role',2),
('2025-05-08 10:40:45.073','2025-08-02 14:24:35.490',NULL,'admin','接口管理','',1,'Paperclip','/api',4,2,'Api',2),
('2025-08-01 13:38:43.498','2025-08-01 17:00:49.305',NULL,'admin','userEdit','',1,'','',0,3,'',1),
('2025-08-01 13:39:44.348','2025-08-01 13:39:44.348',NULL,'admin','userCreate','',1,'','',0,3,'',1),
('2025-08-01 13:49:23.200','2025-08-01 13:49:23.200',NULL,'admin','userDelete','',1,'','',0,3,'',1),
('2025-08-01 17:10:47.738','2025-08-01 17:10:47.738',NULL,'admin','menuEdit','',1,'','',0,5,'',1);
INSERT INTO `go-easy-admin`.sys_menu (created_at,updated_at,deleted_at,create_by,name,name_code,is_show,icon,`path`,sort,parent_id,component,menu_type) VALUES
('2025-08-01 17:11:04.946','2025-08-01 17:11:04.946',NULL,'admin','menuCreate','',1,'','',0,5,'',1),
('2025-08-01 17:11:19.752','2025-08-01 17:11:19.752',NULL,'admin','menuDelete','',1,'','',0,5,'',1),
('2025-08-01 17:25:47.492','2025-08-01 17:26:24.826',NULL,'admin','menuStatus','',1,'','',0,5,'',1),
('2025-08-01 17:26:51.075','2025-08-01 17:27:13.134',NULL,'admin','menuSort','',1,'','',0,5,'',1),
('2025-08-01 17:31:55.862','2025-08-01 17:31:55.862',NULL,'admin','roleCreate','',1,'','',0,6,'',1),
('2025-08-01 17:32:09.260','2025-08-01 17:32:09.260',NULL,'admin','roleEdit','',1,'','',0,6,'',1),
('2025-08-01 17:32:23.799','2025-08-01 17:38:35.999',NULL,'admin','roleDelete','',2,'','',0,6,'',1),
('2025-08-01 17:32:40.848','2025-08-01 17:32:40.848',NULL,'admin','roleMenu','',1,'','',0,6,'',1),
('2025-08-01 17:32:50.672','2025-08-01 17:32:50.672',NULL,'admin','roleApi','',1,'','',0,6,'',1),
('2025-08-01 17:52:39.656','2025-08-01 17:52:39.656',NULL,'admin','apiEdit','',1,'','',0,18,'',1);
INSERT INTO `go-easy-admin`.sys_menu (created_at,updated_at,deleted_at,create_by,name,name_code,is_show,icon,`path`,sort,parent_id,component,menu_type) VALUES
('2025-08-01 17:53:07.032','2025-08-01 17:53:07.032',NULL,'admin','apiDelete','',1,'','',0,18,'',1),
('2025-08-04 15:15:06.087','2025-08-04 15:18:22.204',NULL,'admin','LDAP配置','',1,'Tools','/ldap',5,2,'Ldap',2),
('2025-08-01 17:52:52.998','2025-08-01 17:52:52.998',NULL,'admin','apiCreate','',1,'','',0,18,'',1);

INSERT INTO `go-easy-admin`.sys_role (created_at,updated_at,deleted_at,create_by,name,`desc`,status) VALUES
('2025-04-21 14:37:25.470','2025-08-01 17:56:25.135',NULL,'admin','admin','默认管理员',1);

INSERT INTO `go-easy-admin`.sys_user (created_at,updated_at,deleted_at,create_by,username,password,nick_name,avatar,email,phone,status) VALUES
('2025-04-18 19:36:08','2025-07-30 17:54:38.006',NULL,NULL,'admin','25285442ebc7d3a0c20047e01d341c31','','xxx',NULL,NULL,1);

INSERT INTO `go-easy-admin`.system_apis (created_at,updated_at,deleted_at,create_by,`path`,`method`,`desc`,api_group) VALUES
('2025-05-08 11:21:21.264','2025-07-31 15:15:52.716',NULL,'admin','/sys/user/create','POST','创建用户','用户管理'),
('2025-05-08 11:22:19.791','2025-07-31 15:16:06.488',NULL,'admin','/sys/user/delete/:id','POST','删除用户','用户管理'),
('2025-05-08 11:26:48.554','2025-07-31 15:16:20.391',NULL,'admin','/sys/user/update/:id','POST','更新用户','用户管理'),
('2025-05-08 11:40:32.108','2025-07-31 15:16:30.785',NULL,'admin','/sys/user/list','GET','用户列表','用户管理'),
('2025-05-08 11:43:11.724','2025-07-31 15:16:41.586',NULL,'admin','/sys/user/get/:id','GET','用户详情','用户管理'),
('2025-05-08 13:47:50.288','2025-07-31 15:16:56.271',NULL,'admin','/sys/menu/create','POST','创建菜单','菜单管理'),
('2025-07-31 16:16:21.637','2025-07-31 16:16:21.637',NULL,'admin','/sys/menu/delete/:id','POST','删除菜单','菜单管理'),
('2025-07-31 16:16:41.611','2025-07-31 16:16:41.611',NULL,'admin','/sys/menu/update/:id','POST','更新菜单','菜单管理'),
('2025-07-31 16:17:05.409','2025-07-31 16:17:05.409',NULL,'admin','/sys/menu/list','GET','菜单列表','菜单管理'),
('2025-07-31 16:17:22.405','2025-07-31 16:17:22.405',NULL,'admin','/sys/menu/get/:id','GET','菜单详情','菜单管理');
INSERT INTO `go-easy-admin`.system_apis (created_at,updated_at,deleted_at,create_by,`path`,`method`,`desc`,api_group) VALUES
('2025-07-31 16:17:45.147','2025-07-31 16:17:45.147',NULL,'admin','/sys/role/create','POST','创建角色','角色管理'),
('2025-07-31 16:18:06.224','2025-07-31 16:18:06.224',NULL,'admin','/sys/role/delete/:id','POST','删除角色','角色管理'),
('2025-07-31 16:18:23.203','2025-07-31 16:18:23.203',NULL,'admin','/sys/role/update/:id','POST','更新角色','角色管理'),
('2025-07-31 16:18:41.561','2025-07-31 16:18:41.561',NULL,'admin','/sys/role/list','GET','角色列表','角色管理'),
('2025-07-31 16:19:01.499','2025-07-31 16:19:01.499',NULL,'admin','/sys/role/get/:id','GET','角色详情','角色管理'),
('2025-07-31 16:19:56.460','2025-07-31 16:19:56.460',NULL,'admin','/sys/apis/create','POST','创建接口','接口管理'),
('2025-07-31 16:20:11.095','2025-07-31 16:20:11.095',NULL,'admin','/sys/apis/delete/:id','POST','删除接口','接口管理'),
('2025-07-31 16:20:25.594','2025-07-31 16:20:25.594',NULL,'admin','/sys/apis/update/:id','POST','更新接口','接口管理'),
('2025-07-31 16:20:50.193','2025-07-31 16:20:50.193',NULL,'admin','/sys/apis/list','GET','接口列表','接口管理'),
('2025-07-31 16:21:05.894','2025-07-31 16:21:05.894',NULL,'admin','/sys/apis/get/:id','GET','接口详情','接口管理');
INSERT INTO `go-easy-admin`.system_apis (created_at,updated_at,deleted_at,create_by,`path`,`method`,`desc`,api_group) VALUES
('2025-07-31 16:21:39.854','2025-07-31 16:21:39.854',NULL,'admin','/sys/apis/get/group','GET','接口归属','接口管理'),
('2025-07-31 16:30:00.502','2025-07-31 16:30:00.502',NULL,'admin','/sys/rbac/create/:id','POST','接口授权','角色管理'),
('2025-07-31 16:30:24.850','2025-07-31 16:30:24.850',NULL,'admin','/sys/rbac/role/get/:id','GET','授权详情','角色管理');
