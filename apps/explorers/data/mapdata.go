package data

type SystemData struct {
	Name           string `json:"name"`
	SysRow         int    `json:"sys_row"`
	SysCol         int    `json:"sys_col"`
	ShortSystem    string `json:"short_system"`
	Description    string `json:"description"`
	SuppressPlanet bool   `json:"suppress_planet"`
}

type MapData struct {
	MinCol  int          `json:"min_col"`
	MinRow  int          `json:"min_row"`
	MaxCol  int          `json:"max_col"`
	MaxRow  int          `json:"max_row"`
	Systems []SystemData `json:"systems"`
}

var ExplorersMapData = makeExplorersMapData()

func makeExplorersMapData() *MapData {
	systems := []SystemData{
		{Name: "Terschaps", SysCol: 16, SysRow: 12},
		{Name: "Mimmar Shari", SysCol: 16, SysRow: 13},
		{Name: "Mapiikhaar", SysCol: 16, SysRow: 14,
			ShortSystem: "Mapiikhaar",
			Description: "Pre-Imperial research station. We got data for it from Iraar Lar."},
		{Name: "Iishmarmep", SysCol: 16, SysRow: 16},
		{Name: "Riimishkin", SysCol: 16, SysRow: 21},

		{Name: "Hasaggan", SysCol: 17, SysRow: 11},
		{Name: "Gipuurkhir", SysCol: 17, SysRow: 13},
		{Name: "Kuuhuurga", SysCol: 17, SysRow: 17},
		{Name: "Forquee", SysCol: 17, SysRow: 19,
			ShortSystem: "Forquee",
			Description: "&ldquo;Leave Forquee alone!&rdquo; Learned Karma techniques " +
				"(e.g. purging Fury) from the aliens(Droyne?). Flying domed " +
				"city. Under Virus attack when we left."},
		{Name: "Kemkeaanguu", SysCol: 17, SysRow: 21},

		{Name: "Did", SysCol: 18, SysRow: 13,
			ShortSystem: "Did",
			Description: "Message to deliver from Imperial Office of Calendar " +
				"Compliance. 10 MCr reward!"},
		{Name: "Arlur", SysCol: 18, SysRow: 14},
		{Name: "Uure", SysCol: 18, SysRow: 18,
			ShortSystem: "Uure",
			Description: "Father's Lodge. Ancient crash site for Imperial archaeological " +
				"support ship. Our first lesson on Karma, Madness, and Fury.<p>" +
				"On our second pass through, Virus bombarded the archaeological " +
				"site, and the Lodge initiated a solar flare to destroy the " +
				"Virus fleet."},
		{Name: "InmuuKi", SysCol: 18, SysRow: 20},

		{Name: "Irshuushdaar", SysCol: 19, SysRow: 15,
			ShortSystem: "Irshuushdaar",
			Description: "Where the space wreck was sent to jump to, presumably by the " +
				"Marquis of Rohintash's fury infused ally."},
		{Name: "Muukher", SysCol: 19, SysRow: 16},
		{Name: "Irdar Ga", SysCol: 19, SysRow: 17,
			ShortSystem: "Irdar_Ga",
			Description: "Ships that go here don't come back, because they're eaten " +
				"by nanites. The nanites are building (have built?) something."},
		{Name: "Vlair", SysCol: 19, SysRow: 18,
			ShortSystem: "Vlair",
			Description: "First time in system: Storm, Omar Factors<p>" +
				"Second time in system: Mine of madness and sarlacc maw"},
		{Name: "Ziger", SysCol: 19, SysRow: 21},

		{Name: "Gier Iir", SysCol: 20, SysRow: 13},
		{Name: "Gimi Kuuid", SysCol: 20, SysRow: 17,
			ShortSystem: "Gimi_Kuuid",
			Description: "Underwater spherical displacement &mdash; signs of alternate 'madness' universe with dark changes (dictator vs. administrator, etc.).<p>" +
				"Gained hexagon key in a pawn shop here.<p>" +
				"Revolution underway on our way back through here on [date]. Virus is also in-system since [date].<p>" +
				"One of the planets hosting a Galactic Traveler office."},
		{Name: "Garuu Uurges", SysCol: 20, SysRow: 18},
		{Name: "Daaruugka", SysCol: 20, SysRow: 19},

		{Name: "Girgulash", SysCol: 21, SysRow: 16,
			ShortSystem: "Girgulash",
			Description: "Trango is a moon of Girgulash, home of Trango shipyards, " +
				"headed by Messorius Thraxton. Thraines, a city on Girgulash " +
				"proper, hosted an Imperial Archive. <p>This is where we " +
				"found the message for the Office of Calendar Compliance " +
				"on Did.<p>First encounter with wireheads; they bombed a " +
				"dome on Trango. Also, our first encounter with old " +
				"Imperial computers that seemed to recognize Stefan as someone " +
				"called Admiral Rogers."},
		{Name: "Khida", SysCol: 21, SysRow: 18,
			ShortSystem: "Khida",
			Description: "Stefan, Kyle Vesta, and Dr. Denmark are from here.<p>" +
				"Dr. Denmark was kidnapped here, kicking off our group's " +
				"association with him.<p>" +
				"Khida Secundus Defensive Facility, Grandma Vesta, Ling " +
				"Standard Products<p>" +
				"The last time we were here, a 'bone' ship (Madness?) jumped " +
				"in, and we helped destroy it."},
		{Name: "Khui", SysCol: 21, SysRow: 19},

		{Name: "Lis", SysCol: 22, SysRow: 11},
		{Name: "Sham", SysCol: 22, SysRow: 12},
		{Name: "Amem", SysCol: 22, SysRow: 13},
		{Name: "Ugar", SysCol: 22, SysRow: 16,
			ShortSystem: "Ugar",
			Description: "Puddle of blood, murder mystery. Cetagandan/tainted " +
				"money plot. On return (15th system): Mr. Data?"},
		{Name: "Vlir", SysCol: 22, SysRow: 17,
			ShortSystem: "Vlir",
			Description: "Western town with the tainted money. Earned favors from the " +
				"Tong (Star Tong?) - Stefan can ask questions where there are " +
				"contacts."},
		{Name: "Duuksha", SysCol: 22, SysRow: 20},

		{Name: "Udipeni", SysCol: 23, SysRow: 16,
			ShortSystem: "Udipeni",
			Description: "Home of Threnody and Max.<p>Bottom cult is taking hold, though " +
				"the military is resistant and Max's efforts last time through " +
				"may have helped. <p>Signs of anagathics use (and possible " +
				"withdrawl) by government officials. <p>" +
				"Found the [steampunk computer] in an old Imperium Library."},
		{Name: "Nagilun", SysCol: 23, SysRow: 17,
			ShortSystem: "Nagilun",
			Description: "Home of Marian Dove.<p>" +
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
		{Name: "Kagershi", SysCol: 23, SysRow: 18,
			ShortSystem: "Kagershi",
			Description: "Cerberus has a base set up here, taking over the system. " +
				"Staging ground for massive Cerberus fleet, but we managed to " +
				"sabotage their fuel supply, making it just impure enough to be " +
				"incompatible with their jump stabilizers.<p>In an old mining " +
				"facility contested by Cerberus and Virus-controlled robots, " +
				"we managed to sneak in and obtain an anomalous sphere of jump " +
				"space contained in real space, contained by a lanthanum ring. " +
				"This is one of the pieces we need to collect before facing the " +
				"Fracture."},
		{Name: "Gowandon", SysCol: 23, SysRow: 19,
			ShortSystem: "Gowandon",
			Description: "Cerberus headquarters. Also, Ling Standard Products is " +
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

		{Name: "Uuduud", SysCol: 24, SysRow: 11},
		{Name: "Uuruum", SysCol: 24, SysRow: 14},
		{Name: "Irar Lar", SysCol: 24, SysRow: 17,
			ShortSystem: "Irar_Lar",
			Description: "The face in Threnody's photos turn out to be Father's " +
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
		{Name: "Kuundin", SysCol: 24, SysRow: 18,
			ShortSystem: "Kuundin",
			Description: "According to the logs from the Corona, this is where " +
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
		{Name: "Kamar Inag", SysCol: 24, SysRow: 20},
		{Name: "Duumar Di", SysCol: 24, SysRow: 21},

		{Name: "Uumbuu", SysCol: 25, SysRow: 12},
		{Name: "Ervlan", SysCol: 25, SysRow: 13},
		{Name: "Gaiid", SysCol: 25, SysRow: 14},
		{Name: "Digapir", SysCol: 25, SysRow: 16},
		{Name: "Shiirla", SysCol: 25, SysRow: 19},
		{Name: "Dinkhaluurk", SysCol: 25, SysRow: 21},
		{Name: "black hole", SysCol: 19, SysRow: 13,
			Description: "Imperial observatory with Gertie was programmed not to " +
				"be able to look at this parsec.",
			SuppressPlanet: true,
		},
	}
	minRow := systems[0].SysRow
	maxRow := systems[0].SysRow
	minCol := systems[0].SysCol
	maxCol := systems[0].SysCol

	for _, sys := range systems {
		if sys.SysRow < minRow {
			minRow = sys.SysRow
		}
		if sys.SysRow > maxRow {
			maxRow = sys.SysRow
		}
		if sys.SysCol < minCol {
			minCol = sys.SysCol
		}
		if sys.SysCol > maxCol {
			maxCol = sys.SysCol
		}
	}

	return &MapData{
		MinRow:  minRow,
		MaxRow:  maxRow,
		MinCol:  minCol,
		MaxCol:  maxCol,
		Systems: systems,
	}
}
