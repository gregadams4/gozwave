package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/Sirupsen/logrus"
)

type CommandClasses struct {
	XMLName        xml.Name       `xml:"zw_classes"`
	CommandClasses []CommandClass `xml:"cmd_class"`
}

type CommandClass struct {
	Name     string    `xml:"name,attr"`
	Help     string    `xml:"help,attr"`
	Key      string    `xml:"key,attr"`
	Version  string    `xml:"version,attr"`
	ReadOnly bool      `xml:"read_only,attr"`
	Comment  string    `xml:"comment,attr"`
	Commands []Command `xml:"cmd"`
}

func (cc *CommandClass) BuildCommandClassVersionStr() string {
	if cc.Version != "1" {
		return fmt.Sprintf("V%s", cc.Version)
	}
	return ""
}
func (cc *CommandClass) BuildCommandClassVar() string {
	var output string
	words := strings.Split(strings.ToLower(cc.Name), "_")
	words = words[2:]
	for _, word := range words {
		output += strings.Title(word)
	}
	if cc.Version != "1" {
		output += fmt.Sprintf("V%s", cc.Version)
	}
	return output
	// return strings.Join(strings.Split(cc.Help, " ")[2:], "")
}

func (cc *CommandClass) RequiresBinary() bool {
	for _, command := range cc.Commands {
		for _, param := range command.Params {
			if param.BuildParamType() != "byte" {
				return true
			}
		}
	}
	return false
}

type Command struct {
	Name          string         `xml:"name,attr"`
	Help          string         `xml:"help,attr"`
	Key           string         `xml:"key,attr"`
	Comment       string         `xml:"comment,attr"`
	Params        []Param        `xml:"param"`
	VariantGroups []VariantGroup `xml:"variant_group"`
	ParamIndex    int
}

func (c *Command) CurIndex() int {
	return c.ParamIndex
}

func (c *Command) IncIndex() int {
	c.ParamIndex++
	return c.ParamIndex - 1
}

func (c *Command) IncIndexTwo() int {
	c.ParamIndex += 2
	return c.ParamIndex
}

func (c *Command) IncIndexFour() int {
	c.ParamIndex += 4
	return c.ParamIndex
}

func (c *Command) IsReport() bool {
	return strings.Contains(strings.ToLower(c.Name), "report") || strings.Contains(strings.ToLower(c.Name), "notification")
}

func (c *Command) BuildCommandVar() string {
	var output string
	words := strings.Split(strings.ToLower(c.Name), "_")
	for _, word := range words {
		output += strings.Title(word)
	}
	return output
}

func (c *Command) BuildCommandLength() int {
	total := 0
	for _, p := range c.Params {
		total += p.CalcIncrement()
	}
	return total
}

func (c *Command) Count() int {
	count := 2

	count += len(c.Params)

	for _, vg := range c.VariantGroups {
		count += len(vg.Params)
	}
	return count
}

func (c *Command) HasVGs() bool {
	return len(c.VariantGroups) != 0
}

type Param struct {
	Name         string      `xml:"name,attr"`
	Key          string      `xml:"key,attr"`
	Type         string      `xml:"type,attr"`
	TypeHashCode byte        `xml:"typehashcode,attr"`
	Comment      string      `xml:"comment,attr"`
	BitMask      BitMask     `xml:"bitmask"`
	Variant      Variant     `xml:"variant"`
	ValueAttrib  ValueAttrib `xml:"valueattrib"`
	Word         Word        `xml:"word"`
	Constants    []Const     `xml:"const"`
}

func (p *Param) BuildParamVar() string {
	name := strings.Replace(p.Name, " ", "", -1)
	name = strings.Replace(name, "-", "", -1)
	name = strings.Replace(name, ".", "", -1)
	name = strings.Replace(name, "+", "Plus", -1)
	return name
}

