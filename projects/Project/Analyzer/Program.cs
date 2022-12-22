using Project;
using CommandLine;
using System.Diagnostics;
namespace Project
{
    class Options
    {
        [Option('r', "read", Required = true, HelpText = "Input files to be processed.")]
        public IEnumerable<string>? InputFiles { get; set; }

        // Omitting long name, defaults to name of property, ie "--verbose"
        [Option(
          Default = false,
          HelpText = "Prints all messages to standard output.")]
        public bool Verbose { get; set; }
    }

    class Program
    {
        static void Main(string[] args)
        {
            CommandLine.Parser.Default.ParseArguments<Options>(args)
            .WithParsed(AnalyzeFiles)
            .WithNotParsed((IEnumerable<Error> errs) =>
            {
                foreach (var err in errs)
                {
                    Trace.TraceError(err.ToString());
                }
            });
        }

        static void AnalyzeFiles(Options opts)
        {
            var files = opts.InputFiles;
            files
            ?.Select(file => File.ReadAllText(file))
            ?.Select(input => Project.Frontend.GetResult(input))
            ?.Select(result =>
            {
                if (result.Errors.Count() != 0)
                {
                    foreach (var err in result.Errors)
                    {
                        Trace.TraceError(err.ToString());
                    }
                    return null;
                }
                var analyzer = new Analyzer();
                analyzer.analyze(result);
                return analyzer;
            })
            ?.Select(analyzer => { analyzer?.dump(); return 0; });
        }

    }
}