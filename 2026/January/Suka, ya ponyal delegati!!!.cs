using System;
using System.Runtime.InteropServices.JavaScript;
using System.Threading.Tasks;
namespace ConsoleApp1
{
    public class StockExchangeMonitor
    {
        public delegate void PriceChnge(int Price);
        public PriceChnge priceChangeHandler {get; set;}
        public void Start()
        {
            while (true)
            {
                int bankPrice = (new Random()).Next(5, 100);
                priceChangeHandler(bankPrice);
                Thread.Sleep(3000);
            }
        }
    }
    class Program
    {
        static void Main(string[] args)
        {
            StockExchangeMonitor stockExchangeMonitor = new StockExchangeMonitor();
            stockExchangeMonitor.priceChangeHandler = ShowPrice;
            stockExchangeMonitor.Start();
        }
        public static void ShowPrice(int price)
        {
            Console.WriteLine(price);
        }
    }
}