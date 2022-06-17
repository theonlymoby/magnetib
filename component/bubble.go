package component

import "github.com/hajimehoshi/ebiten"

type Bubble struct {
	Params        Object
	PositiveImage *ebiten.Image
	NegativeImage *ebiten.Image
	Positive      bool
	Width         int
	Height        int
}

func (b *Bubble) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(b.Params.X, b.Params.Y)
	if b.Positive {
		screen.DrawImage(b.PositiveImage, o)
	} else {
		screen.DrawImage(b.NegativeImage, o)
	}
}
