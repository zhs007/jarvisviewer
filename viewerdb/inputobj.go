package viewerdb

import "github.com/graphql-go/graphql"

// typeInputViewerData - InputViewerData
//		you can see assistantdb.graphql
var typeInputViewerData = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "InputViewerData",
		Fields: graphql.InputObjectConfigFieldMap{
			"type": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"token": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"graph": &graphql.InputObjectFieldConfig{
				Type: typeGraphData,
			},
		},
	},
)
