package mayaEvent

import (
    "fmt"
    "maya/util/constants/signal"
)
func TriggerUpArrow(){
    fmt.Printf(signals.CURSOR_UP)
}
func TriggerDownArrow(){
    fmt.Printf(signals.CURSOR_DOWN)
}
func TriggerLeftArrow(){
    fmt.Printf(signals.CURSOR_BACKWARD)
}
func TriggerRightArrow(){
    fmt.Printf(signals.CURSOR_FORWARD)
}
func TriggerHome(){
    fmt.Printf(signals.CURSOR_HOME)
}
func TriggerEnd(){
    //fmt.Printf(signals.ERASE_CHARACTER)
}
func TriggerDel(){
    fmt.Printf(signals.DELETE_CHARACTER)
}
func TriggerClearScreen() string {
    return signals.ERASE_SCREEN
}
func TriggerCharacterInsert(input string){
    fmt.Printf(signals.INSERT_CHARACTER)
    fmt.Printf(input)
}
