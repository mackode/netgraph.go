package main

import (
  "fmt"
  "github.com/wcharczuk/go-chart/v2"
  "os"
  "time"
)

const GRAPH_FILE = "netgraph.png"
const GRAPH_WIDTH = 1920
const GRAPH_HEIGHT = 1000

func drawChart(ring *Dpoints) {
  up, down, err := fetchUpDown()
  if err != nil {
    panic(err)
  }
  ring.Add(up, down)
  times, ups, downs := ring.All()
  xAxisCfg := chart.XAxis{
    ValueFormatter: func(v interface{}) string {
      return time.Unix(0, int64(v.(float64))).Format("03:04:05")
    },
  }
  yAxisCfg := chart.YAxis{
    Range: &chart.LogarithmicRange{
      Max: 100000,
    },
    ValueFormatter: func(v interface{}) string {
      return fmt.Sprintf("%.2f MBps", v.(float64) / 1000.0)
    },
  }
  upseries := chart.TimeSeries{
    XValues: times,
    YValues: ups,
    Style: chart.Style{
      StrokeColor: chart.ColorCyan,
      StrokeWidth: 10,
      FillColor: chart.ColorGreen.WithAlpha(64),
    },
  }
  downseries := chart.TimeSeries{
    XValues: times,
    YValues: downs,
    Style: chart.Style{
      StrokeColor: chart.ColorRed,
      StrokeWidth: 10,
      FillColor: chart.ColorBlue.WithAlpha(64),
    },
  }
  graph := chart.Chart{
    XAxis: xAxisCfg,
    YAxis: yAxisCfg,
    Height: GRAPH_HEIGHT,
    Width: GRAPH_WIDTH,
    Series: []chart.Series{
      upseries,
      downseries,
    },
  }
  f, _ := os.Create(GRAPH_FILE)
  defer f.Close()
  graph.Render(chart.PNG, f)
}
