package stringster_test

import (
	"github.com/matryer/is"
	"github.com/neonima/fftw/pkg/stringster"
	"testing"
)

func TestRemoveSpecialChars(t *testing.T) {
	tests := []struct {
		title    string
		input    string
		expected string
	}{
		{
			title:    "it should remove all special chars",
			input:    "',#£\\'",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			is := is.New(t)
			res := stringster.RemoveSpecialChars(test.input)
			is.Equal(test.expected, res)

		})
	}
}

func TestRemoveAccent(t *testing.T) {
	tests := []struct {
		title    string
		input    string
		expected string
	}{
		{
			title:    "it should replace all accent",
			input:    "éléphants",
			expected: "elephants",
		},
		{
			title:    "it should replace all accent",
			input:    "žůžo, ïju'e",
			expected: "zuzo, iju'e",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			is := is.New(t)
			res := stringster.RemoveAccent(test.input)
			is.Equal(test.expected, res)

		})
	}
}

func TestRenameFile(t *testing.T) {
	tests := []struct {
		title    string
		input    string
		expected string
	}{
		{
			title:    "it should replace all accent",
			input:    "hhg/Mon Fichier trés chiant.png",
			expected: "hhg/mon_fichier_tres_chiant.png",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			is := is.New(t)
			res := stringster.RenameFile(test.input)
			is.Equal(test.expected, res)

		})
	}
}
