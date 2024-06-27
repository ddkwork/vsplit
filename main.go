package main

import (
	"vsplit"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("vsplit", func(w *unison.Window) {
		w.Content().AddChild(vsplit.New().Layout())
	})
}
