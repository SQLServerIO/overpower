package planetattack

import (
	"database/sql"
	"mule/hexagon"
)

type ShipView struct {
	db         *sql.DB
	Gid        int
	Fid        int
	Sid        int
	Controller int
	Loc        hexagon.Coord
	Trail      []hexagon.Coord
	Size       int
}
