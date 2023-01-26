using HtmlAgilityPack;

namespace Prototype
{
    class HTMLAnalyzer : IAnalyzer
    {
        private readonly JSAnalyzer _jsAnalyzer;
        private List<NodeInfo> _info = new List<NodeInfo> { };

        public HTMLAnalyzer(JSAnalyzer jsAnalyzer)
        {
            _jsAnalyzer = jsAnalyzer;
        }

        private string _programName = "Anon";
        public void Analyze(string program, string programName = "Anon")
        {
            _programName = programName;

            var dom = new HtmlDocument();
            dom.LoadHtml(program);
            var body = dom.DocumentNode.SelectSingleNode("//body");
            TraverseTree(body);
        }

        private void TraverseTree(HtmlNode node)
        {
            if (node.NodeType == HtmlNodeType.Element)
            {
                var info = new NodeInfo
                {
                    Position = new Position
                    {
                        Line = node.Line,
                        Collumn = node.LinePosition,
                        Length = 0,
                        FileName = _programName
                    },
                    Intent = NodeInfo.IntentType.Give,
                    DataKind = "html-element"
                };
                info.Data["Tag"] = node.OriginalName;
                foreach (var attr in node.Attributes)
                {
                    info.Data[attr.OriginalName] = attr.Value;
                    AnalyzeAttribute(attr);
                }
                _info.Add(info);
                foreach (var child in node.ChildNodes)
                {
                    TraverseTree(child);
                }
            }
        }

        private void AnalyzeAttribute(HtmlAttribute attr)
        {
            switch (attr.OriginalName)
            {
                case "oninput":
                case "onclick":
                    {
                        _jsAnalyzer.Analyze(attr.Value, _programName);
                        break;
                    }
            }
        }

        public IEnumerable<NodeInfo> DumpAnalysis()
        {
            return _info.Concat(_jsAnalyzer.DumpAnalysis());
        }
    }
}