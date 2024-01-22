package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Enigma struct {
	numcycles        int
	rotors           []*Rotor
	reflector        *Reflector
	plugboard        *Plugboard
	rotorsettings    []interface{}
	reflectorsetting string
	plugboardsetting []string
}

func NewEnigma() *Enigma {
	enigma := &Enigma{
		numcycles:        0,
		rotors:           make([]*Rotor, 0),
		reflector:        nil,
		plugboard:        nil,
		rotorsettings:    []interface{}{"VIII", 0, "V", 0, "II", 0, "IV", 0},
		reflectorsetting: "C",
		plugboardsetting: []string{},
	}

	enigma.plugboard = NewPlugboard(enigma.plugboardsetting)
	for i := 0; i < len(enigma.rotorsettings); i += 2 {
		settings := []interface{}{enigma.rotorsettings[i], enigma.rotorsettings[i+1]}
		enigma.rotors = append(enigma.rotors, NewRotor(settings))
	}
	enigma.reflector = NewReflector(enigma.reflectorsetting)

	return enigma
}

func (e *Enigma) printSetup() {
	for _, r := range e.rotors {
		fmt.Printf("%s\t%s\n", r.setting, r.sequence)
	}
}

func (e *Enigma) reset() {
	e.numcycles = 0
	for _, r := range e.rotors {
		r.reset()
	}
}

func (e *Enigma) encode(c string) string {
	c = strings.ToUpper(c)
	if !isAlpha(c) {
		return c
	}
	e.rotors[0].rotate()
	if strings.Contains(e.rotors[1].base, e.rotors[1].notch[0]) {
		e.rotors[1].rotate()
	}
	for i := 0; i < len(e.rotors)-1; i++ {
		if e.rotors[i].turnover {
			e.rotors[i].turnover = false
			e.rotors[i+1].rotate()
		}
	}
	index := e.plugboard.forward(c)
	for _, r := range e.rotors {
		index = r.forward(index)
	}
	index = e.reflector.forward(index)
	for i := len(e.rotors) - 1; i >= 0; i-- {
		index = e.rotors[i].reverse(index)
	}
	c = e.plugboard.reverse(index)
	return c
}

// TODO: Implementar metodo de decode
func (e *Enigma) decode(d string) string {
	d = strings.ToUpper(d)
	return d
}

func isAlpha(s string) bool {
	return s != "" && strings.IndexFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	}) == -1
}
