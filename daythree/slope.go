package daythree

// TobogganMap is a struct for holding information about a sled run
type TobogganMap struct {
	tmap [][]rune
	tree rune
	free rune
}

// Course is the direction the Toboggan goes down the slop
type Course struct {
	X int
	Y int
}

// Courses is a set of directions down a slope
type Courses []Course

// NewTobogganMap creates a new toboggan map
func NewTobogganMap(lines []string) *TobogganMap {
	tmap := make([][]rune, len(lines))
	for y := range lines {
		tmap[y] = make([]rune, len(lines[y]))
		for x, v := range lines[y] {
			tmap[y][x] = v
		}
	}

	return &TobogganMap{
		tmap: tmap,
		tree: rune('#'),
		free: rune('.'),
	}
}

// PlotCourse finds how many Trees are in the path
func (t *TobogganMap) PlotCourse(course Course) int {
	trees := 0

	x := course.X
	yLen := len(t.tmap)

	for y := course.Y; y < yLen; y += course.Y {
		xIndex := x % len(t.tmap[y])

		if t.tmap[y][xIndex] == t.tree {
			trees++
		}

		x += course.X
	}

	return trees
}

// ProbabilityOfSuddenArboreal the number of trees in each course multiplied together
func (t *TobogganMap) ProbabilityOfSuddenArboreal(courses Courses) int {
	trees := 1
	for _, c := range courses {
		trees *= t.PlotCourse(c)
	}

	return trees
}
