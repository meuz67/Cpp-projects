using System;
namespace ConsoleApp1;
class Program
{
    static void DoOperation(int a, int b, Operation op)
    {
        Console.WriteLine(op(a,b));
    }
    static int Add(int a, int b)
    {
        return a + b;
    }
    static int Subtract(int a, int b)
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
    delegate int Operation(int a, int b);
    static void Main(string[] args)
    {
        Operation op;
        op = Subtract;
        DoOperation(1,2,op);
    }
}