using System;
namespace ConsoleApp1
{
    internal class Program
    {
        static void PrintArr(int[] arr)
        {
            for (int i = 0; i < arr.Length; i++)
            {
                Console.Write(arr[i] + " ");
            }

        }
        static int[] RemoveAt(ref int[] arr, int at)
        {
            int[] newArr = new int[arr.Length - 1];
            for (int i = 0; i < at; i++)
                newArr[i] = arr[i];

            for (int i = at + 1; i < arr.Length; i++)
                newArr[i - 1] = arr[i];
            arr = newArr;
            return arr;   
        }
        static void Main()
        {
            int[] arr = { 1, 2, 3, 4 };
            arr = RemoveAt(ref arr, 2);
            PrintArr(arr);
        }
    }
}
