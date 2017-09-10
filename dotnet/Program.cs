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
