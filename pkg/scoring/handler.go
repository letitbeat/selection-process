package scoring

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/letitbeat/selection-process/pkg/tag"
	"github.com/letitbeat/selection-process/pkg/user"
)

// Handler for score endpoint
type Handler struct {
	DBURI string
}

// Score endpoint that performs the scoring for the given task Id
func (h *Handler) Score(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	taskID := string(params["taskID"])

	var users []user.User
	err := json.NewDecoder(request.Body).Decode(&users)
	if err != nil {
		writeErr(response, err)
	}

	r := tag.Repository{DBURI: h.DBURI}
	taskTags, err := r.GetByTaskID(taskID)

	var usersRsp []user.User
	scorer := Scorer{Ref: cleanTags(taskTags.Tags)}
	for _, user := range users {
		user.Score = scorer.Score(user.Tags)
		user.Tags = []string{}
		usersRsp = append(usersRsp, user)
	}

	if err != nil {
		writeErr(response, err)
	}

	err = json.NewEncoder(response).Encode(usersRsp)
	if err != nil {
		writeErr(response, err)
	}
}

// cleanTags in case of duplicates
func cleanTags(tags []string) map[string]bool {
	m := make(map[string]bool)
	for _, tag := range tags {
		m[tag] = true
	}
	return m
}

func writeErr(response http.ResponseWriter, err error) {
	msg := fmt.Sprintf(`{"message" : "%s"}`, err.Error())
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(msg))
}
