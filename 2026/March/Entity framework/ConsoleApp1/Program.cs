using ConsoleApp1.Data;
using ConsoleApp1.Models;

namespace ConsoleApp1;

class Program
{
    static void Main(string[] args)
    {
        using ContosoPizzaContext context = new ContosoPizzaContext(); 
        var products = context.Products
            .Where(p => p.price < 10.00M)
            .OrderBy(p => p.Id);
        foreach (var product in products)
        {
            Console.WriteLine("Id = " + product.Id);
            Console.WriteLine("Name = " + product.name);
            Console.WriteLine("Price = " + product.price);
            for (int i = 0; i < 20; i++)
            {
                Console.Write("-");
            }
            Console.WriteLine();
        }
    }
}