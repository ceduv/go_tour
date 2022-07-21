package likereflect

import (
	"errors"
	"fmt"
	"strings"
)

var errEmptyStruct = errors.New("empty struct")

// structure go simplifié.
type GoStruct struct {
	name   string
	fields []*GoField
}

// *************************************
// 				Methodes de GoStruct
// *************************************

func (s *GoStruct) IdenticalTo(sBis *GoStruct, checkOrder bool) bool {
	// ! TODO: implement.
	return false
}

// R nombre de fields.
func (s *GoStruct) NumField() int {
	if s == nil {
		return 0
	}

	return len(s.fields)
}

// R le field select by index.
func (s *GoStruct) Field(i int) *GoField {
	if s == nil || i < 0 || i >= len(s.fields) {
		return nil
	}

	return s.fields[i]
}

// R le field select by name.
func (s *GoStruct) FieldByName(fName string) *GoField {
	if s == nil {
		return nil
	}

	for _, v := range s.fields {
		if v.name == fName {
			return v
		}
	}

	return nil
}

// Formate les fields d une *Gostruct
// ex: s.Name{field1: value, field2: value2, ...}
func (s *GoStruct) PrintDeclaration() (string, error) {
	if s == nil {
		return "", errEmptyStruct
	}

	fields := make([]string, 0, s.NumField())

	for _, f := range s.fields {
		fields = append(fields, fmt.Sprintf("%s: %+v", f.name, f.value))
	}
	// * OldVersion
	// * res := fmt.Sprintf("%v{", s.name)
	// * for i := 0; i < s.NumField(); i++ {
	// * 	res += fmt.Sprintf("%+v: %+v", s.fields[i].name, s.fields[i].value)
	// * 	if i < s.NumField()-1 {
	// * 		res += ","
	// * 	}
	// * }
	// *return res + "}"

	return fmt.Sprintf("%v{", s.name) + strings.Join(fields, ", ") + "}", nil
}

// todo : comment.
func (s *GoStruct) PrintStruct() string {
	if s == nil {
		return ""
	}

	res := fmt.Sprintf("type %s struct {\n", s.name)

	for _, v := range s.fields {
		tags := []string{}
		for tk, tv := range v.tags {
			tags = append(tags, fmt.Sprintf(`"%s":"%s"`, tk, tv))
		}

		res += fmt.Sprintf("\t%s %s `%s`\n", v.name, v.gType, strings.Join(tags, " "))
		res += fmt.Sprintf("\t%s %s `...`\n", v.name, v.gType)
	}

	return res + "}"
}

// Permet d ajouter les valeur passées aux fields
// si le name existe maj les valeurs.
func (s *GoStruct) addValue(name string, value int, tags map[string]string) {
	if s == nil {
		return
	}

	for i, v := range s.fields {
		if v.name == name {
			s.fields[i] = &GoField{name, value, "int", tags}

			return
		}
	}

	s.fields = append(s.fields, &GoField{name, value, "int", tags})
}

// permet de passer une valeur et call addvalues avec les valeurs de field.
func (s *GoStruct) addT1(value int) error {
	if s == nil {
		return errEmptyStruct
	}

	s.addValue("t1", value, map[string]string{"watch": "", "json": "t1"})

	return nil
}

func (s *GoStruct) addT2(value int) error {
	if s == nil {
		return errEmptyStruct
	}

	s.addValue("t2", value, map[string]string{"watch": "always", "json": "t2"})

	return nil
}

func (s *GoStruct) addT3(value int) error {
	if s == nil {
		return errEmptyStruct
	}

	s.addValue("t3", value, nil)

	return nil
}
