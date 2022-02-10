package main

import(
	"context"
	"fmt"

	"github.com/skypies/util/gcp/ds"
	"github.com/skypies/util/gcp/singleton"

	"github.com/abworrall/grapher/pkg/gamegraph"
)

var(
	GoogleCloudProjectId = "worrall-io"
	SingletonName = "gamegraph"
)

func GetSProvider(ctx context.Context) singleton.SingletonProvider {
	dsProvider, err := ds.NewCloudDSProvider(ctx, GoogleCloudProjectId)
	if err != nil {
		panic(fmt.Errorf("GetSProvider: could not get a clouddsprovider (projectId=%s): %v\n", GoogleCloudProjectId, err))
	}

	return singleton.NewProvider(dsProvider)
}

func ReadGameGraph(ctx context.Context) *gamegraph.GameGraph {
	sp := GetSProvider(ctx)
	gg := gamegraph.New()
	
	if err := sp.ReadSingleton (ctx, SingletonName, nil, gg); err != nil {
		panic(fmt.Errorf("ReadGameGraph: could not readsingleton from project %s: %v\n", GoogleCloudProjectId, err))
	}

	/*gg.Add("A", 1, 71, 100)
	gg.Add("D",  1, 71, 35, 100)
	gg.Add("M",  1, 71, 23, 1, 57)
	gg.Add("P",  1, 35, 100)*/
	
	return gg
}

func WriteGameGraph(ctx context.Context, gg *gamegraph.GameGraph) {
	sp := GetSProvider(ctx)

	if err := sp.WriteSingleton (ctx, SingletonName, nil, gg); err != nil {
		panic(fmt.Errorf("WriteGameGraph: could not writesingleton from project %s: %v\n", GoogleCloudProjectId, err))
	}
}

