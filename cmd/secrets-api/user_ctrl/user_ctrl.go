package user_ctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	apiErr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
	"github.com/sstalschus/secrets-api/internal/user"
)

type Controller struct {
	usersService user.IService
	logger       log.Provider
	apiErr       apiErr.Provider
}

func NewController(
	usersService user.IService,
	logger log.Provider,
	apiErr apiErr.Provider,
) *Controller {
	return &Controller{
		usersService: usersService,
		logger:       logger,
		apiErr:       apiErr,
	}
}

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to read body", log.Body{})
		return
	}

	var user internal.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to unmarshal body", log.Body{})
		return
	}

	errResponse := c.validateBody(&user)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	errResponse = c.usersService.CreateUser(r.Context(), &user)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (c Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := internal.GetField(r.Context(), "user_id")

	user, errResponse := c.usersService.GetUser(r.Context(), userID)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	userRes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.logger.Error(r.Context(), "Internal Server - Fail to unmarshal body", log.Body{})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userRes)
}

func (c Controller) validateBody(user *internal.User) (apiErr *apiErr.Message) {
	if user.Email == "" || user.Name == "" || user.Password == "" {
		apiErr = c.apiErr.BadRequest("Missing params", fmt.Errorf(""))
	}
	return apiErr
}
