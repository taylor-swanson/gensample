package web

import (
	"net/http"
	"strconv"

	"github.com/elastic/go-ucfg"

	"gensample/internal/context"
	"gensample/internal/generator"
)

var (
	httpMethods = [...]string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
	}
	httpStatuses = [...]int{
		http.StatusContinue,
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent,
		http.StatusPartialContent,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusNotModified,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusRequestEntityTooLarge,
		http.StatusTeapot,
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
	}
	httpVersions = [...]string{
		"HTTP/1.0",
		"HTTP/1.1",
		"HTTP/2",
	}
	userAgents = [...]string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:98.0) Gecko/20100101 Firefox/98.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3; rv:98.0) Gecko/20100101 Firefox/98.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/98.0 Mobile/15E148 Safari/605.1.15",
		"Mozilla/5.0 (Android 12; Mobile; rv:68.0) Gecko/68.0 Firefox/98.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/99.0.4844.59 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.88 Mobile Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Safari/605.1.15",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPad; CPU OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Mobile/15E148 Safari/604.1",
	}
)

const (
	NameHTTPStatus  = "http_status"
	NameHTTPMethod  = "http_method"
	NameHTTPVersion = "http_version"
	NameUserAgent   = "user_agent"
)

type httpStatus struct{}

func (g *httpStatus) Generate(ctx *context.Context) string {
	return strconv.Itoa(httpStatuses[ctx.Rand.IntN(len(httpStatuses))])
}

func NewHTTPStatusGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	return &httpStatus{}, nil
}

type httpMethod struct{}

func (g *httpMethod) Compile() error {
	return nil
}

func (g *httpMethod) Generate(ctx *context.Context) string {
	return httpMethods[ctx.Rand.IntN(len(httpMethods))]
}

func NewHTTPMethodGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	return &httpMethod{}, nil
}

type httpVersion struct{}

func (g *httpVersion) Compile() error {
	return nil
}

func (g *httpVersion) Generate(ctx *context.Context) string {
	return httpVersions[ctx.Rand.IntN(len(httpVersions))]
}

func NewHTTPVersionGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	return &httpVersion{}, nil
}

type userAgent struct{}

func (g *userAgent) Generate(ctx *context.Context) string {
	return userAgents[ctx.Rand.IntN(len(userAgents))]
}

func NewUserAgentGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	return &userAgent{}, nil
}

// TODO: URLGenerator

func init() {
	generator.Register(NameHTTPStatus, NewHTTPStatusGenerator)
	generator.Register(NameHTTPMethod, NewHTTPMethodGenerator)
	generator.Register(NameHTTPVersion, NewHTTPVersionGenerator)
	generator.Register(NameUserAgent, NewUserAgentGenerator)
}
