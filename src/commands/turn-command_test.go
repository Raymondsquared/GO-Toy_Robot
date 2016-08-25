package commands

import (
	"consts"
    "models"
    "testing"
)

func TestValidTurnLeftCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{1, 1, consts.NORTH, 1, 1, consts.WEST},
		{1, 1, consts.EAST, 1, 1, consts.NORTH},
		{1, 1, consts.SOUTH, 1, 1, consts.EAST},
		{1, 1, consts.WEST, 1, 1, consts.SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, consts.LEFT)
			
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

func TestValidTurnRightCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{1, 1, consts.NORTH, 1, 1, consts.EAST},
		{1, 1, consts.EAST, 1, 1, consts.SOUTH},
		{1, 1, consts.SOUTH, 1, 1, consts.WEST},
		{1, 1, consts.WEST, 1, 1, consts.NORTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, consts.RIGHT)
			
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

func TestValid4TimesTurnCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF consts.DIRECTION
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{1, 1, consts.NORTH, 1, 1, consts.NORTH},
		{1, 1, consts.EAST, 1, 1, consts.EAST},
		{1, 1, consts.SOUTH, 1, 1, consts.SOUTH},
		{1, 1, consts.WEST, 1, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, consts.RIGHT)
			TurnCommand(&robot, consts.RIGHT)
			TurnCommand(&robot, consts.RIGHT)
			TurnCommand(&robot, consts.RIGHT)
			TurnCommand(&robot, consts.LEFT)
			TurnCommand(&robot, consts.LEFT)
			TurnCommand(&robot, consts.LEFT)
			TurnCommand(&robot, consts.LEFT)
			
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

func TestInvalidRobotTurnCommand(t *testing.T) {
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
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, consts.RIGHT)
			TurnCommand(&robot, consts.RIGHT)
			TurnCommand(&robot, consts.LEFT)
			
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