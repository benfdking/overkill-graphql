// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) __Service_sdl(ctx context.Context, field graphql.CollectedField, obj *fedruntime.Service) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "_Service",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SDL, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var _ServiceImplementors = []string{"_Service"}

func (ec *executionContext) __Service(ctx context.Context, sel ast.SelectionSet, obj *fedruntime.Service) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, _ServiceImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("_Service")
		case "sdl":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec.__Service_sdl(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalN_Service2githubᚗcomᚋ99designsᚋgqlgenᚋpluginᚋfederationᚋfedruntimeᚐService(ctx context.Context, sel ast.SelectionSet, v fedruntime.Service) graphql.Marshaler {
	return ec.__Service(ctx, sel, &v)
}

// endregion ***************************** type.gotpl *****************************
