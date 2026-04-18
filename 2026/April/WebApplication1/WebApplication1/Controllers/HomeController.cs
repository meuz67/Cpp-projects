using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using WebApplication1.Data;
using WebApplication1.Models;

namespace WebApplication1.Controllers;

public class HomeController : Controller
{
    private readonly ApplicationDbContext _context;
    public HomeController(ApplicationDbContext context)
    {
        _context = context;
    }
    public IActionResult Index()
    {
        var phones = _context.Phones.ToList();
        phones.OrderBy(x => x.Id);
        return View(phones);
    }
    [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
    public IActionResult Error()
    {
        return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
    }
    public IActionResult Create()
    {
        return View();
    }

    [HttpPost]
    public IActionResult Create(Phone phone)
    {
        CheckCategory(phone);
        if (ModelState.IsValid)
        {
            _context.Phones.Add(phone);
            _context.SaveChanges();
            return RedirectToAction("Index");
        }
        return View(phone);
    }

    public IActionResult Edit(int? id)
    {
        return View(_context.Phones.Find(id));
    }

    [HttpPost]
    public IActionResult Edit(Phone phone)
    {
        CheckCategory(phone);
        if (ModelState.IsValid)
        {
            _context.Phones.Update(phone);
            _context.SaveChanges();
            return RedirectToAction("Index");
        }
        return View(phone);
    }
    [HttpPost]
    public IActionResult Delete(int? id)
    {
        if (id == null)
        {
            return NotFound();
        }
        var phone = _context.Phones.Find(id);
        _context.Phones.Remove(phone);
        _context.SaveChanges();
        return RedirectToAction("Index");
    }
    private void CheckCategory(Phone phone)
    {
        bool exist = _context.Phones.Any(x => x.Name == phone.Name && x.Id != phone.Id);
        if (exist)
        {
            ModelState.AddModelError("Name", "Phone name already exists");
        }
    }
}