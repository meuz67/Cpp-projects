using System;
using System.Threading.Tasks;
class Program
{
    public static void findSum(int[] arr, int task)
    {
        for (int i = 0; i < arr.Length; i++)
        {
            for (int j = 0; j < arr.Length; j++)
            {
                if ((arr[i] + arr[j]) == task)
                {
                    Console.WriteLine($"{arr[i]} and {arr[j]}");
                }
            }
        }
    } 
    static void Main(string[] args)
    {
        Random rand = new Random();
        int task = 15;
        int[] arr = new int[10];
        for (int i = 0; i < arr.Length; i++)
        {
            arr[i] = rand.Next(1, 11);
            Console.Write($"{arr[i]} ");
        }
        Console.WriteLine();
        findSum(arr, task);
    } 
}