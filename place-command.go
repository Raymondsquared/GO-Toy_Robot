package main

func PlaceCommand(robot *Robot, tableTop *TableTop, x int, y int, direction DIRECTION) { 
    
    var isValid = true;
    
    if tableTop == nil {
        isValid = false;
    }
    
    if x < 0 || y < 0 {
        isValid = false;
    }
    
     if x > tableTop.Width || y > tableTop.Length {
        isValid = false;
    }

    robot.IsValid = isValid
    
    if isValid == true {
        robot.TableTop = *tableTop
        robot.X = x
        robot.Y = y
        robot.Direction = direction
    }

}