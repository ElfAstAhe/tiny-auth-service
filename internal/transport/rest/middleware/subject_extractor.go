package middleware

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

type SubjectExtractor struct {
	jwtHTTPHelper *helper.JWTHTTPHelper
	authHelper    *auth.Helper
	log           logger.Logger
	ignore        []string
}

func NewSubjectExtractorMiddleware(jwtHTTPHelper *helper.JWTHTTPHelper, authHelper *auth.Helper, logger logger.Logger) *SubjectExtractor {
	return &SubjectExtractor{
		jwtHTTPHelper: jwtHTTPHelper,
		authHelper:    authHelper,
		log:           logger.GetLogger("HTTP-JWT-Extractor"),
	}
}

func (je *SubjectExtractor) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		je.log.Info("AuthExtractorMiddleware.Handle start")
		defer je.log.Info("AuthExtractorMiddleware.Handle finish")

		// check for ignorance
		// ..

		// here and so on we extract subject and gen 401 or continue pipeline
		// extract subject
		subj, err := je.authHelper.SubjectFromHTTPRequest(r)
		if err != nil {
			je.log.Errorf("AuthExtractorMiddleware.Handle error [%v]", err)

			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(rw, r.WithContext(auth.WithSubject(r.Context(), subj)))
	})
}
