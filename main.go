// Code generated by a developer but lint is complaining so we added this line DO NOT EDIT.
// TODO: remove the above line
//
// Package goengine_dev is a set of packages that provide many tools for event sourcing.
//
// goengine_dev contains the following packages:
//
// The aggregate package provides all components needed to create aggregates and record onto them.
//
// TODO write something nice about every sub package
package goengine_dev

// In order to avoid lots having to manually write similar code use `go generate` to generate the code
//go:generate go run tools/eventstore_inmemory_gen.go

// In order to make sure that we have the same mocks we can regenerate them using `go generate`
//go:generate go run vendor/github.com/vektra/mockery/cmd/mockery/mockery.go -dir eventstore/ -name=EventStore -outpkg mocks -output ./mocks -case underscore
//go:generate go run vendor/github.com/vektra/mockery/cmd/mockery/mockery.go -dir eventstore/ -name=PayloadFactory -outpkg mocks -output ./mocks -case underscore
//go:generate go run vendor/github.com/vektra/mockery/cmd/mockery/mockery.go -dir messaging/ -name=Message -outpkg mocks -output ./mocks -case underscore

// blank imports help docs.
import (
	// aggregate package
	_ "github.com/hellofresh/goengine/aggregate"
	// messaging package
	_ "github.com/hellofresh/goengine/messaging"
	// mocks package
	_ "github.com/hellofresh/goengine/mocks"
)
