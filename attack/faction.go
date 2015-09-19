package attack

type Faction struct {
	Name        string
	FactionID   int
	TurnDone    bool
	BuildOrders []Order
	View        SectorView
}

func NewFaction() *Faction {
	return &Faction{
		BuildOrders: []Order{},
	}
}

type SectorView struct {
	Faction    int
	Turn       int
	PlanetGrid map[[2]int]PlanetView
	ShipGrid   map[[2]int][]ShipView
	TrailGrid  map[[2]int][]ShipTrail
}

type Order struct {
	Location [2]int
	Size     int
	Target   [2]int
}

func NewOrder() *Order {
	return &Order{}
}

func NewSectorView() *SectorView {
	return &SectorView{
		PlanetGrid: map[[2]int]PlanetView{},
		ShipGrid:   map[[2]int][]ShipView{},
		TrailGrid:  map[[2]int][]ShipTrail{},
	}
}

type PlanetView struct {
	Name        string
	Location    [2]int
	Yours       bool
	Inhabitants [2]int
	Resources   int
}

func NewPlanetView() *PlanetView {
	return &PlanetView{}
}

type ShipView struct {
	FactionID int
	Size      int
	Location  [2]int
	Yours     bool
	Target    [2]int
}

func NewShipView() *ShipView {
	return &ShipView{}
}

func (sv *SectorView) AddPlanet(pl *Planet) {
	pv := NewPlanetView()
	pv.Name = pl.Name
	pv.Location = pl.Location
	if pl.Faction() == sv.Faction {
		pv.Yours = true
		pv.Inhabitants = pl.Inhabitants
		pv.Resources = pl.Resources
	}
	sv.PlanetGrid[pl.Location] = *pv
}

func (sv *SectorView) AddShip(cl *Ship) {
	cv := NewShipView()
	cv.FactionID = cl.FactionID
	cv.Size = cl.Size
	loc := cl.Location()
	cv.Location = loc
	if cv.FactionID == sv.Faction {
		cv.Yours = true
		cv.Target = cl.Target()
	}
	if list, ok := sv.ShipGrid[loc]; ok {
		sv.ShipGrid[loc] = append(list, *cv)
	} else {
		sv.ShipGrid[loc] = []ShipView{*cv}
	}
}

func (sv *SectorView) AddShipTrail(loc [2]int, trail ShipTrail) {
	if list, ok := sv.TrailGrid[loc]; ok {
		sv.TrailGrid[loc] = append(list, trail)
	} else {
		sv.TrailGrid[loc] = []ShipTrail{trail}
	}
}

func (s *Sector) MakeView(factionID int) *SectorView {
	sv := NewSectorView()
	sv.Faction = factionID
	sv.Turn = s.Turn
	for _, pl := range s.PlanetGrid {
		sv.AddPlanet(pl)
	}
	for loc, list := range s.ShipGrid {
		if s.VisibleArea(factionID, loc) {
			for _, cl := range list {
				sv.AddShip(cl)
			}
		}
	}
	for loc, list := range s.TrailGrid {
		if s.VisibleArea(factionID, loc) {
			for _, trail := range list {
				sv.AddShipTrail(loc, trail)
			}
		}
	}
	return sv
}

func (s *Sector) VisibleArea(factionID int, loc [2]int) bool {
	for loc2, pl := range s.PlanetGrid {
		if pl.Faction() == factionID && HexDist(loc, loc2) < VISRANGE {
			return true
		}
	}
	return false
}