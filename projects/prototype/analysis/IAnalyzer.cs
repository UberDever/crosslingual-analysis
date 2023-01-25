namespace Prototype
{
    interface IAnalyzer
    {
        void Analyze(string program, string programName = "Anon");
        void DumpAnalysis();
    };

    class Position
    {
        public int Line { get; init; }
        public int Col { get; init; }
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
        public string Data { get; set; } = "";

    };
}