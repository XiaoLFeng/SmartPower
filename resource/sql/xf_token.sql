-- Description: 令牌表
create table xf_token
(
    tuuid      varchar(36)                not null
        constraint xf_token_pk
            primary key,
    uuid       varchar(36)               not null
        constraint xf_token_xf_user_uuid_fk
            references xf_user
            on update cascade on delete cascade,
    token      varchar(36)                not null,
    created_at datetime default getdate() not null,
    expired_at datetime                   not null
)
go

exec sp_addextendedproperty 'MS_Description', N'令牌主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_token', 'COLUMN', 'tuuid'
go

exec sp_addextendedproperty 'MS_Description', N'用户主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_token', 'COLUMN', 'uuid'
go

exec sp_addextendedproperty 'MS_Description', N'授权令牌', 'SCHEMA', 'dbo', 'TABLE', 'xf_token', 'COLUMN', 'token'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_token', 'COLUMN', 'created_at'
go

exec sp_addextendedproperty 'MS_Description', N'过期时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_token', 'COLUMN', 'expired_at'
go

create unique index xf_token_token_uindex
    on xf_token (token)
go

