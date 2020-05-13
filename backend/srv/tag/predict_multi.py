# -*- coding: UTF-8 -*-

# 地址
from config import path_hyper_parameters
# 数据预处理, 删除文件目录下文件
from text_preprocess import PreprocessTextMulti, load_json
# 模型图
from graph import TextCNNGraph as Graph

import numpy as np


def pred_input(path_hyper_parameter=path_hyper_parameters):
    # 输入预测
    # 加载超参数
    hyper_parameters = load_json(path_hyper_parameter)
    pt = PreprocessTextMulti()
    # 模式初始化和加载
    graph = Graph(hyper_parameters)
    graph.load_model()
    ra_ed = graph.word_embedding
    ques = '我要打王者荣耀'
    # str to token
    ques_embed = ra_ed.sentence2idx(ques)
    if hyper_parameters['embedding_type'] == 'bert':
        x_val_1 = np.array([ques_embed[0]])
        x_val_2 = np.array([ques_embed[1]])
        x_val = [x_val_1, x_val_2]
    else:
        x_val = ques_embed
    # 预测
    pred = graph.predict(x_val)
    print(pred)
    # 取id to label and pred
    pre = pt.prereocess_idx(pred[0])
    ls_nulti = []
    for ls in pre[0]:
        if ls[1] >= 0.5:
            ls_nulti.append(ls)
    print(pre[0])
    print(ls_nulti)
    while True:
        print("请输入: ")
        ques = input()
        ques_embed = ra_ed.sentence2idx(ques)
        print(ques_embed)
        if hyper_parameters['embedding_type'] == 'bert':
            x_val_1 = np.array([ques_embed[0]])
            x_val_2 = np.array([ques_embed[1]])
            x_val = [x_val_1, x_val_2]
        else:
            x_val = ques_embed
        pred = graph.predict(x_val)
        pre = pt.prereocess_idx(pred[0])
        ls_nulti = []
        for ls in pre[0]:
            if ls[1] >= 0.5:
                ls_nulti.append(ls)
        print(pre[0])
        print(ls_nulti)

if __name__=="__main__":
    # 测试集预测
    # pred_tet(path_test=path_byte_multi_news_valid, rate=1) # sample条件下设为1,否则训练语料可能会很少

    # 可输入 input 预测
    pred_input()
