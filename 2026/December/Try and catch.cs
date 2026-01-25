using System;
using System.Runtime.InteropServices;
namespace  program
{
    class Person
    {
        private int age;
        public string Name { get; set; } = "";
        public int Age
        {
            get => age;
            set
            {
                if (value < 18) 
                    throw new Exception("Less than 18");
                else
                    age = value;
            }
        }
    }
    class Program
    {
        static void Main(string[] args)
        {
            try
            {
                Person Tom = new Person { Name = "Tom", Age = 17 };
            }
            catch (Exception e)
            {
                Console.WriteLine(e.Message);
            }
        }
    }
}