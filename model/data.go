package model

import (
	"fmt"

	"github.com/cshabsin/thegrid/example/view"
)

func ExplorersMapData() *MapData {
	systems := []struct {
		name           string
		col, row       int
		shortSystem    string
		description    string
		suppressPlanet bool
	}{
		{name: "Terschaps", col: 16, row: 12},
		{name: "Mimmar Shari", col: 16, row: 13},
		{name: "Mapiikhaar", col: 16, row: 14,
			shortSystem: "Mapiikhaar",
			description: "Pre-Imperial research station. We got data for it from Iraar Lar."},
		{name: "Iishmarmep", col: 16, row: 16},
		{name: "Riimishkin", col: 16, row: 21},

		{name: "Hasaggan", col: 17, row: 11},
		{name: "Gipuurkhir", col: 17, row: 13},
		{name: "Kuuhuurga", col: 17, row: 17},
		{name: "Forquee", col: 17, row: 19,
			shortSystem: "Forquee",
			description: "&ldquo;Leave Forquee alone!&rdquo; Learned Karma techniques " +
				"(e.g. purging Fury) from the aliens(Droyne?). Flying domed " +
				"city. Under Virus attack when we left."},
		{name: "Kemkeaanguu", col: 17, row: 21},

		{name: "Did", col: 18, row: 13,
			shortSystem: "Did",
			description: "Message to deliver from Imperial Office of Calendar " +
				"Compliance. 10 MCr reward!"},
		{name: "Arlur", col: 18, row: 14},
		{name: "Uure", col: 18, row: 18,
			shortSystem: "Uure",
			description: "Father's Lodge. Ancient crash site for Imperial archaeological " +
				"support ship. Our first lesson on Karma, Madness, and Fury.<p>" +
				"On our second pass through, Virus bombarded the archaeological " +
				"site, and the Lodge initiated a solar flare to destroy the " +
				"Virus fleet."},
		{name: "InmuuKi", col: 18, row: 20},

		{name: "Irshuushdaar", col: 19, row: 15,
			shortSystem: "Irshuushdaar",
			description: "Where the space wreck was sent to jump to, presumably by the " +
				"Marquis of Rohintash's fury infused ally."},
		{name: "Muukher", col: 19, row: 16},
		{name: "Irdar Ga", col: 19, row: 17,
			shortSystem: "Irdar_Ga",
			description: "Ships that go here don't come back, because they're eaten " +
				"by nanites. The nanites are building (have built?) something."},
		{name: "Vlair", col: 19, row: 18,
			shortSystem: "Vlair",
			description: "First time in system: Storm, Omar Factors<p>" +
				"Second time in system: Mine of madness and sarlacc maw"},
		{name: "Ziger", col: 19, row: 21},

		{name: "Gier Iir", col: 20, row: 13},
		{name: "Gimi Kuuid", col: 20, row: 17,
			shortSystem: "Gimi_Kuuid",
			description: "Underwater spherical displacement &mdash; signs of alternate 'madness' universe with dark changes (dictator vs. administrator, etc.).<p>" +
				"Gained hexagon key in a pawn shop here.<p>" +
				"Revolution underway on our way back through here on [date]. Virus is also in-system since [date].<p>" +
				"One of the planets hosting a Galactic Traveler office."},
		{name: "Garuu Uurges", col: 20, row: 18},
		{name: "Daaruugka", col: 20, row: 19},

		{name: "Girgulash", col: 21, row: 16,
			shortSystem: "Girgulash",
			description: "Trango is a moon of Girgulash, home of Trango shipyards, " +
				"headed by Messorius Thraxton. Thraines, a city on Girgulash " +
				"proper, hosted an Imperial Archive. <p>This is where we " +
				"found the message for the Office of Calendar Compliance " +
				"on Did.<p>First encounter with wireheads; they bombed a " +
				"dome on Trango. Also, our first encounter with old " +
				"Imperial computers that seemed to recognize Stefan as someone " +
				"called Admiral Rogers."},
		{name: "Khida", col: 21, row: 18,
			shortSystem: "Khida",
			description: "Stefan, Kyle Vesta, and Dr. Denmark are from here.<p>" +
				"Dr. Denmark was kidnapped here, kicking off our group's " +
				"association with him.<p>" +
				"Khida Secundus Defensive Facility, Grandma Vesta, Ling " +
				"Standard Products<p>" +
				"The last time we were here, a 'bone' ship (Madness?) jumped " +
				"in, and we helped destroy it."},
		{name: "Khui", col: 21, row: 19},

		{name: "Lis", col: 22, row: 11},
		{name: "Sham", col: 22, row: 12},
		{name: "Amem", col: 22, row: 13},
		{name: "Ugar", col: 22, row: 16,
			shortSystem: "Ugar",
			description: "Puddle of blood, murder mystery. Cetagandan/tainted " +
				"money plot. On return (15th system): Mr. Data?"},
		{name: "Vlir", col: 22, row: 17,
			shortSystem: "Vlir",
			description: "Western town with the tainted money. Earned favors from the " +
				"Tong (Star Tong?) - Stefan can ask questions where there are " +
				"contacts."},
		{name: "Duuksha", col: 22, row: 20},

		{name: "Udipeni", col: 23, row: 16,
			shortSystem: "Udipeni",
			description: "Home of Threnody and Max.<p>Bottom cult is taking hold, though " +
				"the military is resistant and Max's efforts last time through " +
				"may have helped. <p>Signs of anagathics use (and possible " +
				"withdrawl) by government officials. <p>" +
				"Found the [steampunk computer] in an old Imperium Library."},
		{name: "Nagilun", col: 23, row: 17,
			shortSystem: "Nagilun",
			description: "Home of Marian Dove.<p>" +
				"Thressalar - Museum attack, underground facility with ancient " +
				"computer, weird vines and grey goo.<p>" +
				"Kurnak, neighboring nation, has a negative tariff and embeds " +
				"surveillance nanites in all cargoes that pass through.<p>" +
				"Doctor Swede (who looks just like Doc Denmark) performed much " +
				"of the terraforming of New Nagilun. New Nagilun still has a low-" +
				"level poisonous biosphere and lethal wildlife. Screechers keep " +
				"wildlife away.<p>Most recent visit: asteroid with strange " +
				"physical properties (gravity distortions?) on collision course " +
				"with planet. Rescued researcher from kidnappers and took him " +
				"to space to help enter the asteroid and determine what to do. " +
				"Raced Cerberus there, and won."},
		{name: "Kagershi", col: 23, row: 18,
			shortSystem: "Kagershi",
			description: "Cerberus has a base set up here, taking over the system. " +
				"Staging ground for massive Cerberus fleet, but we managed to " +
				"sabotage their fuel supply, making it just impure enough to be " +
				"incompatible with their jump stabilizers.<p>In an old mining " +
				"facility contested by Cerberus and Virus-controlled robots, " +
				"we managed to sneak in and obtain an anomalous sphere of jump " +
				"space contained in real space, contained by a lanthanum ring. " +
				"This is one of the pieces we need to collect before facing the " +
				"Fracture."},
		{name: "Gowandon", col: 23, row: 19,
			shortSystem: "Gowandon",
			description: "Cerberus headquarters. Also, Ling Standard Products is " +
				"headquartered at Rhona Minor. LSP runs mining operations in " +
				"the ring, essentially using independent miners in indentured " +
				"servitude. We rescued numerous Aivar from this system.<p>" +
				"On Gowandon proper, we " +
				"explored an underwater arcology and found the Deeps, who " +
				"were especially affected by the virus ravaging the planet. " +
				"We managed to develop a cure to the virus using " +
				"Imperial-quality lab facilities, even as " +
				"the arcology itself tried to launch itself into orbit.<p>" +
				"Met up with some of Admiral Rogers (head of the " +
				"&ldquo;wizards&rdquo;) minions, and they spilled some info on " +
				"their long-term strategy.<p>Imperial Scout base also granted " +
				"us access on the strength of Stefan's identity as Admiral " +
				"Rogers. Gained more background and a skill point."},

		{name: "Uuduud", col: 24, row: 11},
		{name: "Uuruum", col: 24, row: 14},
		{name: "Irar Lar", col: 24, row: 17,
			shortSystem: "Irar_Lar",
			description: "The face in Threnody's photos turn out to be Father's " +
				"Observatory, where Father is fighting off Virus. We managed to " +
				"hook up with the friendly half of the installation and make our " +
				"way through to the data room, where Father was able to deliver " +
				"various data and useful tech for us. As we escaped, Father " +
				"blew the installation, and <a href=\"" +
				"http://digitalblasphemy.com/cgi-bin/mobilev.cgi?i=aftermath2k141&r=640x480" +
				"\">the whole planet</a>. We had a LSP " +
				"observer along on our visit, their IT guy. It's unclear how " +
				"complete his report will be to them.<p>" +
				"Mary (Marian's doppelganger) went to Irar Lar on the Long " +
				"Shot to look for clues to where the Fracture will be."},
		{name: "Kuundin", col: 24, row: 18,
			shortSystem: "Kuundin",
			description: "According to the logs from the Corona, this is where " +
				"Anagathics come from. We found a lab ship and boarded it, " +
				"discovering that things had gone terribly wrong. We found " +
				"formulas for variants on Anagathics, but all the formulas " +
				"depend on a &ldquo;precursor&rdquo; compound that we have " +
				"no supply of.<p>" +
				"Next, we visited the planet itself, where we discovered that " +
				"the Droyne of the planet had ascended to become a planet-" +
				"wide spirit that protects the native populace.<p>" +
				"The <em>Fruitful Discovery</em>, a LSP science vessel, " +
				"jumped in and approached the planet to refuel and experiment " +
				"on the locals. We got them off the planet, then protected them " +
				"from the Virus ship that jumped into the system. We agreed to " +
				"coordinate our jumps to Irar Lar, which was both our next " +
				"destination."},
		{name: "Kamar Inag", col: 24, row: 20},
		{name: "Duumar Di", col: 24, row: 21},

		{name: "Uumbuu", col: 25, row: 12},
		{name: "Ervlan", col: 25, row: 13},
		{name: "Gaiid", col: 25, row: 14},
		{name: "Digapir", col: 25, row: 16},
		{name: "Shiirla", col: 25, row: 19},
		{name: "Dinkhaluurk", col: 25, row: 21},
		{name: "black hole", col: 19, row: 13,
			description: "Imperial observatory with Gertie was programmed not to " +
				"be able to look at this parsec.",
			suppressPlanet: true,
		},
	}
	minRow := 999
	maxRow := 0
	minCol := 999
	maxCol := 0
	for _, sys := range systems {
		if minRow > sys.row {
			minRow = sys.row
		}
		if minCol > sys.col {
			minCol = sys.col
		}
		if maxRow < sys.row {
			maxRow = sys.row
		}
		if maxCol < sys.col {
			maxCol = sys.col
		}
	}
	fmt.Println("maxRow:", maxRow)
	fmt.Println("maxCol:", maxCol)
	fmt.Println("minRow:", minRow)
	fmt.Println("minCol:", minCol)
	numRows := maxRow - minRow + 1
	numCols := maxCol - minCol + 1
	fmt.Println("numRows", numRows, "; numCols", numCols)
	hexGrid := make([][]view.Entity, numCols)
	for col := range hexGrid {
		hexGrid[col] = make([]view.Entity, numRows)
		for row := range hexGrid[col] {
			hexGrid[col][row] = emptySystem{sysCol: col + minCol, sysRow: row + minRow}
		}
	}
	for _, sys := range systems {
		sysData := &systemData{
			name:           sys.name,
			sysCol:         sys.col,
			sysRow:         sys.row,
			description:    sys.description,
			suppressPlanet: sys.suppressPlanet,
		}
		if sys.shortSystem != "" {
			sysData.href = "http://scripts.mit.edu/~ringrose/explorers/index.php?title=" + sys.shortSystem
		}
		hexGrid[sys.col-minCol][sys.row-minRow] = sysData
	}
	mapData := &MapData{
		FirstCol: minCol,
		FirstRow: minRow,
		HexGrid:  hexGrid,
	}

	return mapData
}
