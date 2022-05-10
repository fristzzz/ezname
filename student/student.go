package student

type Student struct {
	Faculty string
	Class   string
	Name    string
	Id      string
}

const (
	F    string = "Faculty"
	C    string = "Class"
	Name string = "Name"
	Id   string = "Id"
)

func NewStudent(f, cls, name, id string) *Student {
	return &Student{
		Faculty: f,
		Class:   cls,
		Name:    name,
		Id:      id,
	}
}
