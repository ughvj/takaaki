package types

type InspectedHistories []InspectedHistory
type InspectedHistory struct {
	NyutaikunUserId int
}

func (ih *InspectedHistory) Refs() []interface{} {
	return []interface{}{
		&ih.NyutaikunUserId,
	}
}

func GetStaticInspectedHistoryCase0() *InspectedHistories {
	ret := InspectedHistories{
	}
	return &ret
}

func GetStaticInspectedHistoryCase1() *InspectedHistories {
	ret := InspectedHistories{
		InspectedHistory{
			NyutaikunUserId: 100,
		},
	}
	return &ret
}

func GetStaticInspectedHistoryCase2() *InspectedHistories {
	ret := InspectedHistories{
		InspectedHistory{
			NyutaikunUserId: 200,
		},
		InspectedHistory{
			NyutaikunUserId: 300,
		},
	}
	return &ret
}
