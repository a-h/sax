echo "Build .NET"
dotnet build ./dotnet/dotnet.csproj
echo "Build Go"
cd go && go build && cd ..
echo "Run .NET with timing"
time dotnet ./dotnet/bin/Debug/netcoreapp2.0/dotnet.dll
echo "Run Go with timing"
time ./go/go