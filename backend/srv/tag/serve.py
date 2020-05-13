from flask import Flask, request
import numpy as np
import json
from config import path_hyper_parameters
from text_preprocess import PreprocessTextMulti, load_json
from graph import TextCNNGraph as Graph

app = Flask(__name__)


@app.route('/')
def index():
    text = request.args['text']

    ques_embed = ra_ed.sentence2idx(text)
    x_val_1 = np.array([ques_embed[0]])
    x_val_2 = np.array([ques_embed[1]])
    x_val = [x_val_1, x_val_2]

    pred = graph.predict(x_val)
    pre = pt.prereocess_idx(pred[0])
    ls_nulti = []
    for ls in pre[0]:
        if ls[1] >= 0.5:
            ls_nulti.append(ls[0])

    return json.dumps({"tag": ls_nulti}, ensure_ascii=False)


if __name__ == "__main__":
    hyper_parameters = load_json(path_hyper_parameters)
    pt = PreprocessTextMulti()
    graph = Graph(hyper_parameters)
    graph.load_model()
    ra_ed = graph.word_embedding
    ques_embed = ra_ed.sentence2idx("sjtujj")
    x_val_1 = np.array([ques_embed[0]])
    x_val_2 = np.array([ques_embed[1]])
    x_val = [x_val_1, x_val_2]
    pred = graph.predict(x_val)
    app.run()
