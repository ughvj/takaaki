package types

type Students []Student
type Student struct {
	Id int
	Name string
	NyutaikunUserId int
}

func (ss *Students) FindByNyutaikunUserId(id int) *Student {
	for _, s := range *ss {
		if s.NyutaikunUserId == id {
			return &s
		}
	}
	return nil
}

func (s *Student) Refs() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Name,
		&s.NyutaikunUserId,
	}
}

func GetStaticStudents() *Students {
	ret := Students {
		Student{
			Id: 1,
			Name: "伊藤　博文",
			NyutaikunUserId: 100,
		},
		Student{
			Id: 2,
			Name: "黒田　清隆",
			NyutaikunUserId: 200,
		},
		Student{
			Id: 3,
			Name: "山縣　有朋",
			NyutaikunUserId: 300,
		},
		Student{
			Id: 4,
			Name: "松方　正義",
			NyutaikunUserId: 400,
		},
		Student{
			Id: 5,
			Name: "大隈　重信",
			NyutaikunUserId: 500,
		},
	}

	return &ret
}
