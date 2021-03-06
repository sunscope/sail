package library

import (
	"errors"
	"fmt"
	"regexp"
	"sort"

	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/devinmcgloin/sail/pkg/sketch/accrew"
	"github.com/devinmcgloin/sail/pkg/sketch/delaunay"
	"github.com/devinmcgloin/sail/pkg/sketch/gradients"
	"github.com/devinmcgloin/sail/pkg/sketch/harmonograph"
	"github.com/devinmcgloin/sail/pkg/sketch/primitives"
	"github.com/devinmcgloin/sail/pkg/sketch/sampling"
)

// Sketches defines all the sketches that the system can render
var options = map[string]sketch.Renderable{
	"accrew/dot-clouds":            accrew.DotCloud{},
	"accrew/disjoint-line-clouds":  accrew.DisjointLineCloud{},
	"accrew/joint-line-clouds":     accrew.JointLineCloud{},
	"accrew/dot-lines":             accrew.DotLines{},
	"delaunay/ring":                delaunay.Ring{},
	"delaunay/mesh":                delaunay.Mesh{},
	"sampling/uniform-rectangle":   sampling.UniformRectangleDot{},
	"sampling/radial-rectangle":    sampling.RadialRectangleDot{},
	"sampling/dot-walk":            sampling.DotWalk{},
	"primitive/line-coloring":      primitives.LineColoring{},
	"primitive/bars":               primitives.Bars{},
	"primitive/rotated-lines":      primitives.RotatedLines{},
	"primitive/falling-rectangles": primitives.FallingRectangles{},
	"gradients/skyspace":           gradients.Skyspace{},
	"harmonograph/single":          harmonograph.Single{},
	"harmonograph/dual":            harmonograph.Dual{},
	"harmonograph/variable":        harmonograph.Variable{},
	"harmonograph/offset":          harmonograph.Offset{},
}

// Lookup finds a sketch based on the sketchID
func Lookup(sketchID string) (sketch.Renderable, error) {
	sketch, ok := options[sketchID]
	if !ok {
		return nil, errors.New("invalid sketch ID")
	}
	return sketch, nil
}

// Exists returns true if the sketch is defined, false otherwise.
func Exists(sketchID string) bool {
	_, ok := options[sketchID]
	return ok
}

// List prints all avaliable sketches
func List(regex string) []string {
	var sketchIDs []string

	for sketchID := range options {
		matched, err := regexp.MatchString(regex, sketchID)
		if err != nil {
			fmt.Printf("%s -> %s\n", sketchID, err)
			continue
		}

		if matched && err == nil {
			sketchIDs = append(sketchIDs, sketchID)
		}
	}
	sort.Strings(sketchIDs)
	return sketchIDs
}
