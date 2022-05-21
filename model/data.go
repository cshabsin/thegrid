package model

type systemData struct {
	name           string
	relRow, relCol int
	description    string
	hasStar        bool
}

var AllSystems = []systemData{
	{name: "Terschaps", relRow: 1, relCol: 0},
	{name: "Mimmar Shari", relRow: 2, relCol: 0},
	{
		name: "Mapiikhaar", relRow: 3, relCol: 0,
		description: "Pre-Imperial research station. We got data for it from Iraar Lar.",
	},
	{name: "Iishmarmep", relRow: 5, relCol: 0},
	{name: "Riimishkin", relRow: 10, relCol: 0},
	{name: "Hasaggan", relRow: 0, relCol: 1},
	{name: "Gipuurkhir", relRow: 2, relCol: 1},
	{name: "Kuuhuurga", relRow: 6, relCol: 1},
	{
		name: "Forquee", relRow: 8, relCol: 1,
		description: "“Leave Forquee alone!” Learned Karma techniques (e.g. purging Fury) from the aliens(Droyne?). Flying domed city. Under Virus attack when we left.",
	},
	{name: "Kemkeaanguu", relRow: 10, relCol: 1},
	{
		name: "Did", relRow: 2, relCol: 2,
		description: "Message to deliver from Imperial Office of Calendar Compliance. 10 MCr reward!",
	},
	{name: "Arlur", relRow: 3, relCol: 2},
	{
		name: "Uure", relRow: 7, relCol: 2,
		description: "Father's Lodge. Ancient crash site for Imperial archaeological support ship. Our first lesson on Karma, Madness, and Fury.\n	On our second pass through, Virus bombarded the archaeological site, and the Lodge initiated a solar flare to destroy the Virus fleet.",
	},
	{name: "InmuuKi", relRow: 9, relCol: 2},
	{
		name: "Irshuushdaar", relRow: 4, relCol: 3,
		description: "Where the space wreck was sent to jump to, presumably by the Marquis of Rohintash's fury infused ally.",
	},
	{name: "Muukher", relRow: 5, relCol: 3},
	{
		name: "Irdar Ga", relRow: 6, relCol: 3,
		description: "Ships that go here don't come back, because they're eaten by nanites. The nanites are building (have built?) something.",
	},
	{
		name: "Vlair", relRow: 7, relCol: 3,
		description: "First time in system: Storm, Omar Factors\n\n	Second time in system: Mine of madness and sarlacc maw",
	},
	{name: "Ziger", relRow: 10, relCol: 3},
	{name: "Gier Iir", relRow: 2, relCol: 4},
	{
		name: "Gimi Kuuid", relRow: 6, relCol: 4,
		description: "Underwater spherical displacement — signs of alternate 'madness' universe with dark changes (dictator vs. administrator, etc.).\n\n	Gained hexagon key in a pawn shop here.\n\n	Revolution underway on our way back through here on [date]. Virus is also in-system since [date].\n\n	One of the planets hosting a Galactic Traveler office.",
	},
	{name: "Garuu Uurges", relRow: 7, relCol: 4},
	{name: "Daaruugka", relRow: 8, relCol: 4},
	{
		name: "Girgulash", relRow: 5, relCol: 5,
		description: "Trango is a moon of Girgulash, home of Trango shipyards, headed by Messorius Thraxton. Thraines, a city on Girgulash proper, hosted an Imperial Archive.\n\n	This is where we found the message for the Office of Calendar Compliance on Did.\n\n	First encounter with wireheads; they bombed a dome on Trango. Also, our first encounter with old Imperial computers that seemed to recognize Stefan as someone called Admiral Rogers.",
	},
	{
		name: "Khida", relRow: 7, relCol: 5,
		description: "Stefan, Kyle Vesta, and Dr. Denmark are from here.\n\n\tDr. Denmark was kidnapped here, kicking off our group's association with him.\n\n\tKhida Secundus Defensive Facility, Grandma Vesta, Ling Standard Products\n\n\tThe last time we were here, a 'bone' ship (Madness?) jumped in, and we helped destroy it.",
	},
	{name: "Khui", relRow: 8, relCol: 5},
	{name: "Lis", relRow: 0, relCol: 6},
	{name: "Sham", relRow: 1, relCol: 6},
	{name: "Amem", relRow: 2, relCol: 6},
	{
		name: "Ugar", relRow: 5, relCol: 6,
		description: "Puddle of blood, murder mystery. Cetagandan/tainted money plot. On return (15th system): Mr. Data?",
	},
	{
		name: "Vlir", relRow: 6, relCol: 6,
		description: "Western town with the tainted money. Earned favors from the Tong (Star Tong?) - Stefan can ask questions where there are contacts.",
	},
	{name: "Duuksha", relRow: 9, relCol: 6},
	{
		name: "Udipeni", relRow: 5, relCol: 7,
		description: "Home of Threnody and Max.\n\n\tBottom cult is taking hold, though the military is resistant and Max's efforts last time through may have helped.\n\n\tSigns of anagathics use (and possible withdrawl) by government officials.\n\n\tFound the [steampunk computer] in an old Imperium Library.",
	},
	{
		name: "Nagilun", relRow: 6, relCol: 7,
		description: "Home of Marian Dove.\n\n\tThressalar - Museum attack, underground facility with ancient computer, weird vines and grey goo.\n\n\tKurnak, neighboring nation, has a negative tariff and embeds surveillance nanites in all cargoes that pass through.\n\n\tDoctor Swede (who looks just like Doc Denmark) performed much of the terraforming of New Nagilun. New Nagilun still has a low-level poisonous biosphere and lethal wildlife. Screechers keep wildlife away.\n\n\tMost recent visit: asteroid with strange physical properties (gravity distortions?) on collision course with planet. Rescued researcher from kidnappers and took him to space to help enter the asteroid and determine what to do. Raced Cerberus there, and won.",
	},
	{
		name: "Kagershi", relRow: 7, relCol: 7,
		description: "Cerberus has a base set up here, taking over the system. Staging ground for massive Cerberus fleet, but we managed to sabotage their fuel supply, making it just impure enough to be incompatible with their jump stabilizers.\n\n\tIn an old mining facility contested by Cerberus and Virus-controlled robots, we managed to sneak in and obtain an anomalous sphere of jump space contained in real space, contained by a lanthanum ring. This is one of the pieces we need to collect before facing the Fracture.",
	},
	{
		name: "Gowandon", relRow: 8, relCol: 7,
		description: "Cerberus headquarters. Also, Ling Standard Products is headquartered at Rhona Minor. LSP runs mining operations in the ring, essentially using independent miners in indentured servitude. We rescued numerous Aivar from this system.\n\n\tOn Gowandon proper, we explored an underwater arcology and found the Deeps, who were especially affected by the virus ravaging the planet. We managed to develop a cure to the virus using Imperial-quality lab facilities, even as the arcology itself tried to launch itself into orbit.\n\n\tMet up with some of Admiral Rogers (head of the “wizards”) minions, and they spilled some info on their long-term strategy.\n\n\tImperial Scout base also granted us access on the strength of Stefan's identity as Admiral Rogers. Gained more background and a skill point.",
	},
	{name: "Uuduud", relRow: 0, relCol: 8},
	{name: "Uuruum", relRow: 3, relCol: 8},
	{
		name: "Irar Lar", relRow: 6, relCol: 8,
		description: "The face in Threnody's photos turn out to be Father's Observatory, where Father is fighting off Virus. We managed to hook up with the friendly half of the installation and make our way through to the data room, where Father was able to deliver various data and useful tech for us. As we escaped, Father blew the installation, and the whole planet. We had a LSP observer along on our visit, their IT guy. It's unclear how complete his report will be to them.\n\n\tMary (Marian's doppelganger) went to Irar Lar on the Long Shot to look for clues to where the Fracture will be.",
	},
	{
		name: "Kuundin", relRow: 7, relCol: 8,
		description: "According to the logs from the Corona, this is where Anagathics come from. We found a lab ship and boarded it, discovering that things had gone terribly wrong. We found formulas for variants on Anagathics, but all the formulas depend on a “precursor” compound that we have no supply of.\n\n\tNext, we visited the planet itself, where we discovered that the Droyne of the planet had ascended to become a planet-wide spirit that protects the native populace.\n\n\tThe Fruitful Discovery, a LSP science vessel, jumped in and approached the planet to refuel and experiment on the locals. We got them off the planet, then protected them from the Virus ship that jumped into the system. We agreed to coordinate our jumps to Irar Lar, which was both our next destination.",
	},
	{name: "Kamar Inag", relRow: 9, relCol: 8},
	{name: "Duumar Di", relRow: 10, relCol: 8},
	{name: "Uumbuu", relRow: 1, relCol: 9},
	{name: "Ervlan", relRow: 2, relCol: 9},
	{name: "Gaiid", relRow: 3, relCol: 9},
	{name: "Digapir", relRow: 5, relCol: 9},
	{name: "Shiirla", relRow: 8, relCol: 9},
	{name: "Dinkhaluurk", relRow: 10, relCol: 9},
	{name: "black hole", relRow: 4, relCol: 1},
}
