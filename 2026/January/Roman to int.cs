using System;
using System.Runtime.InteropServices;
using System.Collections.Generic;
namespace program
{
    class Program
    {
        static void Main(string[] args)
        {
            string input = Console.ReadLine().ToUpper();
            try
            {
                Console.WriteLine(RomanToInt(input));
            }
            catch (Exception e)
            {
                Console.WriteLine(e.Message);
            }
        }
        static int RomanToInt(string s)
        {
            int result = 0;
            Dictionary<char, int> romanDictionary = new Dictionary<char, int>()
            {
                { 'I', 1 }, { 'V', 5 }, { 'X', 10 }, { 'L', 50 }, { 'C', 100 }, { 'D', 500 }, { 'M', 1000 }
            };
            for (int i = 0; i < s.Length; i++)
            {
                int value = romanDictionary[s[i]];
                if (i + 1 < s.Length && romanDictionary[s[i+1]] > value)
                {
                    result -= value;
                }
                else
                {
                    result += value;
                }
            }
            return result;
        }
    }
}