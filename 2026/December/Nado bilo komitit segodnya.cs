using System;
using System.Net.Security;

namespace ConsoleApp1
{
    class Counter
    {
        public int Value { get; set; }

        public static Counter operator +(Counter c1, Counter c2)
        {
            return new Counter { Value = c1.Value + c2.Value };
        }

        public static bool operator >(Counter c1, Counter c2)
        {
            return c1.Value > c2.Value;
        }

        public static bool operator <(Counter c1, Counter c2)
        {
            return c1.Value < c2.Value;
        }

        public static int operator +(Counter c1, int c2)
        {
            return c1.Value + c2;
        }
    }
    class Program
    {
        static void Main(string[] args)
        {
            Counter counter1 = new Counter {Value = 55 };
            Counter counter2 = new Counter {Value = 66 };
            int result = counter1 + 55;
            Console.WriteLine(result);
        }
    }
}