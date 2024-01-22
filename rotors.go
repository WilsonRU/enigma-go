package main

import "strings"

type Rotor struct {
	setting    string
	ringoffset int
	base       string
	settings   map[string][][]string
	turnovers  []string
	notch      []string
	sequence   string
	turnover   bool
}

func NewRotor(settings []interface{}) *Rotor {
	rotor := &Rotor{
		setting:    settings[0].(string),
		ringoffset: settings[1].(int),
		base:       "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		settings: map[string][][]string{
			"I":    {{"EKMFLGDQVZNTOWYHXUSPAIBRCJ"}, {"R"}, {"Q"}},
			"II":   {{"AJDKSIRUXBLHWTMCQGZNPYFVOE"}, {"F"}, {"E"}},
			"III":  {{"BDFHJLCPRTXVZNYEIWGAKMUSQO"}, {"W"}, {"V"}},
			"IV":   {{"ESOVPZJAYQUIRHXLNFTGKDCMWB"}, {"K"}, {"J"}},
			"V":    {{"VZBRGITYUPSDNHLXAWMJQOFECK"}, {"A"}, {"Z"}},
			"VI":   {{"JPGVOUMFYQBENHZRDKASXLICTW"}, {"AN"}, {"ZM"}},
			"VII":  {{"NZJHGRCXMYSWBOUFAIVLPEKQDT"}, {"AN"}, {"ZM"}},
			"VIII": {{"FKQHTLXOCBJSPDZRAMEWNIUYGV"}, {"AN"}, {"ZM"}},
		},
		turnovers: nil,
		notch:     nil,
		sequence:  "",
		turnover:  false,
	}

	rotor.turnovers = rotor.settings[rotor.setting][1]
	rotor.notch = rotor.settings[rotor.setting][2]
	rotor.reset()

	return rotor
}

func (r *Rotor) reset() {
	r.base = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r.sequence = r.sequenceSettings()
	r.ringSettings()
}

func (r *Rotor) sequenceSettings() string {
	return r.settings[r.setting][0][0]
}

func (r *Rotor) ringSettings() {
	for i := 0; i < r.ringoffset; i++ {
		r.rotate()
	}
}

func (r *Rotor) forward(index int) int {
	return strings.Index(r.sequence, string(r.base[index]))
}

func (r *Rotor) reverse(index int) int {
	if index < 0 || index >= len(r.sequence) {
		return 0
	}
	return strings.Index(r.base, string(r.sequence[index]))
}

func (r *Rotor) rotate() {
	r.base = r.base[1:] + r.base[:1]
	r.sequence = r.sequence[1:] + r.sequence[:1]
	if strings.Contains(r.turnovers[0], string(r.base[0])) {
		r.turnover = true
	}
}
