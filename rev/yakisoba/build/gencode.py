import os
import random

flag = b'ctf4b{sp4gh3tt1_r1pp3r1n0}\0'

def gen_ret():
    c = random.randint(0, 0xff)
    if c == 0:
        return 1
    else:
        return c

def gen_group(target, complexity):
    obj = list(range(0x100))
    random.shuffle(obj)
    r = set(obj[:min(0x100, complexity)])
    r.add(target)
    return r

def gen_check(flag, depth=0, complexity=5, tab=0):
    sp = " " * tab
    code  = "%sswitch(flag[%d]) {\n" % (sp, depth)
    for t in gen_group(flag[depth], complexity):
        if depth == len(flag) - 1:
            if t == flag[depth]:
                code += "%s  case %d: return 0;\n" % (sp, t)
            else:
                code += "%s  case %d: return %d;\n" % (sp, t, gen_ret())
        else:
            if t == flag[depth]:
                code += "%s  case %d:\n" % (sp, t)
                code += gen_check(flag, depth + 1, complexity, tab + 2)
                code += "%s  break;\n" % sp
            else:
                code += "%s  case %d: return %d;\n" % (sp, t, gen_ret())
    code += "%s  default: return %d;\n" % (sp, gen_ret())
    code += "%s}\n" % sp
    return code

code = gen_check(flag)
with open("template.c", "r") as f:
    template = f.read()
output = template.replace("%%%%", code)
with open("main.c", "w") as f:
    f.write(output)

os.system("make")
