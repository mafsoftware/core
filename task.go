package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Task encapsulates the data needed to perform a task.
type Task[T any] struct {
	Input    T
	Request  *http.Request
	Params   map[string]string
	UserInfo UserInfo
}

// TaskVoid is used in place of Task for requests that do not have a body (e.g. HTTP GET, etc.)
type TaskVoid struct {
	Request  *http.Request
	Params   map[string]string
	UserInfo UserInfo
}

// Query invokes Task's Request.URL.Query().
func (t TaskVoid) Query() map[string][]string { return t.Request.URL.Query() }

// Query invokes Task's Request.URL.Query().
func (t Task[T]) Query() map[string][]string { return t.Request.URL.Query() }

// TaskOptions contains options for the given task.
type TaskOptions struct {
	Access            Access
	RequiresSecretKey bool
	RefreshingToken   bool
}

// AccessOptions contains frequently used TaskOptions.
type AccessOptions struct {
	Open        TaskOptions
	User        TaskOptions
	Admin       TaskOptions
	SecretAdmin TaskOptions
	SecretOwner TaskOptions
}

type TaskHandleFunc[I, O any] func(input Task[I]) (*O, *Error)

type TaskVoidHandleFunc[O any] func(input TaskVoid) (*O, *Error)

type Validateable interface {
	Validate(UserInfo) *Error
}

// AddTask registers task to be executed when API endpoint is reached.
func AddTask[I Validateable, O any](app *App, method, path string, options TaskOptions, fn TaskHandleFunc[I, O]) {
	app.Router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logAppServe(app.Build, method, path)
		userInfo, err := app.validateRequestWithOptions(r, options)
		if err != nil {
			writeError(w, err)
			return
		}
		input, err := decodeFrom[I](r)
		if err != nil {
			writeError(w, err)
			return
		}
		err = (*input).Validate(userInfo)
		if err != nil {
			writeError(w, err)
			return
		}
		task := Task[I]{*input, r, paramsFrom(ps), userInfo}
		output, err := fn(task)
		if err != nil {
			writeError(w, err)
			return
		}
		writeSuccess(w, output)
	})
}

// AddTaskVoid registers a get-only task (no input data) to be executed when API endpoint is reached.
func AddTaskVoid[O any](app *App, method, path string, options TaskOptions, fn TaskVoidHandleFunc[O]) {
	app.Router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logAppServe(app.Build, method, path)
		userInfo, err := app.validateRequestWithOptions(r, options)
		if err != nil {
			writeError(w, err)
			return
		}
		task := TaskVoid{r, paramsFrom(ps), userInfo}
		output, err := fn(task)
		if err != nil {
			writeError(w, err)
			return
		}
		writeSuccess(w, output)
	})
}

// Convenience

// TaskGET is a convenience shortcut for AddTaskVoid.
func TaskGET[O any](app *App, path string, options TaskOptions, fn TaskVoidHandleFunc[O]) {
	AddTaskVoid(app, http.MethodGet, path, options, fn)
}

// TaskPOST is a convenience shortcut for AddTask.
func TaskPOST[I Validateable, O any](app *App, path string, options TaskOptions, fn TaskHandleFunc[I, O]) {
	AddTask(app, http.MethodPost, path, options, fn)
}

// TaskPUT is a convenience shortcut for AddTask.
func TaskPUT[I Validateable, O any](app *App, path string, options TaskOptions, fn TaskHandleFunc[I, O]) {
	AddTask(app, http.MethodPut, path, options, fn)
}

// TaskDELETE is a convenience shortcut for AddTask.
func TaskDELETE[I Validateable, O any](app *App, path string, options TaskOptions, fn TaskHandleFunc[I, O]) {
	AddTask(app, http.MethodDelete, path, options, fn)
}

// Utils

func logAppServe(build, method, path string) {
	log.Println("com.mafsoftware.carbon:" + build + " : " + method + path)
}

func paramsFrom(ps httprouter.Params) map[string]string {
	params := make(map[string]string)
	for _, v := range ps {
		params[v.Key] = v.Value
	}
	return params
}

func decodeFrom[T any](r *http.Request) (*T, *Error) {
	var request *T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return nil, newError(ErrorCodeIncorrectRequestBody)

	}
	return request, nil
}

func writeSuccess(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func writeError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	_, _ = fmt.Fprintln(w, err.JSON())
}

/*
EXAMPLE OF HOW TO USE DYNAMIC MIDDLEWARE

func Middleware[I, O any](fn TaskHandleFunc[I, O]) TaskHandleFunc[I, O] {
	return func(input Task[I]) (*O, *Error) {
		// return nil, newError("es-419", ErrorCodeAccessTokenMalformed)
		return fn(input)
	}
}
*/
