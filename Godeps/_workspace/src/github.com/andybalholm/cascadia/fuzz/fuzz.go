package fuzz

import "github.com/devacto/grobot/Godeps/_workspace/src/github.com/andybalholm/cascadia"

func Fuzz(data []byte) int {
	sel, err := cascadia.Compile(string(data))
	if err != nil {
		if sel != nil {
			panic("sel != nil on error")
		}
		return 0
	}
	return 1
}
