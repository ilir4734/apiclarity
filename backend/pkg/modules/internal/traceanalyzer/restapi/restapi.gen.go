// Package restapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.1-0.20220609223533-7da811e1cf30 DO NOT EDIT.
package restapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	externalRef0 "github.com/openclarity/apiclarity/api3/common"
)

// Annotation defines model for Annotation.
type Annotation struct {
	Annotation string `json:"annotation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Severity   string `json:"severity"`
}

// Annotations defines model for Annotations.
type Annotations struct {
	Items *[]Annotation `json:"items,omitempty"`

	// Total event annotations count
	Total int `json:"total"`
}

// Redacted defines model for redacted.
type Redacted = bool

// GetApiFindingsParams defines parameters for GetApiFindings.
type GetApiFindingsParams struct {
	// Should findings include sensitive data ?
	Sensitive *externalRef0.Sensitive `form:"sensitive,omitempty" json:"sensitive,omitempty"`
}

// GetEventAnnotationsParams defines parameters for GetEventAnnotations.
type GetEventAnnotationsParams struct {
	Redacted *Redacted `form:"redacted,omitempty" json:"redacted,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get findings for an API and module
	// (GET /apiFindings/{apiID})
	GetApiFindings(w http.ResponseWriter, r *http.Request, apiID externalRef0.ApiID, params GetApiFindingsParams)
	// Delete all API findings for an API
	// (POST /apiFindings/{apiID}/reset)
	ResetApiFindings(w http.ResponseWriter, r *http.Request, apiID externalRef0.ApiID)
	// Get Annotations for an event
	// (GET /eventAnnotations/{eventID})
	GetEventAnnotations(w http.ResponseWriter, r *http.Request, eventID int64, params GetEventAnnotationsParams)
	// Start Trace Analysis for an API
	// (POST /{apiID}/start)
	StartTraceAnalysis(w http.ResponseWriter, r *http.Request, apiID externalRef0.ApiID)
	// Stop Trace Analysis for an API
	// (POST /{apiID}/stop)
	StopTraceAnalysis(w http.ResponseWriter, r *http.Request, apiID externalRef0.ApiID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetApiFindings operation middleware
func (siw *ServerInterfaceWrapper) GetApiFindings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiID" -------------
	var apiID externalRef0.ApiID

	err = runtime.BindStyledParameterWithLocation("simple", false, "apiID", runtime.ParamLocationPath, chi.URLParam(r, "apiID"), &apiID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiID", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiFindingsParams

	// ------------- Optional query parameter "sensitive" -------------
	if paramValue := r.URL.Query().Get("sensitive"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sensitive", r.URL.Query(), &params.Sensitive)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sensitive", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiFindings(w, r, apiID, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ResetApiFindings operation middleware
func (siw *ServerInterfaceWrapper) ResetApiFindings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiID" -------------
	var apiID externalRef0.ApiID

	err = runtime.BindStyledParameterWithLocation("simple", false, "apiID", runtime.ParamLocationPath, chi.URLParam(r, "apiID"), &apiID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ResetApiFindings(w, r, apiID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetEventAnnotations operation middleware
func (siw *ServerInterfaceWrapper) GetEventAnnotations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "eventID" -------------
	var eventID int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventID", runtime.ParamLocationPath, chi.URLParam(r, "eventID"), &eventID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "eventID", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEventAnnotationsParams

	// ------------- Optional query parameter "redacted" -------------
	if paramValue := r.URL.Query().Get("redacted"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "redacted", r.URL.Query(), &params.Redacted)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "redacted", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetEventAnnotations(w, r, eventID, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// StartTraceAnalysis operation middleware
func (siw *ServerInterfaceWrapper) StartTraceAnalysis(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiID" -------------
	var apiID externalRef0.ApiID

	err = runtime.BindStyledParameterWithLocation("simple", false, "apiID", runtime.ParamLocationPath, chi.URLParam(r, "apiID"), &apiID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.StartTraceAnalysis(w, r, apiID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// StopTraceAnalysis operation middleware
func (siw *ServerInterfaceWrapper) StopTraceAnalysis(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiID" -------------
	var apiID externalRef0.ApiID

	err = runtime.BindStyledParameterWithLocation("simple", false, "apiID", runtime.ParamLocationPath, chi.URLParam(r, "apiID"), &apiID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.StopTraceAnalysis(w, r, apiID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/apiFindings/{apiID}", wrapper.GetApiFindings)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apiFindings/{apiID}/reset", wrapper.ResetApiFindings)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/eventAnnotations/{eventID}", wrapper.GetEventAnnotations)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/{apiID}/start", wrapper.StartTraceAnalysis)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/{apiID}/stop", wrapper.StopTraceAnalysis)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWUU/bQAz+Kydvj1HSDbSHvnWDTX2YhIA3xINJnHKQ3B2+S6Wsyn+f7pI2aZMB09AE",
	"ElKlJGef/dn+bHcDqS6NVqSchfkGDDKW5IjDF1OGqaPMv0sFc3ioiGuIQGFJMO/lEdj0lkr0ihnlWBUO",
	"5jkWliJwtfG6N1oXhAqaptlqBx8LpbRDJ7UK/lkbYicpyHBP1hmyjqVaQRPBvVTZpKCFNyGwtCaWrp4Q",
	"NhEwPVSSfbhXILNtmNEQxsBE5/96F6G+uaPUeTd9THYclHRU7r98ZMphDh+SvhZJl6FkkJ5m5wmZsQ7f",
	"2mHRJt2mLE2bKrj0x4LWpJzowVuR6ko52JmRytGKeBR7a3UcmNeTKtdjh4uz5bcCfVrEJWNKYqGwqH8R",
	"i586qwoSi7Ol9ytdQc9QhwjWxLa1PYtn8ScfrDak0EiYw1E8i48gAoPuNmQwQSO/S5VJtbLJBo1cnjT+",
	"fEVuAqy9tyLXLPLuhtC5QCUGqMoAIxKoMn8ci2VpCipJOcrETS0I09tOCQIwDhleZjCHH+QWPZqAsu+p",
	"q66RPPK+jwJgGBbBcUXDpuooEsfJwQ+NPPK8KbVKugTFNZbFNJmCn6aJ/tVgH1JyQcpKJ9cEzbWPwBqt",
	"bMv0z7OZf6RaOVKhEGhMIdOQq+TOtm394kGeLXfJD5Q9qH4otNiqiK+VygqCoNcNrv8P2cjzLnFTkE+Z",
	"NQvuNSKwVVki1y3deiJ7VrdMDtTtGOovTHVIwmTbBjHahuc+k8+9+LVzeUy643HLh0jebo1PqCBHAosi",
	"lHai3G2Nw8QfLJ9kE04eG4aePgczeGCg9SCsoVTmMm1XytTEOz3w/CyqdOAeJUuuuUTXLqsvx5O7a3qB",
	"DkbU7m/KS02oJzb1H+ZOv8hHDTxKuepSHcq6bVfrkPfadd/DhRcPi2nlkCLxqG7hQtDfqr+RDp+9vv69",
	"qNKUrH27M+YJ9hwyUZvHiKjN3/FQm3cavtOwfpI73lzzOwAA///F+U6MtA4AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../../../../../../api3/common/openapi.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}