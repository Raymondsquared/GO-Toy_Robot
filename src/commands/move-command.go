package commands

import (
    "consts"
    "models"
)

///<summary>
/// Concrete implementation of ICommand interface for Command Pattern
///</summary>
///<remarks>
/// Handling all move commmands, send the command to the receiver to do some actions
/// MOVE will move the toy robot one unit forward in the direction it is currently facing.
///</remarks>
func MoveCommand(robot *models.Robot) { 
    if robot.IsValid == true {
        switch robot.Direction {
            case consts.NORTH:
                if robot.Y < robot.TableTop.Length {
                    robot.Y++
                }
            case consts.EAST:
                if robot.X < robot.TableTop.Width {
                    robot.X++
                }
            case consts.SOUTH:
                if robot.Y > 0 {
                    robot.Y--
                }
            case consts.WEST:
                if robot.X > 0 {
                    robot.X--
                }
        }
    }
}