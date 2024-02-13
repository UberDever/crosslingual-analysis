struct counter {
    int count;
};

struct counter counter_new() {
    return (struct counter){0};
}

int counter_get(struct counter self) {
    return self.count;
}

void counter_reset(struct counter* self) {
    self->count = 0;
}

void counter_inc(struct counter* self) {
    self->count += 1;
}