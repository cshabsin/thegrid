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
	hexes := map[string][2]int{
		"Khida":      {21, 18},
		"Gimi Kuuid": {20, 17},
		"Vlair":      {19, 18},
		"Uure":       {18, 18},
		"Forquee":    {17, 19},
		"Vlir":       {22, 17},
		"Nagilun":    {23, 17},
		"Udipeni":    {23, 16},
		"Ugar":       {22, 16},
		"Girgulash":  {21, 16},
		"Kagershi":   {23, 18},
		"Gowandon":   {23, 19},
		"Kuundin":    {24, 18},
		"IrarLar":    {24, 17},
	}

	var pathData PathData
	add := func(startCoord [2]int, startOffsetX, startOffsetY int, endCoord [2]int, endOffsetX, endOffsetY int) {
		if startCoord == [2]int{0, 0} {
			panic("startCoord is nil")
		}
		if endCoord == [2]int{0, 0} {
			panic("endCoord is nil")
		}
		startCoord[0] -= mapData.MinCol // TODO: normalize the coordinate systems somehow
		endCoord[0] -= mapData.MinCol
		startCoord[1] -= mapData.MinRow
		endCoord[1] -= mapData.MinRow
		pathData.Segments = append(pathData.Segments, PathSegment{
			StartCoord:  startCoord,
			StartOffset: [2]int{startOffsetX, startOffsetY},
			EndCoord:    endCoord,
			EndOffset:   [2]int{endOffsetX, endOffsetY},
		})
	}

	add(hexes["Khida"], -10, -4, hexes["Gimi Kuuid"], 0, 32)
	add(hexes["Gimi Kuuid"], 0, 32, hexes["Vlair"], 0, -12)
	add(hexes["Vlair"], 0, -12, hexes["Uure"], 0, -12)
	add(hexes["Uure"], 0, -12, hexes["Forquee"], 0, -12)
	add(hexes["Forquee"], 0, -12, hexes["Uure"], -20, -20)
	add(hexes["Uure"], -20, -20, hexes["Vlair"], 0, -35)
	add(hexes["Vlair"], 0, -35, hexes["Gimi Kuuid"], -5, -15)
	add(hexes["Gimi Kuuid"], -5, -15, hexes["Khida"], 0, -20)
	add(hexes["Khida"], 0, -20, hexes["Vlir"], 0, -20)
	add(hexes["Vlir"], 0, -20, hexes["Nagilun"], 0, -20)
	add(hexes["Nagilun"], 0, -20, hexes["Udipeni"], -23, 23)
	add(hexes["Udipeni"], -23, 23, hexes["Ugar"], 0, -10)
	add(hexes["Ugar"], 0, -10, hexes["Girgulash"], 25, 0)
	add(hexes["Girgulash"], 25, 0, hexes["Ugar"], -15, 30)
	add(hexes["Ugar"], -15, 30, hexes["Nagilun"], -10, 5)
	add(hexes["Nagilun"], -10, 5, hexes["Kagershi"], -10, 0)
	add(hexes["Kagershi"], -10, 0, hexes["Gowandon"], -10, -10)
	add(hexes["Gowandon"], -10, -10, hexes["Kuundin"], -10, 0)
	add(hexes["Kuundin"], -10, 0, hexes["IrarLar"], -10, 0)
	add(hexes["IrarLar"], -10, 0, hexes["Nagilun"], 20, 0)

	return &pathData
}
