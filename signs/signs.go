package signs

// GetSign Get a Sign
func GetSign(month, day int) string {
	//var m = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	var signs = []string{"魔羯", "水瓶", "双鱼", "白羊", "金牛", "双子", "巨蟹", "狮子", "处女", "天秤", "天蝎", "射手"}
	var days = []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	if day < days[month-1] {
		return signs[month-1]
	}
	return signs[month]
}
