package data

import "fmt"

type Date int64

func NewDate(day int64) Date {
	return Date(day)
}

var monthName = []string{
	"芽吹き", "若葉", "陽盛り", "実り", "霜降り", "雪籠り",
}

var seasonName = []string{
	"春", "初夏", "夏", "秋", "初冬", "冬",
}

func (d Date) String() string {
	year := d / (24 * 6)
	dayOfYear := d % (24 * 6)
	month := dayOfYear / 24
	dayOfMonth := dayOfYear % 24

	return fmt.Sprintf(
		"王都暦%d年 %sの月 %d日 %s",
		year+568, monthName[month], dayOfMonth+1, seasonName[month],
	)
}
