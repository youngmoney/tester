package main

import ()

func Match(path string, matchers *[]Test) *Test {
	for _, m := range *matchers {
		if m.MatchPathRegex.MatchString(path) {
			return &m
		}
	}
	return nil
}
