namespace Prototype
{
    interface IAnalyzer
    {
        void Analyze(string program, string programName = "Anon");
        IEnumerable<NodeInfo> DumpAnalysis();
    };

    class Position
    {
        public int Line { get; init; }
        public int Collumn { get; init; }
        public int Length { get; init; }
        public string FileName { get; init; } = "";
    }

    class NodeInfo
    {
        public enum IntentType
        {
            None,
            Want,
            Give
        }

        public Position Position { get; init; } = new Position { };
        public IntentType Intent { get; set; }
        public string DataKind { get; set; } = "";
        public Dictionary<string, string> Data { get; set; } = new Dictionary<string, string> { };

    };
}