using System;
namespace ConsoleApp1
{
    internal class Program
    {
        static string UpAndLow(string text)
        {
            char[] result = new char[text.Length];
            for (int i = 0; i < result.Length; i++) 
            { 
                if(i%2 == 0)
                {
                    result[i] = Char.ToUpper(text[i]);
                }
                else
                {
                    result[i] = Char.ToLower(text[i]);
                }
            }
            return new string(result);
        }
        static void Main()
        {
            string text = Console.ReadLine();
            string result = UpAndLow(text);
            Console.WriteLine(result);
        }
    }
}