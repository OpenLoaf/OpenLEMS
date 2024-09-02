package main

import (
	_ "application/internal/logic"

	"application/internal/cmd"
	"context"
)

func main() {
	cmd.Main.Run(context.Background())
}
