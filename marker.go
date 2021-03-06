package colorpicker

import (
	"image/color"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
)

var (
	markerFillColor   = color.RGBA{50, 50, 50, 120}
	markerStrokeColor = color.RGBA{50, 50, 50, 200}
)

type marker struct {
	*canvas.Circle
	center fyne.Position
	radius float64
}

func newMarker(radius float64, strokeWidth int) *marker {
	marker := &marker{
		Circle: &canvas.Circle{
			FillColor:   markerFillColor,
			StrokeColor: markerStrokeColor,
			StrokeWidth: float32(strokeWidth),
		},
		radius: radius,
	}
	marker.setPosition(fyne.NewPos(0, 0))
	return marker
}

func (m *marker) setPosition(p fyne.Position) {
	m.center = p
	m.Position1 = fyne.NewPos(p.X-int(m.radius), p.Y-int(m.radius))
	m.Position2 = fyne.NewPos(p.X+int(m.radius), p.Y+int(m.radius))
}

func (m *marker) setPositionY(y int) {
	m.setPosition(fyne.NewPos(m.center.X, y))
}

type circleBarMarker struct {
	*marker
	cx, cy float64
}

func newCircleBarMarker(w, h int, hueBarWidth float64) *circleBarMarker {
	fw := float64(w)
	fh := float64(h)
	fr := hueBarWidth / 2
	marker := &circleBarMarker{
		marker: newMarker(fr, 2),
		cx:     fw / 2,
		cy:     fh / 2,
	}
	markerCenter := fyne.NewPos(int(math.Round(fw-fr)), int(math.Round(fh/2)))
	marker.marker.setPosition(markerCenter)
	return marker
}

func (m *circleBarMarker) setCircleMarkerPosition(p fyne.Position) {
	v := newVectorFromPoints(m.cx, m.cy, float64(p.X), float64(p.Y))
	nv := v.normalize()
	center := newVector(m.cx, m.cy)
	markerCenter := center.add(nv.multiply(m.cx - m.radius)).toPosition()
	m.marker.setPosition(markerCenter)
}

func (m *circleBarMarker) setCircleMarekerPositionFromHue(hue float64) {
	rad := -2 * math.Pi * hue
	center := newVector(m.cx, m.cy)
	dir := newVector(1, 0).rotate(rad).multiply(m.cx - m.radius)
	markerCenter := center.add(dir).toPosition()
	m.marker.setPosition(markerCenter)
}

func (m *circleBarMarker) calcHueFromCircleMarker(p fyne.Position) float64 {
	v := newVectorFromPoints(m.cx, m.cy, float64(p.X), float64(p.Y))
	baseV := newVector(1, 0)
	rad := math.Acos(baseV.dot(v) / (v.norm() * baseV.norm()))
	if float64(p.Y)-m.cy >= 0 {
		rad = math.Pi*2 - rad
	}
	hue := rad / (math.Pi * 2)
	return hue
}
