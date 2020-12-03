package student

import (
	"strconv"

	"../args"
)

type Student struct {
	Name    string
	Subject map[string]float64
}

func (student *Student) String() string {
	myStr := "Nombre: " + student.Name + "<br>"
	for key, val := range student.Subject {
		myStr += "Subject:" + key + "=>" + "Grade:" + strconv.FormatFloat(val, 'f', 2, 64)
	}
	return myStr
}

func (student *Student) GradeAvrg() string {
	myStr := ""
	studentGradeAverage := float64(0)
	counter := float64(0)
	for _, grade := range student.Subject {
		studentGradeAverage += grade
		counter++
	}
	studentGradeAverage /= counter
	myStr = strconv.FormatFloat(studentGradeAverage, 'f', 2, 64)
	return myStr
}

func (student *Student) GradeAvrgF() float64 {
	studentGradeAverage := float64(0)
	counter := float64(0)
	for _, grade := range student.Subject {
		studentGradeAverage += grade
		counter++
	}
	studentGradeAverage /= counter
	return studentGradeAverage
}

func (student *Student) SelectString() string {
	myStr := "<option value=\"" + student.Name + "\">" + student.Name + "</option>\n"
	return myStr
}

type StudentAdmin struct {
	Students []Student
}

func (students *StudentAdmin) exists(student Student) int {
	// fmt.Println("Validando...")
	for i, val := range students.Students {
		// fmt.Println("i: ", i, "val: ", val)
		if val.Name == student.Name {
			return i
		}
	}
	return -1
}

func (students *StudentAdmin) Add(student Student, args *args.Args) {
	if i := students.exists(student); i >= 0 {
		// fmt.Println("Ya existe")
		students.Students[i].Subject[args.Subject] = args.Grade
	} else {
		students.Students = append(students.Students, student)
	}
}

func (students *StudentAdmin) SelectString() string {
	myStr := ""
	for _, val := range students.Students {
		// fmt.Println(i, val)
		myStr += val.SelectString()
	}
	return myStr
}
