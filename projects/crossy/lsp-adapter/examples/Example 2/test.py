import ctypes
import sys

usage = """
    test.py <scenario: fast | full> <number of increments>
"""

lib = ctypes.CDLL('lib_dir/lib.so')

class Counter(ctypes.Structure):
    _fields_ = [
        ('count', ctypes.c_int)
    ]

lib.counter_new.argtypes = []
lib.counter_new.restype = ctypes.POINTER(Counter)

lib.counter_get.argtypes = [ctypes.POINTER(Counter)]
lib.counter_get.restype = ctypes.c_int

lib.counter_reset.argtypes = [ctypes.POINTER(Counter)]
lib.counter_reset.restype = None

lib.counter_inc.argtypes = [ctypes.POINTER(Counter)]
lib.counter_inc.restype = None

lib.counter_free.argtypes = [ctypes.POINTER(Counter)]
lib.counter_free.restype = None

args = sys.argv[1:]

scenario = args[0]
increments = int(args[1])

counter = lib.counter_new()

for i in range(increments):
    lib.counter_inc(counter)

if scenario == 'full':
    print("Before reset:", lib.counter_get(counter))
    lib.counter_reset(counter)
    print("After reset:", lib.counter_get(counter))
    for i in range(increments):
        lib.counter_inc(counter)

print("Counter:", lib.counter_get(counter))

lib.counter_free(counter)