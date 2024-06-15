package blog

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-blog/types"
	"simple-blog/utils"
)

type Handler struct {
	blogStore types.BlogStore
}

func NewBlogHandler(blogStore types.BlogStore) *Handler {
	return &Handler{blogStore: blogStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/blog", h.createBlog).Methods(http.MethodPost)
	router.HandleFunc("/blog", h.getBlogs).Methods(http.MethodGet)
	router.HandleFunc("/blog/{id}", h.getBlogById).Methods(http.MethodGet)
	router.HandleFunc("/blog/{id}", h.updateBlog).Methods(http.MethodPatch)
	router.HandleFunc("/blog/{id}", h.deleteBlog).Methods(http.MethodDelete)
}

func (h *Handler) createBlog(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateBlogPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	blog, err := h.blogStore.CreateBlog(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"message": "blog created",
		"blog":    blog,
	})

}

func (h *Handler) getBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.blogStore.GetBlogs()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"blogs": blogs,
	})
}

func (h *Handler) getBlogById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	intId, err := utils.ConvertBlogIdToInt(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	blog, err := h.blogStore.GetBlogById(intId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"blog": blog,
	})
}

func (h *Handler) updateBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	intId, err := utils.ConvertBlogIdToInt(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	var payload types.CreateBlogPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	blog, err := h.blogStore.GetBlogById(intId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.blogStore.UpdateBlog(blog.Id, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"message": "blog updated",
		"blog":    res,
	})

}

func (h *Handler) deleteBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	intId, err := utils.ConvertBlogIdToInt(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	delErr := h.blogStore.DeleteBlog(intId)
	if delErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, delErr)
	}
	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"message": "blog deleted",
	})
}
