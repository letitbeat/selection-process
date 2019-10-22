package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/letitbeat/selection-process/pkg/user"
)

type Handler struct {
	DBURI         string
	ScoringServer string
}

func (h *Handler) GetByID(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id := string(params["id"])

	r := Repository{DBURI: h.DBURI}

	task, err := r.GetByID(id)

	if err != nil {
		writeErr(response, err)
	}

	var applicants []string
	for _, applicant := range task.Applicants {
		applicants = append(applicants, applicant.ID)
	}

	users, err := score(id, applicants, h.DBURI, h.ScoringServer)
	if err != nil {
		writeErr(response, err)
	}

	// sorting by score before
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	err = json.NewEncoder(response).Encode(convertTask(task, users))
	if err != nil {
		writeErr(response, err)
	}
}

func score(taskID string, applicants []string, dbURI, scoringSrv string) ([]user.User, error) {

	r := user.Repository{DBURI: dbURI}
	users, err := r.GetByIDs(applicants)

	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(users)
	if err != nil {
		return nil, err
	}

	scoringURI := fmt.Sprintf("%s/scoring/%s", scoringSrv, taskID)
	request, err := http.NewRequest("POST", scoringURI, payload)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var usersResp []user.User
	err = json.NewDecoder(resp.Body).Decode(&usersResp)
	if err != nil {
		return nil, err
	}
	return usersResp, nil
}

type task struct {
	ID         string      `json:"taskId"`
	Applicants []user.User `json:"applicants"`
	Desc       string      `json:"description"`
	Country    string      `json:"country"`
	Tags       []string    `json:"tags"`
}

// converts to expected struct
func convertTask(t Task, users []user.User) task {
	return task{
		ID:         t.ID,
		Applicants: users,
		Desc:       t.Description,
		Country:    t.Country,
	}
}

func writeErr(response http.ResponseWriter, err error) {
	msg := fmt.Sprintf(`{"message" : "%s"}`, err.Error())
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(msg))
}