func (p *Param) CalcIncrement() int {
	if p.Type == "CONST" {
		return 1
	}
	if p.Type == "BYTE" {
		return 1
	}
	if p.Type == "WORD" {
		return 2
	}
	if p.Type == "DWORD" {
		return 4
	}
	if p.Type == "STRUCT_BYTE" {
		return 1
	}
	if p.Type == "VARIANT" {
		return 1
	}
	if p.Type == "BITMASK" {
		return 1
	}
	logrus.Infof("no handler for %s", p.Type)
	return 1
}

func (p *Param) BuildParamType() string {
	if p.Type == "CONST" {
		return "byte"
	}
	if p.Type == "BYTE" {
		return "byte"
	}
	if p.Type == "WORD" {
		return "uint16"
	}
	if p.Type == "DWORD" {
		return "uint64"
	}
	if p.Type == "STRUCT_BYTE" {
		return "byte"
	}
	//believe this will actually need to be []byte as it varies dependent on variant paramoffs value.
	if p.Type == "VARIANT" {
		return "byte"
	}
	if p.Type == "BITMASK" {
		return "byte"
	}
	log.Printf("no handler - %s", p.Type)
	return "byte"
}

type VariantGroup struct {
	Name         string  `xml:"name,attr"`
	Key          string  `xml:"key,attr"`
	VariantKey   string  `xml:"variantKey,attr"`
	ParamOffs    string  `xml:"paramOffs,attr"`
	SizeMask     string  `xml:"sizemask,attr"`
	SizeOffs     string  `xml:"sizeoffs,attr"`
	TypeHashCode string  `xml:"typehashcode,attr"`
	Comment      string  `xml:"comment,attr"`
	Params       []Param `xml:"param"`
}

func (vg *VariantGroup) BuildVGVar() string {
	return strings.Title(vg.Name)
}

type Const struct {
	FlagName string `xml:"flagname,attr"`
	Key      string `xml:"key,attr"`
	FlagMask string `xml:"flagmask,attr"`
}

type Word struct {
	Key        string `xml:"key,attr"`
	HasDefines bool   `xml:"hasdefines,attr"`
	ShowHex    bool   `xml:"showhex"`
}

type Variant struct {
	Paramoffs int    `xml:"paramoffs,attr"`
	ShowHex   bool   `xml:"showhex"`
	Signed    bool   `xml:"signed"`
	SizeMask  string `xml:"sizemask"`
	SizeOffs  int    `xml:"sizeoffs"`
}

type ValueAttrib struct {
	Key        string     `xml:"key,attr`
	HasDefines bool       `xml:"hasdefines,attr"`
	ShowHex    bool       `xml:"showhex,attr"`
	BitFlags   []BitFlag  `xml:"bitflag"`
	BitFields  []BitField `xml:"bitfield"`
}

type BitMask struct {
	Key       string `xml:"key,attr"`
	ParamOffs int    `xml:"paramoffs,attr"`
	LenMask   string `xml:"lenmask,attr"`
	LenOffs   int    `xml:"lenoffs,attr"`
}

type BitFlag struct {
	Key      string `xml:"key,attr"`
	FlagName string `xml:"flagname,attr"`
	FlagMask string `xml:"flagmask,attr"`
}

type BitField struct {
	Key       string `xml:"key,attr"`
	FieldName string `xml:"fieldname,attr"`
	FieldMask string `xml:"fieldmask,attr"`
	Shifter   int    `xml:"shifter"`
}

// func (c *Command) IsReport() bool {
// 	if strings.Contains(strings.ToLower(c.Name), "report") {
// 		return true
// 	}
// 	return false
// }

// type Frame struct {
// 	Name   string
// 	Size   int
// 	Values []Value
// }

// type Value struct {
// 	Type        string
// 	Name        string
// 	FuncVarName string
// }

// func commandClassMap(command string) string {
// 	switch command {
// 	case "Event":
// 		return "Notification"
// 	default:
// 		return command
// 	}
// }
var templ2 = `package {{.Package}}

type ZWaveCommand byte

const (
    {{- range .CommandClasses}}
    {{.BuildCommandClassVar}} = {{.Key}}
    {{- end}}
)
`

