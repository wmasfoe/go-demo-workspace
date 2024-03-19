package data

type Student struct {
	Id     int64    `json:"id"`
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Skills []string `json:"skills"`
}

func GetStudent(id int64) Student {
	var std = Student{}
	for _, student := range Students {
		if student.Id == id {
			std = student
		}
	}
	return std
}

var Students = []Student{
	{1, "John", 20, []string{"C", "Python"}},
	{2, "Doe", 22, []string{"Go", "Python"}},
	{3, "Smith", 24, []string{"Go", "Python", "Java"}},
	{4, "Tom", 26, []string{"Go", "Python", "Java", "C"}},
	{5, "Jerry", 28, []string{"Go", "Python", "Java", "C", "C++"}},
}
