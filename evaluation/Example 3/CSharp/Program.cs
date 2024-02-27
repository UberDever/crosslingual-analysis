using VB;

namespace CSharp
{
    class A : BaseVB
    {
        public A(int field_)
        {
            field = field_;
        }
        public void Print()
        {
            System.Console.WriteLine(this.field);
        }
    }
    class Program
    {
        static void Main(string[] args)
        {
            var a = new A(10);
            a.Print();
        }
    }
}
