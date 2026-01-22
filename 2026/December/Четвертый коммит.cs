namespace ConsoleApp1
{
    class bank
    {
        public delegate void MyDelegate(string message);
        private int sum {  get; set; }
        public MyDelegate BankDelegate { get; set; }
        public void MoreMoney(int moreMoney)
        {
            sum+= moreMoney;
            BankDelegate?.Invoke($"You have {sum}");
        }
        public void LessMoney(int lessMoney)
        {
            sum-= lessMoney;
            BankDelegate?.Invoke($"You have {sum}");
        }
    }
    internal class Program
    {
        static void Main(string[] args)
        {
            bank account = new bank();
            account.BankDelegate = WriteMessage;
            account.MoreMoney(5);
        }
        public static void WriteMessage(string message)
        {
            Console.WriteLine(message);
        }
    }
}
