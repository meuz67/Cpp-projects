namespace BlazorApp1;

public class ContactService
{
    public List<Contact> Contacts = new List<Contact>();
    public void Add(Contact contact)
    {
        Contacts.Add(contact);
    }
}