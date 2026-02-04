using System;
using System.Runtime.InteropServices;
namespace program
{
    abstract class Point
    {
        public double x;
        public double y;

        protected Point(double x, double y)
        {
            this.x = x;
            this.y = y;
        }
    }

    class A : Point
    {
        public A(double x, double y) : base(x, y){}
    }
    class B : Point
    {
        public B(double x, double y) : base(x, y)
        {
        }
    }
class Program
    {
        static double Distance(A a, B b)
        {
            double result = Math.Sqrt(Math.Pow((b.x - a.x), 2) + Math.Pow((b.y - a.y), 2));
            return result;
        }
        static void Main(string[] args)
        {
            A a = new A(9.6, 15.4);
            B b = new B(41.0, 54.2);
            Console.WriteLine(Distance(a, b));
        }
    }
}