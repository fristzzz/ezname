package main

import (
	"fmt"
	"strings"

	stu "./student"
)

func GetName(s []stu.Student, id, form string) (newName string) {
	fmt.Printf("id found:%+s\n", id)
	var astu stu.Student
	for i := 0; i < len(s); i++ {
		if strings.EqualFold(s[i].Id, id) {
			astu = s[i]
			break
		}
	}
	newName = strings.ReplaceAll(form, "姓名", astu.Name)
	newName = strings.ReplaceAll(newName, "班级", astu.Faculty+astu.Class)
	newName = strings.ReplaceAll(newName, "学号", astu.Id)
	return

}
