package commands

import (
    "consts"
    "models"
)

///<summary>
/// Concrete implementation of ICommand interface for Command Pattern
///</summary>
///<remarks>
/// Handling all turning commmands (left / right), send the command to the receiver to do some actions
/// LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing the position of the robot.
///</remarks>
func TurnCommand(robot *models.Robot, turn consts.TURN) { 
    if robot.IsValid == true {
        switch robot.Direction {
            case consts.NORTH:
                if turn == consts.LEFT {
                    robot.Direction = consts.WEST
                }
                if turn == consts.RIGHT {
                    robot.Direction = consts.EAST
                }
            case consts.EAST:
                if turn == consts.LEFT {
                    robot.Direction = consts.NORTH
                }
                if turn == consts.RIGHT {
                    robot.Direction = consts.SOUTH
                }
            case consts.SOUTH:
                if turn == consts.LEFT {
                    robot.Direction = consts.EAST
                }
                if turn == consts.RIGHT {
                    robot.Direction = consts.WEST
                }
            case consts.WEST:
                if turn == consts.LEFT {
                    robot.Direction = consts.SOUTH
                }
                if turn == consts.RIGHT {
                    robot.Direction = consts.NORTH
                }
        }
    }
}