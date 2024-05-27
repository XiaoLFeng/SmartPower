-- Description: 用户表
create table dbo.xf_user
(
    uuid       varchar(36)               not null
        constraint xf_user_pk
            primary key,
    username   varchar(30)               not null,
    email      nvarchar(50)               not null,
    phone      nvarchar(15)               not null,
    password   nvarchar(100)              not null,
    role       varchar(36)                not null
        constraint xf_user_fy_role_ruuid_fk
            references dbo.xf_role
            on update cascade,
    created_at datetime default getdate() not null,
    updated_at datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'用户表', 'SCHEMA', 'dbo', 'TABLE', 'xf_user'
go

exec sp_addextendedproperty 'MS_Description', N'用户 id', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'uuid'
go

exec sp_addextendedproperty 'MS_Description', N'用户名', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'username'
go

exec sp_addextendedproperty 'MS_Description', N'邮箱', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'email'
go

exec sp_addextendedproperty 'MS_Description', N'手机号', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'phone'
go

exec sp_addextendedproperty 'MS_Description', N'用户密码', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'password'
go

exec sp_addextendedproperty 'MS_Description', N'角色主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'role'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'created_at'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_user', 'COLUMN', 'updated_at'
go

