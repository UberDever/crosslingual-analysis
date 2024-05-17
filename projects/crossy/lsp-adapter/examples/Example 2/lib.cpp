class Counter
{
public:
    Counter() : count{0} {}

    int Get()
    {
        return count;
    }

    void Reset()
    {
        count = 0;
    }

    void Inc()
    {
        count++;
    }

private:
    int count;
};

extern "C"
{
    Counter *counter_new()
    {
        return new Counter();
    }

    void counter_free(Counter *self)
    {
        delete self;
    }

    int counter_get(Counter *self)
    {
        return self->Get();
    }

    void counter_reset(Counter *self)
    {
        self->Reset();
    }

    void counter_inc(Counter *self)
    {
        self->Inc();
    }
}