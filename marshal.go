package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Doc struct {
	XMLName     xml.Name        `xml:"Comprobante"`
	Tipo        string          `xml:"tipoDeComprobante,attr"`
	Version     string          `xml:"version,attr"`
	Emisor      CFDIEmisor      `xml:"Emisor"`
	Receptor    CFDIReceptor    `xml:"Receptor"`
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

func parseXml(path string) {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Doc
	xml.Unmarshal(b, &q)
	fmt.Println(q)
}
