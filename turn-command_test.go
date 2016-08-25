package main

import "testing"

func TestValidTurnLeftCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF DIRECTION
		outX int 
		outY int
		outF DIRECTION
	}{
		{1, 1, NORTH, 1, 1, WEST},
		{1, 1, EAST, 1, 1, NORTH},
		{1, 1, SOUTH, 1, 1, EAST},
		{1, 1, WEST, 1, 1, SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, LEFT)
			
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
		inF DIRECTION
		outX int 
		outY int
		outF DIRECTION
	}{
		{1, 1, NORTH, 1, 1, EAST},
		{1, 1, EAST, 1, 1, SOUTH},
		{1, 1, SOUTH, 1, 1, WEST},
		{1, 1, WEST, 1, 1, NORTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, RIGHT)
			
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
		inF DIRECTION
		outX int 
		outY int
		outF DIRECTION
	}{
		{1, 1, NORTH, 1, 1, NORTH},
		{1, 1, EAST, 1, 1, EAST},
		{1, 1, SOUTH, 1, 1, SOUTH},
		{1, 1, WEST, 1, 1, WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			TurnCommand(&robot, RIGHT)
			TurnCommand(&robot, RIGHT)
			TurnCommand(&robot, RIGHT)
			TurnCommand(&robot, RIGHT)
			TurnCommand(&robot, LEFT)
			TurnCommand(&robot, LEFT)
			TurnCommand(&robot, LEFT)
			TurnCommand(&robot, LEFT)
			
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
			TurnCommand(&robot, RIGHT)
			TurnCommand(&robot, RIGHT)
			TurnCommand(&robot, LEFT)
			
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