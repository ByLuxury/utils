package utils
import "time"

// ConvertLocalTimeByTime 时间处理
func ConvertLocalTimeByTime(t time.Time) (string, error) {
	Bdate := t.Format(time.RFC3339)
	dateTime, err := time.Parse(time.RFC3339, Bdate)
	return dateTime.In(time.Local).Format("2006-01-02 15:04:05"), err
}

// ConvertLocalTimeByString ...
func ConvertLocalTimeByString(t string) (string, error) {
	dateTime, err := time.Parse("2006-01-02 15:04:05", t)
	return dateTime.In(time.Local).Format("2006-01-02 15:04:05"), err
}

