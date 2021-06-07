package apiserver

import (
	"net/http"

	"github.com/furrygem/authentication_server/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// the main struct that will contain configured modules instances
type apiserver struct {
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
	//sessionStore        *sessions.CookieStore
	//sessionStoreAuthKey []byte
}

// entrypoint to server
func Start(c *Config) error {
	var (
		server *apiserver
		st     *store.Store
		err    error
	)
	logger := logrus.New()
	server = &apiserver{
		logger: logger,
		router: mux.NewRouter(),
	} // create new apiserver instance
	/*creating new store instance and cathing error*/
	st, err = store.Open(c.StoreConfig)
	if err != nil {
		server.logger.Error("Error initializing store instance")
		return err
	}
	st.Logger = logger
	server.store = st
	if err = http.ListenAndServe(c.BindAddr, server); err != nil {
		server.logger.Error("Error while serving")
		return err
	}
	return nil

}

func (s *apiserver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.logger.Infof("Serving client | remote_addr: %s | endpoint: %s | Method: %s", r.RemoteAddr, r.RequestURI, r.Method)
	s.router.ServeHTTP(rw, r)
}
