using System;
using Avalonia.Controls;
using Avalonia.Interactivity;
using System.Diagnostics;
namespace AvaloniaApplication1;

public partial class MainWindow : Window
{
    public MainWindow()
    {
        InitializeComponent();  
    }

    private void CelsiusToFahrenheit(object sender, RoutedEventArgs e)
    {
        if (double.TryParse(Celsius.Text, out double celsius))
        {
            var F = celsius * 9d / 5d + 32;
            Fahrenheit.Text = F.ToString("0.0");
        }
        else
        {
            Celsius.Text = "0";
            Fahrenheit.Text = "0";
        }
    }

    private void FahrenheitToCelsius(object sender, RoutedEventArgs e)
    {
        if (double.TryParse(Fahrenheit.Text, out double fahrenheit))
        {
            var C = (fahrenheit - 32) / 1.8;
            Celsius.Text = C.ToString("0.0");
        }
        else
        {
            Celsius.Text = "0";
            Fahrenheit.Text = "0";
        }
    }
}