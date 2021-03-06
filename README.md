# Gomu
Gomu is a Golang micro-service template repository

## How-to
This repository is set as a template one.
You can click on the green button 'Use this template' on top-right or by forking it.

This README.md can be deleted to use your project's one.

## Template functionalities

- HTTP & HTTPS server
- Base MVC model (Models, Views & Controllers)
- Version & Healthcheck endpoints
- Application-wide and controller-scope middlewares
- Logging, MIME & HTTP Method middlewares
- Errors
- Flags parser
- JSON logging facility on standard output
- Mux router

## Documentation

#### Controllers
Controllers are structs with methods that handle each sub-endpoint logic like /create or /delete
If you need some additionak context, you can add fields in controller struct.
Each controller must implements `ServeHTTP()` method to satisfy the `http.Handler` interface

Exemple:
```golang
...

// HealthController struct represents the /health controller
type HealthController struct{}

// NewHealthController function returns a new HealthController struct pointer
func NewHealthController() *HealthController {
	return &HealthController{}
}

func (ctl *HealthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Write HTTP headers
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write HTTP response status
	w.WriteHeader(http.StatusOK)

	// Write payload
	fmt.Fprint(w, "healthy")
}
```

It's recommended to apply for each controller the `X-Content-Type-Options` HTTP header to prevents MIME type sniffing, for more informations see: [https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options)

#### Flags
Flags are application execution arguments. The template provides a flag parser to easily add flags.

You can specify a flag name, description, if it must be set or unique. If you want to add custome flag types, just copy/paste an existing `*_flag.go` file and adapt code, also in `FlagSet.Parse()` method.

The `Flag` interface allows arbitrary flags to be parsed at execution time.

```golang
// Flag interface implemented by application flags
type Flag interface {
	Parsed() bool
	IsUnique() bool
	IsRequired() bool
	Who() string
	Type() int
}
```

#### Middlewares
Middlewares are code snippets that runs before controller logic, to handle particular cases.

The template provides some basic middlewares like HTTP method or MIME type security.
It provides also a logging middleware to log every made request to the service.

Example:
```golang
...

// HTTPMethodMiddleware controller-scope middleware checks if a request method is allowed or not
func HTTPMethodMiddleware(next http.Handler, allowedMethods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowedMethods {
			if r.Method != method {
				http.Error(w, errors.Wrap(ierrors.ErrHTTP, fmt.Sprintf("Method %s is not allowed", method)).Error(), http.StatusMethodNotAllowed)

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
```

The notion of middleware scope intents to factorize source code and avoid long middlewares chains.
For example, the logging middleware must wraps **any** controller to log every sub-endpoint access, even good or inexistants ones. So, we wraps a middleware function into a `http.Handler` to be used with [gorilla/mux](https://github.com/gorilla/mux)

The controller-scope middlewares, are just middlewares functions to chain in router handler, if they are not needed for any sub-endpoint.

Example:
```golang
...

// Register controllers to router with controllers scope middlewares
R.Handle(createResourceView.Path(), mdw.HTTPMethodMiddleware(createResourceController, createResourceView.Methods()...))

// Register application-wide middlewares
R.Use(mdw.LoggingMiddleware(logger))

...
```

Each file must contains only one middleware and filename must be prefixed with `wide_` or `ctl_` to mark a visual difference.

#### Models
Models are simple struct that contains fields with JSON tags to be marshalled/unmarshalled

Each struct must be declared in a unique file in /internal/models

Example:
```golang
type Contact struct {
    Name    string    `json:"name"`
    Email   string    `json:"email"`
    Phone   string    `json:"phone"`
}
```

#### Views
Views are structs with two methods: Path() and Methods()

+   `Path()` method returns the sub-endpoint suffix, like /create or /{id}
+   `Methods()` method returns the allowed HTTP methods on the view sub-endpoint

#### Miscellaneous

The error handling is quite simple and you can see [this page](https://itnext.io/golang-error-handling-best-practice-a36f47b0b94c) to retreive good informations.

The package that is recommended to use is [pkg/errors](https://github.com/pkg/errors)

A template Dockerfile is provided and talk for itself, you can see it in /deployment

#### Further code
+   Template OpenAPI 3.0.x Specification
+   Template Helm chart