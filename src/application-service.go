package main

import (
    "consts"
    "models"
    "strings"
    "strconv"
    "factories"
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
	
	var command = strings.Replace(inputTexts, "\n", "", -1)
	
	var input = models.Input {}
	
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
					input.Command = consts.PLACE_COMMAND
					input.TableTop = tableTop
					input.X = x
					input.Y = y
					input.Direction = direction
					isValid = true
				}
			}
		}
	} else { //MOVE, LEFT, RIGHT, REPORT and EXIT
		switch command {
			case "MOVE":
				input.Command = consts.MOVE_COMMAND
				isValid = true
			case "LEFT":
				input.Command = consts.LEFT_COMMAND
				isValid = true
			case "RIGHT":
				input.Command = consts.RIGHT_COMMAND
				isValid = true
			case "REPORT":
				input.Command = consts.REPORT_COMMAND
				isValid = true
			case "EXIT":
				shouldExit = true
			default:
		}	
	}
	
	cmd := factories.CommandFactory(robot, input)
	
	if cmd != nil {
		var inv = commands.NewInvoker(cmd)
		inv.Invoke()
	}
	
	return isValid, shouldExit
}