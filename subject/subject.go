package subject

import "fmt"

type Subject struct {
	Name  string
	Grade float64
}

type RegisteredSubject struct {
	Name    string
	Average float64
	Sum     float64
	Counter int
}

func (subject *RegisteredSubject) SelectString() string {
	myStr := "<option value=\"" + subject.Name + "\">" + subject.Name + "</option>\n"
	return myStr
}

type SubjectAdmin struct {
	Subjects []RegisteredSubject
}

func (subjects *SubjectAdmin) exists(subject Subject) int {
	// fmt.Println("Validando...")
	for i, val := range subjects.Subjects {
		// fmt.Println("i: ", i, "val: ", val)
		if val.Name == subject.Name {
			return i
		}
	}
	return -1
}

func (subjects *SubjectAdmin) Add(subject Subject) {
	if i := subjects.exists(subject); i >= 0 {
		fmt.Println("Ya existe")
		fmt.Println(subjects.Subjects[i])
		subjects.Subjects[i].Sum += subject.Grade
		subjects.Subjects[i].Counter++
		subjects.Subjects[i].Average = subjects.Subjects[i].Sum / float64(subjects.Subjects[i].Counter)
		fmt.Println(subjects.Subjects[i])
	} else {
		newRegSubject := RegisteredSubject{Name: subject.Name, Average: subject.Grade, Sum: subject.Grade, Counter: 1}
		subjects.Subjects = append(subjects.Subjects, newRegSubject)
	}
}

func (subjects *SubjectAdmin) SelectString() string {
	myStr := ""
	for _, val := range subjects.Subjects {
		// fmt.Println(i, val)
		myStr += val.SelectString()
	}
	return myStr
}
