module github.com/mehmetali10/opcua

go 1.20

require (
	github.com/google/uuid v1.3.0
	github.com/gopcua/opcua v0.0.0-00010101000000-000000000000
	github.com/pascaldekloe/goe v0.1.1
	github.com/pkg/errors v0.9.1
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	golang.org/x/term v0.8.0
)

require golang.org/x/sys v0.8.0 // indirect

// retract (
// 	v0.2.5 // https://github.com/gopcua/opcua/issues/538
// 	v0.2.4 // https://github.com/gopcua/opcua/issues/538
// )

replace github.com/gopcua/opcua => github.com/mehmetali10/opcua v1.0.0
