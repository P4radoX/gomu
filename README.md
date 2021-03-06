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

#### Models

#### Views

#### Miscellaneous
