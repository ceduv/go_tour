package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type change struct {
	field string
	old   any
	new   any
}

func printChanges(changes []change) {
	for _, c := range changes {
		fmt.Printf("  - %s: %+v => %+v\n", c.field, c.old, c.new)
	}
}

// *****************************************************************************
// *****************************************************************************
// REFLECT

type goFieldValue struct {
	GType string
	Value any
}

type goField struct {
	Name  string
	Value any
	GType string
	// Value *goFieldValue
	Tags map[string]string
}

type goStruct struct {
	Name   string
	Fields []*goField
}

func (s *goStruct) NumField() int {
	if s == nil {
		return 0
	}

	return len(s.Fields)
}

func (s *goStruct) Field(i int) *goField {
	if s == nil || i >= len(s.Fields) {
		return nil
	}

	return s.Fields[i]
}

func (s *goStruct) FieldByName(fName string) *goField {
	if s == nil {
		return nil
	}

	for _, v := range s.Fields {
		if v.Name == fName {
			return v
		}
	}

	return nil
}

func (s *goField) TagLookup(key string) (string, bool) {
	if s == nil {
		return "", false
	}

	v, ok := s.Tags[key]

	return v, ok
}

func (s *goStruct) printDeclaration() string {
	// TODO: implement
	// s.Name{field1: value, field2: value2, ...}
	// ex: S{t1: 2, t2: 2, t3: 0}
	return ""
}

func (s *goStruct) printStruct() string {
	if s == nil {
		return ""
	}

	res := fmt.Sprintf("type %s struct {\n", s.Name)

	for _, v := range s.Fields {
		tags := []string{}
		for tk, tv := range v.Tags {
			tags = append(tags, fmt.Sprintf(`"%s":"%s"`, tk, tv))
		}

		res += fmt.Sprintf("\t%s %s `%s`\n", v.Name, v.GType, strings.Join(tags, " "))
	}

	return res + "}"
}

func (s *goStruct) addValue(name string, value int, tags map[string]string) {
	for i, v := range s.Fields {
		if v.Name == name {
			s.Fields[i] = &goField{name, value, "int", tags}

			return
		}
	}

	s.Fields = append(s.Fields, &goField{name, value, "int", tags})
}

func (s *goStruct) addT1(value int) {
	s.addValue("t1", value, map[string]string{"watch": "", "json": "t1"}) // t1 int `watch:""`
}

func (s *goStruct) addT2(value int) {
	s.addValue("t2", value, map[string]string{"watch": "always", "json": "t2"})
}

func (s *goStruct) addT3(value int) {
	s.addValue("t3", value, nil)
}

func _compareValues(gf1, gf2 *goField) bool {
	if gf1.GType != gf2.GType {
		return false
	}

	switch gf1.GType {
	case "time.Time":
		t1 := gf1.Value.(time.Time)
		t2 := gf2.Value.(time.Time)

		return t1.Unix() == t2.Unix()
	case "int":
		return gf1.Value.(int) == gf2.Value.(int)
	default:
		return reflect.DeepEqual(gf1.Value, gf2.Value)
	}
}

func _compare2(ss1, ss2 goStruct, res, always []change) ([]change, []change) {
	for _, v := range ss1.Fields {
		ss2Field := ss2.FieldByName(v.Name)
		// if ss2Field == nil {
		// 	// Message d'erreur, panic...
		// }

		watchValue, ok := v.TagLookup("watch")

		switch {
		case !ok:
			// do nothing
		case watchValue == "":
			if !_compareValues(v, ss2Field) {
				res = append(res, change{field: v.Name, old: v.Value, new: ss2Field.Value})
			}
		case watchValue == "always":
			always = append(always, change{field: v.Name, old: v.Value, new: ss2Field.Value})
		}
	}

	return res, always
}

func compare2(ss1, ss2 goStruct) []change {
	res, always := _compare2(ss1, ss2, []change{}, []change{})

	if len(res) > 0 && len(always) > 0 {
		res = append(res, always...)
	}

	return res
}

func newS2(t1, t2, t3 int) goStruct {
	s := goStruct{Name: "S2", Fields: []*goField{}}
	s.addT1(t1)
	s.addT2(t2)
	s.addT3(t3)

	return s
}

func main2() {
	var changes []change

	s1 := newS2(1, 2, 0)
	fmt.Println(s1.printStruct())

	s3 := newS2(2, 2, 0)

	changes = compare2(s1, s3)
	fmt.Println("===> s3:")
	printChanges(changes)
}

// *****************************************************************************
// *****************************************************************************
// CLASSIQUE

type S struct {
	t1 int `watch:""`
	t2 int `watch:"always"`
	t3 int
}

func _compare(ss1, ss2 S, res, always []change) ([]change, []change) {
	if ss1.t1 != ss2.t1 {
		res = append(res, change{field: "t1", old: ss1.t1, new: ss2.t1})
	}

	always = append(always, change{field: "t2", old: ss1.t2, new: ss2.t2})

	return res, always
}

func compare(ss1, ss2 S) []change {
	res, always := _compare(ss1, ss2, []change{}, []change{})

	if len(res) > 0 && len(always) > 0 {
		res = append(res, always...)
	}

	return res
}

func main1() {
	s1 := S{t1: 1, t2: 2, t3: 0}
	s2 := S{t1: 1, t2: 3, t3: 0}
	s3 := S{t1: 2, t2: 2, t3: 0}
	s4 := S{t1: 2, t2: 3, t3: 0}
	s5 := S{t1: 1, t2: 3, t3: 1}
	s6 := S{t1: 2, t2: 3, t3: 1}

	var changes []change

	changes = compare(s1, s2)
	fmt.Printf("===> s2: %+v\n", changes)
	changes = compare(s1, s3)
	fmt.Printf("===> s3: %+v\n", changes)
	changes = compare(s1, s4)
	fmt.Printf("===> s4: %+v\n", changes)
	changes = compare(s1, s5)
	fmt.Printf("===> s5: %+v\n", changes)
	changes = compare(s1, s6)
	fmt.Printf("===> s6: %+v\n", changes)
}

func main() {
	// main1()
	// fmt.Println("*****************************************************************************")
	main2()
}
