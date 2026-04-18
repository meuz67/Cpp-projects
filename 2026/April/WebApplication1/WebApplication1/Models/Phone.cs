using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models;

public class Phone
{
    [Key]
    public int Id { get; set; }
    [Required]
    [MaxLength(50)]
    public string Name { get; set; }
    [Required]
    [Range(0, int.MaxValue)]
    public decimal Price { get; set; }
}