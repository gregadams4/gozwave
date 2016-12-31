package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// http://products.z-wavealliance.org/products/1729
// http://www.pepper1.net/zwavedb/

type deviceDataValue struct {
	Value string `xml:"value,attr"`
}

type deviceData struct {
	ManufacturerID  deviceDataValue `xml:"manufacturerId"`
	ProductType     deviceDataValue `xml:"productType"`
	ProductID       deviceDataValue `xml:"productId"`
	LibType         deviceDataValue `xml:"libType"`
	ProtoVersion    deviceDataValue `xml:"protoVersion"`
	ProtoSubVersion deviceDataValue `xml:"protoSubVersion"`
	AppVersion      deviceDataValue `xml:"appVersion"`
	AppSubVersion   deviceDataValue `xml:"appSubVersion"`
	BasicClass      deviceDataValue `xml:"basicClass"`
	GenericClass    deviceDataValue `xml:"genericClass"`
	SpecificClass   deviceDataValue `xml:"specificClass"`
	Optional        deviceDataValue `xml:"optional"`
	Listening       deviceDataValue `xml:"listening"`
	Routing         deviceDataValue `xml:"routing"`
	BeamSensor      string          `xml:"beamSensor"`
	RfFrequency     string          `xml:"rfFrequency"`
}

type commandClass struct {
	ID         string `xml:"id,attr"`
	Controlled string `xml:"controlled,attr"`
	InNIF      string `xml:"inNIF,attr"`
	Secure     string `xml:"secure,attr"`
	NonSecure  string `xml:"nonSecure,attr"`
	Version    string `xml:"version,attr"`
}

func (c commandClass) IDasHex() string {
	return "0x" + c.ID
}

type deviceDescription struct {
	BrandName   string `xml:"brandName"`
	ProductName string `xml:"productName"`
	Description string `xml:"description"`
}

func (dd deviceDescription) EscapedDescription() string {
	return removeNewLines(dd.Description)
}

func (dd deviceDescription) EscapedProductName() string {
	return removeNewLines(dd.ProductName)
}
func (dd deviceDescription) EscapedBrandName() string {
	return removeNewLines(dd.BrandName)
}

type assocGroup struct {
	Number      int    `xml:"number,attr"`
	MaxNodes    int    `xml:"maxNodes,attr"`
	Description string `xml:"description"`
}

type hexVal string

type configParam struct {
	ID          int           `xml:"number,attr"`
	Type        string        `xml:"type,attr"`
	Size        int           `xml:"size,attr"`
	Default     hexVal        `xml:"default,attr"`
	Name        string        `xml:"name>lang"`
	Description string        `xml:"description>lang"`
	Values      []configValue `xml:"value"`
}

func (cp configParam) EscapedDescription() string {
	return removeNewLines(cp.Description)
}

func (cp configParam) EscapedName() string {
	return removeNewLines(cp.Name)
}

type configValue struct {
	From        hexVal `xml:"from,attr"`
	To          hexVal `xml:"to,attr"`
	Unit        string `xml:"unit,attr"`
	Reserved    bool   `xml:"reserved,attr"`
	Description string `xml:"description>lang"`
}

func (cv configValue) EscapedDescription() string {
	return removeNewLines(cv.Description)
}

type zWaveDevice struct {
	Filename          string
	Name              string
	XMLName           xml.Name          `xml:"ZWaveDevice"`
	DescriptorVersion string            `xml:"descriptorVersion"`
	DeviceData        deviceData        `xml:"deviceData"`
	DeviceDescription deviceDescription `xml:"deviceDescription"`
	CommandClasses    []commandClass    `xml:"commandClasses>commandClass"`
	AssocGroups       []assocGroup      `xml:"assocGroups"`
	ConfigParams      []configParam     `xml:"configParams>configParam"`
}

func (hv hexVal) HexStrToInt() int {
	v, err := strconv.ParseInt(string(hv), 16, 64)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return int(v)
}

