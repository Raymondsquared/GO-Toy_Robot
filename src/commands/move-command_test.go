package commands

import (
    "consts"
    "models"
    "testing"
)

func TestValidMoveCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{1, 1, consts.NORTH, 1, 2, consts.NORTH},
		{1, 1, consts.EAST, 2, 1, consts.EAST},
		{1, 1, consts.SOUTH, 1, 0, consts.SOUTH},
		{1, 1, consts.WEST, 0, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			var pr = NewPlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			pr.Execute()
			
			var mr = NewMoveCommand(&robot)
			mr.Execute()
			
			if robot.X != c.outX {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != c.outY {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != c.outF {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != true {
				t.Errorf("should return valid robot")
			}
		}
	}
}

func TestOutOfBoundMoveCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{1, 1, consts.NORTH, 1, 5, consts.NORTH},
		{1, 1, consts.EAST, 5, 1, consts.EAST},
		{1, 1, consts.SOUTH, 1, 0, consts.SOUTH},
		{1, 1, consts.WEST, 0, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			var pr = NewPlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			pr.Execute()
			
			var mr = NewMoveCommand(&robot)
			mr.Execute()
			mr.Execute()
			mr.Execute()
			mr.Execute()
			mr.Execute()
			mr.Execute()
			mr.Execute()
			
			if robot.X != c.outX {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != c.outY {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != c.outF {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != true {
				t.Errorf("should return valid robot")
			}
		}
	}
}

func TestNotPlacedRobotMoveCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
	}{
		{-1, 1, consts.NORTH},
		{1, -2, consts.EAST},
		{-3, -3, consts.SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			var pr = NewPlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			pr.Execute()
			
			var mr = NewMoveCommand(&robot)
			mr.Execute()
			mr.Execute()
			mr.Execute()
			
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