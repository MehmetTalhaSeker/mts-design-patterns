package main

type ThirdPartyLibFn func(string)

func Execute(fn ThirdPartyLibFn) {
	fn("BAR BAZ FOO")
}
