import sys

with open('jawiki-latest-all-titles') as f:
    lines_strip = [line.strip() for line in f.readlines()]

args = sys.argv
serchlist = args[1].split(",")

def serchresult(x):
    l_count = len([line for line in lines_strip if x in line])
    print(x +":"+ str(l_count))
    return l_count

loglist = list(map(serchresult, serchlist))