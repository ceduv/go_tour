package likereflect

import (
	"reflect"
	"time"
)

// type goFieldValue struct {
// 	GType string
// 	Value any
// }

// champs simplifié d une structure simplifié
// t1 1 int map[json:t1 watch:].
type GoField struct {
	name  string
	value any
	gType string
	// todo : Value *goFieldValue.
	tags map[string]string
}

func (f *GoField) GetName() string {
	return f.name
}

func (f *GoField) GetValue() any {
	return f.value
}

// verifie que les types comparés soient similaires
// Param des lignes ici v et ss2 go field.
func CompareFields(gf1, gf2 *GoField) bool {
	if gf1.gType != gf2.gType {
		return false
	}

	switch gf1.gType {
	case "time.Time":
		t1 := gf1.value.(time.Time)
		t2 := gf2.value.(time.Time)

		return t1.Unix() == t2.Unix()
	case "int":
		return gf1.value.(int) == gf2.value.(int)
	default:
		return reflect.DeepEqual(gf1.value, gf2.value)
	}
}

// todo : Comment.
func (gf1 *GoField) CompareTo(gf2 *GoField) bool {
	if gf1.gType != gf2.gType {
		return false
	}

	switch gf1.gType {
	case "time.Time":
		t1 := gf1.value.(time.Time)
		t2 := gf2.value.(time.Time)

		return t1.Unix() == t2.Unix()
	case "int":
		return gf1.value.(int) == gf2.value.(int)
	default:
		return reflect.DeepEqual(gf1.value, gf2.value)
	}
}

// On cherche dans les tags la valeur associe a la clé
// ici v = ""/always/vide  ok = true true false.
func (f *GoField) TagLookup(key string) (string, bool) {
	if f == nil {
		return "", false
	}

	v, ok := f.tags[key]

	return v, ok
}
