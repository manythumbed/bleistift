package bleistift

import "bitbucket.org/zombiezen/gopdf/pdf"

func project(a, b pdf.Point, d pdf.Unit) pdf.Point {
	return plus(a, multiply(minus(b, a), d))
}

func plus(a, b pdf.Point) pdf.Point {
	return pdf.Point{pdf.Unit{a.X + b.X}, pdf.Unit{a.Y + b.Y}}
}

func minus(a, b pdf.Point) pdf.Point {
	return pdf.Point{a.X - b.X, a.Y - b.Y}
}

func multiply(a pdf.Point, s pdf.Unit) pdf.Point {
	return pdf.Point{s * a.X, s * a.Y}
}
