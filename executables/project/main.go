/**
 *
 */
package main

import (
	"github.com/wplib/go-lib/project"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR:", r)
		}
	}()
	pf := project.NewProjectFile()
	pf.LoadJSON()
	for i, c:= range pf.GetComponents() {
		fmt.Printf("\n[%d] %-22v %v",i,c.GetType()+":",c.GetLocation())
	}
}


