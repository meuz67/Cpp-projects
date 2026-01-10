using System;

namespace ConsoleApp1
{
    internal class Program
    {
        static void HelloWorld(string msg) { 
            Console.WriteLine(msg); 
        }
        static void Main(string[] args)
        {
            HelloWorld("Console.WriteLine");
        }
    }
}