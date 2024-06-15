package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"simple-blog/service/blog"
)

type Server struct {
	addr string
	db   *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// blog handler
	blogStore := blog.NewStorage(s.db)
	blogHandler := blog.NewBlogHandler(blogStore)
	blogHandler.RegisterRoutes(subRouter)

	fmt.Println("Server is running on port ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
