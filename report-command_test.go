package main

import "testing"

func TestValidReportCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF DIRECTION
	}{
		{1, 1, NORTH},
		{1, 2, EAST},
		{2, 1, SOUTH},
		{3, 3, WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			ReportCommand(&robot)
			
			if robot.X != c.inX {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != c.inY {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != c.inF {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != true {
				t.Errorf("should return valid robot")
			}
		}
	}
}

func TestInvalidRobotReportCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF DIRECTION
	}{
		{-1, 1, NORTH},
		{1, -1, EAST},
		{-1,-1, SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			ReportCommand(&robot);
			
			if robot.X != 0 {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != 0 {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != UNKNOWN_DIRECTION {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != false {
				t.Errorf("should return invalid robot")
			}
		}
	}
}