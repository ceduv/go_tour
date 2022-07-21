package change

import (
	"fmt"

	likereflect "exo_matt/like_reflect"
)

// structure
// sert a instancier les Changements.
type Change struct {
	field string
	old   any
	new   any
}

func LoopOverSlice() {
	v := []int{1, 2, 3}

	for idx, vv := range v {
		fmt.Printf("===> %d %+v", idx, vv)
	}

	for idx := 0; idx < len(v); idx++ {
		fmt.Printf("===> %d %+v", idx, v[idx])
	}
}

// check si les fieldbyName correspondent entre les 2 entitées
// Selon watchvalue || ok on rentre dans un case
// R 2[] de type Change(struct).
func _compare(ss1, ss2 *likereflect.GoStruct, res, always []Change) ([]Change, []Change) {
	for idx := 0; idx < ss1.NumField(); idx++ {
		ss1Field := ss1.Field(idx)
		ss2Field := ss2.Field(idx)

		if ss2Field == nil {
			panic(" problem field == nil")
		}

		watchValue, ok := ss1Field.TagLookup("watch")

		switch {
		case !ok:
			// * do nothing
		case watchValue == "":
			if !likereflect.CompareFields(ss1Field, ss2Field) {
				res = append(res, Change{field: ss1Field.GetName(), old: ss1Field.GetValue(), new: ss2Field.GetValue()})
			}
		case watchValue == "always":
			always = append(always, Change{field: ss1Field.GetName(), old: ss1Field.GetValue(), new: ss2Field.GetValue()})
		}
	}

	return res, always
}

// compare les donnees, les stocks dans res & always
// R un [] des Changements ex: [{t1 1 2} {t2 2 2}].
func Compare(ss1, ss2 *likereflect.GoStruct) []Change {
	// ! todo : call identicalto ss1/ss2
	// ! todo : renvoi un bool
	res, always := _compare(ss1, ss2, []Change{}, []Change{})

	if len(res) > 0 && len(always) > 0 {
		res = append(res, always...)
	}

	return res
}

// display les differences entre s1 et s2.
func PrintChanges(changes []Change) {
	for _, c := range changes {
		fmt.Printf("  - %s: %+v => %+v\n", c.field, c.old, c.new)
	}
}
