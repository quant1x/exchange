# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [0.5.11] - 2024-12-27
### Changed
- 新增2025年交易日历

## [0.5.10] - 2024-08-06
### Changed
- 更新依赖库num版本到0.3.6
- update changelog

## [0.5.9] - 2024-07-05
### Changed
- 更新依赖库num版本到0.3.5

## [0.5.8] - 2024-06-20
### Changed
- 更新依赖库num版本到0.3.4
- update changelog

## [0.5.7] - 2024-06-14
### Changed
- 更新依赖库版本
- update changelog

## [0.5.6] - 2024-05-30
### Changed
- 修复计算分钟数的bug
- 新增chglog配置文件
- update changelog
- update changelog

## [0.5.5] - 2024-05-20
### Changed
- 调整私有变量名为驼峰命名规则
- 时间常量增加交易开始时间
- update changelog

## [0.5.4] - 2024-05-16
### Changed
- 更新num版本到0.3.2
- update changelog

## [0.5.3] - 2024-05-11
### Changed
- 更新num版本到0.3.1
- update changelog

## [0.5.2] - 2024-04-16
### Changed
- 删除数据集长度对比多余的表达式
- 更新num版本到0.2.9
- update changelog

## [0.5.1] - 2024-04-16
### Changed
- 指数列表增加沪深300
- 指数列表增加科创50
- 更新依赖库num版本到0.2.7
- 修订DateRange入参end为非交易日时会返回超过end的交易日的bug
- 新增计算评估收益率的函数
- update changelog

## [0.5.0] - 2024-04-12
### Changed
- 新增判断证券代码的类型, 除去北交所代码外的指数,板块,ETF以及个股
- update changelog

## [0.4.8] - 2024-04-10
### Changed
- 新增数据两个日期之间的所有交易日
- 简化DateRange调用方法, 抽象出transactionDateRange函数
- 更新依赖库版本
- update changelog

## [0.4.7] - 2024-03-30
### Changed
- 更新依赖库版本
- update changelog

## [0.4.6] - 2024-03-21
### Changed
- 更新依赖库版本
- update changelog

## [0.4.5] - 2024-03-18
### Changed
- 调整尾盘集合竞价数据的结束时间,给快照留30s的更新收盘数据的时间buffer
- update changelog

## [0.4.4] - 2024-03-17
### Changed
- 更新依赖库版本
- update changelog

## [0.4.3] - 2024-03-12
### Changed
- 更新依赖库版本及go版本
- update changelog

## [0.4.2] - 2024-03-12
### Changed
- 更新依赖库版本
- update changelog

## [0.4.1] - 2024-03-11
### Changed
- 更新依赖库版本
- update changelog

## [0.4.0] - 2024-02-28
### Changed
- 更新依赖库版本
- update changelog

## [0.3.9] - 2024-02-28
### Changed
- 更新依赖库版本
- 更新依赖库版本
- update changelog

## [0.3.8] - 2024-02-12
### Changed
- 调整变量名
- 新增操作接口
- 定义一个未实现的获取Operator接口实例的函数
- !4 添加根据timestamp获取TimeKind的函数
* 添加根据timestamp获取TimeKind的函数
- 更新依赖库版本
- 从engine迁移成交数据相关时间常量
- 更新依赖库版本
- update changelog

## [0.3.7] - 2024-01-27
### Changed
- 更新依赖库版本
- update changelog

## [0.3.6] - 2024-01-25
### Changed
- 更新依赖库版本
- update changelog

## [0.3.5] - 2024-01-25
### Changed
- 更新依赖库版本
- update changelog

## [0.3.4] - 2024-01-25
### Changed
- 更新依赖库版本
- update changelog

## [0.3.3] - 2024-01-24
### Changed
- 更新依赖库版本
- update changelog

## [0.3.2] - 2024-01-23
### Changed
- 更新依赖库版本
- update changelog

## [0.3.1] - 2024-01-23
### Changed
- 更新依赖库版本
- update changelog

## [0.3.0] - 2024-01-23
### Changed
- 强化panic日志
- update changelog

