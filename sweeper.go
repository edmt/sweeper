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
	app.Name = "sweeper"
	app.Usage = "Me llama usted, entonces voy, Don Barredora es quien yo soy 🎵"
	app.Version = "0.1.4"

	app.Flags = []cli.Flag{
		cli.StringFlag{"baseDir", "undefined", "Directorio base para iniciar la búsqueda"},
		cli.StringFlag{"year", "undefined", "Año para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"month", "undefined", "Mes para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"day", "undefined", "Día para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"backUpDir", "undefined", "Directorio base para respaldo"},
	}
	app.Action = func(c *cli.Context) {
		globPatternList := GetGlobPatternList(
			c.String("baseDir"),
			c.String("year"),
			c.String("month"),
			c.String("day"))
		l4g.Info("Directorios encontrados: %d", len(globPatternList))
		for _, globPattern := range globPatternList {
			files, _ := ListFiles(globPattern)
			l4g.Info("%d archivos en directorio %s", len(files), globPattern)
			for _, filePath := range files {
				l4g.Debug("Procesando archivo: %s", filePath)
				Replace(filePath, c)
			}
		}
	}
	l4g.Info("Process ID: %d", os.Getpid())
	app.Run(os.Args)
	l4g.Info("sweeper stopped")
	time.Sleep(time.Second)
}
