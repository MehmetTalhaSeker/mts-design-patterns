package main

import "fmt"

func myExecuteFunc(db DB) ThirdPartyLibFn {
	return func(s string) {
		fmt.Println("myExec ", s)
		db.Store(s)
	}
}
