echo "Build .NET"
dotnet build ./dotnet/dotnet.csproj

echo "Build Go (standard library)"
cd go && go build && cd ..

echo "Build Go (github.com/tamerh/xml-stream-parser)"
cd go-xmlparser && go build && cd ..

echo "Run .NET with timing"
time dotnet ./dotnet/bin/Debug/netcoreapp2.0/dotnet.dll

echo "Run Go (standard library) with timing"
time ./go/go -profile=true

echo "Run Go (github.com/tamerh/xml-stream-parser) with timing"
time ./go-xmlparser/go-xmlparser -profile=true