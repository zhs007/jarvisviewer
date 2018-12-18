package viewerdb

import (
	"github.com/graphql-go/graphql"
)

// typeGraphNode - GraphNode
//		you can see assistantdb.graphql
var typeGraphNode = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GraphNode",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"size": &graphql.Field{
				Type: graphql.Int,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// typeGraphLink - GraphLink
//		you can see assistantdb.graphql
var typeGraphLink = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GraphLink",
		Fields: graphql.Fields{
			"src": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"dest": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"size": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// typeGraphData - GraphData
//		you can see assistantdb.graphql
var typeGraphData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GraphData",
		Fields: graphql.Fields{
			"nodes": &graphql.Field{
				Type: graphql.NewList(typeGraphNode),
			},
			"links": &graphql.Field{
				Type: graphql.NewList(typeGraphLink),
			},
		},
	},
)

// typeViewerData - ViewerData
//		you can see assistantdb.graphql
var typeViewerData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ViewerData",
		Fields: graphql.Fields{
			"type": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"graph": &graphql.Field{
				Type: typeGraphData,
			},
		},
	},
)
