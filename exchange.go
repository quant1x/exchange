package exchange

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/num"
	"strings"
	"sync"
)

//const (
//	MarketShangHai string = "sh" // 上海
//	MarketShenZhen string = "sz" // 深圳
//	MarketBeiJing  string = "bj" // 北京
//	MarketHongKong string = "hk" // 香港
//	MarketUSA      string = "us" // 美国
//)

const (
	CN = "cn" // A股
	HK = "hk" // 港股
	US = "us" // 美股
)

type exchanges struct {
	sessions map[string]TradingSession `name:"交易时段" yaml:"sessions"`
}

var (
	onceExchange sync.Once
	mapExchange  = map[string]TradingSession{}
)

// 加载配置文件
func lazyLoadExchanges() {
	cn := ExchangeSessions(
		ExchangeTime(CallAuctionAndCancel, "09:15:00", "09:20:00"), // 竞价, 可撤单
		ExchangeTime(CallAuction, "09:20:00", "09:25:00"),          // 竞价, 不可撤单
		ExchangeTime(TradingAndCancel, "09:30:00", "11:30:00"),     // 交易, 可撤单
		ExchangeTime(TradingAndCancel, "13:00:00", "14:57:00"),     // 交易, 可撤单
		ExchangeTime(CallAuction|Trading, "14:57:00", "15:00:00"),  // 竞价, 交易, 不可撤单
	)
	mapExchange[CN] = cn
}

func GetExchange(name string) TradingSession {
	onceExchange.Do(lazyLoadExchanges)
	return mapExchange[name]
}

type MarketType = uint8

const (
	MarketIdShenZhen MarketType = iota // 深圳
	MarketIdShangHai MarketType = 1    // 上海
	MarketIdBeiJing  MarketType = 2    // 北京
	MarketIdHongKong MarketType = 21   // 香港
	MarketIdUSA      MarketType = 22   // 美国

	StockDelisting = "DELISTING" // 退市
)

const (
	MarketShangHai string = "sh" // 上海
	MarketShenZhen string = "sz" // 深圳
	MarketBeiJing  string = "bj" // 北京
	MarketHongKong string = "hk" // 香港
	MarketUSA      string = "us" // 美国
)

const (
	MARKET_CN_FIRST_DATE     = "19901219"   // 上证指数的第一个交易日
	MARKET_CH_FIRST_LISTTIME = "1990-12-19" // 个股上市日期
)

var (
	// 全部市场简写
	marketFlags = []string{"sh", "sz", "SH", "SZ", "bj", "BJ", "hk", "HK", "us", "US"}
	// A股市场代码
	marketAShareFlags = []string{"sh", "sz", "SH", "SZ", "bj", "BJ"}
)

func GetSecurityCode(market MarketType, symbol string) (securityCode string) {
	switch market {
	case MarketIdUSA:
		return MarketUSA + symbol
	case MarketIdHongKong:
		return MarketHongKong + symbol[:5]
	case MarketIdBeiJing:
		return MarketBeiJing + symbol[:6]
	case MarketIdShenZhen:
		return MarketShenZhen + symbol[:6]
	default:
		return MarketShangHai + symbol[:6]
	}
}

// GetMarket 判断股票ID对应的证券市场匹配规则
//
//	['50', '51', '60', '90', '110'] 为 sh
//	['00', '12'，'13', '18', '15', '16', '18', '20', '30', '39', '115'] 为 sz
//	['5', '6', '9'] 开头的为 sh， 其余为 sz
func GetMarket(symbol string) string {
	symbol = strings.TrimSpace(symbol)
	market := "sh"
	if api.StartsWith(symbol, marketFlags) {
		market = strings.ToLower(symbol[0:2])
	} else if api.EndsWith(symbol, marketFlags) {
		length := len(symbol)
		// 后缀一个点号+2位字母, 代码在最前面
		market = strings.ToLower(symbol[length-2:])
	} else if api.StartsWith(symbol, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		market = "sh"
	} else if api.StartsWith(symbol, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		market = "sz"
	} else if api.StartsWith(symbol, []string{"5", "6", "9", "7"}) {
		market = "sh"
	} else if api.StartsWith(symbol, []string{"88"}) {
		market = "sh"
	} else if api.StartsWith(symbol, []string{"4", "8"}) {
		market = "bj"
	}
	return market
}

// GetMarketId 获得市场ID
func GetMarketId(symbol string) uint8 {
	market := GetMarket(symbol)
	marketId := MarketIdShangHai
	if market == "sh" {
		marketId = MarketIdShangHai
	} else if market == "sz" {
		marketId = MarketIdShenZhen
	} else if market == "bj" {
		marketId = MarketIdBeiJing
	}
	return marketId
}

func GetMarketFlag(marketId MarketType) string {
	switch marketId {
	case MarketIdShenZhen:
		return MarketShenZhen
	case MarketIdBeiJing:
		return MarketBeiJing
	case MarketIdHongKong:
		return MarketHongKong
	case MarketIdUSA:
		return MarketUSA
	default:
		return MarketShangHai
	}
}

