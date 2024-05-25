-- Description: 用电企业表
create table dbo.xf_companies
(
    cods           nvarchar(18)               not null
        constraint xf_companies_pk
            primary key,
    name           nvarchar(40)               not null,
    address        nvarchar(1024)             not null,
    representative nvarchar(100)              not null,
    uuid           nvarchar(36)               not null
        constraint xf_companies_xf_user_uuid_fk
            references dbo.xf_user
            on update cascade,
    created_at     datetime default getdate() not null,
    updated_at     datetime
)
    go

exec sp_addextendedproperty 'MS_Description', N'用电企业表', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies'
go

exec sp_addextendedproperty 'MS_Description', N'企业唯一信用社会代码', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies',
     'COLUMN', 'cods'
go

exec sp_addextendedproperty 'MS_Description', N'企业名字', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies', 'COLUMN', 'name'
go

exec sp_addextendedproperty 'MS_Description', N'企业所在地址', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies', 'COLUMN',
     'address'
go

exec sp_addextendedproperty 'MS_Description', N'法定代表人', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies', 'COLUMN',
     'representative'
go

exec sp_addextendedproperty 'MS_Description', N'指定用户', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies', 'COLUMN', 'uuid'
go

exec sp_addextendedproperty 'MS_Description', N'用户表 uuid 外键', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies',
     'CONSTRAINT', 'xf_companies_xf_user_uuid_fk'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies', 'COLUMN',
     'created_at'
go

exec sp_addextendedproperty 'MS_Description', N'更新时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies', 'COLUMN',
     'updated_at'
go

