using System;
using System.Collections.Generic;
using System.Runtime.CompilerServices;
namespace ConsoleApp1
{
    internal class Program
    {
        public delegate void PrintResult(string message);
        static double Centre(int[] arr1, int[] arr2)
        {
            double sum = 0;
            int[] arr3 = new int[arr1.Length + arr2.Length];
            for (int i = 0; i < arr1.Length; i++)
                arr3[i] = arr1[i];

            for (int i = 0; i < arr2.Length; i++)
                arr3[i + arr1.Length] = arr2[i];
            for (int i = 0; i < arr3.Length; i++)
            {
                sum += arr3[i];
            }
            double result = sum / arr3.Length;
            return result;
        }
        static void Main(string[] args)
        {
            PrintResult printer = Print;
            int[] arr1 = { 1, 2, 3, 4, };
            int[] arr2 = { 3, 4, 5, 6, };
            PrintArr(arr1);
            PrintArr(arr2);
            Print(Centre(arr1, arr2).ToString());
        }
        static void PrintArr(int[] arr)
        {
            foreach (int i in arr)
            {
                Console.Write(i + " ");
            }
            Console.WriteLine();
        }
        static void Print(string Message) 
        {  
            Console.WriteLine(Message); 
        }
    }
}

