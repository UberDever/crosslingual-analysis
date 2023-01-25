using StackExchange.Redis;
using Hime.Redist;
using Javascript;
using System.Text.RegularExpressions;
using Microsoft.Extensions.Configuration;

using Matchers = System.Collections.Generic.List<System.Collections.Generic.List<Prototype.Option>>;

namespace Prototype
{
    struct OptionRaw
    {
        public string Pattern { get; init; }
        public List<int> Paths { get; init; }
        public string ActionName { get; init; }
    }

    struct Option
    {
        public Regex Pattern { get; init; }
        public List<int> Paths { get; init; }
        public Action Action { get; init; }
    }


    class TreeMatcher
    {
        private readonly Matchers _matchers;
        private int _currentMatcher = 0;
        private (int, int) _chainStart = (0, 0);
        private (int, int) _chainEnd = (0, 0);
        private Stack<int> _paths = new Stack<int> { };

        public TreeMatcher(Matchers matchers)
        {
            _matchers = matchers;
        }

        public void Eat(ASTNode node)
        {
            if (node.Value is null)
            {
                return;
            }

            var options = _matchers?[_currentMatcher];
            var value = node.Value;
            bool haveEaten = false;
            if (!_paths.TryPeek(out int currentOption))
            {
                currentOption = -1;
                _chainStart = (node.Position.Line, node.Position.Column);
            }
            for (int i = 0; i < options?.Count; i++)
            {
                var option = options[i];
                if (option.Pattern.Match(value).Success && option.Paths.Contains(currentOption))
                {
                    _currentMatcher++;
                    _paths.Push(i);
                    haveEaten = true;
                    if (_currentMatcher == _matchers?.Count)
                    {
                        _chainEnd = (node.Position.Line, node.Position.Column);
                        ChainCompleted();
                    }
                    break;
                }
            }

            if (!haveEaten)
            {
                _currentMatcher = 0;
                _paths.Clear();
                _chainStart = (0, 0);
                _chainEnd = (0, 0);
            }
        }

        private void ChainCompleted()
        {
            _currentMatcher = 0;
            var chain = _paths.Reverse().Select((optionI, matcherI) => _matchers[matcherI][optionI].Pattern);
            foreach (var n in chain)
            {
                Console.Write(n.ToString() + " ");
            }
            Console.WriteLine($"{_chainStart} : {_chainEnd}");
        }
    }

    class JSAnalyzer : IAnalyzer
    {
        private readonly ConnectionMultiplexer _multiplexer;
        private readonly IConfiguration _configuration;
        private TreeMatcher matcher;


        public JSAnalyzer(ConnectionMultiplexer multiplexer, IConfiguration configuration)
        {
            _multiplexer = multiplexer;
            _configuration = configuration;

            var matchers = configuration.GetSection("patterns").Get<List<List<OptionRaw>>>();
            var convertedMatchers = matchers?.Select(options =>
                options.Select(option =>
                {
                    var methodInfo = this.GetType().GetMethod(option.ActionName);
                    if (methodInfo is not null)
                    {
                        return new Option
                        {
                            Pattern = new Regex(option.Pattern),
                            Paths = option.Paths,
                            Action = (Action)Delegate.CreateDelegate(
                                type: typeof(Action),
                                method: methodInfo,
                                firstArgument: this
                            )
                        };
                    }
                    else
                    {
                        return new Option
                        {
                            Pattern = new Regex(option.Pattern),
                            Paths = option.Paths,
                            Action = () => { }
                        };
                    }
                }).ToList()).ToList() ?? new Matchers();
            matcher = new TreeMatcher(convertedMatchers);
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

        private void onGetElementById()
        {
            Console.WriteLine("id");
        }

        private void onGetElementByClassName()
        {
            Console.WriteLine("class");
        }
        private void onGetElementByTagName()
        {
            Console.WriteLine("tag");
        }


        public void TraverseAST(ASTNode node)
        {
            // Console.WriteLine(root.ToString());
            matcher.Eat(node);
            foreach (var child in node.Children)
            {
                TraverseAST(child);
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