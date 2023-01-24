using StackExchange.Redis;
using Hime.Redist;
using Javascript;
using System.Text.RegularExpressions;

namespace Prototype
{
    struct Match
    {
        private Regex pattern { get; init; }
        private List<int> paths { get; init; }
        private string actionName { get; init; }
    }

    class TreeMatcher
    {
        string json = @"[
            [{
                'pattern': '.',
                'paths': [-1],
                'action': ''
            }],
            [
                {
                    'pattern': 'getElement?ById',
                    'paths': [0],
                    'action': ''
                },
                {
                    'pattern': 'getElement?ByClassName',
                    'paths': [0],
                    'action': ''
                },
                {
                    'pattern': 'getElement?ByTagName',
                    'paths': [0],
                    'action': ''
                },
            ],
            [
                {
                    'pattern': '(',
                    'paths': [0, 1, 2],
                    'action': ''
                }
            ],
            [
                {
                    'pattern': '',
                    'paths': [0],
                    'action': ''
                }
            ],
            [
                {
                    'pattern': '(',
                    'paths': [0, 1, 2],
                    'action': ''
                }
            ]
        ]";

        private List<Regex[]> matchings;
        private int currentMatch = 0;
        private Regex varMatch = new Regex(":.*:");
        private bool ChainCompleted { get; }

        public TreeMatcher()
        {
            matchings = new List<Regex[]>();
            var options = new string[][] {
                new string[] { "." },
                new string[] { "getElement?ById", "getElement?ByClassName", "getElement?ByTagName" },
                new string[] { "(" },
                new string[] { """ ".*" """ },
                new string[] { ")" }};
            matchings = options.Select(cases =>
            {
                return cases.Select(option => new Regex(option, RegexOptions.Compiled)).ToArray<Regex>();
            }).ToList<Regex[]>();

        }

        public void eat(ASTNode node)
        {
            var options = matchings[currentMatch];
            var value = node.Value;
            var symbol = node.Symbol;
            bool haveEaten = false;
            foreach (var option in options)
            {
                if (option.Match(value).Success)
                {
                    currentMatch++;
                    haveEaten = true;
                    if (currentMatch == options.Length)
                    {
                        currentMatch = 0;
                    }
                    break;
                }
                else
                {

                }
            }

            if (!haveEaten)
            {
                currentMatch = 0;
            }
        }


    }


    class JSAnalyzer : IAnalyzer
    {
        private readonly ConnectionMultiplexer _multiplexer;



        public JSAnalyzer(ConnectionMultiplexer multiplexer)
        {
            _multiplexer = multiplexer;
        }

        public void Analyze(string program)
        {
            var lexer = new JavascriptLexer(program);
            var parser = new JavascriptParser(lexer);
            var result = parser.Parse();
            if (result.IsSuccess)
            {
                TraverseAST(result.Root);
                // DumpTree(result.Root, new bool[] { });
            }
            else
            {
                foreach (var err in result.Errors)
                {
                    Console.WriteLine(err);
                }
                System.Environment.Exit(-1);
            }

        }

        public void TraverseAST(ASTNode root)
        {
            Console.WriteLine(root.ToString());
            foreach (var node in root.Children)
            {
                TraverseAST(node);
            }
        }

        private void DumpTree(ASTNode node, bool[] crossings)
        {
            for (int i = 0; i < crossings.Length - 1; i++)
                Console.Write(crossings[i] ? "|   " : "    ");
            if (crossings.Length > 0)
                Console.Write("+-> ");
            Console.WriteLine(node.ToString());
            for (int i = 0; i != node.Children.Count; i++)
            {
                bool[] childCrossings = new bool[crossings.Length + 1];
                Array.Copy(crossings, childCrossings, crossings.Length);
                childCrossings[childCrossings.Length - 1] = (i < node.Children.Count - 1);
                DumpTree(node.Children[i], childCrossings);
            }
        }
    }
}