## [0.2.9] - 2024-01-23
### Changed
- 迁移cache的map和pool到gox
- update changelog

## [0.2.8] - 2024-01-22
### Changed
- 更新gox版本
- update changelog

## [0.2.7] - 2024-01-22
### Changed
- update changelog
- 更新gox版本, 修复RollingOnce死锁的bug
- update changelog

## [0.2.6] - 2024-01-22
### Changed
- 更新gox版本
- 更新gox版本, 完善RollingOnce

## [0.2.5] - 2024-01-19
### Changed
- 修订涨停价格函数名拼写错误的bug
- update changelog

## [0.2.4] - 2024-01-19
### Changed
- 增加约定的指数和板块列表
- 新增通达信协议日期为YYYYMMDD格式的十进制整型的功能函数
- update changelog

## [0.2.3] - 2024-01-18
### Changed
- Touch函数迁移到gox
- update changelog

## [0.2.2] - 2024-01-18
### Changed
- 新增Touch函数
- update changelog

## [0.2.1] - 2024-01-17
### Changed
- 新增泛型的sync.Map和sync.Pool
- update changelog

## [0.2.0] - 2024-01-15
### Changed
- 新增是否收盘的函数
- 新增是否收盘的函数
- update changelog

## [0.1.9] - 2024-01-14
### Changed
- 增加校验交易日期范围, 开始日期不能晚于结束日期
- update changelog

## [0.1.8] - 2024-01-13
### Changed
- 更新依赖库版本号
- update changelog

## [0.1.7] - 2024-01-13
### Changed
- 更新依赖库版本号
- update changelog

## [0.1.6] - 2024-01-13
### Changed
- update changelog
- 修复结束日期早于开始日期的异常
- update changelog

## [0.1.5] - 2024-01-12
### Changed
- 合并exchange和market

## [0.1.4] - 2024-01-12
### Changed
- 新增一个交易日范围的函数
- update changelog

## [0.1.3] - 2024-01-11
### Changed
- 更新依赖库版本
- 从gotdx拆分和协议无关的代码
- 调整包路径
- update changelog

## [0.1.2] - 2024-01-09
### Changed
- 时间范围增加交易类型
- 调整计算分钟数的方法
- 优化获取毫秒数的函数
- 更新gox版本
- 优化时段的判断
- update changelog

## [0.1.1] - 2024-01-02
### Changed
- 新增: 值Range功能
- update changelog

## [0.1.0] - 2024-01-02
### Changed
- Initial commit
- 初始化exchange模块


