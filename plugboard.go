package main

import "strings"

type Plugboard struct {
	base    string
	mapping map[string]string
}

func NewPlugboard(mapping []string) *Plugboard {
	plug := &Plugboard{
		base:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		mapping: make(map[string]string),
	}

	for _, m := range plug.base {
		plug.mapping[string(m)] = string(m)
	}

	for _, m := range mapping {
		plug.mapping[string(m[0])] = string(m[1])
		plug.mapping[string(m[1])] = string(m[0])
	}

	return plug
}

func (plug *Plugboard) forward(c string) int {
	return strings.Index(plug.base, plug.mapping[c])
}

func (plug *Plugboard) reverse(index int) string {
	return plug.mapping[string(plug.base[index])]
}
