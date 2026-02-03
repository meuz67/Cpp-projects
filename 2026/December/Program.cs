using System;
using System.Threading;
namespace ConsoleApp1
{
    class Program
    {
        static int x = 0;
        static object locker = new();
        static void Count()
        {
            x = 1;
            lock (locker)
            {
                for (int i = 0; i < 6; i++)
                {
                    Console.WriteLine($"{Thread.CurrentThread.Name} : {x}");
                    x++;
                    Thread.Sleep(1000);
                }
            }
        }
        static void Main(string[] args)
        {
            int x = 1;
            for (int i = 0; i < 10; i++)
            {
                Thread myThread = new Thread(Count);
                myThread.Name = "Thread" + i;
                myThread.Start();
            }
        }
    }
}