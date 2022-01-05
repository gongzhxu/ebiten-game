package face

import (
	"embed"
)

//go:embed images/*
var Images embed.FS

//go:embed models/*
var Models embed.FS
