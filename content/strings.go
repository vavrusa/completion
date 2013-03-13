package content

import (
	"bytes"
	"fmt"
)

func (f File) String() string {
	return f.Name
}

func (s SourceLocation) String() string {
	return fmt.Sprintf("%s:%d:%d", s.File, s.Line, s.Column)
}

func (f FullyQualifiedName) String() string {
	if f.Absolute != "" {
		return f.Absolute
	}
	return f.Relative
}

func (t Type) String() (ret string) {
	switch t.Flags & FLAG_TYPE_MASK {
	case FLAG_TYPE_POINTER:
		return t.Specialization[0].String() + "*"
	case FLAG_TYPE_ARRAY:
		return t.Specialization[0].String() + "[]"
	case FLAG_TYPE_METHOD:
		return t.Methods[0].String()
	}
	ret += fmt.Sprintf("%s%s", t.Flags, t.Name.String())
	if len(t.Specialization) > 0 {
		ret += "<"
		for i := range t.Specialization {
			if i > 0 {
				ret += ", "
			}
			ret += t.Specialization[i].String()
		}
		ret += ">"
	}
	for i, e := range t.Extends {
		if i == 0 {
			ret += "\n\textends "
		} else {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", e)
	}
	for i, e := range t.Implements {
		if i == 0 {
			ret += "\n\timplements "
		} else {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", e)
	}
	for _, e := range t.Fields {
		ret += "\n\t"
		ret += fmt.Sprintf("%s", e)
	}
	for _, e := range t.Methods {
		ret += "\n\t"
		ret += fmt.Sprintf("%s", e)
	}
	for _, e := range t.Types {
		ret += "\n\t"
		ret += fmt.Sprintf("%s", e)
	}

	return ret
}

func (v Variable) String() string {
	return fmt.Sprintf("%s %s", v.Type, v.Name)
}

func (f Field) String() (ret string) {
	return ret + fmt.Sprintf("%s%s %s", f.Flags, f.Type, f.Name)
}

func (m Method) String() (ret string) {
	ret = fmt.Sprintf("%s", m.Flags)
	for i := range m.Returns {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", m.Returns[i])
	}
	ret += fmt.Sprintf(" %s(", m.Name)
	for i := range m.Parameters {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", m.Parameters[i])
	}
	ret += ")"
	return
}

func (a Flags) String() (ret string) {
	switch a & FLAG_ACC_MASK {
	case FLAG_ACC_PUBLIC:
		ret += "public "
	case FLAG_ACC_PRIVATE:
		ret += "private "
	case FLAG_ACC_PROTECTED:
		ret += "protected "
	}
	if a&FLAG_STATIC != 0 {
		ret += "static "
	}
	if a&FLAG_FINAL != 0 {
		ret += "final "
	}
	if a&FLAG_RESTRICT != 0 {
		ret += "restrict "
	}
	if a&FLAG_CONST != 0 {
		ret += "const "
	}
	if a&FLAG_VOLATILE != 0 {
		ret += "volatile "
	}
	if a&FLAG_REFERENCE != 0 {
		ret += "&"
	}

	switch a & FLAG_TYPE_MASK {
	case FLAG_TYPE_ARRAY:
		ret += "[]"
	case FLAG_TYPE_PACKAGE:
		ret += "package "
	case FLAG_TYPE_POINTER:
		ret += "*"
	case FLAG_TYPE_CLASS:
		ret += "class "
	case FLAG_TYPE_INTERFACE:
		ret += "interface "
	}
	return
}

func (cr CompletionResult) String() (ret string) {
	b := bytes.NewBuffer(nil)
	if len(cr.Types) > 0 {
		b.WriteString("Types:\n")
		for _, t := range cr.Types {
			b.WriteString(fmt.Sprintf("\t%s\n", t))
		}
	}
	if len(cr.Fields) > 0 {
		b.WriteString("Fields:\n")
		for _, t := range cr.Fields {
			b.WriteString(fmt.Sprintf("\t%s\n", t))
		}
	}
	if len(cr.Methods) > 0 {
		b.WriteString("Methods:\n")
		for _, t := range cr.Methods {
			b.WriteString(fmt.Sprintf("\t%s\n", t))
		}
	}
	ret = string(b.Bytes())
	return
}
