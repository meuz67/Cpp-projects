using System;
using System.Net.Mail;
using Functions;
using classes;
namespace ConsoleApp
{
    class ConsoleApp1
    {
        static void Main(string[] args)
        {
            Person person = new Person();
            person.name = "John";
            person.surname = "Doe";
            person.thirdname = "John";
            person.Full();
        }
    }        
}