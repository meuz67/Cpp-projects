using System;
using System.Threading.Tasks;

namespace ConsoleApp1
{
    class Program
    {
        static async Task Counter()
        {
            for (int i = 0; i < 10; i++)
            {
                Console.WriteLine($"first counter: {i}");
                await Task.Delay(1000);
            }
        }

        static async Task FromTenToZero()
        {
            for (int i = 10; i > 0; i--)
            {
                Console.WriteLine($"second counter: {i}");
                await Task.Delay(1000);
            } 
        }
        static async Task Main(string[] args)
        {
            Task t = FromTenToZero();
            Task b = Counter();
            await Task.WhenAll(t, b);
        }
    }
}