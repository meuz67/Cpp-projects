using System;
using System.Collections.Generic;

namespace program
{
    class Program
    {
        static void Main(string[] args)
        {
            string text = Console.ReadLine();
            Console.WriteLine(Refrased(text, -3));
        }

        static string Refrased(string text, int change)
        {
            string updatedText = "";
            string alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
            try
            {
                foreach (var c in text)
                {
                    int temp = alphabet.IndexOf(c);
                    updatedText += alphabet[temp + change];
                }

                return updatedText;
            }
            catch(Exception e)
            {
                throw new Exception("Mistake");
            }
        }
    }
}