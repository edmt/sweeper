package main

import (
	l4g "code.google.com/p/log4go"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"strings"
)

const (
	SCHEMA_CORTO                       = "http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/TimbreFiscalDigital/TimbreFiscalDigital.xsd"
	SCHEMA_CONSITIO                    = "http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/TimbreFiscalDigital/TimbreFiscalDigital.xsd"
	SCHEMA_CONSITIO_CONCFD             = "http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd"
	ELEMENTO_TFD_CON_DEFINICION_TIPOTF = "<tfd:TimbreFiscalDigital xmlns:tfd=\"http://www.sat.gob.mx/TimbreFiscalDigital\" xsi:schemaLocation=\"http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/TimbreFiscalDigital/TimbreFiscalDigital.xsd\""
	DECLARACION_NAMESPACE_TFD          = "xmlns:tfd=\"http://www.sat.gob.mx/TimbreFiscalDigital\""
	ELEMENTO_TFD                       = "<tfd:TimbreFiscalDigital"
	ELEMENTO_TFD_CON_DEFINICION_OK     = "<tfd:TimbreFiscalDigital xmlns:tfd=\"http://www.sat.gob.mx/TimbreFiscalDigital\" xsi:schemaLocation=\"http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd\""
	CADENA_TEMPORAL                    = "ddl_tmp_01"
	CADENA_VACIA                       = ""
	NAMESPACE_CFDV2                    = "http://www.sat.gob.mx/cfd/2"
)

func Replace(filename string, c *cli.Context) bool {
	contents, errOnRead := ioutil.ReadFile(filename)
	if errOnRead == nil {
		new_content, hasChanged := fixSchemaLocation(string(contents))
		if hasChanged == true {
			backUpFilePath := Format(filename, c.String("baseDir"), c.String("backUpDir"))
			l4g.Debug("Respalda archivo %s en %s", filename, backUpFilePath)
			BackUp(filename, backUpFilePath)
			new_content_in_bytes := []byte(new_content)
			l4g.Debug("Intenta escribir en archivo %s", filename)
			errOnWrite := ioutil.WriteFile(filename, new_content_in_bytes, 0644)
			if errOnWrite != nil {
				l4g.Error("Error al escribir. Archivo: %s, Error: %s", filename, errOnWrite.Error())
			} else {
				l4g.Debug("Archivo %s exitosamente corregido", filename)
			}
			return true
		} else {
			l4g.Debug("No es necesario modificar el archivo %s", filename)
		}
	} else {
		l4g.Error("Error al leer archivo %s", filename)
	}
	return false
}

func fixSchemaLocation(contents string) (string, bool) {
	if !strings.Contains(contents, ELEMENTO_TFD_CON_DEFINICION_OK) && !strings.Contains(contents, NAMESPACE_CFDV2) {
		caso_timbre := false
		new_content := contents
		if strings.Contains(new_content, ELEMENTO_TFD_CON_DEFINICION_TIPOTF) {
			caso_timbre = true
			// Para el caso de Timbre Fiscal, donde se está definiendo en el lugar adecuado, pero con un schema location que no aparece en el anexo 20.
			new_content = strings.Replace(new_content, ELEMENTO_TFD_CON_DEFINICION_TIPOTF, CADENA_TEMPORAL, -1)
		}
		// Buscar todas las definiciones adicionales posibles en cualquier parte del documento
		new_content = strings.Replace(new_content, DECLARACION_NAMESPACE_TFD, CADENA_VACIA, -1)
		new_content = strings.Replace(new_content, SCHEMA_CORTO, CADENA_VACIA, -1)
		new_content = strings.Replace(new_content, SCHEMA_CONSITIO, CADENA_VACIA, -1)
		new_content = strings.Replace(new_content, SCHEMA_CONSITIO_CONCFD, CADENA_VACIA, -1)

		if caso_timbre {
			// Reemplazar cadena temporal por la apertura y definición del nodo TimbreFiscalDigital
			new_content = strings.Replace(new_content, CADENA_TEMPORAL, ELEMENTO_TFD_CON_DEFINICION_OK, -1)
		} else {
			// Reemplazar elemento del timbre por elemento con definición correcta (corta)
			new_content = strings.Replace(new_content, ELEMENTO_TFD, ELEMENTO_TFD_CON_DEFINICION_OK, -1)
		}
		return new_content, true
	}
	return contents, false
}

func Format(str, oldFormat, newFormat string) string {
	replacer := strings.NewReplacer(oldFormat, newFormat)
	return replacer.Replace(string(str))
}
