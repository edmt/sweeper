package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Doc struct {
	XMLName     xml.Name        `xml:"Comprobante"`
	Tipo        string          `xml:"tipoDeComprobante,attr"`
	Version     string          `xml:"version,attr"`
	Emisor      CFDIEmisor      `xml:"Emisor"`
	Receptor    CFDIReceptor    `xml:"Receptor"`
	Conceptos   []CFDIConcepto  `xml:"Conceptos>Concepto"`
	Complemento CFDIComplemento `xml:"Complemento"`
}

type CFDIEmisor struct {
	XMLName xml.Name `xml:"Emisor"`
	RFC     string   `xml:"rfc,attr"`
}

type CFDIReceptor struct {
	XMLName xml.Name `xml:"Receptor"`
	RFC     string   `xml:"rfc,attr"`
}

type CFDIConcepto struct {
	XMLName     xml.Name `xml:"Concepto"`
	Descripcion string   `xml:"descripcion,attr"`
}

type CFDIComplemento struct {
	XMLName             xml.Name               `xml:"Complemento"`
	TimbreFiscalDigital TFDTimbreFiscalDigital `xml:"TimbreFiscalDigital"`
}

type TFDTimbreFiscalDigital struct {
	XMLName           xml.Name `xml:"TimbreFiscalDigital"`
	NumeroCertificado string   `xml:"noCertificadoSAT,attr"`
	FechaTimbrado     string   `xml:"FechaTimbrado,attr"`
}

func (t TFDTimbreFiscalDigital) String() string {
	return fmt.Sprintf("%s\t%s", t.NumeroCertificado, t.FechaTimbrado)
}

func (d Doc) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s",
		d.Emisor.RFC,
		d.Receptor.RFC,
		d.Complemento.TimbreFiscalDigital.NumeroCertificado,
		d.Complemento.TimbreFiscalDigital.FechaTimbrado,
		d.Version)
}

func (c CFDIConcepto) ContainsKeyword() bool {
	desc := strings.ToLower(c.Descripcion)
	return strings.Contains(desc, "magna") ||
		strings.Contains(desc, "premium") ||
		strings.Contains(desc, "diesel")
}

func (d Doc) ContainsGasKeyword() bool {
	for _, concept := range d.Conceptos {
		if concept.ContainsKeyword() {
			return true
		}
	}
	return false
}

func parseXml(path string) {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	rawContent, _ := ioutil.ReadAll(xmlFile)

	var query Doc
	xml.Unmarshal(rawContent, &query)
	if query.ContainsGasKeyword() {
		fmt.Println(query)
	}
}
