package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/canvas"
  "fyne.io/fyne/v2/container"
  "os"
  "time"
)

func main() {
  a := app.New()
  w := a.NewWindow("Netgraph")
  width := float32(GRAPH_WIDTH)
  height := float32(GRAPH_HEIGHT)
  w.Resize(fyne.NewSize(width, height))
  w.SetFixedSize(true)
  ring := NewRing(30)
  img := updateChart(ring, width, height)
  con := container.NewWithoutLayout(img)
  w.SetContent(con)
  w.Canvas().SetOnTypedKey(
    func(ev *fyne.KeyEvent) {
      key := string(ev.Name)
      switch key {
        case "Q":
          os.Exit(0)
      }
    }
  )

  go func() {
    for {
      select {
        case <- time.After(5 * time.Second):
          img = updateChart(ring, width, height)
          con.Refresh()
      }
    }
  }()
  w.ShowAndRun()
}

func updateChart(ring *Dpoints, width, height float32) *canvas.Image {
  drawChart(ring)
  img := canvas.NewImageFromFile(GRAPH_FILE)
  img.FillMode = canvas.ImageFillOriginal
  img.Resize(fyne.NewSize(width, height))
  return img
}
