var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () =>
{
    return $"The time is {DateTime.UtcNow.ToString("h:mm tt")}, UTC.";
})

app.Run();

