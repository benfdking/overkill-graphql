// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Product struct {
		ID    func(childComplexity int) int
		Name  func(childComplexity int) int
		Price func(childComplexity int) int
	}

	Query struct {
		__resolve__service func(childComplexity int) int
	}

	User struct {
		ID        func(childComplexity int) int
		Purchases func(childComplexity int) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Product.id":
		if e.complexity.Product.ID == nil {
			break
		}

		return e.complexity.Product.ID(childComplexity), true

	case "Product.name":
		if e.complexity.Product.Name == nil {
			break
		}

		return e.complexity.Product.Name(childComplexity), true

	case "Product.price":
		if e.complexity.Product.Price == nil {
			break
		}

		return e.complexity.Product.Price(childComplexity), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.purchases":
		if e.complexity.User.Purchases == nil {
			break
		}

		return e.complexity.User.Purchases(childComplexity), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "graph/products.graphql", Input: `type Product {
  id: ID!
  name: String!
  price: String!
}

type User {
  id: ID!
  purchases: [Product]
}
`, BuiltIn: false},
	{Name: "federation/directives.graphql", Input: `
scalar _Any
scalar _FieldSet

directive @external on FIELD_DEFINITION
directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
directive @key(fields: _FieldSet!) repeatable on OBJECT | INTERFACE
directive @extends on OBJECT | INTERFACE
`, BuiltIn: true},
	{Name: "federation/entity.graphql", Input: `
type _Service {
  sdl: String
}

extend type Query {
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
