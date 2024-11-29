package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tberk-s/Go-To-Do-API/internal/todo"
)

// Item ...
type Item struct {
	Item string `json:"item"` // This is item
}

// Server ...
type Server struct {
	mux *http.ServeMux
}

// New ...
func New(todoSvc *todo.Service) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(write http.ResponseWriter, _ *http.Request) {
		todoItems, err := todoSvc.GetAll()
		if err != nil {
			log.Println(err)
			write.WriteHeader(http.StatusInternalServerError)

			return
		}
		write.WriteHeader(http.StatusOK)
		b, err := json.Marshal(todoItems)
		if err != nil {
			log.Println(err)
		}
		_, err = write.Write(b)
		if err != nil {
			log.Println(err)
		}

	})

	mux.HandleFunc("POST /todo", func(write http.ResponseWriter, req *http.Request) {
		var todo Item
		if err := json.NewDecoder(req.Body).Decode(&todo); err != nil {
			log.Println(err)
			write.WriteHeader((http.StatusBadRequest))

			return
		}

		if err := todoSvc.Add(todo.Item); err != nil {
			errMessage := fmt.Sprintf("Error : %s", err)
			log.Println(err)
			write.WriteHeader(http.StatusBadRequest)
			_, _ = write.Write([]byte(errMessage))
		}
		write.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("GET /search", func(write http.ResponseWriter, req *http.Request) {
		query := req.URL.Query().Get("q")

		if query == "" {
			write.WriteHeader(http.StatusBadRequest)

			return
		}

		result, err := todoSvc.Search(query)
		if err != nil {
			log.Println(err.Error())
			write.WriteHeader(http.StatusInternalServerError)

			return
		}

		b, _ := json.Marshal(result)
		write.Write(b)
	})

	mux.HandleFunc("DELETE /todo", func(write http.ResponseWriter, req *http.Request) {
		task := req.URL.Query().Get("task")
		if task == "" {
			write.WriteHeader(http.StatusBadRequest)

			return
		}

		if err := todoSvc.Delete(task); err != nil {
			log.Println(err)
			write.WriteHeader(http.StatusBadRequest)
			write.Write([]byte(err.Error()))

			return
		}
		message := fmt.Sprintf("Deleted task: %s\n", task)
		write.WriteHeader(http.StatusOK)
		number, writeErr := write.Write([]byte(message))
		log.Println(number)
		if writeErr != nil {
			log.Println(writeErr.Error())
			write.WriteHeader(http.StatusInternalServerError)

			return
		}
	})

	return &Server{
		mux: mux,
	}
}

// Serve ...
func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
