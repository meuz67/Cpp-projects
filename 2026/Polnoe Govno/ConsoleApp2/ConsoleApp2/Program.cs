using System;
using System.Net.Mail;
using Functions;
namespace ConsoleApp
{
    class ConsoleApp1
    {
        static void Main(string[] args)
        {
            DateTime now = DateTime.Now;
            Console.WriteLine(now.IsDayOfWeek(DayOfWeek.Monday));
        }
    }        
}