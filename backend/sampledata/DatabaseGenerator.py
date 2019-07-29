# encoding:utf-8
import random
import sys

reload(sys)
sys.setdefaultencoding("utf8")

alpha = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
         'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
         'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
         'u', 'v', 'w', 'x', 'y', 'z']
names = []


def randStr(length, candidate):
    string = ""
    for i in range(0, length):
        string += alpha[random.randint(0, candidate - 1)]
    return string


def Unicode():
    val = random.randint(0x4e00, 0x9fbf)
    return unichr(val)


def readFile(filename):
    fin = open(filename)
    for line in fin:
        names.append(line.rstrip('\n').decode("utf-8"))
    fin.close()


def generateGood():
    readFile("goodbase.txt")

    fout = open("good.csv", 'w')
    for i in range(0, 200000):
        fout.write(str(i) + ",")
        length = random.randint(1, 5)
        for j in range(0, length):
            fout.write(names[random.randint(0, len(names) - 1)].encode("utf-8"))
        fout.write("," + str(float(random.randint(0, 999999)) / 100) + ",")
        length = random.randint(10, 50)
        for j in range(0, length):
            fout.write(names[random.randint(0, len(names) - 1)].encode("utf-8"))
        fout.write("," + randStr(24, 16) + "\n")
    fout.close()


def generateUser():
    readFile("userbase.txt")

    fout = open("user.csv", 'w')
    for i in range(0, 10000):
        fout.write(str(i) + ",")
        fout.write(names[random.randint(0, len(names) - 1)].encode("utf-8") + ",")
        fout.write(randStr(24, 16) + ",")
        fout.write(randStr(11, 10) + ",")
        fout.write("517" + randStr(9, 10) + ",")
        for j in range(0, 3):
            fout.write(Unicode())
        fout.write("," + str(random.randint(1, 1000) / 999 + 1) + "," + str(random.randint(1, 1000) / 999 + 1) + "\n")
    fout.close()


def generateSellInfo():
    fout = open("sellinfo.csv", 'w')
    for i in range(0, 200000):
        fout.write(str(i) + ",")
        fout.write(str(random.randint(1, 999) / 200 + 1) + ",")
        fout.write(str(random.randint(1, 1565390608)) + ",")
        fout.write(str(random.randint(1, 1655390608)) + ",")
        fout.write(str(random.randint(0, 9999)) + ",")
        fout.write(str(random.randint(0, 199999)) + "\n")
    fout.close()


generateGood()
generateUser()
generateSellInfo()
