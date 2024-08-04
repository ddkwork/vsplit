package main

import (
	"vsplit"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("vsplit", func(w *unison.Window) {
		w.Content().AddChild(vsplit.New().Layout())
	})
}
