using System.Diagnostics.Metrics;

namespace ConsoleApp1;

class Car
{ 
    virtual public void Drive()
    {
        Console.WriteLine("Driving");
    }
}

class Person
{
    public void Drive(Car car)
    {
        car.Drive();
    }
}

class SportCar : Car
{
    public override void Drive()
    {
        Console.WriteLine("Driving fast");
    }
}
class Program
{
    static void Main(string[] args)
    {
        Car car = new Car();
        Person person = new Person();
        person.Drive(new SportCar());
    }
}