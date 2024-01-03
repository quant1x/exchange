package exchange

import (
	"fmt"
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/pkg/yaml"
	"regexp"
	_ "runtime"
	"strings"
	"time"
	_ "unsafe"
)

// 值范围正则表达式
var (
	stringRangePattern = "[~-]\\s*"
	stringRangeRegexp  = regexp.MustCompile(stringRangePattern)
)

var (
	ErrTimeFormat     = exception.New(errnoConfig+2, "时间范围格式错误")
	formatOfTimestamp = time.TimeOnly
)

// TimeKind 时段类型
type TimeKind uint64

const (
	CallAuction TimeKind = 1 << iota // 集合竞价
	CanCancel                        // 可撤单
	Trading                          // 交易时段
)

const (
	CallAuctionAndCancel TimeKind = CallAuction | CanCancel // 集合竞价可撤单
	TradingAndCancel     TimeKind = Trading | CanCancel     // 交易可撤单
)

const (
	secondsPerMinute      = 60
	secondsPerHour        = 60 * secondsPerMinute
	secondsPerDay         = 24 * secondsPerHour
	milliSecondsPerSecond = 1000
)

//go:linkname walltime runtime.walltime
func walltime() (int64, int32)

// 调用公开结构的私有方法
//
//go:linkname abstime time.Time.abs
func abstime(t time.Time) uint64

var (
	// 获取偏移的秒数
	zoneName, offsetInSecondsEastOfUTC = time.Now().Zone()
)

// 获取当前的时间戳, 毫秒数
func timestamp() int64 {
	sec, nsec := walltime()
	sec += int64(offsetInSecondsEastOfUTC)
	milli := sec*milliSecondsPerSecond + int64(nsec)/1e6%milliSecondsPerSecond
	return milli
}

func getTradingTimestamp() string {
	now := time.Now()
	return now.Format(formatOfTimestamp)
}

// TimeRange 时间范围
type TimeRange struct {
	kind    TimeKind // 时间类型
	begin   string   // 开始时间
	end     string   // 结束时间
	tmBegin int64    // 开始的秒数
	tmEnd   int64    // 结束的秒数
}

func ExchangeTime(kind TimeKind, begin, end string) TimeRange {
	tmBegin, err := time.Parse(time.TimeOnly, begin)
	if err != nil {
		panic(err)
	}
	tmEnd, err := time.Parse(time.TimeOnly, end)
	if err != nil {
		panic(err)
	}
	tr := TimeRange{
		kind:    kind,
		begin:   begin,
		end:     end,
		tmBegin: tmBegin.Unix(),
		tmEnd:   tmEnd.Unix(),
	}
	if tr.begin > tr.end {
		tr.begin, tr.end = tr.end, tr.begin
		tr.tmBegin, tr.tmEnd = tr.tmEnd, tr.tmBegin
	}
	return tr
}

func (this TimeRange) Minutes() int {
	if (this.kind & Trading) == 0 {
		return 0
	}
	seconds := this.tmEnd - this.tmBegin
	minutes := seconds / 60
	remaining := seconds % 60
	if remaining > 0 {
		minutes++
	}
	return int(minutes)
}

func (this TimeRange) String() string {
	return fmt.Sprintf("{begin: %s, end: %s}", this.begin, this.end)
}

func (this TimeRange) v2String() string {
	return fmt.Sprintf("%s~%s", this.begin, this.end)
}

func (this *TimeRange) Parse(text string) error {
	text = strings.TrimSpace(text)
	arr := stringRangeRegexp.Split(text, -1)
	if len(arr) != 2 {
		return ErrTimeFormat
	}
	this.begin = strings.TrimSpace(arr[0])
	this.end = strings.TrimSpace(arr[1])
	if this.begin > this.end {
		this.begin, this.end = this.end, this.begin
	}
	return nil
}

// UnmarshalText 设置默认值调用
//
//	由于begin和end字段不可访问, 默认值调用实际无效
func (this *TimeRange) UnmarshalText(text []byte) error {
	_ = text
	//TODO implement me
	panic("implement me")
}

// UnmarshalYAML YAML自定义解析
func (this *TimeRange) UnmarshalYAML(node *yaml.Node) error {
	var key, value string
	if len(node.Content) == 0 {
		value = node.Value
	} else if len(node.Content) == 2 {
		key = node.Content[0].Value
		value = node.Content[1].Value
	}
	_ = key
	return this.Parse(value)
}

func (this *TimeRange) IsTrading(timestamp ...string) bool {
	var tm string
	if len(timestamp) > 0 {
		tm = strings.TrimSpace(timestamp[0])
	} else {
		tm = getTradingTimestamp()
	}
	if tm >= this.begin && tm <= this.end {
		return true
	}
	return false
}
