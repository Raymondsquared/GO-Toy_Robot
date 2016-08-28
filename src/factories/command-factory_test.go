package factories

import (
	"consts"
    "models"
    "testing"
)

func TestValidPlaceCommand(t *testing.T) {
	cases := []struct {
		inCommand consts.COMMAND
		inX int
		inY int
		inF consts.DIRECTION
	}{
		{consts.PLACE_COMMAND, 1, 1, consts.WEST},
		{consts.PLACE_COMMAND, 1, 1, consts.UNKNOWN_DIRECTION},
		{consts.PLACE_COMMAND, -1, 1, consts.WEST},
		{consts.PLACE_COMMAND, 1, -1, consts.WEST},
		{consts.PLACE_COMMAND, 6, 7, consts.WEST},
	}
	for _, c := range cases {
		
		tableTop, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
		    input := models.Input {c.inCommand, tableTop, c.inX, c.inY, c.inF}
		    
			var result = CommandFactory(&robot, input)
			
			if result == nil {
				t.Errorf("should return success from valid place command")
			}
		}
	}
}


func TestOtherValidCommand(t *testing.T) {
	cases := []struct {
		inCommand consts.COMMAND
	}{
		{consts.MOVE_COMMAND},
		{consts.LEFT_COMMAND},
		{consts.RIGHT_COMMAND},
		{consts.REPORT_COMMAND},
	}
	for _, c := range cases {
		
		_, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
		    input := models.Input {}
		    
		    input.Command = c.inCommand
		    
			var result = CommandFactory(&robot, input)
			
			if result == nil {
				t.Errorf("should return success from valid command")
			}
		}
	}
}


func TestInvalidCommand(t *testing.T) {
	cases := []struct {
		inCommand consts.COMMAND
	}{
		{consts.UNKNOWN_COMMAND},
	}
	for _, c := range cases {
		
		_, errTableTop := models.NewTableTop(5, 5)
		
		if errTableTop != nil {
			t.Errorf("cant throw error on creating table")
		} else {
			robot := models.Robot{}
		    
		    input := models.Input {}
		    
		    input.Command = c.inCommand
		    
			var result = CommandFactory(&robot, input)
			
			if result != nil {
				t.Errorf("should return nil from unknown command")
			}
		}
	}
}