var templ = `package {{.Package}}

import (
	"github.com/stampzilla/gozwave/commands"
)

type CommandClass struct {
	ID         commands.ZWaveCommand
	Controlled string
	InNIF      string
	Secure     bool
	NonSecure  bool
	Version    string
}
type parameter struct {
	ID int
	Name string
	Type string
	Description string
}
type AssocGroup struct {
	Number      int
	MaxNodes    int
	Description string
}
type ConfigParam struct {
	ID      int
	Type        string
	Size        int
	Default     int
	Name        string
	Description string
	Values      []ConfigValue
}
type ConfigValue struct {
	From        int
	To          int
	Unit        string
	Reserved    bool
	Description string
}
type Device struct{
	Brand string
	Product string
	Description string

	CommandClasses []*CommandClass
	Parameters []*parameter
	ConfigParams []ConfigParam
	AssocGroups []AssocGroup

	ManufacturerID string
	ProductType string
	ProductID string
}
func New(manufacturerID, productType, productID string) *Device{
	dev := manufacturerID+productType+productID
	switch dev {
{{- range $value := .Devices }}
	case "{{ $value.DeviceData.ManufacturerID.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductID.Value }}":
		return New{{ $value.DeviceData.ManufacturerID.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductID.Value }}() // {{ $value.Filename }} | {{ $value.Name }}
{{- end}}
	}

	return nil
}
{{- range $value := .Devices }}
func New{{ $value.DeviceData.ManufacturerID.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductID.Value }}() *Device{
	return &Device{
		Brand: "{{ $value.DeviceDescription.EscapedBrandName }}",
		Product: "{{ $value.DeviceDescription.EscapedProductName }}",
		Description: "{{ $value.DeviceDescription.EscapedDescription }}",

		ManufacturerID: "{{$value.DeviceData.ManufacturerID.Value }}",
		ProductType: "{{$value.DeviceData.ProductType.Value }}",
		ProductID: "{{$value.DeviceData.ProductID.Value }}",
		CommandClasses: []*CommandClass{
{{- range $cmd := $value.CommandClasses }}
			&CommandClass{
				{{- if ne $cmd.ID ""}}
				ID: {{$cmd.IDasHex}},
				{{- end }}
				{{- if ne $cmd.Controlled "" }}
				Controlled: "{{$cmd.Controlled}}",
				{{- end }}
				{{- if ne $cmd.InNIF ""}}
				InNIF: "{{$cmd.InNIF}}",
				{{- end }}
				{{- if ne $cmd.Secure ""}}
				Secure: {{$cmd.Secure}},
				{{- end }}
				{{- if ne $cmd.NonSecure ""}}
				NonSecure: {{$cmd.NonSecure}},
				{{- end }}
				{{- if ne $cmd.Version ""}}
				Version: "{{$cmd.Version}}",
				{{- end }}
			},
{{- end}}
		},
		ConfigParams: []ConfigParam{
{{- range $cv := $value.ConfigParams }}
			ConfigParam{
				ID: {{$cv.ID}},
				Type: "{{$cv.Type}}",
				Size: {{$cv.Size}},
				{{- if ne $cv.Default ""}}
				Default: {{$cv.Default.HexStrToInt}},
				{{- end }}
				Name: "{{$cv.EscapedName}}",
				Description: "{{$cv.EscapedDescription}}",
				Values: []ConfigValue{
{{- range $v := $cv.Values }}
					ConfigValue{
						{{- if ne $v.From ""}}
						From: {{$v.From.HexStrToInt}},
						{{- end }}
						{{- if ne $v.To ""}}
						To: {{$v.To.HexStrToInt}},
						{{- end }}
						Unit: "{{$v.Unit}}",
						Reserved: {{$v.Reserved}},
						Description: "{{$v.EscapedDescription}}",
					},
{{- end}}
				},
			},
{{- end}}
		},
	}
}
{{- end}}
`

type templates struct {
	Devices map[string]*zWaveDevice
	Package string
}

var outputfile string
var databaseDir string
var packageName string

func init() {
	flag.StringVar(&outputfile, "file", "", "Output file")
	flag.StringVar(&databaseDir, "databasedir", "./database", "Directory with xml files")
	flag.StringVar(&packageName, "package", "devices", "Package name of generated code")
	flag.Parse()
}

func main() {
	devices := &templates{
		Package: packageName,
	}
	files, err := ioutil.ReadDir(databaseDir)
	if err != nil {
		log.Fatal(err)
	}

	devices.Devices = make(map[string]*zWaveDevice)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".xml" {
			continue
		}
		var dev *zWaveDevice
		xmlFile, err := os.Open(filepath.Join(databaseDir, file.Name()))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		decoder := xml.NewDecoder(xmlFile)

		err = decoder.Decode(&dev)
		xmlFile.Close()
		if err != nil {
			log.Println(file.Name())
			log.Println(err)
			return
		}

		dev.Filename = file.Name()
		dev.Name =
			dev.DeviceData.ManufacturerID.Value +
				dev.DeviceData.ProductType.Value +
				dev.DeviceData.ProductID.Value

		//spew.Dump(dev)
		//TODO wee need to be more specific in our template above
		// 135-0115-1000-0001-06-03-16-01-04.xml
		// 132-0115-1000-0001-06-03-16-01-05.xml
		// has same ManufacturerID-ProductType-ProductID ?
		devices.Devices[dev.Name] = dev
	}

	// return
	t := template.New("t")
	t, err = t.Parse(templ)
	if err != nil {
		panic(err)
	}

	if outputfile == "" {
		err = t.Execute(os.Stdout, devices)
	} else {
		file, err := os.Create(outputfile)
		if err != nil {
			log.Println(err)
			return
		}
		if err != nil {
			log.Println("Unable to open file", err)
			return
		}
		err = t.Execute(file, devices)
		file.Close()
	}
	if err != nil {
		panic(err)
	}

}

func removeNewLines(s string) string {
	s = strings.Replace(s, "\n", `\n`, -1)
	s = strings.Replace(s, "\r", `\n`, -1)
	s = strings.Replace(s, "\\", "/", -1)
	return s
}
