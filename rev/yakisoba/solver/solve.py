import angr
import claripy
from logging import getLogger, WARN

getLogger("angr").setLevel(WARN + 1)
getLogger("claripy").setLevel(WARN + 1)

flag = claripy.BVS("flag", 8 * 0x20)
p = angr.Project("../files/yakisoba", load_options={"auto_load_libs": False})
state = p.factory.entry_state(stdin=flag)
simgr = p.factory.simulation_manager(state)

simgr.explore(find=0x4006d2, avoid=0x4006f7)

try:
    found = simgr.found[0]
    print(found.solver.eval(flag, cast_to=bytes))
except IndexError:
    print("Not Found")
