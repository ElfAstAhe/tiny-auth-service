package middleware

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport"
)

type SubjectExtractor struct {
	jwtHTTPHelper  *helper.JWTHTTPHelper
	authHelper     *auth.Helper
	log            logger.Logger
	ignorancePaths *transport.HTTPPathMatchers
}

func NewSubjectExtractorMiddleware(
	ignorancePaths *transport.HTTPPathMatchers,
	jwtHTTPHelper *helper.JWTHTTPHelper,
	authHelper *auth.Helper,
	logger logger.Logger,
) *SubjectExtractor {
	return &SubjectExtractor{
		ignorancePaths: ignorancePaths,
		jwtHTTPHelper:  jwtHTTPHelper,
		authHelper:     authHelper,
		log:            logger.GetLogger("HTTP-JWT-Extractor"),
	}
}

func (je *SubjectExtractor) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		je.log.Info("AuthExtractorMiddleware.Handle start")
		defer je.log.Info("AuthExtractorMiddleware.Handle finish")

		// check for ignorance
		if je.ignorancePaths.Match(r.Method, r.RequestURI) {
			next.ServeHTTP(rw, r)

			return
		}

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
