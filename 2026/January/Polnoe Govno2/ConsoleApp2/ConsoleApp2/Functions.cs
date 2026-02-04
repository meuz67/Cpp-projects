using System;
using System.Net.Mail;
using classes;
namespace Functions
{
    public static class Functions
    {
        public static void Full(this Person person)
        {
            Console.WriteLine($"{person.name} {person.surname} {person.thirdname}");
        }
    }
}