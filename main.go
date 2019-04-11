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

		// Pack folder to .pack
		case mode.IsDir():
			var pack packs.Pack
			pack.LoadFromDir(path)
			pack.WritePack("Packed", "Assets_256")

		// Unpack .pack
		case mode.IsRegular():
			var pack packs.Pack
			pack.LoadFromFile(path)
			pack.Unpack("Unpacked")
		}
	}
}
