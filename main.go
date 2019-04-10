package main

import (
	"fl-library/files/packs"
	"fl-library/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	argc := len(os.Args)

	if argc < 2 {
		log.Fatal("Missing arguments")
	}

	for _, a := range os.Args[1:] {

		// Grab Absolute Path
		path, err := filepath.Abs(a)
		utils.Check(err)

		// Check if dir. if so, pack
		file, err := os.Stat(path)
		utils.Check(err)

		switch mode := file.Mode(); {
		case mode.IsDir():
			// TODO

			var pack packs.Pack
			pack.LoadFromDir(path)
			pack.Display()
			pack.WritePack(".", "Assets_256.pack") // TODO: inc number

		case mode.IsRegular():
			var pack packs.Pack
			pack.LoadFromFile(path)
			pack.Display()
			pack.Unpack(".")
		}
	}
}
