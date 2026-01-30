using System;
using Gdk;
using Gtk;
using Window = Gtk.Window;

namespace ConsoleApp1
{
    class Program
    {
        public class UpgradedButton : Button
        {
            public int numberOfButton;
            public UpgradedButton(string label, int numberOfButton) : base(label)
            {
                SetSizeRequest(100, 100);
                this.numberOfButton = numberOfButton;
            }
        }
        static void Main(string[] args)
        {
            var provider = new Gtk.CssProvider();
            provider.LoadFromData("button {\n    border-radius: 0;\n}");
            string Hello = "Hello";
            Application.Init();
            Window myWin = new Window(Hello);
            Fixed fixedContainer = new Fixed();
            Button button1 = new UpgradedButton("1", 1);
            Button button2 = new UpgradedButton("2", 2);
            Button button3 = new UpgradedButton("3", 3);
            Button button4 = new UpgradedButton("4", 4);
            Button button5 = new UpgradedButton("5", 5);
            Button button6 = new UpgradedButton("6", 6);
            Button button7 = new UpgradedButton("7", 7);
            Button button8 = new UpgradedButton("8", 8);
            Button button9 = new UpgradedButton("9", 9);
        fixedContainer.Put(button1, 0, 0);
            fixedContainer.Put(button2, 100, 0);
            fixedContainer.Put(button3, 200, 0);
            fixedContainer.Put(button4, 0, 100);
            fixedContainer.Put(button5, 100, 100);
            fixedContainer.Put(button6, 200, 100);
            fixedContainer.Put(button7, 0, 200);
            fixedContainer.Put(button8, 100, 200);
            fixedContainer.Put(button9, 200, 200);
            myWin.SetDefaultSize(300, 300);
            foreach (Widget w in fixedContainer.Children)
            {
                if (w is UpgradedButton btn)
                {
                    btn.Clicked += (s, e) => Console.WriteLine($"Clicked {btn.numberOfButton}"); 
                }
            }
            myWin.DeleteEvent += delegate{Application.Quit();};
            myWin.Add(fixedContainer);
            myWin.ShowAll();
            Application.Run();
        }
    }
}