using System;
using System.Diagnostics;
namespace ConsoleApp1
{
    partial class Person
    {
        public string GetFullName()
        {
            return firstname + " " + lastname;
        }
        public void PrintFullName()
        {
            Console.WriteLine(GetFullName());
        }
    }
}