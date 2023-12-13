package cubes

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const regx = "[0-9]"

type Pull struct {
	Red   int
	Green int
	Blue  int
}

func (p Pull) Power() int {
	return p.Blue * p.Red * p.Green
}

type Game struct {
	ID    int
	Pulls []Pull
}

func extractNumber(s string, r *regexp.Regexp) (int, error) {
	return strconv.Atoi(strings.Join(r.FindAllString(s, 10), ""))
}

func NewGame(input string) (Game, error) {
	g := Game{
		Pulls: make([]Pull, 0),
	}
	p := strings.Split(input, ": ")
	if len(p) != 2 && len(p[0]) != 6 {
		return g, errors.New("unexpected input")
	}

	r, err := regexp.Compile(regx)
	if err != nil {
		return g, err
	}

	g.ID, err = extractNumber(p[0], r)
	if err != nil {
		return g, err
	}

	rawPulls := strings.Split(p[1], "; ")
	for _, rawPull := range rawPulls {
		var pull Pull
		rawValues := strings.Split(rawPull, ", ")
		for _, rawValue := range rawValues {
			val, err := extractNumber(rawValue, r)
			if err != nil {
				return g, err
			}
			if strings.Contains(rawValue, "red") {
				pull.Red = val
			}
			if strings.Contains(rawValue, "green") {
				pull.Green = val
			}
			if strings.Contains(rawValue, "blue") {
				pull.Blue = val
			}
		}
		g.Pulls = append(g.Pulls, pull)
	}

	return g, nil
}

func (g Game) Power() int {
	p := Pull{}

	for i, pull := range g.Pulls {
		if i == 0 {
			p = pull
			continue
		}
		if pull.Red > p.Red {
			p.Red = pull.Red
		}
		if pull.Green > p.Green {
			p.Green = pull.Green
		}
		if pull.Blue > p.Blue {
			p.Blue = pull.Blue
		}
	}

	return p.Power()
}
