package middleware

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
)

type AuthExtractor struct {
	authHelper     auth.Helper
	log            logger.Logger
	ignorancePaths *libhttp.PathMatchers
}

func NewAuthExtractor(
	ignorancePaths *libhttp.PathMatchers,
	authHelper auth.Helper,
	logger logger.Logger,
) *AuthExtractor {
	return &AuthExtractor{
		ignorancePaths: ignorancePaths,
		authHelper:     authHelper,
		log:            logger.GetLogger("HTTP-JWT-Extractor"),
	}
}

func (aem *AuthExtractor) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		aem.log.Debug("AuthExtractorMiddleware.Handle start")
		defer aem.log.Debug("AuthExtractorMiddleware.Handle finish")

		// check for ignorance
		if aem.ignorancePaths.Match(r.Method, r.RequestURI) {
			next.ServeHTTP(rw, r)

			return
		}

		// here and so on we extract subject and gen 401 or continue pipeline
		// extract subject
		subj, err := aem.authHelper.SubjectFromHTTPRequest(r)
		if err != nil {
			aem.log.Errorf("AuthExtractorMiddleware.Handle error [%v]", err)

			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(rw, r.WithContext(auth.WithSubject(r.Context(), subj)))
	})
}
