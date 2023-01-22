

namespace Prototype
{
    class Program
    {
        private const string HtmlPath = "tests/form.html.css.js/index.html";
        private const string CssPath = "tests/form.html.css.js/style.css";
        private const string JsPath = "tests/form.html.css.js/script.js";
        static void Main()
        {

            // {
            //     var stylesheet = new StylesheetParser().Parse(File.ReadAllText(CssPath));

            //     foreach (var rule in stylesheet.StyleRules)
            //     {
            //         Console.WriteLine();
            //     }

            // }
            // return;

            // IConfiguration config = Configuration.Default;

            // //Create a new context for evaluating webpages with the given config
            // IBrowsingContext context = BrowsingContext.New(config);

            // //Just get the DOM representation
            // IDocument document = await context.OpenAsync(req => req.Content(File.ReadAllText(HtmlPath)));

            // //Serialize it back to the console
            // Console.WriteLine(document.DocumentElement.OuterHtml);
            // return;

            // // 
            // var lexer = new JavascriptLexer(File.ReadAllText(JsPath));
            // var parser = new JavascriptParser(lexer);
            // var result = parser.Parse();
            // if (result.IsSuccess)
            // {
            //     Console.WriteLine(result.Root);
            // }
            // else
            // {
            //     foreach (var err in result.Errors)
            //     {
            //         Console.WriteLine(err);
            //     }
            // }
        }
    }
}