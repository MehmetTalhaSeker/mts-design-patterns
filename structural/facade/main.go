package main

func main() {
	complexDB := &allDbOperations{}

	justDBGetter := NewDBGetter(*complexDB)

	// just hided all other operations except gets from client -> simplified interface.

	justDBGetter.GetByID("23")
	justDBGetter.Get()

	//justDBGetter.Create() // -> won't work
}
