package unique

import mapset "github.com/deckarep/golang-set"

// UniqueDatabase ...
type UniqueDatabase struct {
	set mapset.Set
}

func (ud UniqueDatabase) init() {
	set := mapset.NewSet()
	ud.set = set
}

// Add ...
func (ud UniqueDatabase) Add(value interface{}) error {
	if ud.set == nil {
		ud.init()
	}
	ud.set.Add(value)
	return nil
}

// Count well it counts
func (ud UniqueDatabase) Count() int {
	if ud.set == nil {
		ud.init()
	}
	return ud.set.Cardinality()
}
