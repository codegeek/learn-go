package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) (league League, err error) {
	err = json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("Problem parsing league, %v", err)
	}
	return
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
