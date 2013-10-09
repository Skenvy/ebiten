// Copyright 2013 Hajime Hoshi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package rects

import (
	"github.com/hajimehoshi/go.ebiten"
	"github.com/hajimehoshi/go.ebiten/graphics"
	"github.com/hajimehoshi/go.ebiten/graphics/matrix"
	"image/color"
	"math/rand"
	"time"
)

type Rects struct {
	rectsTexture graphics.Texture
	rect         *graphics.Rect
	rectColor    *color.RGBA
}

func New() *Rects {
	return &Rects{
		rect:      &graphics.Rect{},
		rectColor: &color.RGBA{},
	}
}

func (game *Rects) Init(tf graphics.TextureFactory) {
	// TODO: fix
	game.rectsTexture = tf.NewTexture(256, 240)
}

func (game *Rects) Update(context ebiten.GameContext) {
	game.rect.X = rand.Intn(context.ScreenWidth())
	game.rect.Y = rand.Intn(context.ScreenHeight())
	game.rect.Width = rand.Intn(context.ScreenWidth() - game.rect.X)
	game.rect.Height = rand.Intn(context.ScreenHeight() - game.rect.Y)

	game.rectColor.R = uint8(rand.Intn(256))
	game.rectColor.G = uint8(rand.Intn(256))
	game.rectColor.B = uint8(rand.Intn(256))
	game.rectColor.A = uint8(rand.Intn(256))
}

func (game *Rects) Draw(g graphics.Context) {
	g.SetOffscreen(game.rectsTexture.ID())

	g.DrawRect(*game.rect, game.rectColor)

	g.SetOffscreen(g.Screen().ID())
	g.DrawTexture(game.rectsTexture.ID(),
		matrix.IdentityGeometry(),
		matrix.IdentityColor())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
