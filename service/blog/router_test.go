package blog

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"simple-blog/types"
	"testing"
)

type mockBlogStore struct{}

func TestBlogServiceHandler(t *testing.T) {
	blogStore := &mockBlogStore{}
	handler := NewBlogHandler(blogStore)
	router := mux.NewRouter()

	t.Run("/CREATE", func(t *testing.T) {
		t.Run("should create a new Blog", func(t *testing.T) {
			payload := types.CreateBlogPayload{
				Title:       "title",
				Description: "description",
			}

			marshalled, _ := json.Marshal(payload)

			req, err := http.NewRequest(http.MethodPost, "/blog", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			router.HandleFunc("/blog", handler.createBlog).Methods("POST")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})

		t.Run("should fail if no payload is provide", func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/blog", nil)
			if err != nil {
				t.Fatal(err)
			}

			router.HandleFunc("/blog", handler.createBlog).Methods("POST")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		})

	})

	t.Run("/GET_ALL_BLOGS should return all blogs", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/blog", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.HandleFunc("/blog", handler.getBlogs).Methods(http.MethodGet)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	// add update unit tests

	t.Run("/GET_BLOG_BY_ID", func(t *testing.T) {
		t.Run("should return a single blog", func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/blog/1", nil)
			if err != nil {
				t.Fatal(err)
			}

			router.HandleFunc("/blog/{id}", handler.getBlogById).Methods(http.MethodGet)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})

		t.Run("should fail if the id is invalid or null", func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/blog/hello", nil)
			if err != nil {
				t.Fatal(err)
			}

			router.HandleFunc("/blog/{id}", handler.getBlogById).Methods(http.MethodGet)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		})

	})

	t.Run("/DELETE_BLOG", func(t *testing.T) {
		t.Run("should delete the blog", func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/blog/2", nil)
			if err != nil {
				t.Fatal(err)
			}

			router.HandleFunc("/blog/{id}", handler.deleteBlog).Methods("DELETE")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})

		t.Run("should fail if the blog id is invalid or null", func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/blog/hello", nil)
			if err != nil {
				t.Fatal(err)
			}

			router.HandleFunc("/blog/{id}", handler.deleteBlog).Methods("DELETE")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		})
	})

}

func (m *mockBlogStore) CreateBlog(payload *types.CreateBlogPayload) (*types.Blog, error) {
	return nil, nil
}

func (m *mockBlogStore) GetBlogs() ([]types.Blog, error) {
	return nil, nil
}

func (m *mockBlogStore) GetBlogById(id int) (*types.Blog, error) {
	return nil, nil
}

func (m *mockBlogStore) UpdateBlog(id int, payload *types.CreateBlogPayload) (*types.Blog, error) {
	return nil, nil
}

func (m *mockBlogStore) DeleteBlog(id int) error {
	return nil
}
