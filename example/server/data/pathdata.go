package data

type PathData struct {
	Segments []PathSegment
}

type PathSegment struct {
	StartCoord  [2]int
	StartOffset [2]int
	EndCoord    [2]int
	EndOffset   [2]int
}

var ExplorersPathData = makePathData(ExplorersMapData)

func makePathData(mapData *MapData) *PathData {
	hexes := map[string][2]int{}
	for _, sys := range mapData.Systems {
		hexes[sys.Name] = [2]int{sys.SysCol - mapData.MinCol, sys.SysRow - mapData.MinRow}
	}

	var pathData PathData
	add := func(start string, startOffsetX, startOffsetY int, end string, endOffsetX, endOffsetY int) {
		startCoord, ok := hexes[start]
		if !ok {
			panic("system " + start + " not found")
		}
		endCoord, ok := hexes[end]
		if !ok {
			panic("system " + end + " not found")
		}
		pathData.Segments = append(pathData.Segments, PathSegment{
			StartCoord:  startCoord,
			StartOffset: [2]int{startOffsetX, startOffsetY},
			EndCoord:    endCoord,
			EndOffset:   [2]int{endOffsetX, endOffsetY},
		})
	}

	add("Khida", -10, -4, "Gimi Kuuid", 0, 32)
	add("Gimi Kuuid", 0, 32, "Vlair", 0, -12)
	add("Vlair", 0, -12, "Uure", 0, -12)
	add("Uure", 0, -12, "Forquee", 0, -12)
	add("Forquee", 0, -12, "Uure", -20, -20)
	add("Uure", -20, -20, "Vlair", 0, -35)
	add("Vlair", 0, -35, "Gimi Kuuid", -5, -15)
	add("Gimi Kuuid", -5, -15, "Khida", 0, -20)
	add("Khida", 0, -20, "Vlir", 0, -20)
	add("Vlir", 0, -20, "Nagilun", 0, -20)
	add("Nagilun", 0, -20, "Udipeni", -23, 23)
	add("Udipeni", -23, 23, "Ugar", 0, -10)
	add("Ugar", 0, -10, "Girgulash", 25, 0)
	add("Girgulash", 25, 0, "Ugar", -15, 30)
	add("Ugar", -15, 30, "Nagilun", -10, 5)
	add("Nagilun", -10, 5, "Kagershi", -10, 0)
	add("Kagershi", -10, 0, "Gowandon", -10, -10)
	add("Gowandon", -10, -10, "Kuundin", -10, 0)
	add("Kuundin", -10, 0, "Irar Lar", -10, 0)
	add("Irar Lar", -10, 0, "Nagilun", 20, 0)

	return &pathData
}
