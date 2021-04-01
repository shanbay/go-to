package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

var intTypes = []string{
	"int",
	"int16",
	"int32",
	"int64",
	"int8",
	"uint",
	"uint16",
	"uint32",
	"uint64",
	"uint8",
}

var floatTypes = []string{"float32", "float64"}

var stringTypes = []string{"string"}

var boolTypes = []string{"bool"}

var allTypesSlice = [][]string{
	intTypes,
	floatTypes,
	stringTypes,
	boolTypes,
}

const (
	writeFileName     = "../to.go"
	writeTestFileName = "../to_test.go"
)

func main() {
	var err error
	t := template.New("oapi-codegen").Funcs(funcMap)
	t, err = parse(t)
	if err != nil {
		panic(err)
	}
	res := []string{"package to\n"}
	testRes := []string{`package to
import "testing"
`}
	for _, types := range allTypesSlice {
		toRes, toTestRes := product(t, types)
		res = append(res, toRes)
		testRes = append(testRes, toTestRes)
	}
	// to.go
	out := []byte(strings.Join(res, "\n"))
	if out, err = imports.Process(writeFileName, out, nil); err != nil {
		fmt.Println(res)
		panic(err)
	}
	if err := os.WriteFile(writeFileName, out, 0644); err != nil {
		panic(err)
	}
	// to_test.go
	out = []byte(strings.Join(testRes, "\n"))
	if out, err = imports.Process(writeTestFileName, out, nil); err != nil {
		panic(err)
	}
	if err := os.WriteFile(writeTestFileName, out, 0644); err != nil {
		panic(err)
	}
}

func product(t *template.Template, types []string) (string, string) {
	res, testsRes := []string{}, []string{}
	exists, testsExists := map[string]struct{}{}, map[string]struct{}{}
	for _, t1 := range types {
		for _, t2 := range types {
			// funcs
			basicTypes := []string{t1, t2, "*" + t1, "*" + t2}
			res = append(res, genFuncs(t, basicTypes, exists))
			testsRes = append(testsRes, genTestFuncs(t, basicTypes, testsExists))
			// slice funcs
			sliceTypes := []string{"[]" + t1, "[]" + t2, "*[]" + t1, "*[]" + t2}
			res = append(res, genFuncs(t, sliceTypes, exists))
			testsRes = append(testsRes, genTestFuncs(t, sliceTypes, testsExists))
		}
	}
	toRes := strings.Join(res, "\n") + "\n"
	toTestRes := strings.Join(testsRes, "\n") + "\n"
	return toRes, toTestRes
}

