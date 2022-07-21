package likereflect

// import (
// 	"fmt"
// 	"reflect"
// 	"time"
// )

// // structure
// // sert a instancier les changements
// type change struct {
// 	field string
// 	old   any
// 	new   any
// }

// // check si les fieldbyName correspondent entre les 2 entitées
// // Selon watchvalue || ok on rentre dans un case
// // R 2[] de type change(struct)
// func _compare2(ss1, ss2 GoStruct, res, always []change) ([]change, []change) {
// 	for idx, v := range ss1.Fields {
// 		ss2Field := ss2.Field(idx)
// 		if ss2Field == nil {
// 			panic(" problem field == nil")
// 		}
// 		// !! v.name = t1 => t2 =>  t3
// 		// !! v = &{t1 1 int map[json:t1 watch:]}
// 		// !! ss2Field = &{t1 2 int map[json:t1 watch:]}

// 		watchValue, ok := v.TagLookup("watch")

// 		switch {
// 		case !ok:
// 			// do nothing
// 		case watchValue == "":
// 			if !_compareValues(v, ss2Field) {
// 				res = append(res, change{field: v.Name, old: v.Value, new: ss2Field.Value})
// 			}
// 		case watchValue == "always":
// 			always = append(always, change{field: v.Name, old: v.Value, new: ss2Field.Value})
// 		}
// 	}

// 	return res, always
// }

// // compare les donnees, les stocks dans res & always
// // R un [] des changements
// // !!! [{t1 1 2} {t2 2 2}]
// func compare2(ss1, ss2 GoStruct) []change {
// 	// !!! call avec 2[]change{}
// 	// !!! pkoi ne pas initialiser dans _compare2
// 	res, always := _compare2(ss1, ss2, []change{}, []change{})

// 	if len(res) > 0 && len(always) > 0 {
// 		res = append(res, always...)
// 	}
// 	return res
// }

// // On cherche dans les tags la valeur associe a la clé
// // ici v = ""/always/vide  ok = true true false
// func (s *GoField) TagLookup(key string) (string, bool) {
// 	if s == nil {
// 		return "", false
// 	}

// 	v, ok := s.Tags[key]
// 	fmt.Println(ok)

// 	return v, ok
// }

// // verifie que les types comparés soient similaires
// // Param des lignes ici v et ss2 go field
// func _compareValues(gf1, gf2 *GoField) bool {
// 	if gf1.GType != gf2.GType {
// 		return false
// 	}

// 	switch gf1.GType {
// 	case "time.Time":
// 		t1 := gf1.Value.(time.Time)
// 		t2 := gf2.Value.(time.Time)
// 		return t1.Unix() == t2.Unix()
// 	case "int":
// 		return gf1.Value.(int) == gf2.Value.(int)
// 	default:
// 		return reflect.DeepEqual(gf1.Value, gf2.Value)
// 	}
// }

// // display les differences entre s1 et s2
// func printChanges(changes []change) {
// 	for _, c := range changes {
// 		fmt.Printf("  - %s: %+v => %+v\n", c.field, c.old, c.new)
// 	}
// }
