package commands

import (
    "consts"
    "models"
    "testing"
)

func TestValidPlaceCommand(t *testing.T) {
	cases := []struct {
		inX, inY int
		intF consts.DIRECTION
	}{
		{0, 0, consts.NORTH},
		{1, 1, consts.SOUTH},
		{5, 5, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.intF)
			
			if robot.X != c.inX {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != c.inY {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != c.intF {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != true {
				t.Errorf("should return valid robot")
			}
		}
	}
}

func TestNegativeXYForPlaceCommand(t *testing.T) {
	cases := []struct {
		inX, inY int
	}{
		{-1, 2},
		{2, -2},
		{-3, -3},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, consts.SOUTH)
			
			if robot.X != 0 {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != 0 {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != consts.UNKNOWN_DIRECTION {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != false {
				t.Errorf("should return invalid robot")
			}
		}
	}
}


func TestOutOfBoundForPlaceCommand(t *testing.T) {
	cases := []struct {
		inX, inY int
	}{
		{3, 2},
		{2, 4},
		{5, 5},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(2, 2)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, consts.SOUTH)
			
			if robot.X != 0 {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != 0 {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != consts.UNKNOWN_DIRECTION {
				t.Errorf("should return valid Direction")
			}
			
			
			if robot.IsValid != false {
				t.Errorf("should return invalid robot")
			}
		}
	}
}