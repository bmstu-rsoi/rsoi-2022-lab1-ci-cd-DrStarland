package person

import "goyave.dev/goyave/v4/validation"

// Goyave provides a powerful, yet easy way to validate all incoming data, no matter
// its type or its format, thanks to a large number of validation rules.

// Incoming requests are validated using rules set, which associate rules
// with each expected field in the request.

// Learn more about validation here: https://goyave.dev/guide/basics/validation.html

// This is the validation rules for the "/echo" route, which is simply
// writing the input as a response.
var (
	PersonRequest = validation.RuleSet{
		"name":    validation.List{"required", "string"},
		"age":     validation.List{"required", "integer", "between:1,150"},
		"address": validation.List{"required", "string"},
		"work":    validation.List{"required", "string"},
	}
	PatchRequest = validation.RuleSet{
		"name":    validation.List{"nullable", "string"},
		"age":     validation.List{"nullable", "integer", "between:1,150"},
		"address": validation.List{"nullable", "string"},
		"work":    validation.List{"nullable", "string"},
	}
)
