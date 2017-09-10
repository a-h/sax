using System;
using System.Xml;

namespace dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
            var filename = "generate.xml";
            var count = CountHouses(filename);
            Console.WriteLine($".NET: Found {count} houses");
        }

        private static int CountHouses(string filename)
        {
            int houses = 0;
            using (XmlReader reader = XmlReader.Create(filename))
            {
                while (reader.Read())
                {
                    switch (reader.NodeType)
                    {
                        case XmlNodeType.Element:
                            if (reader.Name == "House")
                            {
                                // The Go program decodes the element, but the .NET one doesn't.
                                houses++;
                            }
                            break;
                    }
                }
            }
            return houses;
        }
    }
}
