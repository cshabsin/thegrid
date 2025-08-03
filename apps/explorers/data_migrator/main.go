package main

import (
	"context"
	"fmt"

	"github.com/cshabsin/thegrid/apps/explorers/data"
	"github.com/cshabsin/thegrid/apps/explorers/dataservice"
)

func main() {
	ctx := context.Background()
	dataSvc, err := dataservice.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create dataservice client:", err)
		return
	}
	defer dataSvc.Close()

	fmt.Println("Migrating systems...")
	for _, sys := range data.ExplorersMapData.Systems {
		if err := dataSvc.AddSystem(ctx, &sys); err != nil {
			fmt.Printf("Failed to add system %s: %v\n", sys.Name, err)
		}
	}

	fmt.Println("Migrating paths...")
	for _, path := range data.ExplorersPathData.Segments {
		if err := dataSvc.AddPath(ctx, &path); err != nil {
			fmt.Printf("Failed to add path %s: %v\n", path.Name, err)
		}
	}

	fmt.Println("Migration complete.")
}
