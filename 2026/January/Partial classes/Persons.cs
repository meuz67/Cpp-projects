using System;
using System.Diagnostics;
namespace ConsoleApp1
{
    partial class Person
    {
        public string firstname { get; set; }
        public string lastname { get; set; }
        public Person(string firstname, string lastname)
        {
            this.firstname = firstname;
            this.lastname = lastname;
        }
        
    }
}