using System;
class ConsoleApp1
{
    static void PrintArr(int[] arr)
    {
        for (int i = 0; i < arr.Length; i++)
        {
            Console.Write(arr[i] +  " ");
        }

        Console.WriteLine();
    }

    static int[] ArrSize(ref int[] arr, int size)
    {
        int[] temp = new int[size];
        for (int i = 0; i < arr.Length; i++)
        {
            temp[i] = arr[i];
        }
        arr = temp;
        return arr;
    }
        static void Main()
    {
        int[] arr = {1, 4, 6};
        PrintArr(arr);
        arr = ArrSize(ref arr, 5);
        PrintArr(arr);
    }
}