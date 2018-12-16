package model

import (
	"fmt"
	"testing"
)

func TestParseObject4d(t *testing.T) {
	fmt.Println(ParseObject4d("(2018-12-13-19-24-02,12.234251,112.234251,123)"))
}
