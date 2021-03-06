package main

var SampleColors = map[byte]Color{
	'w': {1.0,1.0,1.0},
	'd': {0.0,0.0,0.0},
	'r': {1.0,0.0,0.0},
	'p': {0.5,0.0,0.5},
	'b': {0.0,0.0,1.0},
}
var SampleHeight = 8
var SampleWidth = 8
var SpritesSample = [][][]byte{
	{
		{'r','r','r','r','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','b','b','b','b'},
	},
	{
		{'r','r','r','w','r','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','b','w','b','b','b'},
	},
	{
		{'r','r','w','r','w','r','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','b','w','b','w','b','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','b','r','b','r','b','r','b'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'w','b','w','b','w','b','w','b'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'r','w','r','w','r','w','r','w'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','b','w','b','w','b','w','b'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'r','w','r','w','r','w','r','w'},
	},
	{
		{'w','w','w','w','b','w','b','w'},
		{'w','w','b','w','w','w','w','w'},
		{'b','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','r'},
		{'w','w','w','w','w','r','w','w'},
		{'w','r','w','r','w','w','w','w'},
	},
	{
		{'w','w','w','w','w','b','w','w'},
		{'w','w','w','b','w','w','w','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','r'},
		{'b','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','w','w','w','r','w','w','w'},
		{'w','w','r','w','w','w','w','w'},
	},
	{
		{'w','w','w','w','b','w','w','w'},
		{'w','w','b','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','r','w','w'},
		{'w','w','w','r','w','w','w','w'},
	},
	{
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','b','w','r','w','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','w','b','w','r','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
	},
	{
		{'w','w','w','r','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','r'},
		{'b','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','b','w','w','w'},
	},
	{
		{'w','r','w','w','r','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','b','w','w','b','w'},
	},
	{
		{'r','w','w','r','w','r','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','r','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','b','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','b','w','b','w','w','b'},
	},
	{
		{'r','w','r','w','r','w','r','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','b','w','b','w','b','w','b'},
	},
	{
		{'r','r','r','r','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','w','w','w','w'},
		{'w','w','w','w','b','b','b','b'},
	},
}
