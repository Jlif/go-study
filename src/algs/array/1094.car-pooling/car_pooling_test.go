package _094_car_pooling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question1094 struct {
	para1094
	ans1094
}

type para1094 struct {
	trips    [][]int
	capacity int
}

type ans1094 struct {
	ans bool
}

func TestCarPooling(t *testing.T) {

	cases := []question1094{
		{para1094{trips: [][]int{{2, 1, 5}, {3, 3, 7}}, capacity: 4}, ans1094{false}},
		{para1094{trips: [][]int{{2, 1, 5}, {3, 3, 7}}, capacity: 5}, ans1094{true}},
		{para1094{trips: [][]int{{2, 1, 5}, {3, 5, 7}}, capacity: 3}, ans1094{true}},
	}

	for _, v := range cases {
		assert.New(t).Equal(carPooling(v.para1094.trips, v.para1094.capacity), v.ans)
	}

}
