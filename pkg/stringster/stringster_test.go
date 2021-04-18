package stringster_test

import (
	"github.com/matryer/is"
	"github.com/neonima/ffw/pkg/stringster"
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

func TestGetExtensionFromString(t *testing.T) {
	tests := []struct {
		title    string
		input    string
		expected string
	}{
		{
			title:    "it should get the extension",
			input:    "hhg/1.png",
			expected: ".png",
		},
		{
			title:    "it should also get the extension even for long complicated once",
			input:    "hhg/1.verylooooooon-extension",
			expected: ".verylooooooon-extension",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			is := is.New(t)
			res := stringster.GetExtensionFromString(test.input)
			is.Equal(test.expected, res)

		})
	}
}

func TestIsExtension(t *testing.T) {
	tests := []struct {
		title          string
		inputFile      string
		inputExtension string
		expected       bool
	}{
		{
			title:          "it should return true if the extension match the file",
			inputFile:      "hhg/1.png",
			inputExtension: ".png",
			expected:       true,
		},
		{
			title:          "it should return false if the extension does not match the file",
			inputFile:      "hhg/1.verylooooooon-extension",
			inputExtension: ".png",
			expected:       false,
		},
		{
			title:          "it should return true if no extensions are provided",
			inputFile:      "hhg/1.verylooooooon-extension",
			inputExtension: "",
			expected:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			is := is.New(t)
			res := stringster.IsExtension(test.inputFile, test.inputExtension)
			if len(test.inputExtension) == 0 {
				res = stringster.IsExtension(test.inputFile)
			}
			is.Equal(test.expected, res)

		})
	}
}
