# trying algorithm in python first



input2 = [[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]

mem = {}

def findMax(input, x, y):
    if mem.get("x"+str(x)+"y"+str(y)) != None:
        return mem["x"+str(x)+"y"+str(y)]
    if y == len(input)-1:
        return input[y][x]
    
    mem["x"+str(x)+"y"+str(y+1)] = findMax(input, x,  y+1)
    mem["x"+str(x+1)+"y"+str(y+1)] = findMax(input, x+1,  y+1)
    return input[y][x] + max(mem["x"+str(x)+"y"+str(y+1)], mem["x"+str(x+1)+"y"+str(y+1)])

print(findMax(input,0,0))
    



# def decode(ch, l):
#     if ch == "" and not (l[-1] < 0 or l[-1] >2):
#         print(l)
#         return sum(l)
    
#     if len(l) == 0:
#         if ch[0] == "L":
#             return min(
#                 decode(ch[1:], [2,1]),
#                 decode(ch[1:], [2,0]),
#                 decode(ch[1:], [1,0])
#             )
#         elif ch[0] == "R":
#             return min(
#                 decode(ch[1:], [0,1]),
#                 decode(ch[1:], [0,2]),
#                 decode(ch[1:], [1,2])
#             )
#         else:
#             return min(
#                 decode(ch[1:], [0,0]),
#                 decode(ch[1:], [1,1]),
#                 decode(ch[1:], [2,2])
#             )
#     else:
#         if l[-1] < 0 or l[-1] >2:
#             return float('inf')
#         elif ch[0] == "L":
#             return  min(
#                 decode(ch[1:], l+[l[-1]-1]),
#                 decode(ch[1:], l+[l[-1]-2])
#             )
#         elif ch[0] == "R":
#             return min(
#                 decode(ch[1:], l+[l[-1]+1]),
#                 decode(ch[1:], l+[l[-1]+2])
#             )
#         elif ch[0] == "=":
#             return decode(ch[1:], l+[l[-1]])
#     return float('inf')

def logic(ch, num):
    res = int(ch) + num
    if res >=0 and res <= 2:
        return str(res)
    return "Z"


def decode(ch, l):
    if ch == "" and not (l[-1] < "0" or l[-1] >"2"):
        return l
    
    if len(l) == 0:
        if ch[0] == "L":
            return min(
                decode(ch[1:], "21"),
                decode(ch[1:], "20"),
                decode(ch[1:], "10")
            )
        elif ch[0] == "R":
            return min(
                decode(ch[1:], "01"),
                decode(ch[1:], "02"),
                decode(ch[1:], "12")
            )
        else:
            return min(
                decode(ch[1:], "00"),
                decode(ch[1:], "11"),
                decode(ch[1:], "22")
            )
    else:
        if l[-1] < "0" or l[-1] >"2":
            return "ZZZZZZZ"
        elif ch[0] == "L":
            return  min(
                decode(ch[1:], l+logic(l[-1], -1)),
                decode(ch[1:], l+logic(l[-1], -2))
            )
        elif ch[0] == "R":
            return min(
                decode(ch[1:], l+logic(l[-1], +1)),
                decode(ch[1:], l+logic(l[-1], +2))
            )
        elif ch[0] == "=":
            return decode(ch[1:], l+l[-1])
    return "ZZZZZZZ"


print(decode("LLRR=", [])) #output = 210122
print("+++++++++")
print(decode("==RLL", [])) #output = 000210
print("+++++++++")

print(decode("=LLRR", [])) #output = 221012
print("+++++++++")

print(decode("RRL=R", [])) #output = 012001
print("+++++++++")
