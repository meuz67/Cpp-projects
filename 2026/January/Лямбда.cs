namespace ConsoleApp1
{
    internal class Program
    {
        public static void Main(string[] args) 
        {
            var add = (int x, int y) =>
            {
                return x + y;
            };
            Console.WriteLine(add(1, 2));
        }
    }
}
