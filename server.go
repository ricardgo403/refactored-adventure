package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"./args"
	"./student"
	"./subject"
)

var mySubjects subject.SubjectAdmin
var myStudents student.StudentAdmin
var l = list.New()

func exists(student *student.Student) bool {
	for _, val := range myStudents.Students {
		if val.Name == student.Name {
			for key := range student.Subject {
				// fmt.Println("Key:", key, "=>", "Element:", element)
				if v, found := val.Subject[key]; found {
					fmt.Println(v)
					return true
				}
			}

		}
	}
	return false
}

func loadHTML(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

func addForm(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHTML("./addForm.html"),
	)
}

func add(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		mySubject := make(map[string]float64)
		f, _ := strconv.ParseFloat(req.FormValue("grade"), 64)
		mySubject[req.FormValue("subject")] = f
		student := student.Student{Name: req.FormValue("student"), Subject: mySubject}
		fmt.Println("Student: ", student)
		myStr := ""
		if !exists(&student) {
			myArgs := args.Args{Name: student.Name, Subject: req.FormValue("subject"), Grade: f}
			thisSubject := subject.Subject{Name: req.FormValue("subject"), Grade: f}
			mySubjects.Add(thisSubject)
			fmt.Println(mySubjects)
			myStudents.Add(student, &myArgs)
			fmt.Println(myStudents)
			myStr = " Se ha agregado un alumno nuevo: " + student.String()
		} else {
			myStr = "La calificación ya estaba asignada y no ha sido modificada..."
		}
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			loadHTML("./result.html"),
			myStr,
		)
	}
}

func stdAvrgForm(res http.ResponseWriter, req *http.Request) {
	myStr := myStudents.SelectString()
	// fmt.Println(myStr)
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHTML("./studentForm.html"),
		myStr,
	)
}

func stdAvrg(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		myStudent := req.FormValue("student")
		studentGradeAverage := float64(0)
		for _, val := range myStudents.Students {
			if val.Name == myStudent {
				// fmt.Println("Encontrado...")
				// fmt.Println("Sacando promedio...")
				counter := float64(0)
				for _, grade := range val.Subject {
					studentGradeAverage += grade
					counter++
				}
				studentGradeAverage /= counter
			}
		}
		myStr := "El promedio de " + myStudent + " es " + strconv.FormatFloat(studentGradeAverage, 'f', 2, 64)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			loadHTML("./studentResult.html"),
			myStr,
		)
	}
}

func subAvrgForm(res http.ResponseWriter, req *http.Request) {
	myStr := mySubjects.SelectString()
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHTML("./subjectForm.html"),
		myStr,
	)
}

func subAvrg(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		subject := req.FormValue("subject")
		subjectGradeAverage := float64(0)
		counter := float64(0)
		for _, val := range myStudents.Students {
			// fmt.Println("Encontrado...")
			// fmt.Println("Sacando promedio...")
			if v, found := val.Subject[subject]; found {
				subjectGradeAverage += v
				counter++
			}
		}
		subjectGradeAverage /= counter
		myStr := "El promedio de " + subject + " es " + strconv.FormatFloat(subjectGradeAverage, 'f', 2, 64)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			loadHTML("./subjectResult.html"),
			myStr,
		)
	}
}

func calcGenAvrg() float64 {
	generalAverage := float64(0)
	counter := float64(0)
	for _, val := range myStudents.Students {
		generalAverage += val.GradeAvrgF()
		counter++
	}
	return generalAverage / counter
}

func genAvrg(res http.ResponseWriter, req *http.Request) {
	generalAverage := calcGenAvrg()
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		myStr := "El promedio general es de: " + strconv.FormatFloat(generalAverage, 'f', 2, 64)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			loadHTML("./generalForm.html"),
			myStr,
		)
	}
}

func main() {
	http.HandleFunc("/addForm", addForm)
	http.HandleFunc("/add", add)
	http.HandleFunc("/studentAvrgForm", stdAvrgForm)
	http.HandleFunc("/stdAverage", stdAvrg)
	http.HandleFunc("/subject-gda-avrg", subAvrgForm)
	http.HandleFunc("/subjectAverage", subAvrg)
	http.HandleFunc("/general-avrg", genAvrg)
	fmt.Println("Corriendo servidor de calificaciones...")
	http.ListenAndServe(":9000", nil)
}
