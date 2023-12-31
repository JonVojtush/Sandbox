https://code.visualstudio.com/docs/languages/go#_test

Test
The VS Code Test UI and editor CodeLens elements allow users to easily run tests, benchmarks, profiles for a given function, file, package, or workspace.

Alternatively, the same functionality is available through a set of commands:

Go: Test Function At Cursor
Go: Test File
Go: Test Package
Go: Test All Packages in Workspace

There are many test-related commands that you can explore by typing "Go: test" in the Command Palette.

>Go:test

Test Commands

The first three above can be used to generate test skeletons for the functions in the current package, file, or at the cursor using gotests. The last few can be used to run tests in the current package, file, or at the cursor using go test. There is also a command for getting test coverage.

You can configure the extension to run tests and compute test coverage using:

go.testOnSave
go.coverOnSave
go.testFlags