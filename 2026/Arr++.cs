using System;

class ConsoleApp1
{
    static void Main()
    {
        int[] arr = { 1, 2, 3, 4, 5 };
        for (int i = 0; i < arr.Length; i++)
        {
            arr[i] = arr[i] + 1;
            Console.WriteLine(arr[i]);
        }
    }
}