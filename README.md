# XML performance comparison

   * Create test file with `go run generate.go`
   * Run `bash test.sh` (requires .NET Core 2.0 installed)

# Output on MacOS using Go 1.9

```
$ bash test.sh
Build .NET
Microsoft (R) Build Engine version 15.3.409.57025 for .NET Core
Copyright (C) Microsoft Corporation. All rights reserved.

  dotnet -> /Users/adrian/go/src/github.com/a-h/sax/dotnet/bin/Debug/netcoreapp2.0/dotnet.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:02.69
Build Go
Run .NET with timing
.NET: Found 0 houses

real    0m3.133s
user    0m2.908s
sys     0m0.150s
Run Go with timing
Go: Found 0 houses

real    0m20.326s
user    0m19.833s
sys     0m0.499s
```