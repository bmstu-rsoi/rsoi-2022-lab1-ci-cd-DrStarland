package route

import (
	"github.com/DrStarland/probagoyave/http/controller/hello"
	"github.com/DrStarland/probagoyave/http/controller/person"

	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/cors"
	"goyave.dev/openapi3"
)

// Routing is an essential part of any Goyave application.
// Routes definition is the action of associating a URI, sometimes having
// parameters, with a handler which will process the request and respond to it.

// Routes are defined in routes registrer functions.
// The main route registrer is passed to "goyave.Start()" and is executed
// automatically with a newly created root-level router.

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {
	// Applying default CORS settings (allow all methods and all origins)
	// Learn more about CORS options here: https://goyave.dev/guide/advanced/cors.html
	router.CORS(cors.Default())

	// Register your routes here

	// Route without validation
	router.Get("/hello/{name}", hello.SayHi)

	// Route with validation
	router.Post("/echo", hello.Echo).Validate(hello.EchoRequest)

	subrouter := router.Subrouter("/api/v1")

	subrouter.Get("/persons", person.Index)
	subrouter.Get("/persons/{personID}", person.Show)
	subrouter.Post("/persons", person.Store).Validate(person.PersonRequest)
	subrouter.Patch("/persons/{personID}", person.Update).Validate(person.PatchRequest)
	subrouter.Delete("/persons/{personID}", person.Destroy)
	// GET /persons/{personId} – информация о человеке;
	// GET /persons – информация по всем людям;
	// POST /persons – создание новой записи о человеке;
	// PATCH /persons/{personId} – обновление существующей записи о человеке;
	// DELETE /person/{personId} – удаление записи о человеке.

	spec := openapi3.NewGenerator().Generate(router)
	opts := openapi3.NewUIOptions(spec)
	openapi3.Serve(router, "/openapi", opts)

}
