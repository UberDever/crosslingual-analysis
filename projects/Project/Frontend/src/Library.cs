using Hime.Redist;
using Csharp;

namespace Project
{
    public class Frontend
    {
        static public ParseResult GetResult(string input)
        {
            var lexer = new CSharpLexer(input);
            var parser = new CSharpParser(lexer);
            return parser.Parse();
        }

    }
}