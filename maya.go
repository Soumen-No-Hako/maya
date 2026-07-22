package main

import (
    "fmt"
    "os"
    "os/exec"
    "maya/util/constants/signal"
    "maya/event/trigger"
)

func main(){
    // loadBuffers()
    // Disable canonical mode (buffering) and echo
        stty := exec.Command("stty", "raw", "-echo")
        stty.Stdin = os.Stdin
        if err := stty.Run(); err != nil {
                fmt.Println("Terminal state modification failed:", err)
                return
        }

        // Restore terminal settings to normal upon exit
        defer func() {
                restore := exec.Command("stty", "sane")
                restore.Stdin = os.Stdin
                restore.Run()
        }()
        //Initialise buffers and other things
        //fmt.Println("Press keys to view input")
	fmt.Printf(mayaEvent.TriggerClearScreen())
        buffer := make([]byte, 32)
        for {
                // Read raw input directly from Standard Input
                n, err := os.Stdin.Read(buffer)
                if err != nil {
                        fmt.Println("Read failure:", err)
                        break
                }
                

                seq := buffer[:n]

                // Output the decimal bytes, hexadecimal representation, and string literal
                inp := string(seq)
                //var evnt = mayaEvent.Detect(inp)

                // Switch case for input

                switch inp {
                    case signals.UP_ARROW_INPUT:    mayaEvent.TriggerUpArrow()
                    case signals.DOWN_ARROW_INPUT:  mayaEvent.TriggerDownArrow()
                    case signals.LEFT_ARROW_INPUT:  mayaEvent.TriggerLeftArrow()
                    case signals.RIGHT_ARROW_INPUT: mayaEvent.TriggerRightArrow()
                    case signals.DEL_INPUT:         mayaEvent.TriggerDel()
                    case signals.HOME_INPUT:        mayaEvent.TriggerHome()
                    case signals.END_INPUT:         mayaEvent.TriggerEnd()
                    default:                        mayaEvent.TriggerCharacterInsert(inp)
                }
                /*if len(inp)==1 {
                fmt.Printf("%s",inp)
                } else {
                fmt.Printf("Bytes: %-10v | Hex: %-8x | String: %q\n", seq, seq, string(seq))
                }*/
                // Termination condition
                if n == 1 && seq[0] == 'q' {
                        break
                }
        }

}
