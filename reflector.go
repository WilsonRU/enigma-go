package main

import "strings"

type Reflector struct {
	setting  string
	base     string
	settings map[string]string
	sequence string
}

func NewReflector(setting string) *Reflector {
	return &Reflector{
		setting: setting,
		base:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		settings: map[string]string{
			"A": "EJMZALYXVBWFCRQUONTSPIKHGD",
			"B": "YRUHQSLDPXNGOKMIEBFZCWVJAT",
			"C": "FVPJIAOYEDRZXWGCTKUQSBNMHL",
			"D": "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		},
		sequence: "",
	}
}

func (r *Reflector) sequenceSettings() string {
	return r.settings[r.setting]
}

func (r *Reflector) forward(index int) int {
	return strings.Index(r.sequence, string(r.base[index]))
}
