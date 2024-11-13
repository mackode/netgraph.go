package main

import (
  "container/ring"
  "time"
)

type Dpoints struct {
  rp *ring.Ring
}

type Dpoint struct {
  dt time.Time
  up float64
  down float64
}

func NewRing(n int) *Dpoints {
  return &Dpoints{rp: ring.New(n)}
}

func (d *Dpoints) Add(up, down float64) {
  d.rp.Value = Dpoint{
    dt: time.Now(),
    up: up,
    down: down,
  }
  d.rp = d.rp.Next()
}

func (d Dpoints) All() ([]time.Time, []float64, []float64) {
  ups, downs := []float64, []float64{}
  times := []time.Time{}
  r := d.rp
  n := 0
  for i := 0; i < d.rp.Len(); i++ {
    r = r.Prev()
    if r.Value == nil {
      r = r.Next()
      break
    }
    n++
  }
  for i := 0; i < n; i++ {
    dp := r.Value.(Dpoint)
    times = append(times, dp.dt)
    ups = append(ups, dp.up)
    downs = append(downs, dp.down)
    r = r.Next()
  }
  return times, ups, downs
}
