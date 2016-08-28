package commands

import (
	"consts"
    "models"
    "testing"
)

func TestValidReportCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
	}{
		{1, 1, consts.NORTH},
		{1, 2, consts.EAST},
		{2, 1, consts.SOUTH},
		{3, 3, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			var pr = NewPlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			pr.Execute()
			
			var rr = NewReportCommand(&robot)
			rr.Execute()
			
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
		inF consts.DIRECTION
	}{
		{-1, 1, consts.NORTH},
		{1, -1, consts.EAST},
		{-1,-1, consts.SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			var pr = NewPlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			pr.Execute()
			
			var rr = NewReportCommand(&robot);
			rr.Execute()
			
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