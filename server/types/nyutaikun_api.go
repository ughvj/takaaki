package types

type NyutaikunEntranceAndExitResult struct {
	Success bool `json:"success"`
	Total int `json:"total"`
	Data []NyutaikunEntranceAndExit `json:"data"`
}

type NyutaikunEntranceAndExit struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	EntranceTime string `json:"entrance_time"`
	ExitTime string `json:"exit_time"`
}

func GetStaticNyutaikunEntranceAndExitResult() *NyutaikunEntranceAndExitResult {
	ret := NyutaikunEntranceAndExitResult {
		Success: true,
		Total: 2,
		Data: []NyutaikunEntranceAndExit{
			NyutaikunEntranceAndExit{
				Id: 1,
				UserId: 100,
				EntranceTime: "2025-02-27T08:00:00+09:00",
				ExitTime: "2025-02-27T09:00:00+09:00",
			},
			NyutaikunEntranceAndExit{
				Id: 2,
				UserId: 200,
				EntranceTime: "2025-02-27T09:00:00+09:00",
				ExitTime: "2025-02-27T10:00:00+09:00",
			},
			NyutaikunEntranceAndExit{
				Id: 3,
				UserId: 300,
				EntranceTime: "2025-02-27T10:00:00+09:00",
				ExitTime: "2025-02-27T11:00:00+09:00",
			},
		},
	}

	return &ret
}
