using HtmlAgilityPack;

namespace Prototype
{
    class HTMLAnalyzer : IAnalyzer
    {
        private readonly JSAnalyzer _jsAnalyzer;

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

        }

        public IEnumerable<NodeInfo> DumpAnalysis()
        {
            return new List<NodeInfo> { };
        }
    }
}