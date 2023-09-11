package plugin

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"

	"github.com/mattermost/mattermost-plugin-template/server/constants"
	"github.com/mattermost/mattermost-plugin-template/server/serializers"
)

// Initializes the plugin REST API
func (p *Plugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.Use(p.WithRecovery)

	// 404 handler
	r.Handle(constants.WildRoute, http.NotFoundHandler())
	return r
}

// Add custom routes and corresponding handlers here
func (p *Plugin) InitRoutes() {
	p.Client = InitClient(p)

	s := p.router.PathPrefix(constants.APIPrefix).Subrouter()

	// TODO: Below are for demo purpose remove them later
	s.HandleFunc(constants.PathGetMe, p.handleAuthRequired(p.handleGetMe)).Methods(http.MethodGet)
	s.HandleFunc(constants.PathGetMockPosts, p.handleAuthRequired(p.handleGetDummyPosts)).Methods(http.MethodGet)
}

// TODO: Below is for demo purposes only remove it later
func (p *Plugin) handleGetMe(w http.ResponseWriter, r *http.Request) {
	mattermostUserID := r.Header.Get(constants.HeaderMattermostUserID)

	userDetails, err := p.API.GetUser(mattermostUserID)
	if err != nil {
		p.API.LogError(constants.ErrorFetchingUserDetails, constants.Error, err.Error())
		p.handleError(w, r, &serializers.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	p.writeJSON(w, http.StatusOK, userDetails)
}

// TODO: Below is for demo purposes only remove it later
func (p *Plugin) handleGetDummyPosts(w http.ResponseWriter, r *http.Request) {
	posts, statusCode, err := p.Client.GetMockPosts()
	if err != nil {
		p.API.LogWarn(constants.ErrorFetchingDummyPosts, constants.Error, err.Error())
		p.handleError(w, r, &serializers.Error{Code: statusCode, Message: err.Error()})
		return
	}

	p.writeJSON(w, statusCode, posts)
}

// writeJSON handles writing HTTP JSON response
func (p *Plugin) writeJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(v)
	if err != nil {
		p.API.LogError("Failed to marshal JSON response", constants.Error, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(b); err != nil {
		p.API.LogError("Failed to write JSON response", constants.Error, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
}

// handleError handles writing HTTP response error
func (p *Plugin) handleError(w http.ResponseWriter, r *http.Request, error *serializers.Error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(error.Code)
	message := map[string]string{constants.Error: error.Message}
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handleAuthRequired verifies if the provided request is performed by an authorized source.
func (p *Plugin) handleAuthRequired(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mattermostUserID := r.Header.Get(constants.HeaderMattermostUserID)
		if mattermostUserID == "" {
			p.handleError(w, r, &serializers.Error{Code: http.StatusUnauthorized, Message: constants.NotAuthorized})
			return
		}

		handleFunc(w, r)
	}
}

// WithRecovery handles recovery from panic
func (p *Plugin) WithRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				p.API.LogError("Recovered from a panic",
					"url", r.URL.String(),
					"error", x,
					"stack", string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
