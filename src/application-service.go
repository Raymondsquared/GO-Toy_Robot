package main

import (
    "consts"
    "models"
    "strings"
    "strconv"
    "commands"
)

///<summary>
/// Application Service for processing input and invoke the method in business logic / domain layer
///</summary>


/* 
 * Process inputs to command handler :
 * 
 * accepted commands:
 * PLACE X,Y,F
 * MOVE
 * LEFT
 * RIGHT
 * REPORT
 */
func ProcessUserInput(inputTexts string, robot *models.Robot, tableTop *models.TableTop) (bool, bool) {
	
	var isValid = false
	var shouldExit = false
	
	var command = strings.Replace(inputTexts, "\n", "", -1);
			
	// PLACE X,Y,F 
	if strings.HasPrefix(command, "PLACE") {
		
		var placeCommands = strings.Split(command, " ")
		
		if len(placeCommands) == 2 {
			
			var placeParams = strings.Split(placeCommands[1], ",")
			
			if len(placeParams) == 3 {

				// X, Y						
				x, errX := strconv.Atoi(placeParams[0])
				y, errY := strconv.Atoi(placeParams[1])
				
				//direction
				var direction = consts.UNKNOWN_DIRECTION
			
				switch placeParams[2] {
					case "NORTH":
						direction = consts.NORTH
					case "EAST":
						direction = consts.EAST
					case "SOUTH":
						direction = consts.SOUTH
					case "WEST":
						direction = consts.WEST
				}
				
				if direction != consts.UNKNOWN_DIRECTION && errX == nil && errY == nil {
					commands.PlaceCommand(robot, tableTop, x, y, direction)
					isValid = true
				}
			}
		}
	} else { //MOVE, LEFT, RIGHT, REPORT and EXIT
		switch command {
			case "MOVE":
				commands.MoveCommand(robot)
				isValid = true
			case "LEFT":
				commands.TurnCommand(robot, consts.LEFT)
				isValid = true
			case "RIGHT":
				commands.TurnCommand(robot, consts.RIGHT)
				isValid = true
			case "REPORT":
				commands.ReportCommand(robot)
				isValid = true
			case "EXIT":
				shouldExit = true
			default:
		}	
	}
	
	return isValid, shouldExit
}