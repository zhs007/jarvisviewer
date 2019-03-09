package viewerdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	pb "github.com/zhs007/jarvisviewer/viewerdb/proto"
)

var typeMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"updViewerData": &graphql.Field{
			Type:        typeViewerData,
			Description: "update viewer data",
			Args: graphql.FieldConfigArgument{
				"dat": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(typeInputViewerData),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ankadb.ErrCtxAnkaDB
				}

				curdb := anka.GetDBMgr().GetDB("viewerdb")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				dat := &pb.ViewerData{}
				err := ankadb.GetMsgFromParam(params, "dat", dat)
				if err != nil {
					return nil, err
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeViewerDataKey(dat.Token)), dat)
				if err != nil {
					return nil, err
				}

				return dat, nil
			},
		},
	},
})
