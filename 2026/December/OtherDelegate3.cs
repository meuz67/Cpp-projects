using System;
namespace ConsoleApp1;
class Program
{
    public delegate void AccountHandler(string message);
    public class Account
    {
        public int sum;
        public AccountHandler? taken;
        public Account(int sum)
        {
            this.sum = sum;
        }
        public void RegisterDel(AccountHandler handler)
        {
            taken =  handler;  
        }
        public void Add(int sum)
        {
            this.sum += sum;
            taken.Invoke("You added " + sum);
        }
        public void Take(int sum)
        {
            if (this.sum > sum)
            {
                this.sum -= sum;
                taken.Invoke("You taken " + sum);
            }
            else
            {
                taken.Invoke("You don`t have enough" + sum);
            }
        }
    }

    static void PrintSingleMessage(string message)
    {
        Console.WriteLine(message);
    }
    static void Main(string[] args)
    {
        Account account = new Account(200);
        account.RegisterDel(PrintSingleMessage);
        account.Take(100);
        account.Add(500);
    }
}