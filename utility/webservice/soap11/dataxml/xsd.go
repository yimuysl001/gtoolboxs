package dataxml

import (
	"encoding/xml"
	"errors"
	"reflect"
	"strings"
)

const (
	String  string = "xsd:string"
	Int32          = "xsd:int"
	Int64          = "xsd:long"
	Float32        = "xsd:float"
	Float64        = "Xsd:double"
	Bool           = "xsd:boolean"
)

type XsdSchema struct {
	XMLName            xml.Name         `xml:"http://www.w3.org/2001/XMLSchema xsd:schema"`
	TNS                string           `xml:"xmlns tns,attr,omitempty"`
	XS                 string           `xml:"xmlns xs,attr,omitempty"`
	TargetNamespace    string           `xml:"targetNamespace,attr,omitempty"`
	ElementFormDefault string           `xml:"elementFormDefault,attr,omitempty"`
	Version            string           `xml:"version,attr,omitempty"`
	Elements           []XsdElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Import             []XsdImport      `xml:"xsd:import"`
}

type XsdElement struct {
	XMLName      xml.Name        `xml:"http://www.w3.org/2001/XMLSchema element"`
	Type         string          `xml:"type,attr,omitempty"`
	Nillable     bool            `xml:"nillable,attr"`
	MinOccurs    int             `xml:"minOccurs,attr"`
	MaxOccurs    int             `xml:"maxOccurs,attr,omitempty"`
	Form         string          `xml:"form,attr,omitempty"`
	Name         string          `xml:"name,attr"`
	ComplexTypes *XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type XsdComplexType struct {
	XMLName  xml.Name           `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Name     string             `xml:"name,attr,omitempty"`
	Abstract bool               `xml:"abstract,attr"`
	Sequence []XsdElement       `xml:"sequence>element"`
	Content  *XsdComplexContent `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
}

type XsdComplexContent struct {
	XMLName   xml.Name     `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
	Extension XsdExtension `xml:"http://www.w3.org/2001/XMLSchema extension"`
}

type XsdExtension struct {
	XMLName  xml.Name     `xml:"http://www.w3.org/2001/XMLSchema extension"`
	Base     string       `xml:"base,attr"`
	Sequence []XsdElement `xml:"sequence>element"`
}

type XsdImport struct {
	XMLName        xml.Name `xml:"xsd:import"`
	SchemaLocation string   `xml:"schemaLocation,attr,omitempty"`
	Namespace      string   `xml:"namespace,attr"`
}

func checkBaseTypeKind(k reflect.Kind) (string, error) {
	switch k {
	case reflect.String:
		return String, nil
	case reflect.Int, reflect.Int32:
		return Int32, nil
	case reflect.Int64:
		return Int64, nil
	case reflect.Bool:
		return Bool, nil
	case reflect.Float32:
		return Float32, nil
	case reflect.Float64:
		return Float64, nil
	default:
		return "", errors.New("no match")
	}
}

func getTagsInfo(t reflect.StructField) (string, bool) {
	required := false
	name := t.Name
	tags := strings.Split(t.Tag.Get("wsdl"), ",")
	for k, v := range tags {
		tag := strings.TrimSpace(v)
		if k == 0 {
			if tag != "" {
				name = tag
			}
		} else {
			if tag == "required" {
				required = true
				break
			}
		}
	}
	return name, required
}
