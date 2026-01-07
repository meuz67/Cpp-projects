using System;
using System.Net.Http.Headers;

namespace ConsoleApp1
{
    internal class Program
    {
        static void Main(string[] args)
        {
            int userInput;
            Console.Write("Write size:");
            int size = int.Parse(Console.ReadLine());
            int[] arr = new int[size];
            for (int i = 0; i < arr.Length; i++) 
            { 
                userInput = int.Parse(Console.ReadLine());
                arr[i] = userInput;
            }
            for (int j = 0; j < arr.Length; j++) 
            {
                Console.WriteLine($"Element number {j} is {arr[j]}");
            }
            Console.ReadKey();
        }
    }
}
