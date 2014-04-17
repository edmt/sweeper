package main
 
import (
  "fmt"
  "flag"
  "path/filepath"
  "io/ioutil"
  "strings"
  "os"
  "io"
)
 
const (
  SEARCH_PATH = "%s\\CFDs_Expedidos\\%s\\%s\\%s\\*.xml"
  RFC_LIST_PATH = "C:\\usersdata\\*"
  OUTPUTDIR = "C:\\backup\\"
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
 
func main() {
  year := flag.String("year", "", "Asigna el año para la búsqueda en el árbol de directorios")
  month := flag.String("month", "", "Asigna el mes para la búsqueda en el árbol de directorios")
  day := flag.String("day", "", "Asigna el día para la búsqueda en el árbol de directorios")
  flag.Usage = func() { flag.PrintDefaults() }
  flag.Parse()
  rfcList, _ := getRFCList()
  globPatternList := getGlobPatternList(rfcList, *year, *month, *day)
  for _, globPattern := range globPatternList {
    files, _ := listFiles(globPattern)
    for _, file := range files {
      //task(file)
      //_, fname := filepath.Split(file)
      //copy(file, OUTPUTDIR + fname)
    }
  }
}
 
func listFiles(globPattern string) (matches []string, err error) {
  return filepath.Glob(globPattern)
}
 
func getGlobPatternList(rfcList []string, year, month, day string) (output []string) {
  output = make([]string, len(rfcList))
  for i, value := range rfcList {
    output[i] = fmt.Sprintf(SEARCH_PATH, value, year, month, day)
  }
  return
}
 
func getRFCList() (matches []string, err error) {
  return filepath.Glob(RFC_LIST_PATH)
}
 
func copy(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}
 
 
func task(filename string) {
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
