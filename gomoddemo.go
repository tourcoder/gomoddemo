package gomoddemo

import "fmt" 

/***
 *
 * go mod demo
 *  
 ***/

func PrintHelloTo (name string) string {

   return fmt.Sprintf("Hello, %s", name)

}