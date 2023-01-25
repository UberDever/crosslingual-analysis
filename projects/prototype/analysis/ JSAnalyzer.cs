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
        public string SymbolPattern { get; init; }
        public string ValuePattern { get; init; }
        public List<int> Paths { get; init; }
        public string ActionName { get; init; }
        public bool EndOfChain { get; init; }
    }

    struct Option
    {
        public Regex SymbolPattern { get; init; }
        public Regex ValuePattern { get; init; }
        public List<int> Paths { get; init; }
        public Action<ASTNode> Action { get; init; }
        public bool EndOfChain { get; init; }

        // TODO: Add quantity
    }


    class TreeMatcher
    {
        private readonly Matchers _matchers;
        private int _currentMatcher = 0;
        private Queue<ASTNode> _nodes = new Queue<ASTNode> { };
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
            }
            for (int i = 0; i < options?.Count; i++)
            {
                var option = options[i];
                if (option.ValuePattern.Match(value).Success && option.Paths.Contains(currentOption))
                {
                    _currentMatcher++;
                    _paths.Push(i);
                    _nodes.Enqueue(node);
                    haveEaten = true;
                    if (option.EndOfChain)
                    {
                        ChainCompleted();
                    }
                    break;
                }
            }

            if (!haveEaten)
            {
                _currentMatcher = 0;
                _paths.Clear();
                _nodes.Clear();
            }
        }

        private void ChainCompleted()
        {
            _currentMatcher = 0;
            var chain = _paths.Reverse().Select((optionI, matcherI) => _matchers[matcherI][optionI]);
            foreach (var (option, node) in chain.Zip(_nodes))
            {
                option.Action(node);
            }
        }
    }

    class JSAnalyzer : IAnalyzer
    {
        enum State
        {
            None,
            GetElementById,
            GetElementByTagName,
            GetElementByClassName,
            FunctionDeclaration
        };

        private readonly ConnectionMultiplexer _multiplexer;
        private readonly IConfiguration _configuration;
        private TreeMatcher _matcher;
        private State _state;


        public JSAnalyzer(ConnectionMultiplexer multiplexer, IConfiguration configuration)
        {
            _multiplexer = multiplexer;
            _configuration = configuration;

            var matchers = configuration.GetSection("patterns").Get<List<List<OptionRaw>>>();
            var convertedMatchers = matchers?.Select(options =>
                options.Select(option =>
                {
                    var methodInfo = this.GetType().GetMethod(option.ActionName ?? "");
                    return new Option
                    {
                        SymbolPattern = new Regex(option.SymbolPattern ?? ""),
                        ValuePattern = new Regex(option.ValuePattern ?? ""),
                        Paths = option.Paths,
                        Action = methodInfo is not null ? (Action<ASTNode>)Delegate.CreateDelegate(
                            type: typeof(Action<ASTNode>),
                            method: methodInfo,
                            firstArgument: this
                        ) : (ASTNode node) => { },
                        EndOfChain = option.EndOfChain
                    };
                }).ToList()).ToList() ?? new Matchers();
            _matcher = new TreeMatcher(convertedMatchers);
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

        public void onGetElementById(ASTNode node)
        {
            _state = State.GetElementById;
        }

        public void onGetElementByClassName(ASTNode node)
        {
            _state = State.GetElementByClassName;
        }
        public void onGetElementByTagName(ASTNode node)
        {
            _state = State.GetElementByTagName;
        }

        public void onArguments(ASTNode node)
        {
            switch (_state)
            {
                case State.GetElementById:
                    {
                        Console.WriteLine($"ID {node.ToString()}");
                        break;
                    }
                case State.GetElementByClassName:
                    {
                        Console.WriteLine($"Class {node.ToString()}");
                        break;
                    }
                case State.GetElementByTagName:
                    {
                        Console.WriteLine($"Tag {node.ToString()}");
                        break;
                    }
            }
            _state = State.None;
        }

        public void onFunctionDeclaration(ASTNode node)
        {
            _state = State.FunctionDeclaration;
        }

        public void onFunctionName(ASTNode node)
        {
            System.Diagnostics.Debug.Assert(_state == State.FunctionDeclaration);


        }

        public void TraverseAST(ASTNode node)
        {
            // Console.WriteLine(root.ToString());
            _matcher.Eat(node);
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