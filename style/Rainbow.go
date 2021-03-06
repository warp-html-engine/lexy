package style

// Code generated by mbndr/lexy
// DO NOT EDIT!
// @generated

import "github.com/mbndr/lexy"

var Rainbow = lexy.Style{
	Foreground: "#d1d9e1",
	Background: "#474949",

	TokenColors: map[lexy.TokenType]string{
		lexy.TokenKeyword:  "#cc99cc",
		lexy.TokenLiteral:  "#cc99cc",
		lexy.TokenBuiltin:  "#b5bd68",
		lexy.TokenOperator: "inherit", // TODO
		lexy.TokenComment:  "#969896",
		lexy.TokenString:   "#8abeb7",
		lexy.TokenNumber:   "#f99157",
	},
}
