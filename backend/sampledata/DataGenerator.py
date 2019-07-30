import random


def readFile(filename):
    names = []
    with open(filename) as f:
        for line in f:
            names.append(line.rstrip('\n'))
    return names


def randCHN():
    head = random.randint(0xb0, 0xf7)
    body = random.randint(0xa1, 0xf9)
    val = f'{head:x}{body:x}'
    ret = bytes.fromhex(val).decode('gb2312')
    return ret


def generateGood():
    names = readFile("goodbase.txt")

    with open("good.csv", 'w') as f:
        for i in range(1, 200001):
            f.write(str(i) + ",")
            length = random.randint(1, 5)
            for _ in range(0, length):
                f.write(names[random.randint(0, len(names) - 1)])
            f.write("," + str(float(random.randint(0, 999999)) / 100) + ",")
            length = random.randint(10, 50)
            for _ in range(0, length):
                f.write(names[random.randint(0, len(names) - 1)])
            f.write(
                "," + "".join([random.choice("0123456789abcdef")for i in range(24)]) + "\n")


def generateUser():
    names = readFile("userbase.txt")

    with open("user.csv", 'w') as f:
        for i in range(1, 10001):
            f.write(str(i) + ",")
            f.write(names[random.randint(0, len(names) - 1)] + ",")
            f.write("".join([random.choice("0123456789abcdef")
                             for i in range(24)]) + ",")
            f.write("".join([str(random.randint(0, 9))
                             for num in range(0, 11)]) + ",")
            f.write("517" + "".join([str(random.randint(0, 9))
                                     for num in range(0, 9)]) + ",")
            for _ in range(0, 3):
                f.write(randCHN())
            f.write("," + str(random.randint(1, 1000) / 999 + 1) +
                    "," + str(random.randint(1, 1000) / 999 + 1) + "\n")


def generateSellInfo():
    with open("sellinfo.csv", 'w') as f:
        for i in range(1, 200001):
            f.write(str(i) + ",")
            f.write(str(random.randint(1, 999) / 200 + 1) + ",")
            f.write(str(random.randint(1, 1565390608)) + ",")
            f.write(str(random.randint(1, 1655390608)) + ",")
            f.write(str(random.randint(0, 9999)) + ",")
            f.write(str(random.randint(0, 199999)) + "\n")
        f.close()


generateGood()
generateUser()
generateSellInfo()
