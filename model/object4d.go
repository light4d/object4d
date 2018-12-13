package model

type Object4d struct {
	T           string
	Lng, Lat, H string
	M           int
}

func ParseObject4d(rawpath string) Object4d {
	o := Object4d{}

	return o
}
func (o Object4d) Bucket() string {
	return o.T
}
func (o Object4d) Objectname() string {
	return o.Lat + "_" + o.Lng + "_" + o.H
}
