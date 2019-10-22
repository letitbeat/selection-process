package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler handles the exposed endpoints
type Handler struct {
	DBURI string
}

// GetByID returns a JSON representation of a User object
// given its ID if matched.
func (h *Handler) GetByID(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id := string(params["id"])

	r := Repository{DBURI: h.DBURI}

	user, err := r.GetByID(id)

	if err != nil {
		writeErr(response, err)
	}

	err = json.NewEncoder(response).Encode(user)

	if err != nil {
		writeErr(response, err)
	}
}

func writeErr(response http.ResponseWriter, err error) {
	msg := fmt.Sprintf(`{"message" : "%s"}`, err.Error())
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(msg))
}
