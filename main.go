package main

import (
	"embed"
	"github.com/zxl00/go-cheetah-admin/api"
	"github.com/zxl00/go-cheetah-admin/cmd"
)

//go:embed all:ui/dist/*
var embedFS embed.FS

func main() {
	api.SetFS(embedFS)
	cmd.Run()
}
