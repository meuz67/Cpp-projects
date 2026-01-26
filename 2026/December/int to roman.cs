using System;
using System.Runtime.InteropServices;
namespace program
{
    class Program
    {
        static string ToRoman(int num)
        {
            string result = "";
            Dictionary<int, string> list = new Dictionary<int, string>()
            {
                { 1000, "M" }, { 900, "CM" }, { 500, "D" }, { 400, "CD" }, { 100, "C" }, { 90, "XC" }, { 50, "L" },
                { 40, "XL" }, { 10, "X" }, { 9, "IX" }, { 5, "V" }, { 4, "IV" }, { 1, "I" }
            };
            foreach (var item in list)
            {
                while (num >= item.Key)
                {
                    result += item.Value;
                    num -= item.Key;
                }
            }
            return result; 
        }
        static void Main(string[] args)
        {
            int num = Convert.ToInt32(Console.ReadLine());
            Console.WriteLine(ToRoman(num));
        }
    }
}