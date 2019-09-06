# encoding: UTF-8

import pandas as pd
from sklearn.model_selection import train_test_split

products = pd.read_csv("data/original_data/products.csv")
label = pd.read_csv("data/original_data/categories.csv")

print((len(products)))
print(products.isnull().sum())

products = products.dropna()

train, test = train_test_split(products.values, test_size=0.2, random_state=0)

print(len(train))
print(len(test))

with open("data/good_tags/good_tags_train.csv", "w", encoding='utf-8') as file:
    for (i, row) in enumerate(train):
        tags = str(row[2]).split(',')
        tagstr = ','.join([label["category"].iloc[int(tag)] for tag in tags])
        file.write(tagstr+"|,|"+str(row[1])+'\n')

with open("data/good_tags/good_tags_test.csv", "w", encoding='utf-8') as file:
    for (i, row) in enumerate(test):
        tags = str(row[2]).split(',')
        tagstr = ','.join([label["category"].iloc[int(tag)] for tag in tags])
        file.write(tagstr+"|,|"+str(row[1])+'\n')

with open("data/good_tags/good_tags_label.csv", "w", encoding='utf-8') as file:
    for (i, row) in enumerate(label.values):
        file.write(str(row[1])+'\n')

