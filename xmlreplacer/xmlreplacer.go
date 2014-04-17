package xmlreplacer

import (
  "io/ioutil"
  "strings"
  "fmt"
)

const(
  CADENA0_OLD = " http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/TimbreFiscalDigital/TimbreFiscalDigital.xsd"
  CADENA0_NEW = ""
 
  CADENA1_OLD = " http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/TimbreFiscalDigital/TimbreFiscalDigital.xsd"
  CADENA1_NEW = ""
 
  CADENA2_OLD = " xmlns:tfd=\"http://www.sat.gob.mx/TimbreFiscalDigital\""
  CADENA2_NEW = ""
 
  CADENA3_OLD = "<tfd:TimbreFiscalDigital"
  CADENA3_NEW = "<tfd:TimbreFiscalDigital xmlns:tfd=\"http://www.sat.gob.mx/TimbreFiscalDigital\" xsi:schemaLocation=\"http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/TimbreFiscalDigital/TimbreFiscalDigital.xsd\""
)

func Replace(filename string) {
  contents, err := ioutil.ReadFile(filename)
  if err == nil {
    replacer := strings.NewReplacer(CADENA0_OLD, CADENA0_NEW, CADENA1_OLD, CADENA1_NEW, CADENA2_OLD, CADENA2_NEW, CADENA3_OLD, CADENA3_NEW)
    new_content := replacer.Replace(string(contents))
    new_content_in_bytes := []byte(new_content)
    err = ioutil.WriteFile(filename, new_content_in_bytes, 0644)
    if err != nil {
      fmt.Println(filename)
      fmt.Println(err)
    }
  }
}

func Format(str, oldFormat, newFormat string) string {
  replacer := strings.NewReplacer(oldFormat, newFormat)
  return replacer.Replace(string(str))
}
