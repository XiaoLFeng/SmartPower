-- Description: 系统表
create table xf_info
(
    system_id  varchar(36)    not null
        constraint xf_info_pk
            primary key,
    keyword      varchar(100)   not null,
    value      nvarchar(1024) not null,
    updated_at datetime      not null
)
go

exec sp_addextendedproperty 'MS_Description', N'系统表', 'SCHEMA', 'dbo', 'TABLE', 'xf_info'
go

exec sp_addextendedproperty 'MS_Description', N'系统主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_info', 'COLUMN', 'system_id'
go

exec sp_addextendedproperty 'MS_Description', N'键', 'SCHEMA', 'dbo', 'TABLE', 'xf_info', 'COLUMN', 'keyword'
go

exec sp_addextendedproperty 'MS_Description', N'值', 'SCHEMA', 'dbo', 'TABLE', 'xf_info', 'COLUMN', 'value'
go

exec sp_addextendedproperty 'MS_Description', N'更新时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_info', 'COLUMN', 'updated_at'
go

