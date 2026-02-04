using System;
namespace ConsoleApp1;
class Program
{
    static int Sum(int a, int b)
    {
        return a + b;
    }
    static int Minus(int a, int b)
    {
        return a - b;
    }
    static int Multiply(int a, int b)
    {
        return a * b;
    }
    static int Divide(int a, int b)
    {
        return a / b;
    }
    public delegate int Option(int a, int b);
    static void Main(string[] argsStrings)
    {
        Option sum = Sum;
        int result = sum(5, 10);
        Console.WriteLine(result);
    }
}