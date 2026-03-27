using System.ComponentModel.DataAnnotations.Schema;

namespace ConsoleApp1.Models;

public class Product
{
    public int Id { get; set; }
    public string? name { get; set; }
    [Column(TypeName = "decimal(6,2)")]
    public decimal price { get; set; }
}