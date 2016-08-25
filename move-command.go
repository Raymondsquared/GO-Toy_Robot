package main

func MoveCommand(robot *Robot) { 
    if robot.IsValid == true {
        switch robot.Direction {
            case NORTH:
                if robot.Y < robot.TableTop.Length {
                    robot.Y++
                }
            case EAST:
                if robot.X < robot.TableTop.Width {
                    robot.X++
                }
            case SOUTH:
                if robot.Y > 0 {
                    robot.Y--
                }
            case WEST:
                if robot.X > 0 {
                    robot.X--
                }
        }
    }
}