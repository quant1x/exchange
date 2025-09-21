package exchange

type TradingDay string

func NewTradeDate(date string) TradingDay {
	date = FixTradeDate(date)
	return TradingDay(date)
}

func (t TradingDay) String() string {
	return string(t)
}
