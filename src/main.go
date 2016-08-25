package main

import (
	"os"
    "fmt"
    "models"
    "bufio"
)

///<summary>
/// Entry point of the application
///</summary>
///<remarks>
/// The application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units.
/// There are no other obstructions on the table surface.
/// The robot is free to roam around the surface of the table, but must be prevented from falling to destruction.Any movement that would result in the robot falling from the table must be prevented, however further valid movement commands must still be allowed.
///</remarks>

func main() {
	
	fmt.Printf("Toy Robot Simulator.\n")
	
    argsWithoutProg := os.Args[1:]
	
	tableTop, err := models.NewTableTop(5, 5)
		
	if err != nil {
	 	panic("cant throw error on creating table")
	} 
	
	robot := models.Robot{}
	
    /*
	 * handle file & keyboard input differently
	 * Input can be from a file, or from standard input, as the developer chooses.
	 */
	
	//from file
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "-f" {
		if len(argsWithoutProg) > 1 {
			var file = argsWithoutProg[1]
			fmt.Printf("Reading from file : %s ... \n", file)
		    
		    f, err := os.Open(file)
		    if err != nil {
				fmt.Printf("%v\n", err)
			}
			
		    scanner := bufio.NewScanner(f)
		    
		    for scanner.Scan() {
				lineText := scanner.Text()
				ProcessUserInput(lineText, &robot, tableTop);
		    }
		} else {
			fmt.Printf("missing file name \n")
		}
	} else { //from keyboard
		fmt.Printf("You may type now : \n")
		reader := bufio.NewReader(os.Stdin)
		
		for {
			text, _ := reader.ReadString('\n')
			success, shouldExit := ProcessUserInput(text, &robot, tableTop);
			
			if shouldExit {
				os.Exit(2)	
			}
			
			if !success {
				fmt.Printf("ALLOWED COMMANDS : PLACE X,Y,F & MOVE & LEFT & RIGHT & REPORT & EXIT\n")
			}
    	}
	}
}