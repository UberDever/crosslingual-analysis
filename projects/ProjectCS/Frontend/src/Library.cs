using Hime.Redist;
using MathExp;

namespace Project
{
    public class Frontend
    {
        static public ParseResult GetResult(string input)
        {
            var lexer = new MathExpLexer(input);
            var parser = new MathExpParser(lexer);
            return parser.Parse();
        }

    }
}