using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace Prototype
{
    class Program
    {
        private const string HtmlPath = "tests/form.html.css.js/index.html";
        private const string CssPath = "tests/form.html.css.js/style.css";
        private const string JsPath = "tests/form.html.css.js/script.js";

        private readonly JSAnalyzer _jsAnalyzer;
        private readonly HTMLAnalyzer _htmlAnalyzer;

        public Program(JSAnalyzer jsAnalyzer, HTMLAnalyzer htmlAnalyzer)
        {
            _jsAnalyzer = jsAnalyzer;
            _htmlAnalyzer = htmlAnalyzer;
        }

        public static IHostBuilder CreateHostBuilder(string[] args)
        {
            return Host.CreateDefaultBuilder(args)
                .ConfigureServices(services =>
                {
                    services.AddTransient<Program>();
                    services.AddTransient<JSAnalyzer>();
                    services.AddTransient<HTMLAnalyzer>();
                });
        }

        public void Run()
        {
            IEnumerable<NodeInfo> info = new List<NodeInfo> { };
            _jsAnalyzer.Analyze(File.ReadAllText(JsPath), JsPath);
            info = info.Concat(_jsAnalyzer.DumpAnalysis());
            _htmlAnalyzer.Analyze(File.ReadAllText(HtmlPath), HtmlPath);
            info = info.Concat(_htmlAnalyzer.DumpAnalysis());

            Console.WriteLine("Analysis results:");

            foreach (var i in info)
            {
                Console.WriteLine($"{i.Intent} {i.DataKind} at {i.Position.FileName}:{i.Position.Line}:{i.Position.Collumn}");
                foreach (var (key, val) in i.Data)
                {
                    Console.WriteLine("\twith " + key + " = " + val);
                }
            }

            Console.WriteLine();
            Console.WriteLine("Resulted pairs:");

            var linker = new Linker(info);
            var links = linker.GetLinks();

            foreach (var link in links)
            {
                var want = link.Item1;
                var give = link.Item2;
                Console.WriteLine($"{want.Intent} {want.DataKind} at {want.Position.FileName}:{want.Position.Line}:{want.Position.Collumn}");
                foreach (var (key, val) in want.Data)
                {
                    Console.WriteLine("\twith " + key + " = " + val);
                }

                Console.WriteLine($"{give.Intent} {give.DataKind} at {give.Position.FileName}:{give.Position.Line}:{give.Position.Collumn}");
                foreach (var (key, val) in give.Data)
                {
                    Console.WriteLine("\twith " + key + " = " + val);
                }
                Console.WriteLine();
            }
        }

        static void Main(string[] args)
        {
            var host = CreateHostBuilder(args).Build();
            // IConfiguration config = host.Services.GetRequiredService<IConfiguration>();
            host.Services.GetRequiredService<Program>().Run();

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