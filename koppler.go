package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	questions    = []string{}
	questionPipe = make(chan string)
)

var addr = flag.String("web.address", "127.0.0.1:8080", "web-address and port to listen on")

func writeError(errno int, res http.ResponseWriter, format string, args ...interface{}) {
	res.WriteHeader(errno)
	fmt.Fprintf(res, format, args...)
}

func handleForm(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Println("no post")
		return
	}

	err := req.ParseForm()
	if err != nil {
		// Not sure what to do, so we just return 400
		log.Printf("ParseMultipartForm returned %v", err)
		writeError(http.StatusBadRequest, res, "Bad request.")
		return
	}

	frage := req.PostFormValue("frage")

	if frage == "" {
		log.Println("empty question submitted")
		res.Write([]byte("Oh, das Feld für die Frage war leer! :("))
	} else {
		log.Println("question submitted:", frage)
		questionPipe <- frage
		res.Write([]byte("Vielen Dank für deine Frage! :)"))
	}

}

func collectQuestions() {
	for {
		question := <-questionPipe
		questions = append(questions, question)
	}
}

func exposeQuestions(res http.ResponseWriter, req *http.Request) {
	questionblock := strings.Join(questions, "\n")
	//res.Write([]byte(questionblock))
	fmt.Fprintf(res, questionblock)
}

func main() {
	flag.Parse()
	go collectQuestions()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/submit", handleForm)
	http.HandleFunc("/questions/raw", exposeQuestions)
	http.HandleFunc("/questions",
		func(w http.ResponseWriter, r *http.Request) {
			if err := ExecuteTemplate(w, TemplateInput{Body: "questions.html"}); err != nil {
				log.Println("Could not render template:", err)
				http.Error(w, "Internal error", 500)
				return
			}
		})
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			if err := ExecuteTemplate(w, TemplateInput{Body: "form.html"}); err != nil {
				log.Println("Could not render template:", err)
				http.Error(w, "Internal error", 500)
				return
			}
		})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
