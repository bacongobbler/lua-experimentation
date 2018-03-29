package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

const luaFormulaTypeName = "formula"

// Formula provides instructions and metadata for Draft to install a plugin.
type Formula struct {
	// The canonical name of the software.
	Name string
	// A (short) description of the software.
	Description string
	// The homepage URL for the software.
	Homepage string
	// The URL used to download the binary distribution for this version of the formula. The file must be a gzipped tarball (.tar.gz) or a zipfile (.zip) for unpacking.
	URL string
	// Additional URLs for this version of the formula.
	Mirrors []string
	// To verify the cached download's integrity and security, we verify the SHA-256 hash matches what we've declared in the formula.
	SHA256 string
	// Caveats inform the user about any Draft-specific caveats regarding this formula.
	Caveats string
	// The version of the software.
	Version string
}

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(filepath.Join("Formula", fmt.Sprintf("%s.lua", os.Args[1]))); err != nil {
		panic(err)
	}
	var formula Formula
	if err := gluamapper.Map(L.GetGlobal(luaFormulaTypeName).(*lua.LTable), &formula); err != nil {
		panic(err)
	}
	fmt.Println(formula)
}
