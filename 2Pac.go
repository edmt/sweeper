package main

import (
	l4g "code.google.com/p/log4go"
	"github.com/codegangsta/cli"
	"os"
	"time"
)

const LOG_CONFIGURATION_FILE = "logging-conf.xml"

func init() {
	l4g.LoadConfiguration(LOG_CONFIGURATION_FILE)
}

func main() {
	app := cli.NewApp()
	app.Name = "2Pac"
	app.Usage = ""
	app.Version = "0.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{"baseDir", "undefined", "Directorio base para iniciar la búsqueda"},
		cli.StringFlag{"backUpDir", "undefined", "Directorio base para respaldo"},
	}
	app.Action = func(c *cli.Context) {
		globPatternList := GetGlobPatternList(
			c.String("baseDir"))
		l4g.Info("Directorios encontrados: %d", len(globPatternList))
		for _, globPattern := range globPatternList {
			files, _ := ListFiles(globPattern)
			l4g.Info("%d archivos en directorio %s", len(files), globPattern)
			for _, filePath := range files {
				l4g.Debug("Procesando archivo: %s", filePath)
				parseXml(filePath)
			}
		}
	}
	l4g.Info("Process ID: %d", os.Getpid())
	app.Run(os.Args)
	l4g.Info("%s stopped", app.Name)
	time.Sleep(time.Second)
}
