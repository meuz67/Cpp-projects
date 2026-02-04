using System;
using System.Diagnostics;
using ConsoleApp1;
namespace ConsoleApp1
{
    internal class Program
    {
        static void Main(string[] args)
        {
            Person person = new Person("firstname", "lastname");
            person.PrintFullName();
        }
    }
}