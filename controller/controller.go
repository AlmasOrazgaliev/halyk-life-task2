package controller

import (
	"encoding/json"
	"errors"
	"github.com/AlmasOrazgaliev/halyk-life-task2/models"
	"github.com/AlmasOrazgaliev/halyk-life-task2/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Controller struct {
	repository *repository.Repository
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		repository: repository.NewDB(db),
	}
}

func (c *Controller) HandleAuthors(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := c.repository.GetAuthors()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errResponse(w, http.StatusNotFound, err)
			return
		} else if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusOK, res)
	} else if r.Method == "POST" {
		var author models.Author
		err := json.NewDecoder(r.Body).Decode(&author)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
			return
		}
		err = c.repository.CreateAuthor(&author)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (c *Controller) HandleBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := c.repository.GetBooks()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errResponse(w, http.StatusNotFound, err)
			return
		} else if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusOK, res)
	} else if r.Method == "POST" {
		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
			return
		}
		err = c.repository.CreateBook(&book)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (c *Controller) HandleMembers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := c.repository.GetMembers()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errResponse(w, http.StatusNotFound, err)
			return
		} else if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusOK, res)
	} else if r.Method == "POST" {
		var member models.Member
		err := json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
			return
		}
		err = c.repository.CreateMember(&member)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (c *Controller) HandleAuthorById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	uid := uint(id)
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	author, err := c.repository.GetAuthor(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errResponse(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		errResponse(w, http.StatusInternalServerError, err)
		return
	}
	if r.Method == "GET" {
		response(w, http.StatusOK, author)
	} else if r.Method == "DELETE" {
		err = c.repository.DeleteAuthor(author)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusNoContent, nil)
	} else if r.Method == "PUT" {
		var updatedAuthor models.Author
		err = json.NewDecoder(r.Body).Decode(&updatedAuthor)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		err = c.repository.UpdateAuthor(author, &updatedAuthor)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (c *Controller) HandleBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	uid := uint(id)
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	book, err := c.repository.GetBook(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errResponse(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		errResponse(w, http.StatusInternalServerError, err)
		return
	}
	if r.Method == "GET" {
		response(w, http.StatusOK, book)
	} else if r.Method == "DELETE" {
		err = c.repository.DeleteBook(book)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusNoContent, nil)
	} else if r.Method == "PUT" {
		var updatedBook models.Book
		err = json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		err = c.repository.UpdateBook(book, &updatedBook)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (c *Controller) HandleMemberById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	uid := uint(id)
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	member, err := c.repository.GetMember(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errResponse(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		errResponse(w, http.StatusInternalServerError, err)
		return
	}
	if r.Method == "GET" {
		response(w, http.StatusOK, member)
	} else if r.Method == "DELETE" {
		err = c.repository.DeleteMember(member)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusNoContent, nil)
	} else if r.Method == "PUT" {
		var updatedMember models.Member
		err = json.NewDecoder(r.Body).Decode(&updatedMember)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		err = c.repository.UpdateMember(member, &updatedMember)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (c *Controller) HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	memberId, err := strconv.Atoi(mux.Vars(r)["memberId"])
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	uMemberId := uint(memberId)
	var subscribe models.BookMember
	err = json.NewDecoder(r.Body).Decode(&subscribe)
	subscribe.MemberId = uMemberId
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	if r.Method == "POST" {
		err = c.repository.CreateSubscribe(&subscribe)
		if err != nil {
			response(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}

}

func (c *Controller) HandleAuthorBooks(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	uid := uint(id)
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	if r.Method == "GET" {
		res, err := c.repository.GetAuthorBooks(uid)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errResponse(w, http.StatusNotFound, err)
			return
		} else if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusOK, res)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}

}

func (c *Controller) HandleMemberBooks(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	uid := uint(id)
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	if r.Method == "GET" {
		res, err := c.repository.GetMemberBooks(uid)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errResponse(w, http.StatusNotFound, err)
			return
		} else if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusOK, res)
	} else {
		errResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func errResponse(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error": err.Error()})
}

func response(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
