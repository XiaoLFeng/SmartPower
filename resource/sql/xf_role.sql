-- Description: 角色表
create table xf_role
(
    ruuid        varchar(36)    not null
        constraint xf_role_pk
            primary key,
    name         varchar(30)    not null,
    display_name nvarchar(30)   not null,
    description  nvarchar(1024) not null,
    created_at   datetime       not null,
    updated_at   datetime       not null
)
go

exec sp_addextendedproperty 'MS_Description', N'角色表', 'SCHEMA', 'dbo', 'TABLE', 'xf_role'
go

exec sp_addextendedproperty 'MS_Description', N'角色表主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_role', 'COLUMN', 'ruuid'
go

exec sp_addextendedproperty 'MS_Description', N'角色名', 'SCHEMA', 'dbo', 'TABLE', 'xf_role', 'COLUMN', 'name'
go

exec sp_addextendedproperty 'MS_Description', N'展示名称', 'SCHEMA', 'dbo', 'TABLE', 'xf_role', 'COLUMN', 'display_name'
go

exec sp_addextendedproperty 'MS_Description', N'描述', 'SCHEMA', 'dbo', 'TABLE', 'xf_role', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_role', 'COLUMN', 'created_at'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_role', 'COLUMN', 'updated_at'
go

create unique index xf_role_name_uindex
    on xf_role (name)
go

