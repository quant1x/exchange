package exchange

import "sync"

const (
	MarketShangHai string = "sh" // 上海
	MarketShenZhen string = "sz" // 深圳
	MarketBeiJing  string = "bj" // 北京
	MarketHongKong string = "hk" // 香港
	MarketUSA      string = "us" // 美国
)

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
		ExchangeTime(CallAuctionAndCancel, "09:15:00", "09:20:00"),
		ExchangeTime(CallAuction, "09:20:00", "09:25:00"),
		ExchangeTime(TradingAndCancel, "09:30:00", "11:30:00"),
		ExchangeTime(TradingAndCancel, "13:00:00", "14:57:00"),
		ExchangeTime(CallAuction, "14:57:00", "15:00:00"),
	)
	mapExchange[CN] = cn
}

func GetExchange(name string) TradingSession {
	onceExchange.Do(lazyLoadExchanges)
	return mapExchange[name]
}
