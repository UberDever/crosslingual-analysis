using Links = System.Collections.Generic.List<System.Tuple<Prototype.NodeInfo, Prototype.NodeInfo>>;

namespace Prototype
{
    class Linker
    {
        private readonly IEnumerable<NodeInfo> _info;
        public Linker(IEnumerable<NodeInfo> info)
        {
            _info = info;
        }

        public Links GetLinks()
        {
            var wants = new List<NodeInfo>();
            var gives = new List<NodeInfo>();
            foreach (var info in _info)
            {
                if (info.Intent == NodeInfo.IntentType.Want)
                {
                    wants.Add(info);
                }
                else
                {
                    gives.Add(info);
                }
            }

            var links = new Links();
            foreach (var want in wants)
            {
                foreach (var give in gives)
                {
                    if (want.DataKind == give.DataKind)
                    {
                        if (want.Data.All(pair =>
                            give.Data.ContainsKey(pair.Key) &&
                            give.Data[pair.Key] == pair.Value))
                        {
                            links.Add(new Tuple<NodeInfo, NodeInfo>(want, give));
                        }
                    }
                }
            }

            return links;
        }
    }
}