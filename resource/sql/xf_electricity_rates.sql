-- Desc: 电费计价峰谷表
create table dbo.xf_electricity_rates
(
    id          bigint identity
        constraint xf_electricity_rates_pk
            primary key,
    period_at   varchar(6)                 not null,
    valley_rate decimal(10, 2)             not null,
    peak_rate   decimal(10, 2)             not null,
    created_at  datetime default getdate() not null,
    updated_at  datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'电费计价峰谷', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates'
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'周期月度表', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates', 'COLUMN',
     'period_at'
go

exec sp_addextendedproperty 'MS_Description', N'谷值率', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates', 'COLUMN',
     'valley_rate'
go

exec sp_addextendedproperty 'MS_Description', N'峰值率', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates', 'COLUMN',
     'peak_rate'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates', 'COLUMN',
     'created_at'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'dbo', 'TABLE', 'xf_electricity_rates', 'COLUMN',
     'updated_at'
go

