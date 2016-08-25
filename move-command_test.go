package main

import "testing"

func TestValidMoveCommand(t *testing.T) {
	cases := []struct {
		inX int 
		inY int
		inF DIRECTION
		outX int 
		outY int
		outF DIRECTION
	}{
		{1, 1, NORTH, 1, 2, NORTH},
		{1, 1, EAST, 2, 1, EAST},
		{1, 1, SOUTH, 1, 0, SOUTH},
		{1, 1, WEST, 0, 1, WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			MoveCommand(&robot)
			
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
		inF DIRECTION
		outX int 
		outY int
		outF DIRECTION
	}{
		{1, 1, NORTH, 1, 5, NORTH},
		{1, 1, EAST, 5, 1, EAST},
		{1, 1, SOUTH, 1, 0, SOUTH},
		{1, 1, WEST, 0, 1, WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			MoveCommand(&robot)
			MoveCommand(&robot)
			MoveCommand(&robot)
			MoveCommand(&robot)
			MoveCommand(&robot)
			MoveCommand(&robot)
			
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
		inF DIRECTION
	}{
		{-1, 1, NORTH},
		{1, -2, EAST},
		{-3, -3, SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := Robot{}
		    
			PlaceCommand(&robot, tableTop, c.inX, c.inY, c.inF)
			MoveCommand(&robot)
			MoveCommand(&robot)
			MoveCommand(&robot)

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