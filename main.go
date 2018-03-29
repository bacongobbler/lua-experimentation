package main

import (
	"fmt"
	
	"github.com/yuin/gopher-lua"
)

const luaFormulaTypeName = "formula"

type Formula struct {
	Name string
	Description string
	Homepage string
	SHA256 string
	URL string
	Version string
}

func(f *Formula) String() string {
	return fmt.Sprintf(
		"%s@%s\n%s\nHomepage: %s\nURL: %s\nsha256: %s",
		f.Name,
		f.Version,
		f.Description,
		f.Homepage,
		f.URL,
		f.SHA256,
	)
}

// Registers my Formula type to given L.
func registerFormulaType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaFormulaTypeName)
	L.SetGlobal("formula", mt)
	// static attributes
	L.SetField(mt, "new", L.NewFunction(newFormula))
	// methods
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), formulaMethods))
}

// Constructor
func newFormula(L *lua.LState) int {
	formula := &Formula{}
	ud := L.NewUserData()
	ud.Value = formula
	L.SetMetatable(ud, L.GetTypeMetatable(luaFormulaTypeName))
	L.Push(ud)
	return 1
}

// Checks whether the first lua argument is a *LUserData with *Formula and returns this *Formula.
func checkFormula(L *lua.LState) *Formula {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Formula); ok {
		return v
	}
	L.ArgError(1, "formula expected")
	return nil
}

var formulaMethods = map[string]lua.LGFunction{
	"name": formulaGetSetName,
	"description": formulaGetSetDescription,
	"homepage": formulaGetSetHomepage,
	"sha256": formulaGetSetSHA256,
	"url": formulaGetSetURL,
	"version": formulaGetSetVersion,
}

// Getter and setter for the Formula#Name
func formulaGetSetName(L *lua.LState) int {
	f := checkFormula(L)
	if L.GetTop() == 2 {
		f.Name = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.Name))
	return 1
}

// Getter and setter for the Formula#Description
func formulaGetSetDescription(L *lua.LState) int {
	f := checkFormula(L)
	if L.GetTop() == 2 {
		f.Description = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.Description))
	return 1
}

// Getter and setter for the Formula#Homepage
func formulaGetSetHomepage(L *lua.LState) int {
	f := checkFormula(L)
	if L.GetTop() == 2 {
		f.Homepage = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.Homepage))
	return 1
}

// Getter and setter for the Formula#SHA256
func formulaGetSetSHA256(L *lua.LState) int {
	f := checkFormula(L)
	if L.GetTop() == 2 {
		f.SHA256 = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.SHA256))
	return 1
}

// Getter and setter for the Formula#URL
func formulaGetSetURL(L *lua.LState) int {
	f := checkFormula(L)
	if L.GetTop() == 2 {
		f.URL = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.URL))
	return 1
}

// Getter and setter for the Formula#Version
func formulaGetSetVersion(L *lua.LState) int {
	f := checkFormula(L)
	if L.GetTop() == 2 {
		f.Version = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.Version))
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()
	registerFormulaType(L)
	if err := L.DoFile("main.lua"); err != nil {
			panic(err)
	}
	f := checkFormula(L)
	fmt.Println(f)
}
