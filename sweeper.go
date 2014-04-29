package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sweeper"
	app.Usage = "Me llama usted, entonces voy, Don Barredora es quien yo soy 🎵"
	app.Version = "0.1.3"

	app.Flags = []cli.Flag{
		cli.StringFlag{"baseDir", "", "Directorio base para iniciar la búsqueda"},
		cli.StringFlag{"year", "", "Año para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"month", "", "Mes para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"day", "", "Día para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"backUpDir", "", "Directorio base para respaldo"},
	}
	app.Action = func(c *cli.Context) {
		globPatternList := GetGlobPatternList(
			c.String("baseDir"),
			c.String("year"),
			c.String("month"),
			c.String("day"))

		fmt.Printf("Directorios pendientes de procesar: %d\n", len(globPatternList))
		for _, globPattern := range globPatternList {
			files, _ := ListFiles(globPattern)
			fmt.Printf("%d archivos en directorio %s\n", len(files), globPattern)
			for _, filePath := range files {
				Replace(filePath, c)
			}
		}
	}
	app.Run(os.Args)
}
