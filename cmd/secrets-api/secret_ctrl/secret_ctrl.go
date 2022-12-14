package secret_ctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	apiErr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
	"github.com/sstalschus/secrets-api/internal/secret"
)

type Controller struct {
	secretService secret.IService
	logger        log.Provider
	apiErr        apiErr.Provider
}

func NewController(
	secretService secret.IService,
	logger log.Provider,
	apiErr apiErr.Provider,
) *Controller {
	return &Controller{
		secretService: secretService,
		logger:        logger,
		apiErr:        apiErr,
	}
}

func (c Controller) CreateSecret(w http.ResponseWriter, r *http.Request) {
	userID := internal.GetField(r.Context(), "user_id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to read body", log.Body{})
		return
	}

	var secret internal.Secret

	err = json.Unmarshal(body, &secret)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to unmarshal body", log.Body{})
		return
	}

	errResponse := c.validateBody(&secret)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	errResponse = c.secretService.CreateSecret(r.Context(), &secret, userID)
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

func (c Controller) GetSecrets(w http.ResponseWriter, r *http.Request) {
	userID := internal.GetField(r.Context(), "user_id")

	secrets := c.secretService.GetSecrets(r.Context(), userID)

	secretsRes, err := json.Marshal(secrets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.logger.Error(r.Context(), "Internal Server - Fail to unmarshal body", log.Body{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(secretsRes)
	return
}

func (c Controller) GetSecret(w http.ResponseWriter, r *http.Request) {
	userID := internal.GetField(r.Context(), "user_id")
	secretID := internal.GetFields(r.Context(), "CtxKey", 0)

	secret, errResponse := c.secretService.GetSecret(r.Context(), secretID, userID)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	secretsRes, err := json.Marshal(secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.logger.Error(r.Context(), "Internal Server - Fail to unmarshal body", log.Body{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(secretsRes)
	return
}

func (c Controller) validateBody(secret *internal.Secret) (apiErr *apiErr.Message) {
	if secret.Value == "" || secret.Key == "" {
		apiErr = c.apiErr.BadRequest("Missing params", fmt.Errorf(""))
	}
	return apiErr
}
