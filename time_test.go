package exchange

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTimeRange_Minutes(t *testing.T) {
	tr := ExchangeTime(0, "09:30:00", "11:30:00")
	minutes := tr.Minutes()
	fmt.Println(minutes)
}

func Test_walltime(t *testing.T) {
	a, b := walltime()
	a += int64(offsetInSecondsEastOfUTC)
	fmt.Println(a, b)
	fmt.Println(a%secondsPerDay, b/1e6)
	now := time.Unix(a, int64(b))
	fmt.Println(now.Format(time.DateTime))

	now = time.Now().Local()
	tm := time.Until(now)
	fmt.Println("Milliseconds =>", tm.Milliseconds())
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	//endOfDay := startOfDay.AddDate(0, 0, 1)
	millis := int64(now.Sub(startOfDay).Milliseconds())
	fmt.Println("当天0点到现在的毫秒数:", millis)
	fmt.Println(now.Date())
	fmt.Println(millis)
	day := 24 * 60 * 60 * 1000
	fmt.Println(day)
	fmt.Println(now.UnixMilli() % int64(day))

	fmt.Println("s", now.Unix())
	millis = now.UnixNano() / 1e6
	fmt.Println(millis % int64(day))
	m := now.Hour() * 60 * 60 * 1000
	m += now.Minute() * 60 * 1000
	m += now.Second() * 1000
	//m += int(now.UnixMilli()) % 1e6
	fmt.Println("s =>", m)

	now = time.Now()
	log.Println("时间戳（秒）：", now.Unix())       // 输出：时间戳（秒） ： 1665807442
	log.Println("时间戳（毫秒）：", now.UnixMilli()) // 输出：时间戳（毫秒）： 1665807442207
	log.Println("时间戳（微秒）：", now.UnixMicro()) // 输出：时间戳（微秒）： 1665807442207974
	log.Println("时间戳（纳秒）：", now.UnixNano())  // 输出：时间戳（纳秒）： 1665807442207974500

	tm1 := abstime(now)
	fmt.Println(tm1)
	tm2 := tm1 / 1e9
	fmt.Println(tm2 % secondsPerDay)
}

func Test_abstime(t *testing.T) {
	now := time.Now()
	tm := abstime(now)
	fmt.Println(tm)
	tm1 := tm / 1e9
	fmt.Println(tm1 % secondsPerDay)
}

func Test_timestamp(t *testing.T) {
	a, b := walltime()
	//a += int64(offsetInSecondsEastOfUTC)
	fmt.Println(a, b)
	tm := timestamp()
	fmt.Println(tm, tm%(secondsPerDay*1000))
	fmt.Println(tm/milliSecondsPerSecond - a)
}
