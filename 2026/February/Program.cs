using System;
using System.IO;
using System.Linq;
class Program
{
    static void SelectionSort(int[] unsorted)
    {
        int smallestIndex;
        for (int i = 0; i < unsorted.Length - 1; i++)
        {
            smallestIndex = i;
            for (int j = i + 1; j < unsorted.Length; j++)
            {
                if (unsorted[j] < unsorted[smallestIndex])
                {
                    smallestIndex = j;
                }
            }

            if (smallestIndex != i)
            {
                int temp = unsorted[i];
                unsorted[i] = unsorted[smallestIndex];
                unsorted[smallestIndex] = temp;
            }
    }
    }
    static void Main()
    {
        int[] unsorted = { 1, 254, 4534, 3, 654, 459, 764 };
        SelectionSort(unsorted);
        Console.WriteLine(string.Join(", ", unsorted));
    }
}