// DetectMarket 检测市场代码
func DetectMarket(symbol string) (marketId MarketType, market string, code string) {
	code = strings.TrimSpace(symbol)
	market = MarketShangHai
	if api.StartsWith(code, marketFlags) {
		// 前缀2位字母后面跟代码
		market = strings.ToLower(code[0:2])
		if code[2:3] == "." {
			// SZ.000002
			code = code[3:]
		} else {
			// SZ000002
			code = code[2:]
		}
	} else if api.EndsWith(code, marketFlags) {
		length := len(code)
		// 后缀一个点号+2位字母, 代码在最前面
		// 600000.SH
		market = strings.ToLower(code[length-2:])
		code = code[:length-3]
	} else if api.StartsWith(code, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		// 上海证券交易所
		// 主板: 60xxxx
		// 科创板: 688xxx
		// B股: 900xxx
		// 优先股: 360xxx
		// 科创板存托凭证: 689xxx
		// 申购/配股/投票: 7xxxxx
		// 上海总规则: http://www.sse.com.cn/lawandrules/guide/stock/jyglywznylc/zn/a/20230209/4ae280c58535e0424b3a9c743c47e6b9.docx
		// 0: 国债/指数, 000 上证指数系列和中证指数系列, 00068x科创板指数
		// 1: 债券
		// 2: 回购
		// 3: 期货
		// 4: 备用
		// 5: 基金/权证
		// 6: A股
		// 7: 非交易业务(发行, 权益分配)
		// 8: 备用, 通达信编制板块指数占用880,881
		// 9: B股
		market = MarketShangHai
	} else if api.StartsWith(code, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		// 深圳交易所
		// 主板: 000,001
		// 中小板: 002,003,004
		// 创业板: 30xxxx
		// 优先股: 140xxx
		// 深圳总规则: https://zhuanlan.zhihu.com/p/63064991
		// 0: 股票
		// 1: 国债/基金
		// 2: B股
		// 30: 创业板
		// 36: 投票, 369999用于深交所认证业务的密码激活/密码挂失
		// 37: 增发/可转债申购
		// 38: 配股/可转债优先权
		// 395: 成家量统计指数
		// 399: 指数
		market = MarketShenZhen
	} else if api.StartsWith(code, []string{"5", "6", "9", "7"}) {
		market = MarketShangHai
	} else if api.StartsWith(code, []string{"88"}) {
		// 通达信板块指数, 在上海交易所
		market = MarketShangHai
	} else if api.StartsWith(code, []string{"4", "8"}) {
		// 北京上市公司: 43, 83,87
		// 新三板: 40,43,83,87
		// 三板A: 400,430,830-839,870-873
		// 三板B: 420
		// 优先股: 820
		market = MarketBeiJing
	}
	marketId = MarketIdShangHai
	if market == MarketShangHai {
		marketId = MarketIdShangHai
	} else if market == MarketShenZhen {
		marketId = MarketIdShenZhen
	} else if market == MarketBeiJing {
		marketId = MarketIdBeiJing
	} else if market == MarketHongKong {
		marketId = MarketIdHongKong
	}
	return marketId, market, code
}

// AssertIndexByMarketAndCode 通过市场id和短码判断是否指数
func AssertIndexByMarketAndCode(marketId MarketType, symbol string) (isIndex bool) {
	if marketId == MarketIdShangHai && api.StartsWith(symbol, []string{"000", "880", "881"}) {
		return true
	} else if marketId == MarketIdShenZhen && api.StartsWith(symbol, []string{"399"}) {
		return true
	}
	return false
}

// AssertIndexBySecurityCode 通过证券代码判断是否指数
func AssertIndexBySecurityCode(securityCode string) (isIndex bool) {
	marketId, _, code := DetectMarket(securityCode)
	return AssertIndexByMarketAndCode(marketId, code)
}

// AssertBlockBySecurityCode 断言证券代码是否板块
func AssertBlockBySecurityCode(securityCode *string) (isBlock bool) {
	marketId, flag, code := DetectMarket(*securityCode)
	if marketId != MarketIdShangHai {
		// 板块指数的数据在上海
		return false
	}
	if !api.StartsWith(code, []string{"880", "881"}) {
		return false
	}
	*securityCode = flag + code
	return true
}

// AssertETFByMarketAndCode 通过市场id和代码判断是否ETF
func AssertETFByMarketAndCode(marketId MarketType, symbol string) (isETF bool) {
	if marketId == MarketIdShangHai && api.StartsWith(symbol, []string{"510"}) {
		return true
	}
	return false
}

// AssertStockByMarketAndCode 通过市场id和代码判断是否个股
func AssertStockByMarketAndCode(marketId MarketType, symbol string) (isStock bool) {
	if marketId == MarketIdShangHai && api.StartsWith(symbol, []string{"60", "68", "510"}) {
		return true
	} else if marketId == MarketIdShenZhen && api.StartsWith(symbol, []string{"00", "30"}) {
		return true
	}
	return false
}

// AssertStockBySecurityCode 通过证券代码判断是否个股
func AssertStockBySecurityCode(securityCode string) (isStock bool) {
	marketId, _, code := DetectMarket(securityCode)
	return AssertStockByMarketAndCode(marketId, code)
}

// MarketLimit 涨跌停板限制
func MarketLimit(securityCode string) float64 {
	_, flag, shortCode := DetectMarket(securityCode)
	if flag == MarketBeiJing {
		return 0.30
	}
	if api.StartsWith(shortCode, []string{"30", "68"}) {
		return 0.20
	}
	return 0.10
}

// LimitUp 返回涨停板价格
func LimitUp(securityCode string, price float64) float64 {
	limit := MarketLimit(securityCode)
	lastClose := num.Decimal(price)
	upStopPrice := num.Decimal(lastClose * (1.0000 + limit))
	return upStopPrice
}

// CorrectSecurityCode 修正证券代码
func CorrectSecurityCode(securityCode string) string {
	if len(securityCode) == 0 {
		return ""
	}
	_, mFlag, mSymbol := DetectMarket(securityCode)
	return mFlag + mSymbol
}
