package aoc2022

import (
	"bufio"
	"image"
	"image/color"
	"io"
	"strconv"
	"strings"
)

const (
	d14Width  int = 900
	d14Height int = 300
)

var (
	d14Air  = color.Gray{}
	d14Rock = color.Gray{0x33}
	d14Sand = color.Gray{0xdd}
)

func d14Parse(input io.Reader) (img *image.Gray, ymax int) {
	img = image.NewGray(
		image.Rectangle{
			Min: image.Point{500 - (d14Width / 2), 0},
			Max: image.Point{500 + (d14Width / 2), d14Height},
		})
	for sc := bufio.NewScanner(input); sc.Scan(); {
		xys := strings.Split(sc.Text(), " -> ")
		var sega, segb image.Point
		for i := 0; i < len(xys); i++ {
			sega = segb
			xy := strings.Split(xys[i], ",")
			segb.X, _ = strconv.Atoi(xy[0])
			segb.Y, _ = strconv.Atoi(xy[1])

			if 0 != i {
				dx, dy := 0, 0
				switch {
				case sega.X < segb.X:
					dx = 1
				case sega.X > segb.X:
					dx = -1
				case sega.Y < segb.Y:
					dy = 1
				case sega.Y > segb.Y:
					dy = -1
				}
				for x, y := sega.X, sega.Y; x != segb.X || y != segb.Y; x, y = x+dx, y+dy {
					img.SetGray(x, y, d14Rock)
				}
			}
			img.SetGray(segb.X, segb.Y, d14Rock)
			if segb.Y > ymax {
				ymax = segb.Y
			}
		}
	}
	return
}

func d14DropSand(img *image.Gray, ymax int) (stopped bool) {
	for x, y := 500, 0; !stopped && y <= ymax+1; y++ {
		if d14Air == img.GrayAt(x, y+1) {
			continue
		}
		if d14Air == img.GrayAt(x-1, y+1) {
			x--
			continue
		}
		if d14Air == img.GrayAt(x+1, y+1) {
			x++
			continue
		}

		stopped = true
		img.SetGray(x, y, d14Sand)
	}

	return
}

type Day14Part1 struct{}

func (_ Day14Part1) Solve(input io.Reader) (answer string, err error) {
	acc := 0
	img, ymax := d14Parse(input)

	for ; ; acc++ {
		if stopped := d14DropSand(img, ymax); !stopped {
			break
		}
	}

	return strconv.Itoa(acc), nil
}

type Day14Part2 struct{}

func (_ Day14Part2) Solve(input io.Reader) (answer string, err error) {
	acc := 0
	img, ymax := d14Parse(input)

	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		img.SetGray(x, ymax+2, d14Rock)
	}

	for ; d14Air == img.GrayAt(500, 0); acc++ {
		d14DropSand(img, ymax)
	}

	return strconv.Itoa(acc), nil
}
