package dataxml

import (
	"encoding/xml"
)

type WsdlDefinitions struct {
	XMLName  xml.Name      `xml:"http://schemas.xmlsoap.org/wsdl/ definitions"`
	SoapEnv  string        `xml:"xmlns:SOAP-ENV,attr"`
	TargetNs string        `xml:"targetNamespace,attr"`
	Tns      string        `xml:"xmlns:tns,attr"`
	Soap     string        `xml:"xmlns:soap,attr"`
	Xsd      string        `xml:"xmlns:xsd,attr"`
	Xsi      string        `xml:"xmlns:xsi,attr"`
	Wsdl     string        `xml:"xmlns:wsdl,attr"`
	Types    WsdlType      `xml:"types"`
	Message  []WsdlMessage `xml:"message"`
	PortType WsdlPortType  `xml:"portType"`
	Binding  WsdlBinding   `xml:"binding"`
	Service  WsdlService   `xml:"service"`
}

type WsdlType struct {
	Schemas []XsdSchema `xml:"schema"`
}

type WsdlMessage struct {
	Name string     `xml:"name,attr"`
	Part []WsdlPart `xml:"part"`
}

type WsdlPart struct {
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr,omitempty"`
	Element string `xml:"element,attr,omitempty"`
}

type WsdlPortType struct {
	Name       string                  `xml:"name,attr"`
	Operations []WsdlPortTypeOperation `xml:"operation"`
}

type WsdlPortTypeOperation struct {
	Name   string                       `xml:"name,attr"`
	Input  WsdlPortTypeOperationMessage `xml:"input"`
	Output WsdlPortTypeOperationMessage `xml:"output"`
	//Fault  WsdlPortTypeOperationMessage `xml:"fault,omitempty"`
}

type WsdlPortTypeOperationMessage struct {
	Name    string `xml:"name,attr,omitempty"`
	Message string `xml:"message,attr"`
}

type WsdlBinding struct {
	Name        string                 `xml:"name,attr"`
	Type        string                 `xml:"type,attr"`
	SoapBinding WsdlSoapBinding        `xml:"soap:binding"`
	Operations  []WsdlBindingOperation `xml:"operation"`
}

type WsdlSoapBinding struct {
	XMLName   xml.Name `xml:"soap:binding"`
	Transport string   `xml:"transport,attr"`
	Style     string   `xml:"style,attr"`
}

type WsdlBindingOperation struct {
	Name          string            `xml:"name,attr"`
	SoapOperation WsdlSoapOperation `xml:"soap:operation"`
	Input         WsdlSoapBodyIO    `xml:"input"`
	Output        WsdlSoapBodyIO    `xml:"output"`
	//Fault         WsdlSoapBody      `xml:"fault>fault,omitempty"`
}

type WsdlSoapOperation struct {
	SoapAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr,omitempty"`
}

type WsdlSoapBodyIO struct {
	SoapBody WsdlSoapBody `xml:"soap:body"`
}

type WsdlSoapBody struct {
	Name          string `xml:"name,attr,omitempty"`
	Use           string `xml:"use,attr"`
	EncodingStyle string `xml:"encodingStyle,attr"`
}

type WsdlService struct {
	Name string          `xml:"name,attr"`
	Port WsdlServicePort `xml:"port"`
}

type WsdlServicePort struct {
	XMLName xml.Name           `xml:"port"`
	Name    string             `xml:"name,attr"`
	Binding string             `xml:"binding,attr"`
	Address WsdlServiceAddress `xml:"soap:address"`
}

type WsdlServiceAddress struct {
	XMLName  xml.Name `xml:"soap:address"`
	Location string   `xml:"location,attr"`
}
