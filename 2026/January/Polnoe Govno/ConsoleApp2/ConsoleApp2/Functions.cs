using System;
using System.Net.Mail;

namespace Functions
{
    public static class Functions
    {
        public static  bool IsDayOfWeek(this DateTime date, DayOfWeek day)
        {
            return date.DayOfWeek == day;
        }
    }
}