package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/edmt/sweeper/fs"
	"github.com/edmt/sweeper/xmlreplacer"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sweeper"
	app.Usage = "Me llama usted, entonces voy, Don Barredora es quien yo soy 🎵"
	app.Version = "0.1.2"

	app.Flags = []cli.Flag{
		cli.StringFlag{"baseDir", "", "Directorio base para iniciar la búsqueda"},
		cli.StringFlag{"year", "", "Año para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"month", "", "Mes para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"day", "", "Día para formar el patrón en la búsqueda de directorios"},
		cli.StringFlag{"backUpDir", "", "Directorio base para respaldo"},
	}
	app.Action = func(c *cli.Context) {
		globPatternList := fs.GetGlobPatternList(
			c.String("baseDir"),
			c.String("year"),
			c.String("month"),
			c.String("day"))

		fmt.Printf("Directorios pendientes de procesar: %d\n", len(globPatternList))
		for _, globPattern := range globPatternList {
			files, _ := fs.ListFiles(globPattern)
			fmt.Printf("%d archivos en directorio %s\n", len(files), globPattern)
			for _, filePath := range files {
				whenReplaced := func() {
					backUpFilePath := xmlreplacer.Format(filePath, c.String("baseDir"), c.String("backUpDir"))
					fs.Mkdir(backUpFilePath)
					fs.Cp(filePath, backUpFilePath)
				}
				xmlreplacer.Replace(filePath, whenReplaced)
			}
		}
	}
	app.Run(os.Args)
}
