using Microsoft.EntityFrameworkCore;
using WebApplication1.Models;

namespace WebApplication1.Data;

public class ApplicationDbContext : DbContext
{
    public DbSet<Phone> Phones { get; set; }
    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
        : base(options){}

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Phone>().HasData(
            new Phone() { Id = 1, Name = "Iphone 2G", Price = 42000 },
            new Phone() { Id = 2, Name = "Iphone 3G", Price = 67000 }
        );
    }
}