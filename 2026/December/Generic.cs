namespace ConsoleApp1
{
    internal class Program
    {
        public static void Main(string[] args)
        {
            int[] intArr = { 1, 2, 3 };
            double[] doubleArr = { 1.0, 2.0, 3.0 };
            string[] stringArr = { "1", "2", "3" };
            displayElement(intArr);
            displayElement(doubleArr);
            displayElement(stringArr);
            Console.ReadKey();
        }
        public static void displayElement<T>(T[] arr)
        {
            foreach (T i in arr)
            {
                Console.Write(i + " ");
            }
            Console.WriteLine();
        }
    }
}