[Unreleased]: https://gitee.com/quant1x/exchange.git/compare/v0.5.11...HEAD
[0.5.11]: https://gitee.com/quant1x/exchange.git/compare/v0.5.10...v0.5.11
[0.5.10]: https://gitee.com/quant1x/exchange.git/compare/v0.5.9...v0.5.10
[0.5.9]: https://gitee.com/quant1x/exchange.git/compare/v0.5.8...v0.5.9
[0.5.8]: https://gitee.com/quant1x/exchange.git/compare/v0.5.7...v0.5.8
[0.5.7]: https://gitee.com/quant1x/exchange.git/compare/v0.5.6...v0.5.7
[0.5.6]: https://gitee.com/quant1x/exchange.git/compare/v0.5.5...v0.5.6
[0.5.5]: https://gitee.com/quant1x/exchange.git/compare/v0.5.4...v0.5.5
[0.5.4]: https://gitee.com/quant1x/exchange.git/compare/v0.5.3...v0.5.4
[0.5.3]: https://gitee.com/quant1x/exchange.git/compare/v0.5.2...v0.5.3
[0.5.2]: https://gitee.com/quant1x/exchange.git/compare/v0.5.1...v0.5.2
[0.5.1]: https://gitee.com/quant1x/exchange.git/compare/v0.5.0...v0.5.1
[0.5.0]: https://gitee.com/quant1x/exchange.git/compare/v0.4.8...v0.5.0
[0.4.8]: https://gitee.com/quant1x/exchange.git/compare/v0.4.7...v0.4.8
[0.4.7]: https://gitee.com/quant1x/exchange.git/compare/v0.4.6...v0.4.7
[0.4.6]: https://gitee.com/quant1x/exchange.git/compare/v0.4.5...v0.4.6
[0.4.5]: https://gitee.com/quant1x/exchange.git/compare/v0.4.4...v0.4.5
[0.4.4]: https://gitee.com/quant1x/exchange.git/compare/v0.4.3...v0.4.4
[0.4.3]: https://gitee.com/quant1x/exchange.git/compare/v0.4.2...v0.4.3
[0.4.2]: https://gitee.com/quant1x/exchange.git/compare/v0.4.1...v0.4.2
[0.4.1]: https://gitee.com/quant1x/exchange.git/compare/v0.4.0...v0.4.1
[0.4.0]: https://gitee.com/quant1x/exchange.git/compare/v0.3.9...v0.4.0
[0.3.9]: https://gitee.com/quant1x/exchange.git/compare/v0.3.8...v0.3.9
[0.3.8]: https://gitee.com/quant1x/exchange.git/compare/v0.3.7...v0.3.8
[0.3.7]: https://gitee.com/quant1x/exchange.git/compare/v0.3.6...v0.3.7
[0.3.6]: https://gitee.com/quant1x/exchange.git/compare/v0.3.5...v0.3.6
[0.3.5]: https://gitee.com/quant1x/exchange.git/compare/v0.3.4...v0.3.5
[0.3.4]: https://gitee.com/quant1x/exchange.git/compare/v0.3.3...v0.3.4
[0.3.3]: https://gitee.com/quant1x/exchange.git/compare/v0.3.2...v0.3.3
[0.3.2]: https://gitee.com/quant1x/exchange.git/compare/v0.3.1...v0.3.2
[0.3.1]: https://gitee.com/quant1x/exchange.git/compare/v0.3.0...v0.3.1
[0.3.0]: https://gitee.com/quant1x/exchange.git/compare/v0.2.9...v0.3.0
[0.2.9]: https://gitee.com/quant1x/exchange.git/compare/v0.2.8...v0.2.9
[0.2.8]: https://gitee.com/quant1x/exchange.git/compare/v0.2.7...v0.2.8
[0.2.7]: https://gitee.com/quant1x/exchange.git/compare/v0.2.6...v0.2.7
[0.2.6]: https://gitee.com/quant1x/exchange.git/compare/v0.2.5...v0.2.6
[0.2.5]: https://gitee.com/quant1x/exchange.git/compare/v0.2.4...v0.2.5
[0.2.4]: https://gitee.com/quant1x/exchange.git/compare/v0.2.3...v0.2.4
[0.2.3]: https://gitee.com/quant1x/exchange.git/compare/v0.2.2...v0.2.3
[0.2.2]: https://gitee.com/quant1x/exchange.git/compare/v0.2.1...v0.2.2
[0.2.1]: https://gitee.com/quant1x/exchange.git/compare/v0.2.0...v0.2.1
[0.2.0]: https://gitee.com/quant1x/exchange.git/compare/v0.1.9...v0.2.0
[0.1.9]: https://gitee.com/quant1x/exchange.git/compare/v0.1.8...v0.1.9
[0.1.8]: https://gitee.com/quant1x/exchange.git/compare/v0.1.7...v0.1.8
[0.1.7]: https://gitee.com/quant1x/exchange.git/compare/v0.1.6...v0.1.7
[0.1.6]: https://gitee.com/quant1x/exchange.git/compare/v0.1.5...v0.1.6
[0.1.5]: https://gitee.com/quant1x/exchange.git/compare/v0.1.4...v0.1.5
[0.1.4]: https://gitee.com/quant1x/exchange.git/compare/v0.1.3...v0.1.4
[0.1.3]: https://gitee.com/quant1x/exchange.git/compare/v0.1.2...v0.1.3
[0.1.2]: https://gitee.com/quant1x/exchange.git/compare/v0.1.1...v0.1.2
[0.1.1]: https://gitee.com/quant1x/exchange.git/compare/v0.1.0...v0.1.1

[0.1.0]: https://gitee.com/quant1x/exchange.git/releases/tag/v0.1.0
