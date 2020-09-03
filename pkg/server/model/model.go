package model

type User struct {
	Name	string `json:"name"`
}

// リクエスト =>
// {
//   "year": Int,
//   "month": Int,
//   "day": Int,
// }

type ZellerElements struct {
	Year	int	`json:"year"`
	Month	int	`json:"month"`
	Day		int	`json:"day"`
}

