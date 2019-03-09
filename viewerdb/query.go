package viewerdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	pb "github.com/zhs007/jarvisviewer/viewerdb/proto"
)

var typeQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"viewerData": &graphql.Field{
				Type: typeViewerData,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
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

					token := params.Args["token"].(string)

					curkey := []byte(makeViewerDataKey(token))

					has, err := curdb.Has(curkey)
					if err != nil {
						return nil, err
					}

					dat := &pb.ViewerData{}

					if has {
						err := ankadb.GetMsgFromDB(curdb, curkey, dat)
						if err != nil {
							return nil, err
						}
					}

					return dat, nil
				},
			},
		},
	},
)
