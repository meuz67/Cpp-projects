using System;
using System.Threading.Tasks;
class Program
{
    public delegate void NotifyHandler();
    class Ping
    {
         public event NotifyHandler OnPing;
         public void GetPing()
         {
             Console.WriteLine("Ping received Pong");
             OnPing?.Invoke();
         }
    }
    class Pong
    {
        public event NotifyHandler OnPong;
        public void GetPong()
        {
            Console.WriteLine("Pong received Ping");
            OnPong?.Invoke();
        }
    }
    static void Main(string[] args)
    {
        Ping ping = new Ping();
        Pong pong = new Pong();
        ping.OnPing += pong.GetPong;
        pong.OnPong += ping.GetPing;
        ping.GetPing();
    }
}