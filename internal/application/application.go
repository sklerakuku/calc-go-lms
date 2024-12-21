package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sklerakuku/calc-go-lms/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

// Функция запуска приложения
// тут будем читать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

func HasLetters(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, ":)))) \n")
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		request := new(Request)
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "Internal server error"}`)
			return
		}

		if HasLetters(request.Expression) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "Expression is not valid"}`)
			return
		}

		result, err := calculation.Calc(request.Expression)
		if err != nil {
			log.Printf("NOT OK: err is %s\n", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, `{"error": "Expression is not valid"}`)
		} else {
			log.Printf("OK: result is %f\n", result)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"result": "%f"}`, result)
		}
	default:
		log.Println("not POST request")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Internal server error"}`))
	}
}

/*type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}*/

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func nihao(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "你好！\n")
}

func (a Application) RunServer() {
	log.Println("Server Run...\nhttp://localhost:8080/api/v1/calculate")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/nihao", nihao)
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	//return http.ListenAndServe(":"+a.config.Addr, nil)

	log.Fatal(http.ListenAndServe(":"+a.config.Addr, nil))
}
