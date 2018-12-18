package viewerdb

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/zhs007/ankadb"
)

// viewerDB -
type viewerDB struct {
	schema graphql.Schema
}

// newDBLogic -
func newDBLogic() ankadb.DBLogic {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    typeQuery,
			Mutation: typeMutation,
			// Types:    curTypes,
		},
	)

	return &viewerDB{
		schema: schema,
	}
}

// OnQuery -
func (logic *viewerDB) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:         logic.schema,
		RequestString:  request,
		VariableValues: values,
		Context:        ctx,
	})
	// if len(result.Errors) > 0 {
	// 	fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	// }

	return result, nil
}

// OnQueryStream -
func (logic *viewerDB) OnQueryStream(ctx context.Context, request string, values map[string]interface{}, funcOnQueryStream ankadb.FuncOnQueryStream) error {
	return nil
}
