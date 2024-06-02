# SmartPower - 智电管家

> 请注意：这个项目不是一个持续性的项目，而是一个数据库课程的作业项目。可以参考学习，但不建议用于生产环墅。
> 项目中的代码可能存在一些问题，不建议直接使用。若产生问题，开发者拒绝维护此项目。
> 
> —— 筱锋xiao_lfeng

智电管家是一款基于 `Golang(GoFrame)` 开发的用电管理系统，使用 `SQL Server` 数据库进行数据存储和管理。
该系统旨在为用电企业提供高效、便捷的用电管理解决方案，帮助企业合理规划用电成本，实现智能化管理。

![Alt](https://repobeats.axiom.co/api/embed/3aa7c589e1f6392ca653d14da6fff0fb028e4179.svg "Repobeats analytics image")

## 功能模块

### 1. 用电企业管理

记录并管理企业的基本信息，包括：
- 企业编号
- 企业名称
- 地址
- 电话
- 联系人信息

### 2. 电费信息管理

记录并管理电费的峰价和谷价信息，以便计算企业的电费支出。

### 3. 用电情况管理

记录企业的用电情况，包括：
- 每月谷电量
- 每月峰电量
- 年月
- 合计电量
- 合计电费

## 系统主要功能要求

### 1. 统计月耗电量及电费

系统能够统计并显示各个用电企业的每月耗电量和电费情况，方便企业了解自身的用电成本。

### 2. 统计谷电量和峰电量

系统能够统计并查询各个企业的总谷电量和总峰电量，帮助企业分析用电模式，优化用电策略。

### 3. 统计区域峰、谷电量及电费比例

系统能够统计并显示该地区的总峰电量和总谷电量，以及相应的电费比例，提供区域用电的综合分析。

### 4. 基础数据修改

系统支持修改企业基本信息和电费信息等基础数据，确保数据的及时更新和准确性。

### 5. 用户分级管理

系统支持用户分级管理，提供良好的人机交互界面。不同级别的用户具有不同的操作权限，以保证系统的安全性和数据的保密性。

## 系统优势

智电管家通过全面的用电管理功能，为企业提供以下优势：
- 提高用电管理的效率
- 帮助企业合理规划用电成本
- 实现智能化、精细化管理

## 技术栈

- **后端**：Golang（使用GoFrame框架）
- **数据库**：SqlServer
- **前端**：提供友好的用户界面，支持多种设备访问，使用 Vue3 进行开发

## 贡献

本项目不进行维护，拒绝贡献

## 许可证

本项目基于 MIT 许可证开源，详情请参见 [LICENSE](LICENSE) 文件。

## 联系我们

如果您有任何问题或建议，请通过以下方式与我们联系：

- 若您借用了本项目的内容或或应用于商业活动，请遵循许可的前提下联系我并标明出处，谢谢🙏！
- 邮箱：gm@x-lf.cn
- 个人博客：[凌中的锋雨](https://blog.x-lf.cn)