package main

func TurnCommand(robot *Robot, turn TURN) { 
    if robot.IsValid == true {
        switch robot.Direction {
            case NORTH:
                if turn == LEFT {
                    robot.Direction = WEST
                }
                if turn == RIGHT {
                    robot.Direction = EAST
                }
            case EAST:
                if turn == LEFT {
                    robot.Direction = NORTH
                }
                if turn == RIGHT {
                    robot.Direction = SOUTH
                }
            case SOUTH:
                if turn == LEFT {
                    robot.Direction = EAST
                }
                if turn == RIGHT {
                    robot.Direction = WEST
                }
            case WEST:
                if turn == LEFT {
                    robot.Direction = SOUTH
                }
                if turn == RIGHT {
                    robot.Direction = NORTH
                }
        }
    }
}