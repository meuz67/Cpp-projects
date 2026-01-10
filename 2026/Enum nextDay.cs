using System;

namespace ConsoleApp1
{
    internal class Program
    {
        enum DayOfWeek 
         {
            Monday = 1, 
            Tuesday = 2,
            Wednesday = 3,
            Thursday = 4,
            Friday = 5,
            Saturday = 6,
            Sunday = 7
         }
        static DayOfWeek GetNextDay(DayOfWeek day)
        {
            if(day < DayOfWeek.Sunday)
            {
                return day + 1;
            }
            else
            {
                return day;
            }
        }

        static void Main()
        {
            DayOfWeek day = DayOfWeek.Monday;
            day = GetNextDay(day);
            Console.WriteLine(day);
        }
    }
}