var templ = `package {{.Package}}

{{- if .CommandClass.RequiresBinary}}
import "encoding/binary"
{{- end}}
{{- range .CommandClass.Commands}}
{{- $command := .}}

{{- if $command.IsReport}}
type {{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}} struct {
    *report
    {{- range .Params}}
    {{.BuildParamVar}} {{.BuildParamType}}
    {{- end}}
    data []byte
}

func New{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}(data []byte) *{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}} {
    if len(data) < {{.BuildCommandLength}} {
				//may want to change this to just return nil
				for i := len(data); i < {{.BuildCommandLength}}; i++ {
            data = append(data, 0x00)
        }
    }

    return &{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}{
				{{- range $i, $value := .Params}}
				{{- if eq .BuildParamType "uint16"}}
				{{.BuildParamVar}}: binary.BigEndian.Uint16(data[{{$command.CurIndex}}:{{$command.IncIndexTwo}}]),
				{{- end}}
				{{- if eq .BuildParamType "uint64"}}
				{{.BuildParamVar}}: binary.BigEndian.Uint64(data[{{$command.CurIndex}}:{{$command.IncIndexFour}}]),
				{{- end}}
				{{- if eq .BuildParamType "byte"}}
				{{.BuildParamVar}}: data[{{$command.IncIndex}}],
        {{- end}}
				{{- end}}
        data: data,
    }
}
{{else}}
type {{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}} struct {
	node byte
	{{- range .Params}}
	{{.BuildParamVar}} {{.BuildParamType}}
	{{- end}}
	{{- range .VariantGroups}}
	{{.BuildVGVar}} []{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}{{.Name}}
	{{- end}}
}

{{- range .VariantGroups}}
type {{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}{{.Name}} struct {
	{{- range .Params}}
	{{.BuildParamVar}} {{.BuildParamType}}
	{{- end}}
}
{{- end}}

func New{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}() {{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}} {
	return {{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}{}
}

func (c *{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}) SetNode(node int) {
	c.node = byte(node)
}

func (c *{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}) Set({{- range .Params}}{{.BuildParamVar}} {{.BuildParamType}},{{- end}}{{- range .VariantGroups}}{{.BuildVGVar}} []{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}{{.Name}},{{- end}}) error {
	{{- range .Params}}
	c.{{.BuildParamVar}} = {{.BuildParamVar}}
	{{- end}}

	{{- range .VariantGroups}}
	c.{{.BuildVGVar}} = {{.BuildVGVar}}
	{{- end}}

	return nil
}

func (c *{{$command.BuildCommandVar}}{{$.CommandClass.BuildCommandClassVersionStr}}) Encode() []byte {
	{{- range .Params}}
	{{- if eq .BuildParamType "uint16"}}
	{{.BuildParamVar}}Bytes := []byte{}
	binary.BigEndian.PutUint16({{.BuildParamVar}}Bytes, c.{{.BuildParamVar}})
	{{- end}}
	{{- if eq .BuildParamType "uint64"}}
	{{.BuildParamVar}}Bytes := []byte{}
	binary.BigEndian.PutUint64({{.BuildParamVar}}Bytes, c.{{.BuildParamVar}})
	{{- end}}
	{{- end}}

	{{- if $command.HasVGs}}
	var data = []byte{
		0x13,
		c.node,
		byte({{.Count}}),
		byte({{$.CommandClass.BuildCommandClassVar}}),
		{{$command.Key}},
		{{- range .Params}}
		{{- if eq .BuildParamType "uint16"}}
		{{.BuildParamVar}}Bytes[0],
		{{.BuildParamVar}}Bytes[1],
		{{- end}}
		{{- if eq .BuildParamType "uint64"}}
		{{.BuildParamVar}}Bytes[0],
		{{.BuildParamVar}}Bytes[1],
		{{.BuildParamVar}}Bytes[2],
		{{.BuildParamVar}}Bytes[3],
		{{- end}}
		{{- if eq .BuildParamType "byte"}}
		c.{{.BuildParamVar}},
		{{- end}}
		{{- end}}
	}
	{{- range .VariantGroups}}
	{{- $vg := .}}
	for _, v := range c.{{$vg.BuildVGVar}} {
		{{- range .Params}}
		{{- if eq .BuildParamType "uint16"}}
		{{.BuildParamVar}}Bytes := []byte{}
		binary.BigEndian.PutUint16({{.BuildParamVar}}Bytes, v.{{.BuildParamVar}})
		data = append(data, {{.BuildParamVar}}Bytes...)
		{{- end}}
		{{- if eq .BuildParamType "uint64"}}
		{{.BuildParamVar}}Bytes := []byte{}
		binary.BigEndian.PutUint64({{.BuildParamVar}}Bytes, v.{{.BuildParamVar}})
		data = append(data, {{.BuildParamVar}}Bytes...)
		{{- end}}
		{{- if eq .BuildParamType "byte"}}
		data = append(data, v.{{.BuildParamVar}})
		{{- end}}
		{{- end}}
	}
	data = append(data, 0x25)
	{{- end}}
	return data
	{{else}}
	return []byte{
		0x13,
		c.node,
		byte({{.Count}}),
		byte({{$.CommandClass.BuildCommandClassVar}}),
		{{$command.Key}},
		{{- range .Params}}
		{{- if eq .BuildParamType "uint16"}}
		{{.BuildParamVar}}Bytes[0],
		{{.BuildParamVar}}Bytes[1],
		{{- end}}
		{{- if eq .BuildParamType "uint64"}}
		{{.BuildParamVar}}Bytes[0],
		{{.BuildParamVar}}Bytes[1],
		{{.BuildParamVar}}Bytes[2],
		{{.BuildParamVar}}Bytes[3],
		{{- end}}
		{{- if eq .BuildParamType "byte"}}
		c.{{.BuildParamVar}},
		{{- end}}
		{{- end}}
		0x25,
	}
	{{- end}}
}
{{- end}}
{{- end}}
`

