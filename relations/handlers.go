package relations

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rkeplin/bible-go-api/core"
)

type Repo interface {
	FindAll(verseId int, translation string) ([]TextCollection, error)
}

type Handler struct {
	repo Repo
}

func (h Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}

	queryParams := r.URL.Query()
	translation := queryParams.Get("translation")

	collection, err := h.repo.FindAll(id, translation)

	if err != nil {
		panic(err)
	}

	core.Respond(w, http.StatusOK, collection)
}
