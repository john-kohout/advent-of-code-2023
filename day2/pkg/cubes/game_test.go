package cubes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	var tests = []struct {
		input string
		want  Game
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: Game{
				ID: 1,
				Pulls: []Pull{
					{
						Blue: 3,
						Red:  4,
					},
					{
						Red:   1,
						Green: 2,
						Blue:  6,
					},
					{
						Green: 2,
					},
				},
			},
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: Game{
				ID: 2,
				Pulls: []Pull{
					{
						Blue:  1,
						Green: 2,
					},
					{
						Green: 3,
						Blue:  4,
						Red:   1,
					},
					{
						Green: 1,
						Blue:  1,
					},
				},
			},
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: Game{
				ID: 3,
				Pulls: []Pull{
					{
						Green: 8,
						Blue:  6,
						Red:   20,
					},
					{
						Blue:  5,
						Red:   4,
						Green: 13,
					},
					{
						Green: 5,
						Red:   1,
					},
				},
			},
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: Game{
				ID: 4,
				Pulls: []Pull{
					{
						Green: 1,
						Red:   3,
						Blue:  6,
					},
					{
						Green: 3,
						Red:   6,
					},
					{
						Green: 3,
						Blue:  15,
						Red:   14,
					},
				},
			},
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: Game{
				ID: 5,
				Pulls: []Pull{
					{
						Red:   6,
						Blue:  1,
						Green: 3,
					},
					{
						Blue:  2,
						Red:   1,
						Green: 2,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		got, err := NewGame(tt.input)
		assert.NoError(t, err)
		assert.Equal(t, tt.want, got)
	}
}

func TestGame_Power(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  48,
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  12,
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:  1560,
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:  630,
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:  36,
		},
	}

	for _, tt := range tests {
		g, err := NewGame(tt.input)
		assert.NoError(t, err)
		got := g.Power()
		assert.Equal(t, tt.want, got)
	}
}

func TestPull_Power(t *testing.T) {
	var tests = []struct {
		input Pull
		want  int
	}{
		{
			input: Pull{
				Red:   4,
				Green: 8,
				Blue:  203,
			},
			want: 6496,
		},
	}

	for _, tt := range tests {
		got := tt.input.Power()
		assert.Equal(t, tt.want, got)
	}
}
