package likereflect

import "fmt"

func NewStructure(sName string) *GoStruct {
	return &GoStruct{name: sName}
}

func (s *GoStruct) AppendField(fName, fGType string, fValue any, fTags map[string]string) error {
	if s == nil {
		return errEmptyStruct
	}

	if s.fields == nil {
		s.fields = []*GoField{}
	}

	s.fields = append(s.fields, &GoField{name: fName, value: fValue, gType: fGType, tags: fTags})

	return nil
}

func NewS(sName string, t1, t2, t3 int) *GoStruct {
	s := NewStructure(sName)

	if err := s.addT1(t1); err != nil {
		fmt.Println(err)
	}

	if err := s.addT1(t2); err != nil {
		fmt.Println(err)
	}

	if err := s.addT1(t3); err != nil {
		fmt.Println(err)
	}

	return s
}

// permet d instancier une variable via gostruct.
func NewS_old(t1, t2, t3 int) *GoStruct {
	// var s *GoStruct
	s := &GoStruct{name: "S", fields: []*GoField{}}

	if err := s.addT1(t1); err != nil {
		fmt.Println(err)
	}

	if err := s.addT1(t2); err != nil {
		fmt.Println(err)
	}

	if err := s.addT1(t3); err != nil {
		fmt.Println(err)
	}

	return s
}

// func foo() {
// 	var s *GoStruct
// 	fmt.Println(s)

// 	t := &GoStruct{}
// 	fmt.Println(t)
// }
