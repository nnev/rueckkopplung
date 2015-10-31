package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"flag"
)

var (
	questions = []string{}
	questionPipe = make(chan string)
)

var addr = flag.String("web.address", "127.0.0.1:8080", "web-address and port to listen on")

var formPage = `<!DOCTYPE html>
<html>
<head>
<title>Frage stellen</title>
<meta charset="UTF-8">
</head>
<body>
<form action="/submit" method="POST">
    Frage: <input type="text" name="frage">
    <input type="submit" value="fragen">
</form>
</body>
</html>`

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
	http.HandleFunc("/submit", handleForm)
	http.HandleFunc("/questions", exposeQuestions)
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, formPage)
		})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
