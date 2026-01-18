namespace ConsoleApp1
{
   
    internal class Program
    {
        public delegate void AccountHandler(string message);
        class Account
        {
            public int sum;
            AccountHandler? taken;
            public void RegisterHandler(AccountHandler handler) 
            { 
            taken = handler;
            }
            public void Add(int sum)
            {
                this.sum += sum;
                taken.Invoke($"You have {sum}");
            }
            public void Remove(int sum)
            {   
                if(this.sum >= sum)
                {
                    this.sum -= sum;
                    taken.Invoke($"You have {sum}");
                }
                else
                {
                    taken.Invoke("You don`t have enough money");
                }
            }
        }
        static void printMessage(string message) {
            Console.WriteLine(message);
        }
        static void Main(string[] args)
        {
            Account account = new Account();
            account.RegisterHandler(printMessage);
            account.Add(1);
        }
    }
}
