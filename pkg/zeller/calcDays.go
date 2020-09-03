package zeller

func calcDays(m , y, d int)int{
	if m < 3  {
		y--
		m += 12
	}
	return ( y + y/4 - y/100 + y/400 + ( 13*m + 8 )/5 + d )%7
}

func DerivationDays(m, y, d int)string{
	days := []string{"日曜日","月曜日","火曜日","水曜日","木曜日","金曜日","土曜日"}
	return days[calcDays(m, y, d)]
}