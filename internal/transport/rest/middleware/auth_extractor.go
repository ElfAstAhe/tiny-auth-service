package middleware

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
)

// AuthExtractor implements a standard HTTP middleware interceptor capable of validating inbound tokens.
// It orchestrates cryptographic extraction boundaries and populates runtime context targets with validated Subject states.
type AuthExtractor struct {
	authHelper     auth.Helper           // Framework core security utility managing cryptographic signature verifications
	log            logger.Logger         // Dedicated structured logging handler targeting security auditing boundaries
	ignorancePaths *libhttp.PathMatchers // Registry passport containing abstract criteria patterns for publicly exposed routes
}

// NewAuthExtractor acts as a factory constructor mounting required token parsing and routing exemption dependencies.
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

// Handle intercept HTTP stream cascades to extract security identifiers and enrich context scopes or yield early 401 failures.
func (aem *AuthExtractor) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		aem.log.Debug("AuthExtractorMiddleware.Handle start")
		defer aem.log.Debug("AuthExtractorMiddleware.Handle finish")

		// check for ignorance
		// SECURITY NOTICE: Match uses r.RequestURI which carries raw query parameters.
		// If criteria lookups expect clean endpoint prefixes, consider shifting to r.URL.Path in future updates.
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