func genFuncs(t *template.Template, allTypes []string, exists map[string]struct{}) string {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	for _, from := range allTypes {
		for _, to := range allTypes {
			if from == to {
				continue
			}
			existsKey := from + "_" + to
			if _, ok := exists[existsKey]; ok {
				continue
			}
			exists[existsKey] = struct{}{}

			if err := t.ExecuteTemplate(w, getTmplName(from, to), &fromTo{From: from, To: to}); err != nil {
				panic(err)
			}
		}
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
	return buf.String()
}

func genTestFuncs(t *template.Template, allTypes []string, exists map[string]struct{}) string {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	for _, from := range allTypes {
		for _, to := range allTypes {
			if from == to {
				continue
			}
			existsKey := from + "_" + to
			if _, ok := exists[existsKey]; ok {
				continue
			}
			exists[existsKey] = struct{}{}

			if err := t.ExecuteTemplate(w, "Test", &fromTo{From: from, To: to}); err != nil {
				panic(err)
			}
		}
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
	return buf.String()
}

func getTmplName(from, to string) string {
	if from == to {
		return ""
	}
	tmplF, tmplT := "x", "y"
	if strings.HasPrefix(from, "*") || strings.HasPrefix(to, "*") {
		if strings.TrimPrefix(from, "*") == strings.TrimPrefix(to, "*") {
			tmplF, tmplT = "x", "x"
		}
		if strings.HasPrefix(from, "*") {
			tmplF += "Ptr"
		}
		if strings.HasPrefix(to, "*") {
			tmplT += "Ptr"
		}
	}
	return tmplF + "_" + tmplT
}

type fromTo struct {
	From string
	To   string
}

func (ft *fromTo) IsSlice() bool {
	return strings.HasPrefix(ft.From, "[]") || strings.HasPrefix(ft.From, "*[]")
}

func toName(s string) string {
	switch {
	case strings.HasPrefix(s, "*[]"):
		return strings.Title(s[3:]) + "SlicePtr"
	case strings.HasPrefix(s, "[]"):
		return strings.Title(s[2:]) + "Slice"
	case strings.HasPrefix(s, "*"):
		return strings.Title(s[1:]) + "Ptr"
	default:
		return strings.Title(s)
	}
}

var funcMap = template.FuncMap{
	"ToName": toName,
	"ToNoPtrName": func(s string) string {
		if strings.HasPrefix(s, "*") {
			return toName(s[1:])
		}
		return toName(s)
	},
	"ToPtrName": func(s string) string {
		if strings.HasPrefix(s, "*") {
			return toName(s)
		}
		return toName("*" + s)
	},
	"TrimSlicePrefix": func(s string) string { return strings.TrimLeft(s, "*[]") },
	"IsPtr":           func(s string) bool { return strings.HasPrefix(s, "*") },
	"Join":            func(s ...string) string { return strings.Join(s, "") },
	"GetInitStmt": func(s string) string {
		switch {
		case strings.HasPrefix(s, "*"):
			return " = nil"
		case strings.Contains(s, "[]"):
			return " = " + s + "{*new(" + strings.TrimLeft(s, "*[]") + ")}"
		case strings.Contains(s, "int") || strings.Contains(s, "float"):
			return " = 1"
		case strings.Contains(s, "string"):
			return ` = ""`
		case strings.Contains(s, "bool"):
			return ` = true`
		default:
			return ""
		}
	},
}

var tmpls = map[string]string{
	"xPtr_x": `func {{.From | ToName}}_{{.To | ToName}}(i {{.From}}, opt ...Option) {{.To}} {
	if len(opt) == 1 && opt[0] == UseDefaultEmpty && i == nil {
		return *new({{.To}})
	}
	return *i
}
`,
	"x_xPtr": `func {{.From | ToName}}_{{.To | ToName}}(i {{.From}}) {{.To}} { return &i }
`,
	"x_y": `func {{.From | ToName}}_{{.To | ToName}}(i {{.From}}) {{.To}} {
{{- if not .IsSlice}}
	return {{.To}}(i)
{{- else}}
	res := {{.To}}{}
	for _, item := range i {
		res = append(res, {{.To | TrimSlicePrefix}}(item))
	}
	return res
{{- end}}
}
`,
	"x_yPtr": `func {{.From | ToName}}_{{.To | ToName}}(i {{.From}}) {{.To}} {
	return {{.To | ToNoPtrName}}_{{.To | ToPtrName}}({{.From | ToName}}_{{.To | ToNoPtrName}}(i))
}
`,
	"xPtr_y": `func {{.From | ToName}}_{{.To | ToName}}(i {{.From}}, opt ...Option) {{.To}} {
	return {{.From | ToNoPtrName}}_{{.To | ToName}}({{.From | ToPtrName}}_{{.From | ToNoPtrName}}(i, opt...))
}
`,
	"xPtr_yPtr": `func {{.From | ToName}}_{{.To | ToName}}(i {{.From}}, opt ...Option) {{.To}} {
		return {{.To | ToNoPtrName}}_{{.To | ToPtrName}}({{.From | ToNoPtrName}}_{{.To | ToNoPtrName}}({{.From | ToPtrName}}_{{.From | ToNoPtrName}}(i, opt...)))
}
`,
	"Test": `func Test_{{.From | ToName}}_{{.To | ToName}}(t *testing.T) {
	var i {{.From}} {{.From | GetInitStmt}}
	var ii {{.To}} {{.To | GetInitStmt}}
	ii = {{.From | ToName}}_{{.To | ToName}}(i{{- if .From | IsPtr}}, UseDefaultEmpty{{- end}})
{{- $i := "i"}} {{- if (.From | IsPtr)}} {{- $i = Join (.From | ToName) "_" (.From | ToNoPtrName) "(i, UseDefaultEmpty)"}} {{- end}}
{{- $tfunc := Join (.To | ToName) "_" (.From | ToNoPtrName)}}
{{- if eq (.To | ToName) (.From | ToNoPtrName)}} {{- $tfunc = ""}} {{- end}}
{{- if not .IsSlice}}
	if {{$tfunc}}(ii{{- if .To | IsPtr}}, UseDefaultEmpty{{- end}}) != {{$i}} {
		t.Errorf("error in {{.From | ToName}}_{{.To | ToName}}")
	}
{{- else}}
	if len({{$tfunc}}(ii{{- if .To | IsPtr}}, UseDefaultEmpty{{- end}}) )!= len({{$i}}) {
		t.Errorf("error in {{.From | ToName}}_{{.To | ToName}}")
	}
{{- end}}
}
`,
}

func parse(t *template.Template) (*template.Template, error) {
	for name, s := range tmpls {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		if _, err := tmpl.Parse(s); err != nil {
			return nil, err
		}
	}
	return t, nil
}
