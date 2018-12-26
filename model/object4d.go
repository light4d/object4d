package model

import (
	"regexp"
	"strings"
)

type Object4d struct {
	T           string
	Lng, Lat, H string
	M           int
}

func ParseObject4d(rawpath string) *Object4d {
	o := new(Object4d)
	exp, err := regexp.Compile(`^/\([0-9-]*,[0-9.]*,[0-9.]*,[0-9.]*\)`)
	if err != nil {
		panic(err)
	}

	o4d := exp.FindString(rawpath)

	if o4d != "" {
		t_3d := strings.Split(o4d[2:len(o4d)-1], ",")
		if len(t_3d) > 0 {
			o.T = t_3d[0]
		}
		if len(t_3d) > 1 {
			o.Lng = t_3d[1]
		}
		if len(t_3d) > 2 {
			o.Lat = t_3d[2]
		}

		if len(t_3d) > 3 {
			o.H = t_3d[3]
		}
	} else {
		return nil
	}

	return o
}
func (o Object4d) Bucket() string {
	return o.T
}
func (o Object4d) Objectname() string {
	return o.Lat + "_" + o.Lng + "_" + o.H
}
func (o Object4d) Url() string {
	return "(" + o.T + "," + o.Lng + "," + o.Lat + "," + o.H + ")"
}