//

// {{- end}}
// {{- end}}
// {{- end}}
// `

type templates struct {
	CommandClass *CommandClass
	Package      string
}

type templates2 struct {
	CommandClasses *[]CommandClass
	Package        string
}

func main() {
	var commandClasses *CommandClasses
	xmlFile, err := os.Open("./ZWave_custom_cmd_classes.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)
	err = decoder.Decode(&commandClasses)
	if err != nil {
		panic(err)
	}

	// spew.Dump(commandClasses)

	for _, commandClass := range commandClasses.CommandClasses {
		// if k == "SwitchBinary" {
		if commandClass.Name == "COMMAND_CLASS_HAIL" || commandClass.Name == "COMMAND_CLASS_CRC_16_ENCAP" || commandClass.Name == "COMMAND_CLASS_REMOTE_ASSOCIATION_ACTIVATE" {
			continue
		}
		if len(commandClass.Commands) > 0 {
			commands := templates{CommandClass: &commandClass, Package: "commands"}
			t := template.New("t")
			t, err = t.Parse(templ)
			if err != nil {
				panic(err)
			}

			file, err := os.Create(fmt.Sprintf("./commands/%s.go", commandClass.BuildCommandClassVar()))
			if err != nil {
				log.Println(err)
				return
			}
			if err != nil {
				log.Println("Unable to open file", err)
				return
			}

			err = t.Execute(file, commands)
			if err != nil {
				panic(err)
			}
		}
		// }
	}

	classes := templates2{CommandClasses: &commandClasses.CommandClasses, Package: "commands"}
	t := template.New("t2")
	t, err = t.Parse(templ2)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("./commands/commands.go")
	if err != nil {
		log.Println(err)
		return
	}
	if err != nil {
		log.Println("Unable to open file", err)
		return
	}
	err = t.Execute(file, classes)
	// output, err := json.Marshal(commandClasses)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(string(output))

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
