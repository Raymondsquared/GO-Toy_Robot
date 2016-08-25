package main

import (
    "consts"
    "models"
    "testing"
)

func TestInvalidCommand(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{""},
		{" "},
		{"PUT"},
		{"TEST"},
		{"EXIT A"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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

func TestIgnoreEverythingBeforePlaceCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
		outValid bool
	}{
		{"MOVE", 0, 0, consts.UNKNOWN_DIRECTION, false},
		{"LEFT", 0, 0, consts.UNKNOWN_DIRECTION, false},
		{"RIGHT", 0, 0, consts.UNKNOWN_DIRECTION, false},
		{"REPORT", 0, 0, consts.UNKNOWN_DIRECTION, false},
		{"PLACE 1,2,NORTH", 1, 2, consts.NORTH, true},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
			if robot.X != c.outX {
				t.Errorf("should return valid X")
			}
			
			if robot.Y != c.outY {
				t.Errorf("should return valid Y")
			}
			
			if robot.Direction != c.outF {
				t.Errorf("should return valid Direction")
			}
			
			if robot.IsValid != c.outValid {
				t.Errorf("should return valid robot")
			}
		}
	}
}

func TestValidPlaceCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{"PLACE 1,2,NORTH", 1, 2, consts.NORTH},
		{"PLACE 2,1,EAST", 2, 1, consts.EAST},
		{"PLACE 1,0,SOUTH", 1, 0, consts.SOUTH},
		{"PLACE 0,1,WEST", 0, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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

func TestInvalidPlaceCommand(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{"PLACE -1,2,NORTH"},
		{"PLACE 2,-1,EAST"},
		{"PLACE -1,-,SOUTH"},
		{"PLACE 1,6,WEST"},
		{"PLACE 7,1,WEST"},
		{"PLACE 9,8,WEST"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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
				t.Errorf("should return valid robot")
			}
		}
	}
}

func TestInvalidPlaceCommandFormat(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{"PLACE"},
		{"PLACE "},
		{"PLACE 1"},
		{"PLACE 1,2"},
		{"PLACE 1,2,3"},
		{"PLACE 1,2,NORTHABC"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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

func TestValidMoveCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{"PLACE 1,1,NORTH", 1, 2, consts.NORTH},
		{"PLACE 1,1,EAST", 2, 1, consts.EAST},
		{"PLACE 1,1,SOUTH", 1, 0, consts.SOUTH},
		{"PLACE 1,1,WEST", 0, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			ProcessUserInput("MOVE", &robot, tableTop)
			
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

func TestValidMove5TimesCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{"PLACE 1,1,NORTH", 1, 5, consts.NORTH},
		{"PLACE 1,1,EAST", 5, 1, consts.EAST},
		{"PLACE 1,1,SOUTH", 1, 0, consts.SOUTH},
		{"PLACE 1,1,WEST", 0, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			ProcessUserInput("MOVE", &robot, tableTop)
			ProcessUserInput("MOVE", &robot, tableTop)
			ProcessUserInput("MOVE", &robot, tableTop)
			ProcessUserInput("MOVE", &robot, tableTop)
			ProcessUserInput("MOVE", &robot, tableTop)
			
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



func TestInvalidMoveBeforePlaceFormat(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{"MOVE"},
		{"MOVE"},
		{"MOVE"},
		{"MOVE"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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

func TestValidLeftCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{"PLACE 1,1,NORTH", 1, 1, consts.WEST},
		{"PLACE 1,1,EAST", 1, 1, consts.NORTH},
		{"PLACE 1,1,SOUTH", 1, 1, consts.EAST},
		{"PLACE 1,1,WEST", 1, 1, consts.SOUTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			ProcessUserInput("LEFT", &robot, tableTop)
			
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


func TestInvalidLeftBeforePlaceFormat(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{"LEFT"},
		{"LEFT"},
		{"LEFT"},
		{"LEFT"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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


func TestValidRightCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{"PLACE 1,1,NORTH", 1, 1, consts.EAST},
		{"PLACE 1,1,EAST", 1, 1, consts.SOUTH},
		{"PLACE 1,1,SOUTH", 1, 1, consts.WEST},
		{"PLACE 1,1,WEST", 1, 1, consts.NORTH},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			ProcessUserInput("RIGHT", &robot, tableTop)
			
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


func TestInvalidRightBeforePlaceFormat(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{"RIGHT"},
		{"RIGHT"},
		{"RIGHT"},
		{"RIGHT"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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




func TestValidReportCommand(t *testing.T) {
	cases := []struct {
		inCommand string
		outX int 
		outY int
		outF consts.DIRECTION
	}{
		{"PLACE 1,2,NORTH", 1, 2, consts.NORTH},
		{"PLACE 2,1,EAST", 2, 1, consts.EAST},
		{"PLACE 1,0,SOUTH", 1, 0, consts.SOUTH},
		{"PLACE 0,1,WEST", 0, 1, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			ProcessUserInput("REPORT", &robot, tableTop)
			
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

func TestInvalidReportBeforePlaceFormat(t *testing.T) {
	cases := []struct {
		inCommand string
	}{
		{"REPORT"},
		{"REPORT"},
		{"REPORT"},
		{"REPORT"},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
			ProcessUserInput(c.inCommand, &robot, tableTop)
			
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