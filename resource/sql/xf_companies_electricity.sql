-- Desc: 企业电费计费表
create table xf_companies_electricity
(
    ceuuid                  varchar(36)                not null,
    cods                    varchar(18)                not null
        constraint xf_companies_electricity_xf_companies_cods_fk
            references xf_companies,
    period_at               varchar(6)                 not null,
    valley_electricity      decimal(10, 2)             not null,
    valley_electricity_bill decimal(10, 2)             not null,
    peak_electricity        decimal(10, 2)             not null,
    peak_electricity_bill   decimal(10, 2)             not null,
    total_electricity       decimal(10, 2)             not null,
    total_bill              decimal(10, 2)             not null,
    created_at              datetime default getdate() not null,
    updated_at              datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'企业电费计费表', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity'
go

exec sp_addextendedproperty 'MS_Description', N'企业电费主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity',
     'COLUMN', 'ceuuid'
go

exec sp_addextendedproperty 'MS_Description', N'计费月份周期', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity',
     'COLUMN', 'period_at'
go

exec sp_addextendedproperty 'MS_Description', N'谷用电', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity', 'COLUMN',
     'valley_electricity'
go

exec sp_addextendedproperty 'MS_Description', N'谷电价', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity', 'COLUMN',
     'valley_electricity_bill'
go

exec sp_addextendedproperty 'MS_Description', N'峰用电', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity', 'COLUMN',
     'peak_electricity'
go

exec sp_addextendedproperty 'MS_Description', N'峰电价', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity', 'COLUMN',
     'peak_electricity_bill'
go

exec sp_addextendedproperty 'MS_Description', N'合计电量', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity',
     'COLUMN', 'total_electricity'
go

exec sp_addextendedproperty 'MS_Description', N'合计电费', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity',
     'COLUMN', 'total_bill'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity',
     'COLUMN', 'created_at'
go

exec sp_addextendedproperty 'MS_Description', N'更新时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_companies_electricity',
     'COLUMN', 'updated_at'
go

create unique index xf_companies_electricity_ceuuid_uindex
    on xf_companies_electricity (ceuuid)
go

